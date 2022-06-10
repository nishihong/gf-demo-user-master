// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/gogf/gf-demo-user/v2/internal/service/internal/dao/internal"
)

// adminRoleUsersDao is the data access object for table yjs_admin_role_users.
// You can define custom methods on it to extend its functionality as you wish.
type adminRoleUsersDao struct {
	*internal.AdminRoleUsersDao
}

var (
	// AdminRoleUsers is globally public accessible object for table yjs_admin_role_users operations.
	AdminRoleUsers = adminRoleUsersDao{
		internal.NewAdminRoleUsersDao(),
	}
)

// Fill with you ideas below.