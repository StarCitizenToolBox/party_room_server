package db

import (
	"github.com/lib/pq"
	"party_room_server/protos"
)

type RoomsTable struct {
	BaseTable
	RoomTypeID     string            `gorm:"index" json:"room_type_id"`                  // 房间类型
	RoomSubTypeIds pq.StringArray    `gorm:"type:text[];index" json:"room_sub_type_ids"` // 房间子类型
	Owner          string            `gorm:"index" json:"owner"`                         //房主
	MaxPlayer      int32             `gorm:"index" json:"max_player"`                    // 最大玩家数量
	Status         protos.RoomStatus `gorm:"index" json:"status"`                        // 房间状态
}