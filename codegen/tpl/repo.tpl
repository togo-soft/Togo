package repos

import (
    "Togo/models"
)

type {{.Name}}Repo struct {
}

func New{{.Name}}Repo() {{.Name}}RepoInterface {
	return &{{.Name}}Repo{}
}

func (this *{{.Name}}Repo) Insert(u *models.User) (int64, error) {
	return engine.Insert(u)
}

func (this *{{.Name}}Repo) Delete(id int64) (int64, error) {
	return engine.Delete(&models.{{.Name}}{Id: id})
}

func (this *{{.Name}}Repo) Update(u *models.{{.Name}}) (int64, error) {
	return engine.Update(u)
}

func (this *{{.Name}}Repo) FindOne(id int64) *models.{{.Name}} {
    var u *models.{{.Name}}
	_, err := engine.Id(id).Get(u)
	if err != nil {
		log.Fatal("find user error:", err)
	}
	return u
}

func (this *{{.Name}}Repo) FindMany() []*models.{{.Name}} {
	all := make([]*models.{{.Name}}, 0)
	if err := engine.Find(&all); err != nil {
		log.Fatal("find all user error:", err)
	}
	return all
}