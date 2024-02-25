package v1

import (
	"wms/service"
	"wms/service/request"
	"wms/service/response"

	"github.com/gin-gonic/gin"
)

type spuApi struct{}

var SpuV1 spuApi

func (a *spuApi) SpuList(c *gin.Context) {
	spuListOut := service.SpuSrv.List(c, nil)
	response.OkWithData(spuListOut, c)
}

func (a *spuApi) SpuAdd(c *gin.Context) {
	parm := new(request.AddSpuRequest)
	c.ShouldBindJSON(parm)
	err := service.SpuSrv.Add(c, parm)
	if err != nil {
		response.Fail(-1, err.Error(), parm, c)
		return
	}
	response.Ok(c)
}
