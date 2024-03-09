package v1

import (
	"wms/service/impl"
	"wms/service/request"
	"wms/service/response"

	"github.com/gin-gonic/gin"
)

func (a *attrApi) AttrList(c *gin.Context) {
	outList := impl.AttrSrvImpl.List(c, nil)
	response.OkWithData(outList, c)
}

func (a *attrApi) AttrAdd(c *gin.Context) {
	parm := new(request.AddAttrRequest)
	c.ShouldBindJSON(parm)
	err := impl.AttrSrvImpl.Add(c, parm)
	if err != nil {
		response.Fail(-1, err.Error(), parm, c)
		return
	}
	response.Ok(c)
}

func (a *attrApi) AttrEdit(c *gin.Context) {
	parm := new(request.EditAttrRequest)
	c.ShouldBindJSON(parm)
	err := impl.AttrSrvImpl.Edit(c, parm)
	if err != nil {
		response.Fail(-1, err.Error(), parm, c)
		return
	}
	response.Ok(c)
}
