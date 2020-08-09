package usecases

import (
	"Togo/models"
	"Togo/repos"
)

type {{.Name}}UC struct {
}

var {{.Name}}Repo = repos.New{{.Name}}Repo()

func New{{.Name}}UC() {{.Name}}Interface {
	return &{{.Name}}UC{}
}

{{range .Func}}
func (this *{{$.Name}}UC) {{.Name}}(ctx *gin.Context) (int, *Response) {}
{{end}}