package v1

import (
	"wms/service/impl"
	"wms/service/response"

	"github.com/gin-gonic/gin"
)

type spuApi struct{}

var SpuV1 spuApi

func GetSpu(c *gin.Context) {
	// c.ShouldBindJSON()
}

func (a *spuApi) SpuList(c *gin.Context) {
	// c.ShouldBindJSON()

	spuListOut := impl.SpuSrv.List(c, nil)
	response.OkWithData(spuListOut, c)
}
