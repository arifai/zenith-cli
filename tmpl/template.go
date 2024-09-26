package tmpl

import "embed"

//go:embed *.tmpl
var TemplateFile embed.FS
