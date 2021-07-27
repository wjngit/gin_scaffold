package controller

import (
	lib "gin_scaffold/common/utils/utilLib"
	"gin_scaffold/dao"
	"gin_scaffold/dto"
	"gin_scaffold/middleware"
	"gin_scaffold/public"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
)

type DemoController struct {
}

func DemoRegister(router *gin.RouterGroup) {
	demo := DemoController{}
	router.GET("/index", demo.Index)
	router.Any("/bind", demo.Bind)
	router.GET("/dao", demo.Dao)
	router.GET("/redis", demo.Redis)
}

func (demo *DemoController) Index(c *gin.Context) {
	middleware.ResponseSuccess(c, "")
	return
}

func (demo *DemoController) Dao(c *gin.Context) {
	tx, err0 := lib.GetGormPool("default")
	if err0 != nil {
		middleware.ResponseError(c, 2000, err0)
		return
	}
	area, err := (&dao.Area{}).Find(c, tx, c.DefaultQuery("id", "4"))
	if err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	middleware.ResponseSuccess(c, area)
	return
}

func (demo *DemoController) Redis(c *gin.Context) {
	redisKey := "redis_key"
	lib.RedisConfDo(public.GetTraceContext(c),
		"default",
		"SET",
		redisKey, "redis_value")
	redisValue, err := redis.String(
		lib.RedisConfDo(public.GetTraceContext(c), "default",
			"GET",
			redisKey))
	if err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	middleware.ResponseSuccess(c, redisValue)
	return
}

func (demo *DemoController) Bind(c *gin.Context) {
	params := &dto.DemoInput{}
	if err := params.BindingValidParams(c); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	middleware.ResponseSuccess(c, params)
	return
}
