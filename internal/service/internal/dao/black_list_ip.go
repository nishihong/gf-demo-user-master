// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/gogf/gf-demo-user/v2/internal/service/internal/dao/internal"
)

// blackListIpDao is the data access object for table yjs_black_list_ip.
// You can define custom methods on it to extend its functionality as you wish.
type blackListIpDao struct {
	*internal.BlackListIpDao
}

var (
	// BlackListIp is globally public accessible object for table yjs_black_list_ip operations.
	BlackListIp = blackListIpDao{
		internal.NewBlackListIpDao(),
	}
)

// Fill with you ideas below.
