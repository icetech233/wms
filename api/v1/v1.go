package v1

import (
	"wms/entity"
	"wms/util"

	"github.com/acmestack/gorm-plus/gplus"
	"github.com/gin-gonic/gin"
)

func GetSpu(c *gin.Context) {
	// c.ShouldBindJSON()

}

func SpuList(c *gin.Context) {
	// c.ShouldBindJSON()
	spus, resultDb := gplus.SelectList[entity.Spu](nil)
	if resultDb.Error != nil {
		panic("sql err" + resultDb.Error.Error())
	}
	// spew.Dump(resultDb.Error, resultDb.RowsAffected)
	spuListOut := make([]*Spu, 0, len(spus))
	for _, spu := range spus {
		//b, _ := json.Marshal(spu)
		_s := new(Spu)
		util.CopyStruct(_s, spu)
		spuListOut = append(spuListOut, _s)
		//fmt.Println("spu:", string(b))
	}

	c.JSON(200, spuListOut)
}

type Spu struct {
	SpuID     int64  `json:"-"`
	SpuCode   string `json:"spuCode"`
	SpuDesc   string `json:"spuDesc"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
