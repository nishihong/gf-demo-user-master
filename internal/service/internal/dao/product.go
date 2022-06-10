// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/gogf/gf-demo-user/v2/internal/service/internal/dao/internal"
)

// productDao is the data access object for table yjs_product.
// You can define custom methods on it to extend its functionality as you wish.
type productDao struct {
	*internal.ProductDao
}

var (
	// Product is globally public accessible object for table yjs_product operations.
	Product = productDao{
		internal.NewProductDao(),
	}
)

// Fill with you ideas below.