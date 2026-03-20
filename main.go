package main

import (
	"flag"
	"log"
	"net/http"

	"weggler/internal/runner"
	"weggler/internal/store"
	"weggler/web"
)

func main() {
	addr   := flag.String("addr", ":8080", "Listen address")
	data   := flag.String("data", "data", "Data directory (stores scripts.json, sources.json, groups.json, jobs/)")
	maxJobs := flag.Int("max-jobs", 4, "Maximum concurrent weggli jobs")
	flag.Parse()

	st, err := store.New(*data)
	if err != nil {
		log.Fatalf("store: %v", err)
	}

	r := runner.New(st, *maxJobs)
	r.ResumeRunning()

	srv := web.NewServer(st, r)
	mux := srv.Routes()

	log.Printf("weggler listening on %s  (max-jobs=%d  data=%s)", *addr, *maxJobs, *data)
	if err := http.ListenAndServe(*addr, mux); err != nil {
		log.Fatalf("server: %v", err)
	}
}
