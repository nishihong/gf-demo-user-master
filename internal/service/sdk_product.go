package service

import (
	"context"

	v1 "github.com/gogf/gf-demo-user/v2/api/v1"

	"github.com/gogf/gf-demo-user/v2/internal/service/internal/dao"

	"github.com/gogf/gf/v2/errors/gerror"
)

type (
	// sSdkProduct is service struct of module Product.
	sSdkProduct struct{}
)

var (
	// insSdkProduct is the instance of service Product.
	insSdkProduct = sSdkProduct{}
)

// Product returns the interface of Product service.
func SdkProduct() *sSdkProduct {
	return &insSdkProduct
}

// GetProfile retrieves and returns current user info in session.
func (s *sSdkProduct) GetSdkProductList(ctx context.Context) (list []v1.SdkProductItem) {

	mod := dao.SdkProduct.Ctx(ctx)

	err := mod.Where("status=", 1).Order("sort asc").Scan(&list)
	if err != nil {
		err = gerror.Newf(`ErrorORM`)

		return nil
	}

	return list
}
