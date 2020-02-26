//go:generate inline -o templates.go -p cdr mixtape.tmpl templates/parts.tmpl

package cdr

import (
	"html/template"
	"io"
	"io/ioutil"
)

func Render(m *Mixtape, w io.Writer) error {
	partsRawTmpl, err := ReadAsset("templates/parts.tmpl", false)
	if err != nil {
		return err
	}
	mixtapeRawTmpl, err := ioutil.ReadFile("mixtape.tmpl")
	if err != nil {
		mixtapeRawTmpl, err = ReadAsset("mixtape.tmpl", false)
		if err != nil {
			return err
		}
	}
	t, err := template.New("mixtape").Parse(string(mixtapeRawTmpl) + string(partsRawTmpl))
	if err != nil {
		return err
	}
	t.ExecuteTemplate(w, "mixtape", m)
	return nil
}
