package view

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var (
	layoutPath  = "views/layouts/"
	templateExt = ".gohtml"
)

//NewView provides a template with all the files included in the LAYOUTPATH
func NewView(layout string, files ...string) *View {
	files = append(files, listFilenames()...)

	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	v := View{
		Template: t,
		Layout:   layout}
	return &v
}

//View is a view for templates
type View struct {
	Template *template.Template
	Layout   string
}

//listFilenames returns the names of all files in the LAYOUTPATH with the TEMPLATEEXT.
func listFilenames() []string {
	files, err := filepath.Glob(layoutPath + "*" + templateExt)
	if err != nil {
		panic(err)
	}
	return files

}

//Render calls ExecuteTemplate on the View and sets the correct Content-Type
func (v View) Render(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html")
	handleErr(v.Template.ExecuteTemplate(w, v.Layout, data))
}

func handleErr(e error) {
	if e != nil {
		panic(e)
	}
}
