package service

import (
	"context"
	"wms/service/impl"
	"wms/service/model"
	"wms/service/request"
)

var SpuSrv ISpuService = impl.NewSpuService()

type ISpuService interface {
	List(ctx context.Context, arg any) []*model.Spu
	Add(ctx context.Context, arg *request.AddSpuRequest) error
}
