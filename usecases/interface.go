package usecases

import "Togo/models"

//这一层该做这些事情
type UseCase interface {
	//查询用户信息
	FetchUser(int) *models.User
	//创建用户
	CreateUser(*models.User) bool
	//删除用户
	DeleteUser(int) error
	//修改用户
	UpdateUser(*models.User) error
}
