
type {{.Name}}RepoInterface interface {
	Insert(*models.{{.Name}}) (int64, error)
	Delete(id int64) (int64, error)
	Update(*models.{{.Name}}) (int64, error)
	FindOne(id int64) *models.{{.Name}}
	FindMany() []*models.{{.Name}}
}