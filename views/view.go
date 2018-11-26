package view

import "html/template"

func NewView(files ...string) *View {
	files = append(files, "views/layouts/footer.gohtml")

	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	v := View{
		Template: t}
	return &v
}

//View is a view for templates
type View struct {
	Template *template.Template
}
