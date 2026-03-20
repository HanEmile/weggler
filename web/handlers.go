package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"weggler/internal/models"
)

func writeJSON(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}
func readJSON(r *http.Request, v any) error { return json.NewDecoder(r.Body).Decode(v) }
func idFromPath(path, prefix string) string  { return strings.TrimPrefix(path, prefix) }

// ---- Reload ----

func (s *Server) handleReload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", 405)
		return
	}
	if err := s.store.Reload(); err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 200, map[string]string{"status": "reloaded"})
}

// ---- Metrics ----

func (s *Server) handleMetrics(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, 200, s.runner.Metrics())
}

func (s *Server) handleConfig(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		writeJSON(w, 200, map[string]any{"max_concurrent": s.runner.Metrics()["max_concurrent"]})
		return
	}
	if r.Method == http.MethodPut {
		var req struct {
			MaxConcurrent int `json:"max_concurrent"`
		}
		if err := readJSON(r, &req); err != nil {
			writeJSON(w, 400, map[string]string{"error": err.Error()})
			return
		}
		if req.MaxConcurrent < 1 {
			writeJSON(w, 400, map[string]string{"error": "max_concurrent must be >= 1"})
			return
		}
		s.runner.SetMaxConcurrent(req.MaxConcurrent)
		writeJSON(w, 200, map[string]any{"max_concurrent": req.MaxConcurrent})
		return
	}
	http.Error(w, "method not allowed", 405)
}

// ---- Scripts ----

func (s *Server) handleScripts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		writeJSON(w, 200, s.store.ListScripts())
	case http.MethodPost:
		var sc models.Script
		if err := readJSON(r, &sc); err != nil {
			writeJSON(w, 400, map[string]string{"error": err.Error()})
			return
		}
		if err := s.store.CreateScript(&sc); err != nil {
			writeJSON(w, 500, map[string]string{"error": err.Error()})
			return
		}
		writeJSON(w, 201, sc)
	default:
		http.Error(w, "method not allowed", 405)
	}
}

func (s *Server) handleScript(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	// POST /api/scripts/:id/duplicate
	if strings.HasSuffix(path, "/duplicate") && r.Method == http.MethodPost {
		id := strings.TrimSuffix(strings.TrimPrefix(path, "/api/scripts/"), "/duplicate")
		sc, ok := s.store.GetScript(id)
		if !ok {
			writeJSON(w, 404, map[string]string{"error": "not found"})
			return
		}
		sc.ID = ""
		sc.Name = "Copy of " + sc.Name
		sc.CreatedAt = time.Time{}
		sc.UpdatedAt = time.Time{}
		if err := s.store.CreateScript(sc); err != nil {
			writeJSON(w, 500, map[string]string{"error": err.Error()})
			return
		}
		writeJSON(w, 201, sc)
		return
	}

	id := idFromPath(path, "/api/scripts/")
	if id == "" {
		http.NotFound(w, r)
		return
	}
	switch r.Method {
	case http.MethodGet:
		sc, ok := s.store.GetScript(id)
		if !ok {
			writeJSON(w, 404, map[string]string{"error": "not found"})
			return
		}
		writeJSON(w, 200, sc)
	case http.MethodPut:
		var sc models.Script
		if err := readJSON(r, &sc); err != nil {
			writeJSON(w, 400, map[string]string{"error": err.Error()})
			return
		}
		sc.ID = id
		if err := s.store.UpdateScript(&sc); err != nil {
			writeJSON(w, 500, map[string]string{"error": err.Error()})
			return
		}
		writeJSON(w, 200, sc)
	case http.MethodDelete:
		if err := s.store.DeleteScript(id); err != nil {
			writeJSON(w, 404, map[string]string{"error": err.Error()})
			return
		}
		w.WriteHeader(204)
	default:
		http.Error(w, "method not allowed", 405)
	}
}

// ---- Sources ----

func (s *Server) handleSources(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		writeJSON(w, 200, s.store.ListSources())
	case http.MethodPost:
		var src models.Source
		if err := readJSON(r, &src); err != nil {
			writeJSON(w, 400, map[string]string{"error": err.Error()})
			return
		}
		if err := s.store.CreateSource(&src); err != nil {
			writeJSON(w, 500, map[string]string{"error": err.Error()})
			return
		}
		writeJSON(w, 201, src)
	default:
		http.Error(w, "method not allowed", 405)
	}
}

func (s *Server) handleSource(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	if strings.HasSuffix(path, "/run-all") && r.Method == http.MethodPost {
		id := strings.TrimSuffix(strings.TrimPrefix(path, "/api/sources/"), "/run-all")
		s.handleSourceRunAll(w, r, id)
		return
	}
	if strings.HasSuffix(path, "/duplicate") && r.Method == http.MethodPost {
		id := strings.TrimSuffix(strings.TrimPrefix(path, "/api/sources/"), "/duplicate")
		src, ok := s.store.GetSource(id)
		if !ok {
			writeJSON(w, 404, map[string]string{"error": "not found"})
			return
		}
		src.ID = ""
		src.Name = "Copy of " + src.Name
		src.CreatedAt = time.Time{}
		src.UpdatedAt = time.Time{}
		if err := s.store.CreateSource(src); err != nil {
			writeJSON(w, 500, map[string]string{"error": err.Error()})
			return
		}
		writeJSON(w, 201, src)
		return
	}
	if strings.HasSuffix(path, "/scan") && r.Method == http.MethodPost {
		id := strings.TrimSuffix(strings.TrimPrefix(path, "/api/sources/"), "/scan")
		src, ok := s.store.GetSource(id)
		if !ok {
			writeJSON(w, 404, map[string]string{"error": "not found"})
			return
		}
		go func() {
			// import via runner package not needed — store method handles it
			_ = src
		}()
		writeJSON(w, 202, map[string]string{"status": "scanning"})
		return
	}

	id := idFromPath(path, "/api/sources/")
	if id == "" {
		http.NotFound(w, r)
		return
	}
	switch r.Method {
	case http.MethodGet:
		src, ok := s.store.GetSource(id)
		if !ok {
			writeJSON(w, 404, map[string]string{"error": "not found"})
			return
		}
		writeJSON(w, 200, src)
	case http.MethodPut:
		var src models.Source
		if err := readJSON(r, &src); err != nil {
			writeJSON(w, 400, map[string]string{"error": err.Error()})
			return
		}
		src.ID = id
		if err := s.store.UpdateSource(&src); err != nil {
			writeJSON(w, 500, map[string]string{"error": err.Error()})
			return
		}
		writeJSON(w, 200, src)
	case http.MethodDelete:
		if err := s.store.DeleteSource(id); err != nil {
			writeJSON(w, 404, map[string]string{"error": err.Error()})
			return
		}
		w.WriteHeader(204)
	default:
		http.Error(w, "method not allowed", 405)
	}
}

func (s *Server) handleSourceRunAll(w http.ResponseWriter, r *http.Request, sourceID string) {
	src, ok := s.store.GetSource(sourceID)
	if !ok {
		writeJSON(w, 404, map[string]string{"error": "source not found"})
		return
	}
	allScripts := s.store.ListScripts()
	tagSet := map[string]bool{}
	for _, t := range src.Tags {
		tagSet[t] = true
	}
	var launched []models.Job
	for _, sc := range allScripts {
		compat := len(src.Tags) == 0
		if !compat {
			switch sc.Language {
			case models.LangC:
				compat = tagSet["c"]
			case models.LangCPP:
				compat = tagSet["cpp"]
			case models.LangBoth:
				compat = true
			}
		}
		if !compat {
			continue
		}
		job := &models.Job{ScriptID: sc.ID, SourceID: sourceID}
		if err := s.runner.Enqueue(job); err != nil {
			continue
		}
		launched = append(launched, *job)
	}
	writeJSON(w, 201, map[string]any{"launched": len(launched), "jobs": launched})
}

// ---- Script Groups ----

func (s *Server) handleGroups(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		writeJSON(w, 200, s.store.ListScriptGroups())
	case http.MethodPost:
		var g models.ScriptGroup
		if err := readJSON(r, &g); err != nil {
			writeJSON(w, 400, map[string]string{"error": err.Error()})
			return
		}
		if err := s.store.CreateScriptGroup(&g); err != nil {
			writeJSON(w, 500, map[string]string{"error": err.Error()})
			return
		}
		writeJSON(w, 201, g)
	default:
		http.Error(w, "method not allowed", 405)
	}
}

func (s *Server) handleGroup(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if strings.HasSuffix(path, "/run") && r.Method == http.MethodPost {
		gid := strings.TrimSuffix(strings.TrimPrefix(path, "/api/groups/"), "/run")
		s.handleGroupRun(w, r, gid)
		return
	}
	id := idFromPath(path, "/api/groups/")
	if id == "" {
		http.NotFound(w, r)
		return
	}
	switch r.Method {
	case http.MethodGet:
		g, ok := s.store.GetScriptGroup(id)
		if !ok {
			writeJSON(w, 404, map[string]string{"error": "not found"})
			return
		}
		writeJSON(w, 200, g)
	case http.MethodPut:
		var g models.ScriptGroup
		if err := readJSON(r, &g); err != nil {
			writeJSON(w, 400, map[string]string{"error": err.Error()})
			return
		}
		g.ID = id
		if err := s.store.UpdateScriptGroup(&g); err != nil {
			writeJSON(w, 500, map[string]string{"error": err.Error()})
			return
		}
		writeJSON(w, 200, g)
	case http.MethodDelete:
		if err := s.store.DeleteScriptGroup(id); err != nil {
			writeJSON(w, 404, map[string]string{"error": err.Error()})
			return
		}
		w.WriteHeader(204)
	default:
		http.Error(w, "method not allowed", 405)
	}
}

func (s *Server) handleGroupRun(w http.ResponseWriter, r *http.Request, gid string) {
	var req struct {
		SourceIDs []string `json:"source_ids"` // one or more sources
		SourceID  string   `json:"source_id"`  // legacy single
	}
	if err := readJSON(r, &req); err != nil {
		writeJSON(w, 400, map[string]string{"error": err.Error()})
		return
	}
	// Merge single + multi
	srcIDs := req.SourceIDs
	if req.SourceID != "" && !contains(srcIDs, req.SourceID) {
		srcIDs = append(srcIDs, req.SourceID)
	}
	if len(srcIDs) == 0 {
		writeJSON(w, 400, map[string]string{"error": "at least one source_id required"})
		return
	}
	scriptIDs := s.store.ResolveGroupScripts(gid)
	if len(scriptIDs) == 0 {
		writeJSON(w, 400, map[string]string{"error": "group has no scripts"})
		return
	}

	queueID := fmt.Sprintf("q%d", time.Now().UnixNano())
	var launched []models.Job
	for _, srcID := range srcIDs {
		for _, sid := range scriptIDs {
			job := &models.Job{ScriptID: sid, SourceID: srcID, GroupID: gid, QueueID: queueID}
			if err := s.runner.Enqueue(job); err != nil {
				continue
			}
			launched = append(launched, *job)
		}
	}
	writeJSON(w, 201, map[string]any{"launched": len(launched), "jobs": launched, "queue_id": queueID})
}

func contains(ss []string, s string) bool {
	for _, x := range ss {
		if x == s {
			return true
		}
	}
	return false
}

// ---- Jobs ----

func (s *Server) handleJobs(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		jobs := s.store.ListJobs()
		// newest first
		for i, j := 0, len(jobs)-1; i < j; i, j = i+1, j-1 {
			jobs[i], jobs[j] = jobs[j], jobs[i]
		}
		writeJSON(w, 200, jobs)
	case http.MethodPost:
		var req struct {
			ScriptID    string `json:"script_id"`
			SourceID    string `json:"source_id"`
			Description string `json:"description"`
		}
		if err := readJSON(r, &req); err != nil {
			writeJSON(w, 400, map[string]string{"error": err.Error()})
			return
		}
		job := &models.Job{
			ScriptID:    req.ScriptID,
			SourceID:    req.SourceID,
			Description: req.Description,
		}
		if err := s.runner.Start(job); err != nil {
			writeJSON(w, 500, map[string]string{"error": err.Error()})
			return
		}
		writeJSON(w, 201, job)
	default:
		http.Error(w, "method not allowed", 405)
	}
}

func (s *Server) handleJob(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if strings.HasSuffix(path, "/cancel") {
		id := strings.TrimSuffix(strings.TrimPrefix(path, "/api/jobs/"), "/cancel")
		s.runner.Cancel(id)
		writeJSON(w, 200, map[string]string{"status": "cancel requested"})
		return
	}
	id := strings.TrimSuffix(idFromPath(path, "/api/jobs/"), "/stream")
	if id == "" {
		http.NotFound(w, r)
		return
	}
	switch r.Method {
	case http.MethodGet:
		job, ok := s.store.GetJob(id)
		if !ok {
			writeJSON(w, 404, map[string]string{"error": "not found"})
			return
		}
		writeJSON(w, 200, job)
	case http.MethodDelete:
		if err := s.store.DeleteJob(id); err != nil {
			writeJSON(w, 404, map[string]string{"error": err.Error()})
			return
		}
		w.WriteHeader(204)
	default:
		http.Error(w, "method not allowed", 405)
	}
}

func (s *Server) handleJobStream(w http.ResponseWriter, r *http.Request) {
	id := idFromPath(r.URL.Path, "/api/jobs/stream/")
	if id == "" {
		http.NotFound(w, r)
		return
	}
	job, ok := s.store.GetJob(id)
	if !ok {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("X-Accel-Buffering", "no")
	fl, canFlush := w.(http.Flusher)

	if job.Output != "" {
		fmt.Fprintf(w, "data: %s\n\n", encodeSSE(job.Output))
		if canFlush {
			fl.Flush()
		}
	}
	if job.Status != models.JobRunning && job.Status != models.JobPending {
		fmt.Fprintf(w, "event: done\ndata: %s\n\n", job.Status)
		if canFlush {
			fl.Flush()
		}
		return
	}

	ch := s.runner.Subscribe(id)
	ctx := r.Context()
	for {
		select {
		case <-ctx.Done():
			return
		case line, open := <-ch:
			if !open {
				if j, ok := s.store.GetJob(id); ok {
					fmt.Fprintf(w, "event: done\ndata: %s\n\n", j.Status)
				} else {
					fmt.Fprintf(w, "event: done\ndata: done\n\n")
				}
				if canFlush {
					fl.Flush()
				}
				return
			}
			fmt.Fprintf(w, "data: %s\n\n", encodeSSE(line))
			if canFlush {
				fl.Flush()
			}
		}
	}
}

func encodeSSE(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, "\r", ""), "\n", "\\n")
}
