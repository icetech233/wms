package v1

import (
	"io"
	"wms/service/impl"
	"wms/service/request"
	"wms/service/response"
	"wms/util"

	"github.com/gin-gonic/gin"
)

func (a *attrApi) AttrEdit(c *gin.Context) {
	bb, _ := io.ReadAll(c.Request.Body)
	util.Log(string(bb) + "\n")
	response.Ok(c)
}

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
