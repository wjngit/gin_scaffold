package router

import (
	"gin_scaffold/controller"
	"gin_scaffold/middleware"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	router := gin.Default()
	router.Use(middlewares...)

	//test接口
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	//demo
	demo := router.Group("/demo")
	demo.Use(
		middleware.RecoveryMiddleware(),
		middleware.RequestLog(),
		middleware.IPAuthMiddleware(),
		middleware.TranslationMiddleware())
	{
		controller.DemoRegister(demo)
	}

	//session store
	store := sessions.NewCookieStore([]byte("secret"))

	//非登录接口
	apiNormalGroup := router.Group("/apiNormal")
	apiNormalGroup.Use(
		sessions.Sessions("my_session", store),
		middleware.RecoveryMiddleware(),
		middleware.RequestLog(),
		middleware.IPAuthMiddleware(),
		middleware.TranslationMiddleware())
	{
		controller.ApiRegister(apiNormalGroup)
	}

	//登录接口
	apiAuthGroup := router.Group("/apiAuth")
	apiAuthGroup.Use(
		sessions.Sessions("my_session", store),
		middleware.RecoveryMiddleware(),
		middleware.RequestLog(),
		middleware.SessionAuthMiddleware(),
		middleware.TranslationMiddleware())
	{
		controller.ApiLoginRegister(apiAuthGroup)
	}
	return router
}
