package web

import (
	"net/http"

	"weggler/internal/runner"
	"weggler/internal/store"
)

type Server struct {
	store  *store.Store
	runner *runner.Runner
}

func NewServer(st *store.Store, r *runner.Runner) *Server {
	return &Server{store: st, runner: r}
}

func (s *Server) Routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", s.handleIndex)
	mux.HandleFunc("/static/", s.handleStatic)

	mux.HandleFunc("/api/reload", s.handleReload)
	mux.HandleFunc("/api/metrics", s.handleMetrics)
	mux.HandleFunc("/api/config", s.handleConfig)

	mux.HandleFunc("/api/scripts", s.handleScripts)
	mux.HandleFunc("/api/scripts/", s.handleScript)

	mux.HandleFunc("/api/sources", s.handleSources)
	mux.HandleFunc("/api/sources/", s.handleSource)

	mux.HandleFunc("/api/groups", s.handleGroups)
	mux.HandleFunc("/api/groups/", s.handleGroup)

	mux.HandleFunc("/api/jobs", s.handleJobs)
	mux.HandleFunc("/api/jobs/", s.handleJob)
	mux.HandleFunc("/api/jobs/stream/", s.handleJobStream)

	return mux
}
