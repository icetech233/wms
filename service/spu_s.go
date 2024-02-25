package service

import (
	"context"
	"wms/service/impl"
	"wms/service/model"
)

var SpuSrv ISpuService = &impl.SpuService{}

type ISpuService interface {
	List(ctx context.Context, arg any) []*model.Spu
}
