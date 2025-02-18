package router

import (
	"net/http"

	"github.com/Sutdown/go_study/mod/controller"

	"github.com/Sutdown/go_study/mod/logger"
	"github.com/gin-gonic/gin"
)

// 这段代码通过 Gin 框架实现了一个简单的 HTTP 服务。
// 初始化一个 Gin 框架的路由引擎。
// 使用自定义的中间件（GinLogger 和 GinRecovery）记录请求日志并处理 panic 恢复。
// 定义一个根路由（/），返回简单的响应。

func Setup() *gin.Engine {
	controller.InitTrans("zh")
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 注册业务路由
	r.POST("/signup", controller.SignUpHandler)
	r.POST("/login", controller.LoginHandler)
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
