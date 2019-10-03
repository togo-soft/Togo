package repositories

import (
	"Togo/models"
)

//User操作接口 定义了该层需要做的事情
type UserDataInterface interface {
	Insert(*models.User) bool
	Delete(id int) error
	Update(*models.User) error
	Fetch(id int) *models.User
}
