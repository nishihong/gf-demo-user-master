// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-06-06 15:21:14
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// NodesMachineroom is the golang structure for table nodes_machineroom.
type NodesMachineroom struct {
	Id            int         `json:"id"            ` //
	NodeId        int         `json:"nodeId"        ` // 节点id
	MachineroomId int         `json:"machineroomId" ` // 机房id
	CreatedAt     *gtime.Time `json:"createdAt"     ` //
	UpdatedAt     *gtime.Time `json:"updatedAt"     ` //
}
