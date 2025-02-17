package logic

import (
	"github.com/Sutdown/go_study/mod/dao/mysql"
	"github.com/Sutdown/go_study/mod/models"
	"github.com/Sutdown/go_study/mod/pkg/snowflake"
)

// 存放业务逻辑

func SignUp(p *models.ParamSignUp) (err error) {
	// 1.判断用户存不存在
	err = mysql.CheckUserExist(p.Username)
	if err != nil {
		return err // 数据库查询出错
	}

	// 2.生成UID
	userID := snowflake.GenID()

	// 构造user实例
	user := models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	// 3.密码加密
	// 3.保存到数据库
	mysql.InsertUser(&user)
	return err
}
