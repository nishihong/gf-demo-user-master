// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/gogf/gf-demo-user/v2/internal/service/internal/dao/internal"
)

// userSerialDao is the data access object for table yjs_user_serial.
// You can define custom methods on it to extend its functionality as you wish.
type userSerialDao struct {
	*internal.UserSerialDao
}

var (
	// UserSerial is globally public accessible object for table yjs_user_serial operations.
	UserSerial = userSerialDao{
		internal.NewUserSerialDao(),
	}
)

// Fill with you ideas below.
