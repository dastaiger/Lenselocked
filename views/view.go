package view

import "html/template"

func NewView(layout string, files ...string) *View {
	files = append(files,
		"views/layouts/footer.gohtml",
		"views/layouts/bootstrap.gohtml")

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
