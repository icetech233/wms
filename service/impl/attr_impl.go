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

type AttrService struct {
}

var AttrSrvImpl AttrService

func (s *AttrService) List(ctx context.Context, arg any) []*model.Attr {
	attrs, resultDb := gplus.SelectList[entity.Attr](nil)
	if resultDb.Error != nil {
		panic("sql err" + resultDb.Error.Error())
	}
	attrListOut := make([]*model.Attr, 0, len(attrs))
	for _, attr := range attrs {
		_s := new(model.Attr)
		util.CopyStruct(_s, attr)
		attrListOut = append(attrListOut, _s)
	}

	return attrListOut
}

func (s *AttrService) Add(ctx context.Context, arg *request.AddAttrRequest) error {
	query, u := gplus.NewQuery[entity.Attr]()
	query.Eq(&u.AttrName, arg.AttrName)
	attr, resultDb := gplus.SelectOne(query)
	if resultDb.RowsAffected > 0 {
		return errors.New("编码:" + arg.AttrName + "已存在")
	}
	attr = new(entity.Attr)
	util.CopyStruct(attr, arg)
	//
	resultDb = gplus.Insert(&attr)
	if resultDb.Error != nil {
		return errors.New("insert attr err:" + resultDb.Error.Error())
	}
	fmt.Println("attr id:", attr.AttrID)
	return nil
}
