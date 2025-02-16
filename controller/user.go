package controller

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"

	"github.com/Sutdown/go_study/mod/models"

	"github.com/Sutdown/go_study/mod/logic"
	"github.com/gin-gonic/gin"
)

func SignUpHandle(c *gin.Context) {
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
