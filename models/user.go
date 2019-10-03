package models

// User Model
type User struct {
	ID       int    `form:"id" json:"id"`
	UserName string `form:"username" gorm:"unique;not null" json:"user_name"`
	Password string `form:"password" json:"-"`
	Email    string `form:"email" gorm:"type:varchar(100);unique" json:"email"`
	Avatar   string `form:"avatar" json:"avatar"`
	Status   string `form:"-" json:"-"`
}
