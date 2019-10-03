package repositories

//这是repo层 通过gorm的Mysql存储层实现

import (
	"Togo/models"
	"Togo/utils"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DBC struct {
	Handle *gorm.DB
}

//生成DBC对象
func NewDBC() (UDI UserDataInterface) {
	//实例化
	var dbc = &DBC{Handle: new(gorm.DB)}
	dbc.Connect()
	UDI = dbc
	//连接
	return
}

//连接
func (this *DBC) Connect() {
	var err error
	var conf = utils.GetConfig()
	this.Handle, err = gorm.Open("mysql", conf.MysqlDSN)
	if err != nil {
		panic("failed to connect database")
	}
	//数据库Ping不通
	if err := this.Handle.DB().Ping(); err != nil {
		fmt.Println("数据库连接状况:", err)
	}
	//根据开启Debug
	if conf.Debug {
		this.Handle.LogMode(true)
	}
	//自动化
	this.Handle.AutoMigrate(&models.User{})
}

//添加数据
func (this *DBC) Insert(user *models.User) bool {
	fmt.Println("人工数据完整性检测:", user)
	if err := this.Handle.Create(user).Error; err != nil {
		fmt.Println("添加数据错误:", err)
	}
	//NewRecord用来检测是否是一个新的记录 是新记录返回true 否则false
	return !this.Handle.NewRecord(user)
}

//删除数据
func (this *DBC) Delete(id int) error {
	user := models.User{ID: id}
	return this.Handle.Delete(&user).Error

}

//修改数据
func (this *DBC) Update(user *models.User) error {
	if err := this.Handle.Save(user).Error; err != nil {
		fmt.Println("修改数据失败:", err)
		return err
	}
	return nil
}

//查询数据
func (this *DBC) Fetch(id int) *models.User {
	var user = &models.User{}
	if err := this.Handle.First(user, id).Error; err != nil {
		fmt.Println("查询数据错误:", err, "查询主键:", id)
	}
	return user
}
