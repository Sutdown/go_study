package logic

import (
	"github.com/Sutdown/go_study/mod/dao/mysql"
	"github.com/Sutdown/go_study/mod/models"
	"github.com/Sutdown/go_study/mod/pkg/snowflake"
)

// 存放业务逻辑

func SignUp(p *models.ParamSignUp) {
	// 判断用户存不存在
	mysql.QueryUserByUsername()

	// 生成UID
	snowflake.GenID()

	// 密码加密

	// 保存到数据库
	mysql.InsertUser()
}
