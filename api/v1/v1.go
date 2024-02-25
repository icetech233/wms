package v1

import (
	"wms/service"
	"wms/service/response"

	"github.com/gin-gonic/gin"
)

type spuApi struct{}

var SpuV1 spuApi

func (a *spuApi) SpuList(c *gin.Context) {
	// c.ShouldBindJSON()

	spuListOut := service.SpuSrv.List(c, nil)
	response.OkWithData(spuListOut, c)
}

func (a *spuApi) SpuAdd(c *gin.Context) {
	// c.ShouldBindJSON()

	spuListOut := service.SpuSrv.List(c, nil)
	response.OkWithData(spuListOut, c)
}
