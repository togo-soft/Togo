package usecases

import (
	"Togo/middleware"
	"Togo/models"
	"Togo/repositories"
	"Togo/utils"
	"crypto/sha256"
	"encoding/hex"
	"github.com/gin-gonic/gin"
)

// UserUC 是用户实例层的一个结构 用来实现用户接口 UserInterface
type UserUC struct {
}

// ur 是仓库存储层的一个实例
var ur = repositories.NewUserRepo()

// NewUserUC 会返回实例层用户模块的实例
func NewUserUC() UserInterface {
	return &UserUC{}
}

// Signup 用户注册逻辑
func (this *UserUC) Signup(ctx *gin.Context) (int, *Response) {
	var user = new(models.User)
	//检测注册信息是否解析完成
	if err := ctx.Bind(user); err != nil {
		return StatusClientError, &Response{
			Code:    ErrorParameterParse,
			Message: "解析参数错误",
		}
	}
	//判断数据是否完整
	if user.Username == "" || user.Email == "" || user.Password == "" {
		return StatusClientError, &Response{
			Code:    ErrorParameterDefect,
			Message: "重要参数缺失",
		}
	}
	//业务逻辑
	{
		//昵称为空 使用默认用户名
		if user.Nickname == "" {
			user.Nickname = user.Username
		}
		//sha256 获得密码摘要
		h := sha256.New()
		h.Write([]byte(user.Password))
		user.Password = hex.EncodeToString(h.Sum(nil))
		//账户状态
		user.Status = true
	}
	//插入数据库
	if id, err := ur.Insert(user); err != nil {
		//插入数据库失败
		return StatusServerError, &Response{
			Code:    ErrorDatabaseInsert,
			Message: "数据库插入出错",
			Data:    err,
		}
	} else {
		//操作成功
		return StatusOK, &Response{
			Code:    StatusOK,
			Message: "用户注册成功!",
			Data:    utils.ParseInt64ToString(id),
		}
	}
}

// Signin 用户登陆逻辑
func (this *UserUC) Signin(ctx *gin.Context) (int, *Response) {
	var user = new(models.User)
	//检测登陆信息是否绑定成功
	if err := ctx.Bind(user); err != nil {
		return StatusClientError, &Response{
			Code:    ErrorParameterParse,
			Message: "解析参数错误",
		}
	}
	//判断数据是否完整
	if user.Username != "" || user.Password != "" {
		//业务逻辑
		{
			//转换password
			h := sha256.New()
			h.Write([]byte(user.Password))
			user.Password = hex.EncodeToString(h.Sum(nil))
			//查询是否存在
			if has, user := ur.RecordExist(user.Username, user.Password); has {
				//登陆成功
				token := middleware.NewJWT().NewToken(user.Id, user.Username, user.Email, user.Phone)
				//todo 更新user信息
				//将token在data中返回
				return StatusOK, &Response{
					Code:    StatusOK,
					Message: "ok",
					Data:    token,
				}
			} else {
				//登陆失败
				return StatusServerError, &Response{
					Code:    ErrorDatabaseQuery,
					Message: "数据库未找到该记录",
				}
			}
		}
	} else {
		return StatusClientError, &Response{
			Code:    ErrorParameterDefect,
			Message: "重要参数缺失",
		}
	}
}

// Logout 用户退出登陆逻辑
func (this *UserUC) Logout(ctx *gin.Context) (int, *Response) {
	return StatusOK, &Response{
		Code:    StatusOK,
		Message: "ok",
	}
}

// Profile 用户查看个人信息逻辑
func (this *UserUC) Profile(ctx *gin.Context) (int, *Response) {
	//从前端请求头中包含的UID来作用户区别 JWT只做登陆评判依据
	uid := utils.ParseStringToInt64(ctx.Request.Header.Get("uid"))
	if uid == 0 {
		return StatusServerError, &Response{
			Code:    ErrorParseRemote,
			Message: "服务端解析出错",
		}
	}
	user := ur.FindOneById(uid)
	return StatusOK, &Response{
		Code:    StatusOK,
		Message: "查询成功",
		Data:    user,
	}
}

// Cancellation 用户自注销 删除账号逻辑
func (this *UserUC) Cancellation(ctx *gin.Context) (int, *Response) {
	id := utils.ParseStringToInt64(ctx.Query("id"))
	if id, err := ur.Delete(id); err != nil {
		//删除数据库操作失败
		return StatusServerError, &Response{
			Code:    ErrorDatabaseDelete,
			Message: "数据库删除操作出错",
		}
	} else {
		//操作成功
		return StatusOK, &Response{
			Code:    StatusOK,
			Message: utils.ParseInt64ToString(id),
		}
	}
}

// ModifyInformation 用户修改信息逻辑
func (this *UserUC) ModifyInformation(ctx *gin.Context) (int, *Response) {
	var user = &models.User{}
	err := ctx.Bind(user)
	//检测参数是否正常
	if err != nil || user.Id == 0 {
		return StatusClientError, &Response{
			Code:    ErrorParameterParse,
			Message: "解析参数错误",
		}
	}
	if id, err := ur.Update(user); err != nil {
		//插入数据库失败
		return StatusServerError, &Response{
			Code:    ErrorDatabaseUpdate,
			Message: "数据库操作出错",
		}
	} else {
		//操作成功
		return StatusOK, &Response{
			Code:    StatusOK,
			Message: utils.ParseInt64ToString(id),
		}
	}
}

// FindOne ...
func (this *UserUC) FindOne(ctx *gin.Context) (int, *Response) {
	var user = &models.User{}
	id := utils.ParseStringToInt64(ctx.Query("id"))
	user = ur.FindOneById(id)
	return StatusOK, &Response{
		Code:    StatusOK,
		Message: utils.ParseInt64ToString(id),
		Data:    user,
	}
}

// FindMany ...
func (this *UserUC) FindMany(ctx *gin.Context) *List {
	return &List{
		Code: StatusOK,
		Data: ur.FindMany(),
	}
}
