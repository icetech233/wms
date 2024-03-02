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

type menuService struct {
}

func NewMenuService() *menuService {
	return new(menuService)
}

var MenuSrvImpl menuService

func (s *menuService) List(ctx context.Context, arg any) []*model.Menu {
	entityList, resultDb := gplus.SelectList[entity.Menu](nil)
	if resultDb.Error != nil {
		panic("sql err" + resultDb.Error.Error())
	}
	//
	outList := make([]*model.Menu, 0, len(entityList))
	for _, _entity := range entityList {
		_model := new(model.Menu)
		util.CopyStruct(_model, _entity)
		outList = append(outList, _model)
	}
	return outList
}

func (s *menuService) Add(ctx context.Context, arg *request.AddMenuRequest) error {
	query, u := gplus.NewQuery[entity.Menu]()
	query.Eq(&u.Key, arg.Key)
	menu, resultDb := gplus.SelectOne(query)
	if resultDb.RowsAffected > 0 {
		return errors.New("key:" + arg.Key + "已存在")
	}
	menu = new(entity.Menu)
	util.CopyStruct(menu, arg)
	//
	resultDb = gplus.Insert(&menu)
	if resultDb.Error != nil {
		return errors.New("insert menu err:" + resultDb.Error.Error())
	}
	fmt.Println("menu id:", menu.MenuID)
	return nil
}
