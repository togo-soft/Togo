package repositories

import (
	"log"
	"Togo/models"
)

//操作用户模型 实现了UserRepoInterface 接口
type UserRepo struct {
}

func NewUserRepo() UserRepoInterface {
	return &UserRepo{}
}

//将user信息插入数据库
func (this *UserRepo) Insert(u *models.User) (int64, error) {
	return engine.Insert(u)
}

//根据ID删除user表信息
func (this *UserRepo) Delete(id int64) (int64, error) {
	return engine.Delete(&models.User{Id: id})
}

//更新user信息
func (this *UserRepo) Update(u *models.User) (int64, error) {
	return engine.Update(u)
}

//根据ID查询一个user信息
func (this *UserRepo) FindOneById(id int64) *models.User {
	var u = &models.User{}
	if _, err := engine.ID(id).Get(u); err != nil {
		log.Fatal("find user error:", err)
	}
	return u
}

//根据已有字段 查询用户信息
func (this *UserRepo) FindOneByField(u *models.User) *models.User {
	_, err := engine.Where("username = ?", u.Username).Or("email = ?", u.Email).Or("phone = ?", u.Phone).Get(u)
	if err != nil {
		log.Fatal("find user error:", err)
	}
	return u
}

//根据用户登陆时的名称和密码 检测user信息是否存在 并返回user信息和是否存在记录
func (this *UserRepo) RecordExist(name, password string) (bool, *models.User) {
	var u = new(models.User)
	log.Println("需要检测的信息", name, password)
	has, _ := engine.Where("username = ?", name).Or("email = ?", name).Or("phone = ?", name).And("password = ?", password).Get(u)
	return has, u
}

func (this *UserRepo) FindMany() []*models.User {
	all := make([]*models.User, 0)
	if err := engine.Find(&all); err != nil {
		log.Fatal("find all user error:", err)
	}
	return all
}
