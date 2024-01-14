package db

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	"party_room_server/protos"
)

// Room 大厅房间表
type Room struct {
	BaseTable
	RoomTypeID     string            `gorm:"index"`             // 房间类型 	// 房间类型
	RoomSubTypeIds pq.StringArray    `gorm:"type:text[];index"` // 房间子类型
	Owner          string            `gorm:"index"`             //房主
	MaxPlayer      int32             `gorm:"index"`             // 最大玩家数量
	Status         protos.RoomStatus `gorm:"index"`             // 房间状态
	DeviceUUID     string            `gorm:"index"`             // 房主的 Device UUID
	Announcement   string            // 公告
	Avatar         string            // 房主头像
	RoomUsers      []RoomUser        //该房间下的玩家
}

// RoomUser 大厅房间用户关联表
type RoomUser struct {
	BaseTable
	RoomID     uuid.UUID `gorm:"type:uuid;index;uniqueIndex:idx_player"` // 房间ID
	Room       *Room
	PlayerName string `gorm:"index;uniqueIndex:idx_player"` //玩家名称
	Avatar     string // 玩家头像
	DeviceUUID string `gorm:"index"` // 玩家的 Device UUID
}
