package gen

import (
	"os"
	"strings"
	"text/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var caser = cases.Title(language.English)

func TitleCase(s string) string {
	return strings.ReplaceAll(caser.String(s), "-", "")
}

func Die(err error) {
	if err != nil {
		panic(err)
	}
}
func Save(path string, t *template.Template, data any) {
	f, err := os.Create(path)
	Die(err)
	defer f.Close()
	t.Execute(f, data)
}
