package cdr

import (
	"html/template"
	"io"
	"io/ioutil"
	"log"
)

func Render(m *Mixtape, w io.Writer) error {
	mixtapeRawTmpl, err := ioutil.ReadFile("mixtape.tmpl")
	if err != nil {
		log.Print("Unable to load custom mixtape.tmpl; falling back to default")
		mixtapeRawTmpl = defaultMixtapeTmpl
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
