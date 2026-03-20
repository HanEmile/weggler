package store

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"weggler/internal/models"
)

// Store uses a directory-based layout:
//   <dir>/scripts.json
//   <dir>/sources.json
//   <dir>/groups.json
//   <dir>/jobs/<id>.json
//
// For backward compat it also reads a legacy single-file weggler.json.

type Store struct {
	mu  sync.RWMutex
	dir string

	scripts []*models.Script
	sources []*models.Source
	groups  []*models.ScriptGroup
	jobs    []*models.Job
}

func New(path string) (*Store, error) {
	// Determine directory and optional legacy file path.
	var dir string
	var legacyFile string // non-empty only when path is an existing .json file

	fi, err := os.Stat(path)
	if err == nil && fi.IsDir() {
		// path is already a directory (e.g. --data data)
		dir = path
	} else if filepath.Ext(path) == ".json" {
		// path looks like a json file (legacy --db weggler.json)
		dir = filepath.Dir(path)
		legacyFile = path
	} else {
		// treat as directory name that may not exist yet
		dir = path
	}

	if err := os.MkdirAll(filepath.Join(dir, "jobs"), 0755); err != nil {
		return nil, err
	}

	s := &Store{dir: dir}
	if err := s.load(legacyFile); err != nil {
		return nil, err
	}
	return s, nil
}

// load tries the split-file format first, then falls back to a legacy single file.
func (s *Store) load(legacyFile string) error {
	if err := s.loadSplit(); err != nil {
		return err
	}
	// If split files existed (even empty), we're done.
	// Check whether we actually got anything from split files by seeing if
	// scripts.json exists.
	if _, err := os.Stat(filepath.Join(s.dir, "scripts.json")); err == nil {
		return nil
	}
	// No split files yet. Try legacy single-file migration.
	if legacyFile == "" {
		return nil // fresh install, nothing to migrate
	}
	data, err := os.ReadFile(legacyFile)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return fmt.Errorf("read %s: %w", legacyFile, err)
	}
	var db struct {
		Scripts      []*models.Script      `json:"scripts"`
		Sources      []*models.Source      `json:"sources"`
		ScriptGroups []*models.ScriptGroup `json:"script_groups"`
		Jobs         []*models.Job         `json:"jobs"`
	}
	if err := json.Unmarshal(data, &db); err != nil {
		return fmt.Errorf("parse %s: %w", legacyFile, err)
	}
	s.scripts = db.Scripts
	s.sources = db.Sources
	s.groups = db.ScriptGroups
	s.jobs = db.Jobs
	// Migrate to split format.
	return s.saveAll()
}

// loadSplit reads all split files; missing files are treated as empty (not errors).
func (s *Store) loadSplit() error {
	var err error
	s.scripts, err = loadJSONOrEmpty[[]*models.Script](filepath.Join(s.dir, "scripts.json"))
	if err != nil {
		return err
	}
	s.sources, err = loadJSONOrEmpty[[]*models.Source](filepath.Join(s.dir, "sources.json"))
	if err != nil {
		return err
	}
	s.groups, err = loadJSONOrEmpty[[]*models.ScriptGroup](filepath.Join(s.dir, "groups.json"))
	if err != nil {
		return err
	}
	entries, err := os.ReadDir(filepath.Join(s.dir, "jobs"))
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	for _, e := range entries {
		if e.IsDir() || filepath.Ext(e.Name()) != ".json" {
			continue
		}
		j, err := loadJSONFile[*models.Job](filepath.Join(s.dir, "jobs", e.Name()))
		if err != nil {
			continue
		}
		s.jobs = append(s.jobs, j)
	}
	return nil
}

// loadJSONOrEmpty returns zero value (nil slice) when file does not exist.
func loadJSONOrEmpty[T any](path string) (T, error) {
	var zero T
	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return zero, nil
	}
	if err != nil {
		return zero, err
	}
	var v T
	if err := json.Unmarshal(data, &v); err != nil {
		return zero, fmt.Errorf("parse %s: %w", path, err)
	}
	return v, nil
}

func loadJSONFile[T any](path string) (T, error) {
	var zero T
	data, err := os.ReadFile(path)
	if err != nil {
		return zero, err
	}
	var v T
	if err := json.Unmarshal(data, &v); err != nil {
		return zero, err
	}
	return v, nil
}

func writeJSONFile(path string, v any) error {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	tmp := path + ".tmp"
	if err := os.WriteFile(tmp, data, 0644); err != nil {
		return err
	}
	return os.Rename(tmp, path)
}

func (s *Store) saveScripts() error {
	return writeJSONFile(filepath.Join(s.dir, "scripts.json"), s.scripts)
}
func (s *Store) saveSources() error {
	return writeJSONFile(filepath.Join(s.dir, "sources.json"), s.sources)
}
func (s *Store) saveGroups() error {
	return writeJSONFile(filepath.Join(s.dir, "groups.json"), s.groups)
}
func (s *Store) saveJob(j *models.Job) error {
	return writeJSONFile(filepath.Join(s.dir, "jobs", j.ID+".json"), j)
}
func (s *Store) deleteJobFile(id string) error {
	return os.Remove(filepath.Join(s.dir, "jobs", id+".json"))
}

func (s *Store) saveAll() error {
	if err := s.saveScripts(); err != nil {
		return err
	}
	if err := s.saveSources(); err != nil {
		return err
	}
	if err := s.saveGroups(); err != nil {
		return err
	}
	for _, j := range s.jobs {
		if err := s.saveJob(j); err != nil {
			return err
		}
	}
	return nil
}

func newID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

// ---- Scripts ----

func (s *Store) ListScripts() []*models.Script {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make([]*models.Script, len(s.scripts))
	copy(out, s.scripts)
	return out
}

func (s *Store) GetScript(id string) (*models.Script, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, sc := range s.scripts {
		if sc.ID == id {
			cp := *sc
			return &cp, true
		}
	}
	return nil, false
}

func (s *Store) CreateScript(sc *models.Script) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	sc.ID = newID()
	sc.CreatedAt = time.Now()
	sc.UpdatedAt = time.Now()
	s.scripts = append(s.scripts, sc)
	return s.saveScripts()
}

func (s *Store) UpdateScript(sc *models.Script) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, ex := range s.scripts {
		if ex.ID == sc.ID {
			sc.CreatedAt = ex.CreatedAt
			sc.UpdatedAt = time.Now()
			s.scripts[i] = sc
			return s.saveScripts()
		}
	}
	return fmt.Errorf("script %s not found", sc.ID)
}

func (s *Store) DeleteScript(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, sc := range s.scripts {
		if sc.ID == id {
			s.scripts = append(s.scripts[:i], s.scripts[i+1:]...)
			return s.saveScripts()
		}
	}
	return fmt.Errorf("script %s not found", id)
}

func (s *Store) UpdateScriptTiming(id string, bytesPerSec, linesPerSec float64) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, sc := range s.scripts {
		if sc.ID == id {
			n := float64(sc.TimingSamples)
			sc.AvgBytesPerSec = (sc.AvgBytesPerSec*n + bytesPerSec) / (n + 1)
			sc.AvgLinesPerSec = (sc.AvgLinesPerSec*n + linesPerSec) / (n + 1)
			sc.TimingSamples++
			return s.saveScripts()
		}
	}
	return fmt.Errorf("script %s not found", id)
}

// ---- Sources ----

func (s *Store) ListSources() []*models.Source {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make([]*models.Source, len(s.sources))
	copy(out, s.sources)
	return out
}

func (s *Store) GetSource(id string) (*models.Source, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, src := range s.sources {
		if src.ID == id {
			cp := *src
			return &cp, true
		}
	}
	return nil, false
}

func (s *Store) CreateSource(src *models.Source) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	src.ID = newID()
	src.CreatedAt = time.Now()
	src.UpdatedAt = time.Now()
	s.sources = append(s.sources, src)
	return s.saveSources()
}

func (s *Store) UpdateSource(src *models.Source) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, ex := range s.sources {
		if ex.ID == src.ID {
			src.CreatedAt = ex.CreatedAt
			src.UpdatedAt = time.Now()
			s.sources[i] = src
			return s.saveSources()
		}
	}
	return fmt.Errorf("source %s not found", src.ID)
}

func (s *Store) DeleteSource(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, src := range s.sources {
		if src.ID == id {
			s.sources = append(s.sources[:i], s.sources[i+1:]...)
			return s.saveSources()
		}
	}
	return fmt.Errorf("source %s not found", id)
}

func (s *Store) UpdateSourceStats(id string, sizeBytes, lineCount int64) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, src := range s.sources {
		if src.ID == id {
			src.SizeBytes = sizeBytes
			src.LineCount = lineCount
			now := time.Now()
			src.LastScanned = &now
			return s.saveSources()
		}
	}
	return fmt.Errorf("source %s not found", id)
}

// ---- ScriptGroups ----

func (s *Store) ListScriptGroups() []*models.ScriptGroup {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make([]*models.ScriptGroup, len(s.groups))
	copy(out, s.groups)
	return out
}

func (s *Store) GetScriptGroup(id string) (*models.ScriptGroup, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, g := range s.groups {
		if g.ID == id {
			cp := *g
			return &cp, true
		}
	}
	return nil, false
}

func (s *Store) CreateScriptGroup(g *models.ScriptGroup) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	g.ID = newID()
	if g.ScriptIDs == nil {
		g.ScriptIDs = []string{}
	}
	if g.ChildIDs == nil {
		g.ChildIDs = []string{}
	}
	g.CreatedAt = time.Now()
	g.UpdatedAt = time.Now()
	s.groups = append(s.groups, g)
	return s.saveGroups()
}

func (s *Store) UpdateScriptGroup(g *models.ScriptGroup) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, ex := range s.groups {
		if ex.ID == g.ID {
			g.CreatedAt = ex.CreatedAt
			g.UpdatedAt = time.Now()
			s.groups[i] = g
			return s.saveGroups()
		}
	}
	return fmt.Errorf("group %s not found", g.ID)
}

func (s *Store) DeleteScriptGroup(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, g := range s.groups {
		if g.ID == id {
			s.groups = append(s.groups[:i], s.groups[i+1:]...)
			return s.saveGroups()
		}
	}
	return fmt.Errorf("group %s not found", id)
}

func (s *Store) ResolveGroupScripts(groupID string) []string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	seen := map[string]bool{}
	var result []string
	var walk func(gid string)
	walk = func(gid string) {
		if seen["g:"+gid] {
			return
		}
		seen["g:"+gid] = true
		for _, g := range s.groups {
			if g.ID == gid {
				for _, sid := range g.ScriptIDs {
					if !seen["s:"+sid] {
						seen["s:"+sid] = true
						result = append(result, sid)
					}
				}
				for _, cid := range g.ChildIDs {
					walk(cid)
				}
			}
		}
	}
	walk(groupID)
	return result
}

// ---- Jobs ----

func (s *Store) ListJobs() []*models.Job {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make([]*models.Job, len(s.jobs))
	copy(out, s.jobs)
	return out
}

func (s *Store) GetJob(id string) (*models.Job, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, j := range s.jobs {
		if j.ID == id {
			cp := *j
			return &cp, true
		}
	}
	return nil, false
}

func (s *Store) CreateJob(j *models.Job) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	j.ID = newID()
	j.CreatedAt = time.Now()
	s.jobs = append(s.jobs, j)
	return s.saveJob(j)
}

func (s *Store) UpdateJob(j *models.Job) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, ex := range s.jobs {
		if ex.ID == j.ID {
			j.CreatedAt = ex.CreatedAt
			s.jobs[i] = j
			return s.saveJob(j)
		}
	}
	return fmt.Errorf("job %s not found", j.ID)
}

func (s *Store) DeleteJob(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, j := range s.jobs {
		if j.ID == id {
			s.jobs = append(s.jobs[:i], s.jobs[i+1:]...)
			return s.deleteJobFile(id)
		}
	}
	return fmt.Errorf("job %s not found", id)
}

func (s *Store) AppendJobOutput(id, chunk string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, j := range s.jobs {
		if j.ID == id {
			j.Output += chunk
			return s.saveJob(j)
		}
	}
	return fmt.Errorf("job %s not found", id)
}

func (s *Store) SaveFindings(id string, findings []models.Finding) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, j := range s.jobs {
		if j.ID == id {
			j.Findings = findings
			return s.saveJob(j)
		}
	}
	return fmt.Errorf("job %s not found", id)
}

func (s *Store) CountRunningJobs() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	n := 0
	for _, j := range s.jobs {
		if j.Status == models.JobRunning || j.Status == models.JobPending {
			n++
		}
	}
	return n
}

func (s *Store) CountQueuedJobs() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	n := 0
	for _, j := range s.jobs {
		if j.Status == models.JobQueued {
			n++
		}
	}
	return n
}

func (s *Store) PeekQueued() *models.Job {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, j := range s.jobs {
		if j.Status == models.JobQueued {
			cp := *j
			return &cp
		}
	}
	return nil
}

// Reload re-reads all data files from disk, replacing in-memory state.
// Running jobs are not affected — the runner holds its own references.
func (s *Store) Reload() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	// Clear everything first.
	s.scripts = nil
	s.sources = nil
	s.groups  = nil
	s.jobs    = nil
	return s.loadSplit()
}
