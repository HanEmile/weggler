package runner

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"weggler/internal/models"
	"weggler/internal/store"
)

var headerRe = regexp.MustCompile(`^(.+):(\d+)$`)

// parseFindings parses weggli plain-text output (ANSI already stripped) into
// Finding structs, and simultaneously builds SnippetHTML from the raw ANSI
// output so the browser can render colour highlights.
func parseFindings(plain, raw string, foundAt time.Time) []models.Finding {
	var findings []models.Finding
	var cur *models.Finding
	var snippetLines []string
	var snippetHTMLLines []string

	// Split raw (ANSI) in parallel with plain so line indices match.
	rawLines := strings.Split(raw, "\n")
	plainLines := strings.Split(plain, "\n")
	// Pad to equal length just in case.
	for len(rawLines) < len(plainLines) {
		rawLines = append(rawLines, "")
	}

	flush := func() {
		if cur != nil {
			cur.Snippet = strings.Join(snippetLines, "\n")
			cur.SnippetHTML = strings.Join(snippetHTMLLines, "\n")
			findings = append(findings, *cur)
			cur = nil
			snippetLines = nil
			snippetHTMLLines = nil
		}
	}
	for i, raw := range plainLines {
		line := strings.TrimRight(raw, "\r")
		rawLine := strings.TrimRight(rawLines[i], "\r")

		// A blank line separates findings — but only when we are NOT already
		// collecting snippet lines. Inside a snippet, blank lines are part of
		// the output (weggli uses them between the opening context and the match).
		if strings.TrimSpace(line) == "" {
			if cur == nil {
				continue // blank between findings, nothing to flush yet
			}
			// We're inside a snippet. Peek ahead: if the next non-blank line
			// is a new file:line header, this blank line ends the finding.
			// Otherwise keep it as part of the snippet.
			isEndOfFinding := false
			for j := i + 1; j < len(plainLines); j++ {
				next := strings.TrimRight(plainLines[j], "\r")
				if strings.TrimSpace(next) == "" {
					continue
				}
				// Next non-blank line: is it a new file:line header?
				if !strings.HasPrefix(next, " ") && !strings.HasPrefix(next, "\t") {
					if headerRe.MatchString(next) {
						isEndOfFinding = true
					}
				}
				break
			}
			if isEndOfFinding {
				flush()
			} else {
				// Blank line is part of the snippet
				snippetLines = append(snippetLines, line)
				snippetHTMLLines = append(snippetHTMLLines, "")
			}
			continue
		}

		if !strings.HasPrefix(line, " ") && !strings.HasPrefix(line, "\t") {
			if m := headerRe.FindStringSubmatch(line); m != nil {
				flush()
				n, _ := strconv.Atoi(m[2])
				cur = &models.Finding{File: m[1], Line: n, FoundAt: foundAt}
				snippetLines = nil
				snippetHTMLLines = nil
				continue
			}
		}
		if cur != nil {
			snippetLines = append(snippetLines, line)
			snippetHTMLLines = append(snippetHTMLLines, ansiToHTML(rawLine))
		}
	}
	flush()
	return findings
}

// scanSourceStats counts bytes and lines in C/C++ source files under path.
func scanSourceStats(path string) (totalBytes int64, totalLines int64) {
	_ = filepath.WalkDir(path, func(p string, d os.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return nil
		}
		ext := strings.ToLower(filepath.Ext(p))
		switch ext {
		case ".c", ".cpp", ".h", ".cc", ".cxx", ".c++", ".hh", ".hpp":
		default:
			return nil
		}
		f, err := os.Open(p)
		if err != nil {
			return nil
		}
		defer f.Close()
		scanner := bufio.NewScanner(f)
		scanner.Buffer(make([]byte, 1<<20), 1<<20)
		for scanner.Scan() {
			totalBytes += int64(len(scanner.Bytes())) + 1
			totalLines++
		}
		return nil
	})
	return
}

// Runner manages concurrent job execution with a configurable concurrency limit.
type Runner struct {
	mu            sync.Mutex
	store         *store.Store
	watchers      map[string][]chan string
	cancels       map[string]func()
	maxConcurrent int
	totalStarted  int64
	totalCompleted int64
	totalFindings  int64
}

func New(st *store.Store, maxConcurrent int) *Runner {
	if maxConcurrent < 1 {
		maxConcurrent = 4
	}
	return &Runner{
		store:         st,
		watchers:      make(map[string][]chan string),
		cancels:       make(map[string]func()),
		maxConcurrent: maxConcurrent,
	}
}

func (r *Runner) Metrics() map[string]any {
	r.mu.Lock()
	defer r.mu.Unlock()
	return map[string]any{
		"max_concurrent":  r.maxConcurrent,
		"running":         r.store.CountRunningJobs(),
		"queued":          r.store.CountQueuedJobs(),
		"total_started":   r.totalStarted,
		"total_completed": r.totalCompleted,
		"total_findings":  r.totalFindings,
	}
}

func (r *Runner) SetMaxConcurrent(n int) {
	r.mu.Lock()
	r.maxConcurrent = n
	r.mu.Unlock()
	r.tryDispatch()
}

func (r *Runner) ResumeRunning() {
	jobs := r.store.ListJobs()
	for _, j := range jobs {
		if j.Status == models.JobRunning || j.Status == models.JobPending {
			j.Status = models.JobFailed
			j.Output += "\n[runner restarted, job interrupted]"
			now := time.Now()
			j.FinishedAt = &now
			_ = r.store.UpdateJob(j)
		}
	}
	r.tryDispatch()
}

func (r *Runner) Subscribe(jobID string) <-chan string {
	ch := make(chan string, 256)
	r.mu.Lock()
	defer r.mu.Unlock()
	r.watchers[jobID] = append(r.watchers[jobID], ch)
	return ch
}

func (r *Runner) broadcast(jobID, line string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, ch := range r.watchers[jobID] {
		select {
		case ch <- line:
		default:
		}
	}
}

func (r *Runner) closeWatchers(jobID string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, ch := range r.watchers[jobID] {
		close(ch)
	}
	delete(r.watchers, jobID)
}

func (r *Runner) Cancel(jobID string) {
	r.mu.Lock()
	cancel, ok := r.cancels[jobID]
	r.mu.Unlock()
	if ok {
		cancel()
	}
}

// Enqueue adds a job. If concurrency allows, it starts immediately; otherwise it is queued.
func (r *Runner) Enqueue(job *models.Job) error {
	script, ok := r.store.GetScript(job.ScriptID)
	if !ok {
		return fmt.Errorf("script not found: %s", job.ScriptID)
	}
	source, ok := r.store.GetSource(job.SourceID)
	if !ok {
		return fmt.Errorf("source not found: %s", job.SourceID)
	}

	job.ScriptName       = script.Name
	job.ScriptConfidence = script.Confidence
	job.SourceName       = source.Name
	job.SourcePath       = source.Path
	job.SourceSizeBytes  = source.SizeBytes
	job.SourceLineCount  = source.LineCount
	job.CreatedAt        = time.Now()

	r.mu.Lock()
	canRun := r.store.CountRunningJobs() < r.maxConcurrent
	r.mu.Unlock()

	if canRun {
		job.Status = models.JobPending
		if err := r.store.CreateJob(job); err != nil {
			return err
		}
		r.mu.Lock()
		r.totalStarted++
		r.mu.Unlock()
		go r.run(job, script, source)
	} else {
		job.Status = models.JobQueued
		if err := r.store.CreateJob(job); err != nil {
			return err
		}
	}
	return nil
}

// Start is an alias for Enqueue.
func (r *Runner) Start(job *models.Job) error { return r.Enqueue(job) }

func (r *Runner) tryDispatch() {
	for {
		r.mu.Lock()
		canRun := r.store.CountRunningJobs() < r.maxConcurrent
		r.mu.Unlock()
		if !canRun {
			return
		}
		queued := r.store.PeekQueued()
		if queued == nil {
			return
		}
		script, ok := r.store.GetScript(queued.ScriptID)
		if !ok {
			queued.Status = models.JobFailed
			_ = r.store.UpdateJob(queued)
			continue
		}
		source, ok := r.store.GetSource(queued.SourceID)
		if !ok {
			queued.Status = models.JobFailed
			_ = r.store.UpdateJob(queued)
			continue
		}
		queued.Status = models.JobPending
		_ = r.store.UpdateJob(queued)
		r.mu.Lock()
		r.totalStarted++
		r.mu.Unlock()
		go r.run(queued, script, source)
	}
}

func buildArgs(script *models.Script, cppMode bool) []string {
	var args []string
	if cppMode {
		args = append(args, "--cpp")
	}
	// -C forces colour output even when stdout is a pipe (not a TTY).
	// weggli's --color flag alone only works on TTYs; -C overrides the isatty check.
	args = append(args, "-C")
	for _, f := range splitFlags(script.ExtraFlags) {
		if f != "-C" && f != "--color" { // deduplicate
			args = append(args, f)
		}
	}
	return args
}

func buildCmd(args []string) *exec.Cmd {
	cmd := exec.Command("weggli", args...)
	// Belt-and-suspenders: also set the conventional env vars that
	// termcolor/anstream check before weggli's own -C flag is parsed.
	cmd.Env = append(os.Environ(), "CLICOLOR_FORCE=1", "COLORTERM=truecolor")
	return cmd
}

func (r *Runner) run(job *models.Job, script *models.Script, source *models.Source) {
	job.Status = models.JobRunning
	job.StartedAt = time.Now()
	if err := r.store.UpdateJob(job); err != nil {
		log.Printf("update job: %v", err)
	}

	// Scan source stats if unknown
	if source.SizeBytes == 0 {
		go func() {
			b, l := scanSourceStats(source.Path)
			_ = r.store.UpdateSourceStats(source.ID, b, l)
			// Update the job's denormalized values too
			job.SourceSizeBytes = b
			job.SourceLineCount = l
		}()
	}

	r.mu.Lock()
	r.cancels[job.ID] = func() {}
	r.mu.Unlock()

	type mode struct {
		cpp   bool
		label string
	}
	var modes []mode
	switch script.Language {
	case models.LangCPP:
		modes = []mode{{true, "c++"}}
	case models.LangBoth:
		modes = []mode{{false, "c"}, {true, "c++"}}
	default:
		modes = []mode{{false, "c"}}
	}

	var allFindings []models.Finding
	killed := false

	for _, m := range modes {
		if killed {
			break
		}
		args := buildArgs(script, m.cpp)
		args = append(args, script.Pattern, source.Path)
		cmd := buildCmd(args)
		pr, pw := io.Pipe()
		cmd.Stdout = pw
		cmd.Stderr = pw

		r.mu.Lock()
		r.cancels[job.ID] = func() {
			killed = true
			if cmd.Process != nil {
				_ = cmd.Process.Kill()
			}
		}
		r.mu.Unlock()

		if len(modes) > 1 {
			header := fmt.Sprintf("\n[--- mode: %s ---]\n", m.label)
			_ = r.store.AppendJobOutput(job.ID, header)
			r.broadcast(job.ID, header)
		}

		if err := cmd.Start(); err != nil {
			msg := fmt.Sprintf("[error starting weggli (%s): %v]\n", m.label, err)
			_ = r.store.AppendJobOutput(job.ID, msg)
			r.broadcast(job.ID, msg)
			_ = pw.Close()
			continue
		}

		var modeOut strings.Builder // plain text (ANSI stripped) — for parseFindings
		var modeRaw strings.Builder // raw ANSI — for HTML snippet colouring
		done := make(chan struct{})
		go func() {
			defer close(done)
			sc := bufio.NewScanner(pr)
			sc.Buffer(make([]byte, 1<<20), 1<<20)
			for sc.Scan() {
				rawLine := sc.Text()
				// Plain text for finding parser (needs clean file:line headers)
				modeOut.WriteString(StripANSI(rawLine) + "\n")
				modeRaw.WriteString(rawLine + "\n")
				// HTML for storage and streaming — ANSI codes become <span> tags
				html := ansiToHTML(rawLine) + "\n"
				_ = r.store.AppendJobOutput(job.ID, html)
				r.broadcast(job.ID, html)
			}
		}()

		_ = cmd.Wait()
		_ = pw.Close()
		<-done
		allFindings = append(allFindings, parseFindings(modeOut.String(), modeRaw.String(), time.Now())...)
	}

	now := time.Now()
	dur := now.Sub(job.StartedAt).Seconds()
	job.FinishedAt = &now
	job.DurationSec = dur
	if killed {
		job.Status = models.JobCanceled
	} else {
		job.Status = models.JobDone
	}

	// Record timing data
	src, _ := r.store.GetSource(source.ID) // re-read to get latest stats
	if !killed && src != nil && src.SizeBytes > 0 && dur > 0 {
		_ = r.store.UpdateScriptTiming(script.ID, float64(src.SizeBytes)/dur, float64(src.LineCount)/dur)
	}

	_ = r.store.SaveFindings(job.ID, allFindings)
	job.Findings = allFindings
	if current, ok := r.store.GetJob(job.ID); ok {
		job.Output = current.Output
	}
	_ = r.store.UpdateJob(job)

	r.mu.Lock()
	r.totalCompleted++
	r.totalFindings += int64(len(allFindings))
	delete(r.cancels, job.ID)
	r.mu.Unlock()

	r.broadcast(job.ID, fmt.Sprintf("\n[job %s — %d findings — %.1fs]\n", job.Status, len(allFindings), dur))
	r.closeWatchers(job.ID)
	r.tryDispatch()
}

func splitFlags(s string) []string {
	var result []string
	cur := ""
	inQuote := false
	var quoteChar rune

	for _, c := range s {
		if inQuote {
			// If we are inside quotes, look for the closing quote
			if c == quoteChar {
				inQuote = false
			} else {
				cur += string(c)
			}
		} else {
			// If we are outside quotes, check for quote starts or spaces
			if c == '\'' || c == '"' {
				inQuote = true
				quoteChar = c
			} else if c == ' ' || c == '\t' {
				if cur != "" {
					result = append(result, cur)
					cur = ""
				}
			} else {
				cur += string(c)
			}
		}
	}
	// Append any leftover string
	if cur != "" {
		result = append(result, cur)
	}
	return result
}
