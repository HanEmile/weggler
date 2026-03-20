package models

import "time"

// Language controls which weggli invocation mode(s) are used.
type Language string

const (
	LangC    Language = "c"
	LangCPP  Language = "cpp"
	LangBoth Language = "both"
)

type Script struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Pattern     string    `json:"pattern"`
	Language    Language  `json:"language"`
	ExtraFlags  string    `json:"extra_flags"`
	// Confidence 0-100: 100 = certain hit, 0 = very noisy.
	Confidence int       `json:"confidence"`
	// Timing: bytes-per-second observed across completed jobs.
	AvgBytesPerSec  float64 `json:"avg_bytes_per_sec"`
	AvgLinesPerSec  float64 `json:"avg_lines_per_sec"`
	TimingSamples   int     `json:"timing_samples"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Source.Tags: "c", "cpp" — what languages the codebase contains.
type Source struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Path        string    `json:"path"`
	Tags        []string  `json:"tags"`
	// Measured on last scan
	SizeBytes int64 `json:"size_bytes"`
	LineCount  int64 `json:"line_count"`
	LastScanned *time.Time `json:"last_scanned,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ScriptGroup is a named collection of script IDs and/or child group IDs.
type ScriptGroup struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ScriptIDs   []string  `json:"script_ids"`
	ChildIDs    []string  `json:"child_ids"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type JobStatus string

const (
	JobQueued   JobStatus = "queued"
	JobPending  JobStatus = "pending"
	JobRunning  JobStatus = "running"
	JobDone     JobStatus = "done"
	JobFailed   JobStatus = "failed"
	JobCanceled JobStatus = "canceled"
)

// Finding represents one weggli match.
type Finding struct {
	File        string    `json:"file"`
	Line        int       `json:"line"`
	Snippet     string    `json:"snippet"`
	SnippetHTML string    `json:"snippet_html,omitempty"`
	FoundAt     time.Time `json:"found_at"`
}

type Job struct {
	ID         string     `json:"id"`
	ScriptID   string     `json:"script_id"`
	SourceID   string     `json:"source_id"`
	GroupID    string     `json:"group_id,omitempty"`
	// QueueID groups jobs queued together (e.g. from a group-run or run-all).
	QueueID    string     `json:"queue_id,omitempty"`
	Description string    `json:"description,omitempty"`
	Status     JobStatus  `json:"status"`
	Output     string     `json:"output"`
	ExitCode   int        `json:"exit_code"`
	StartedAt  time.Time  `json:"started_at"`
	FinishedAt *time.Time `json:"finished_at,omitempty"`
	CreatedAt  time.Time  `json:"created_at"`
	Findings   []Finding  `json:"findings"`
	// Runtime metrics
	DurationSec float64 `json:"duration_sec,omitempty"`
	// Denormalized for display
	ScriptName       string `json:"script_name"`
	ScriptConfidence int    `json:"script_confidence"`
	SourceName       string `json:"source_name"`
	SourcePath       string `json:"source_path"`
	SourceSizeBytes  int64  `json:"source_size_bytes"`
	SourceLineCount  int64  `json:"source_line_count"`
}
