package handler

import (
	"Togo/usecases"
	"github.com/gin-gonic/gin"
)

var {{.Name}}UserCase = usecases.New{{.Name}}UC()

{{range .Func}}
func {{.Name}}(this *gin.Context) {
	this.JSON({{$.Name}}UserCase.Signup(this))
}
{{end}}
