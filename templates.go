package cdr

import (
	_ "embed"
)

//go:embed mixtape.tmpl
var defaultMixtapeTmpl []byte

//go:embed templates/parts.tmpl
var partsRawTmpl []byte
