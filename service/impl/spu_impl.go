package impl

import (
	"context"
	"errors"
	"fmt"
	"wms/entity"
	"wms/service/model"
	"wms/service/request"
	"wms/util"

	"github.com/acmestack/gorm-plus/gplus"
)

type spuService struct {
}

func NewSpuService() *spuService {
	return new(spuService)
}

var SpuSrvImpl spuService

func (s *spuService) List(ctx context.Context, arg any) []*model.Spu {

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

func (s *spuService) Add(ctx context.Context, arg *request.AddSpuRequest) error {
	// spu := new(entity.Spu)
	// util.CopyStruct(spu, arg)

	query, u := gplus.NewQuery[entity.Spu]()
	query.Eq(&u.SpuCode, arg.SpuCode)
	spu, resultDb := gplus.SelectOne(query)
	if resultDb.RowsAffected > 0 {
		return errors.New("编码:" + arg.SpuCode + "已存在")
	}
	spu = new(entity.Spu)
	util.CopyStruct(spu, arg)
	//
	resultDb = gplus.Insert(&spu)
	if resultDb.Error != nil {
		return errors.New("insert spu err:" + resultDb.Error.Error())
	}
	fmt.Println("spu id:", spu.SpuID)
	return nil
}
