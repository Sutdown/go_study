package controller

import (
	_ "errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	_ "github.com/Sutdown/go_study/mod/dao/mysql"
	"github.com/Sutdown/go_study/mod/logic"
	"github.com/Sutdown/go_study/mod/models"
	"github.com/go-playground/validator/v10"
)

func SignUpHandler(c *gin.Context) {
	// 1. 获取参数，参数校验
	var p models.ParamSignUp
	if err := c.ShouldBindJSON(&p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("SIgnUp with invalid param", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "请求参数有误",
		})
		// 判断err是不是validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// ResponseError(c, CodeInvalidParam)
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		// ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		c.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)),
		})
		return
	}
	fmt.Println(p)

	// 2. 业务处理
	logic.SignUp(&p)
	// 3. 返回响应
	c.JSON(http.StatusOK, "ok")
}

func LoginHandler(c *gin.Context) {
	// 1.获取请求参数及参数校验
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("Login with invalid param", zap.Error(err))
		// 判断err是不是validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)),
		})
		return
	}
	// 2.业务逻辑处理
	token, err := logic.Login(p)
	if err != nil {
		zap.L().Error("logic.Login failed", zap.String("username", p.Username), zap.Error(err))
		if {
			ResponseError(c, CodeUserExist)
		}
		ResponseErrorWithMsg(c, CodeServerBusy)
		return
	}
		c.JSON(http.StatusOK, gin.H{
			"msg": "密码无效",
		})
		return
	}

	// 3.返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "登陆成功",
	})
}
