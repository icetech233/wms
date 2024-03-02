package router

import (
	"net/http"
	v1 "wms/api/v1"
	"wms/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouterV1(engine *gin.Engine) *gin.Engine {

	engine.Use(gin.Logger(), gin.Recovery())
	engine.Use(middleware.Cors())
	// engine.Use(cors.Default())
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1api := engine.Group("/api/v1")
	{
		//
		g_spu := v1api.Group("/spu")
		g_spu.GET("/list", v1.SpuV1.SpuList)
		g_spu.POST("/add", v1.SpuV1.SpuAdd)
		//
		g_attr := v1api.Group("/attr")
		g_attr.GET("/list", v1.AttrV1.AttrList)
		g_attr.POST("/add", v1.AttrV1.AttrAdd)
		//
		g_menu := v1api.Group("/menu")
		g_menu.GET("/list", v1.MenuV1.MenuAdd)
		g_menu.POST("/add", v1.MenuV1.MenuAdd)
	}

	return engine
}
