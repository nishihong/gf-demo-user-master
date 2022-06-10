package service

import (
	"context"

	"github.com/gogf/gf-demo-user/v2/internal/service/internal/dao"
	"github.com/gogf/gf-demo-user/v2/internal/service/internal/do"
)

type (
	// sProduct is service struct of module Product.
	sProduct struct{}
)

var (
	// insProduct is the instance of service Product.
	insProduct = sProduct{}
)

// Product returns the interface of Product service.
func Product() *sProduct {
	return &insProduct
}

// GetProfile retrieves and returns current user info in session.
func (s *sProduct) GetProductList(ctx context.Context, name string) bool {
	count, err := dao.User.Ctx(ctx).Where(do.User{
		Name: name,
	}).Count()
	if err != nil {
		return false
	}
	return count == 0
}
