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
	var room = db.RoomsTable{
		RoomTypeID:     req.RoomTypeID,
		RoomSubTypeIds: req.RoomSubTypeIds,
		Owner:          req.Owner,
		MaxPlayer:      req.MaxPlayer,
		Status:         protos.RoomStatus_Open,
	}
	db.DB.Create(&room).WithContext(ctx)
	return &protos.RoomData{
		Id:             room.ID.String(),
		RoomTypeID:     room.RoomTypeID,
		RoomSubTypeIds: room.RoomSubTypeIds,
		Owner:          room.Owner,
		MaxPlayer:      room.MaxPlayer,
		Status:         req.Status,
		CurPlayer:      0,
	}, nil
}

func (IndexServiceImpl) GetRoomList(ctx context.Context, req *protos.RoomListPageReqData) (*protos.RoomListData, error) {
	return nil, nil
}
