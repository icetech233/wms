package v1

import (
	"wms/service/impl"
	"wms/service/request"
	"wms/service/response"

	"github.com/gin-gonic/gin"
)

func (a *menuApi) MenuList(c *gin.Context) {
	outList := impl.MenuSrvImpl.List(c, nil)
	response.OkWithData(outList, c)
}

func (a *menuApi) MenuAdd(c *gin.Context) {
	parm := new(request.AddMenuRequest)
	c.ShouldBindJSON(parm)
	err := impl.MenuSrvImpl.Add(c, parm)
	if err != nil {
		response.Fail(-1, err.Error(), parm, c)
		return
	}
	response.Ok(c)
}
