package impl

import (
	"context"
	"wms/entity"
	"wms/service/model"
	"wms/util"

	"github.com/acmestack/gorm-plus/gplus"
)

type SpuService struct {
}

var SpuSrvImpl SpuService

func (s *SpuService) List(ctx context.Context, arg any) []*model.Spu {

	spus, resultDb := gplus.SelectList[entity.Spu](nil)
	if resultDb.Error != nil {
		panic("sql err" + resultDb.Error.Error())
	}
	//
	spuListOut := make([]*model.Spu, 0, len(spus))
	for _, spu := range spus {
		_s := new(model.Spu)
		util.CopyStruct(_s, spu)
		spuListOut = append(spuListOut, _s)
	}

	return spuListOut
}
