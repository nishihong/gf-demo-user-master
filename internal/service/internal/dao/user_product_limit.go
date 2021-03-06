// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/gogf/gf-demo-user/v2/internal/service/internal/dao/internal"
)

// userProductLimitDao is the data access object for table yjs_user_product_limit.
// You can define custom methods on it to extend its functionality as you wish.
type userProductLimitDao struct {
	*internal.UserProductLimitDao
}

var (
	// UserProductLimit is globally public accessible object for table yjs_user_product_limit operations.
	UserProductLimit = userProductLimitDao{
		internal.NewUserProductLimitDao(),
	}
)

// Fill with you ideas below.
