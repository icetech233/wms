package impl

import (
	"context"
	"errors"
	"wms/entity"
	"wms/service/model"
	"wms/service/request"
	"wms/util"

	"github.com/acmestack/gorm-plus/gplus"
)

type AttrService struct {
}

var AttrSrvImpl AttrService

// /attr/edit

func (s *AttrService) Edit(ctx context.Context, arg any) error {

	return nil
}

func (s *AttrService) List(ctx context.Context, arg any) []*model.Attr {
	attrs, resultDb := gplus.SelectList[entity.Attr](nil)
	if resultDb.Error != nil {
		panic("sql err" + resultDb.Error.Error())
	}

	attrids := make([]int64, len(attrs))
	for i, item := range attrs {
		attrids[i] = item.AttrID
	}

	avMap := make(map[int64][]*model.AttrValue)
	// 查询子表
	query, av := gplus.NewQuery[entity.AttrValue]()
	query.In(&av.AttrID, attrids)
	attrvals, _ := gplus.SelectList[entity.AttrValue](query)

	for _, attrval := range attrvals {
		vlist, e1 := avMap[attrval.AttrID]
		if !e1 {
			vlist = make([]*model.AttrValue, 0)
		}
		_m := new(model.AttrValue)
		util.CopyStruct(_m, attrval)
		vlist = append(vlist, _m)
		avMap[attrval.AttrID] = vlist
	}

	attrListOut := make([]*model.Attr, 0, len(attrs))
	for _, attr := range attrs {
		_s := new(model.Attr)
		util.CopyStruct(_s, attr)
		// 补充对应数据
		_s.Val = avMap[attr.AttrID]
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
	util.Log("attr id:", attr.AttrID)
	return nil
}
