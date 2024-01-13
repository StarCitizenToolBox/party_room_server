package service

import (
	"context"
	"errors"
	"party_room_server/build_conf"
	"party_room_server/db"
	"party_room_server/protos"
	"party_room_server/starcitizen_api"
	"strings"
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

	// getUserData
	userData, err := starcitizen_api.GetUserData(req.Owner)
	if err != nil {
		return nil, err
	}
	var room = db.RoomsTable{
		RoomTypeID:     req.RoomTypeID,
		RoomSubTypeIds: req.RoomSubTypeIds,
		Owner:          req.Owner,
		MaxPlayer:      req.MaxPlayer,
		CurPlayer:      0,
		DeviceUUID:     req.DeviceUUID,
		Status:         protos.RoomStatus_Open,
		Announcement:   req.Announcement,
		Avatar:         userData.Data.Profile.Image,
	}
	db.DB.WithContext(ctx).Create(&room)
	return _roomDataToRoomResultData(&room, true), nil
}

func (IndexServiceImpl) GetRoomList(ctx context.Context, req *protos.RoomListPageReqData) (*protos.RoomListData, error) {
	var rooms []*db.RoomsTable
	q := db.DB.WithContext(ctx).Offset(int(req.PageNum)).Limit(50)
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
		roomsResult = append(roomsResult, _roomDataToRoomResultData(room, false))
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

func _roomDataToRoomResultData(room *db.RoomsTable, fullInfo bool) *protos.RoomData {
	r := &protos.RoomData{
		Id:             room.ID.String(),
		RoomTypeID:     room.RoomTypeID,
		RoomSubTypeIds: room.RoomSubTypeIds,
		MaxPlayer:      room.MaxPlayer,
		Status:         room.Status,
		CurPlayer:      room.CurPlayer,
		Announcement:   "-",
		Avatar:         room.Avatar,
	}
	if fullInfo {
		r.Owner = room.Owner
		r.Announcement = room.Announcement
	} else {
		r.Owner = _maskMiddleChars(room.Owner)
	}
	return r
}

func _maskMiddleChars(s string) string {
	length := len(s)
	if length <= 2 {
		return "***"
	}
	firstChar := string(s[0])
	lastChar := string(s[length-1])
	middleChars := strings.Repeat("*", length-2)

	return firstChar + middleChars + lastChar
}
