package service

import (
	"context"

	v1 "github.com/gogf/gf-demo-user/v2/api/v1"
	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/gogf/gf-demo-user/v2/internal/service/internal/dao"
)

type (
	// sSdkProduct is service struct of module Product.
	sSdkProduct struct{}
)

var (
	// insSdkProduct is the instance of service Product.
	insSdkProduct = sSdkProduct{}
)

// SdkProduct returns the interface of SdkProduct service.
func SdkProduct() *sSdkProduct {
	return &insSdkProduct
}

// GetSdkProductList
func (s *sSdkProduct) GetSdkProductList(ctx context.Context) (list []v1.SdkProductItem, err error) {

	mod := dao.SdkProduct.Ctx(ctx)

	err = mod.Where("status=", 1).Order("sort asc").Scan(&list)
	if err != nil {
		err = gerror.Newf(`ErrorORM`)
		return nil, err
	}

	return list, nil
}
