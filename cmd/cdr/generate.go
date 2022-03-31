package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/thebaer/cdr"
	"github.com/urfave/cli/v2"
)

var (
	cmdServe = cli.Command{
		Name:   "preview",
		Usage:  "serve the mixtape site",
		Action: serveAction,
	}
	cmdGenerate = cli.Command{
		Name:   "burn",
		Usage:  "generate the static mixtape site",
		Action: generateAction,
	}
)

func newMixtape(wd string) (*cdr.Mixtape, error) {
	m := &cdr.Mixtape{Tracks: []cdr.Track{}}
	c := 1

	filepath.Walk(wd, func(path string, i os.FileInfo, err error) error {
		if !i.IsDir() && !strings.HasPrefix(i.Name(), ".") && i.Name() != "index.html" && !strings.HasSuffix(i.Name(), ".tmpl") {
			t, err := cdr.NewTrack(i.Name())
			if err != nil {
				log.Printf("Skipping track %s: %v", i.Name(), err)
				return nil
			}
			t.Num = c
			c++
			log.Println("Adding track", t.Title)
			m.Tracks = append(m.Tracks, *t)
		}

		return nil
	})
	return m, nil
}

func generateAction(c *cli.Context) error {
	f, err := os.Create("index.html")
	if err != nil {
		return err
	}
	defer f.Close()

	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	m, err := newMixtape(wd)
	if err != nil {
		return err
	}

	err = cdr.Render(m, f)
	if err != nil {
		return err
	}

	return nil
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

	port := os.Getenv("PORT")
	if port == "" {
		port = "9991"
	}
	log.Printf("Serving on http://localhost:%s", port)
	return http.ListenAndServe(":"+port, nil)
}
