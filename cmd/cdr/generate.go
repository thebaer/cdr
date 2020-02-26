package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/thebaer/cdr"
	"github.com/urfave/cli"
)

var (
	cmdServe = cli.Command{
		Name:   "preview",
		Usage:  "serve the mixtape site",
		Action: serveAction,
	}
)

func newMixtape(wd string) (*cdr.Mixtape, error) {
	m := &cdr.Mixtape{Tracks: []cdr.Track{}}

	filepath.Walk(wd, func(path string, i os.FileInfo, err error) error {
		if !i.IsDir() && !strings.HasPrefix(i.Name(), ".") && i.Name() != "index.html" {
			t, err := cdr.NewTrack(i.Name())
			if err != nil {
				log.Printf("Skipping track %s: %v", i.Name(), err)
				return nil
			}
			log.Println("Adding track", t.Title)
			m.Tracks = append(m.Tracks, *t)
		}

		return nil
	})
	return m, nil
}

func serveAction(c *cli.Context) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	m, err := newMixtape(wd)
	if err != nil {
		return err
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI != "/" {
			log.Printf("GET %s", r.RequestURI)
			http.ServeFile(w, r, filepath.Join(wd, r.RequestURI))
			return
		}
		err := cdr.Render(m, w)
		if err != nil {
			log.Printf("[ERROR] Render failed! %s", err)
		}
		log.Printf("GET /")
	})

	return http.ListenAndServe(":9991", nil)
}
