
type {{.Name}}Interface interface {
	{{range .Func}}
	{{.Name}}(ctx *gin.Context) (int, *Response)
    {{end}}
}