package repositories

import "Togo/models"

//用户实体-仓库操作-接口
type UserRepoInterface interface {
	//添加用户接口
	Insert(user *models.User) (int64, error)
	//删除用户接口
	Delete(id int64) (int64, error)
	//修改用户接口
	Update(user *models.User) (int64, error)

	//根据id查询单个用户接口
	FindOneById(id int64) *models.User
	//根据字段查询用户信息接口
	FindOneByField(u *models.User) *models.User
	//记录是否存在
	RecordExist(name, password string) (bool, *models.User)
	//查询用户列表接口
	FindMany() []*models.User
}
