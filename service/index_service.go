package service

import (
	"context"
	"errors"
	"party_room_server/build_conf"
	"party_room_server/db"
	"party_room_server/protos"
)

type IndexServiceImpl struct {
	protos.UnimplementedIndexServiceServer
}

func (IndexServiceImpl) PingServer(_ context.Context, data *protos.PingData) (*protos.PingData, error) {
	if data.Data == "PING" {
		if data.ClientVersion >= build_conf.MinReqClientVersion {
			return &protos.PingData{Data: "PONG", ClientVersion: data.ClientVersion, ServerVersion: build_conf.ServerVersion}, nil
		}
		return nil, errors.New("band Client Version")
	}
	return nil, errors.New("bad Request")
}

func (IndexServiceImpl) GetRoomTypes(_ context.Context, _ *protos.Empty) (*protos.RoomTypesData, error) {
	return &protos.RoomTypesData{
		RoomTypes: build_conf.RoomTypesData,
	}, nil
}

func (IndexServiceImpl) CreateRoom(ctx context.Context, req *protos.RoomData) (*protos.RoomData, error) {
	// TODO 格式检查 子类型约束检查
	var room = db.RoomsTable{
		RoomTypeID:     req.RoomTypeID,
		RoomSubTypeIds: req.RoomSubTypeIds,
		Owner:          req.Owner,
		MaxPlayer:      req.MaxPlayer,
		DeviceUUID:     req.DeviceUUID,
		Status:         protos.RoomStatus_Open,
	}
	db.DB.WithContext(ctx).Create(&room)
	return _roomDataToRoomResultData(&room), nil
}

func (IndexServiceImpl) GetRoomList(ctx context.Context, req *protos.RoomListPageReqData) (*protos.RoomListData, error) {
	var rooms []*db.RoomsTable
	q := db.DB.WithContext(ctx).Offset(int(req.PageNum)).Limit(50)
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	if req.Status != protos.RoomStatus_All {
		q = q.Where("status = ?", req.Status)
	}
	if req.TypeID != "" {
		q = q.Where("room_type_id = ?", req.TypeID)
	}
	if req.SubTypeID != "" {
		q = q.Where("? = ANY(room_sub_type_ids)", req.SubTypeID)
	}
	// TODO 排序
	r := q.Find(&rooms)
	if r.Error != nil {
		return nil, r.Error
	}
	var roomsResult []*protos.RoomData
	for _, room := range rooms {
		roomsResult = append(roomsResult, _roomDataToRoomResultData(room))
	}

	return &protos.RoomListData{
		PageData: &protos.BasePageRespData{
			Code:       200,
			Message:    "",
			HasNext:    len(rooms) >= 50,
			CurPageNum: req.PageNum,
		},
		Rooms: roomsResult,
	}, nil
}

func _roomDataToRoomResultData(room *db.RoomsTable) *protos.RoomData {
	return &protos.RoomData{
		Id:             room.ID.String(),
		RoomTypeID:     room.RoomTypeID,
		RoomSubTypeIds: room.RoomSubTypeIds,
		Owner:          room.Owner,
		MaxPlayer:      room.MaxPlayer,
		Status:         room.Status,
		CurPlayer:      0, // TODO 获取房间实际人数
	}
}
