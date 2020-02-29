//go:generate inline -o templates.go -p cdr mixtape.tmpl templates/parts.tmpl

package cdr

import (
	"html/template"
	"io"
	"io/ioutil"
	"log"
)

func Render(m *Mixtape, w io.Writer) error {
	partsRawTmpl, err := ReadAsset("templates/parts.tmpl", false)
	if err != nil {
		return err
	}
	mixtapeRawTmpl, err := ioutil.ReadFile("mixtape.tmpl")
	if err != nil {
		log.Print("Unable to load local mixtape.tmpl; falling back to default")
		mixtapeRawTmpl, err = ReadAsset("mixtape.tmpl", false)
		if err != nil {
			return err
		}
	} else {
		log.Print("Generating from local mixtape.tmpl")
	}
	t, err := template.New("mixtape").Parse(string(mixtapeRawTmpl) + string(partsRawTmpl))
	if err != nil {
		log.Printf("[ERROR] Unable to parse: %v", err)
		return err
	}
	err = t.ExecuteTemplate(w, "mixtape", m)
	if err != nil {
		log.Printf("[ERROR] Unable to render: %v", err)
	}
	return nil
}
