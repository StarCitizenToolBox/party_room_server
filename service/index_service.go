package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
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

	// 用户检查
	r, err := _touchUser(ctx, req.Owner, req.DeviceUUID)
	if err != nil {
		return nil, err
	}
	if r != nil {
		return nil, errors.New("您已 加入/创建 其他房间，请先离开对应房间再创建房间")
	}

	// getUserData
	userData, err := starcitizen_api.GetUserData(req.Owner)
	if err != nil {
		return nil, err
	}
	fmt.Println("[IndexServiceImpl] CreateRoom user Data === ", userData)
	var room = db.Room{
		RoomTypeID:     req.RoomTypeID,
		RoomSubTypeIds: req.RoomSubTypeIds,
		Owner:          req.Owner,
		MaxPlayer:      req.MaxPlayer,
		DeviceUUID:     req.DeviceUUID,
		Status:         protos.RoomStatus_Open,
		Announcement:   req.Announcement,
		Avatar:         userData.Data.Profile.Image,
	}
	// 创建房间
	if err := db.DB.WithContext(ctx).Create(&room).Error; err != nil {
		return nil, err
	}
	RoomServer.CreateRoom(room.ID)
	return _roomDataToRoomResultData(&room, true), nil
}

func (IndexServiceImpl) GetRoomList(ctx context.Context, req *protos.RoomListPageReqData) (*protos.RoomListData, error) {
	var rooms []*db.Room
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
	r := q.Preload("RoomUsers").Find(&rooms)
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

func (IndexServiceImpl) TouchUser(ctx context.Context, req *protos.PreUser) (*protos.RoomData, error) {
	room, err := _touchUser(ctx, req.UserName, req.DeviceUUID)
	if err != nil {
		return nil, err
	}
	if room == nil {
		return nil, nil
	}
	return _roomDataToRoomResultData(room, false), err
}

func (IndexServiceImpl) JoinRoom(preUser *protos.PreUser, stream protos.IndexService_JoinRoomServer) error {
	user, err := _joinRoom(stream.Context(), preUser.RoomID, preUser.UserName, preUser.DeviceUUID)
	if err != nil {
		return nil
	}
	ctx, cancel := context.WithCancel(stream.Context())
	// 将用户 stream 加入map
	if err := _onUserJoinRoom(user, stream, cancel); err != nil {
		return err
	}
	select {
	case <-ctx.Done():
		// user offline
		_onUserOfflineRoom(preUser.RoomID, preUser.UserName, preUser.DeviceUUID)
		return nil
	}
}

func _joinRoom(ctx context.Context, roomID string, playerName string, deviceUUID string) (*db.RoomUser, error) {
	roomUUID, err := uuid.Parse(roomID)
	if err != nil {
		return nil, err
	}
	room, err := _touchUser(ctx, playerName, deviceUUID)
	if err != nil {
		return nil, err
	}
	if room != nil && room.ID.String() != roomID {
		return nil, errors.New("您已加入了其他房间，请退出后再加入新的房间。")
	}
	var targetRoom db.Room
	if err := db.DB.WithContext(ctx).Preload("RoomUsers").Where("id = ?", roomID).Find(&targetRoom).Error; err != nil {
		return nil, err
	}
	for _, user := range targetRoom.RoomUsers {
		if user.PlayerName == playerName {
			if user.DeviceUUID != deviceUUID {
				return nil, errors.New("您已在其他设备加入了这个房间。")
			}
			// 用户已加入该房间，跳过资料创建
			user.Room = &targetRoom
			return &user, nil
		}
	}
	userData, err := starcitizen_api.GetUserData(playerName)
	if err != nil {
		return nil, err
	}
	newUser := &db.RoomUser{
		RoomID:     roomUUID,
		PlayerName: playerName,
		DeviceUUID: deviceUUID,
		Avatar:     userData.Data.Profile.Image,
		Status:     protos.RoomUserStatus_RoomUserStatusJoin,
	}
	if err := db.DB.Create(newUser).Error; err != nil {
		return nil, err
	}
	newUser.Room = &targetRoom
	return newUser, nil
}

func _touchUser(ctx context.Context, playerName string, deviceUUID string) (*db.Room, error) {
	// 查找用户已创建的房间
	var userCreatedRooms []db.Room
	if err := db.DB.WithContext(ctx).Order("updated_at desc").Preload("RoomUsers").Where("owner = ?", playerName).Find(&userCreatedRooms).Error; err != nil {
		return nil, err
	}
	if len(userCreatedRooms) != 0 {
		for _, room := range userCreatedRooms {
			if room.DeviceUUID != deviceUUID {
				return nil, errors.New("您的账号在其他设备存在未销毁的房间，请退出房间或等待自动销毁后重试")
			}
		}
		return &(userCreatedRooms[0]), nil
	}

	// 查找用户已加入的房间
	var user []db.RoomUser
	if err := db.DB.WithContext(ctx).Order("updated_at desc").Preload("Room").Find(&user, "player_name = ?", playerName).Error; err != nil {
		return nil, err
	}
	if len(user) == 0 {
		return nil, nil
	}
	for _, roomUser := range user {
		if roomUser.DeviceUUID != deviceUUID {
			return nil, errors.New("您的账号在其他设备存在未退出的房间，无法在新的设备 创建/加入 房间")
		}
	}
	return user[0].Room, nil
}

func _getRoomPlayerList(roomID string) ([]db.RoomUser, error) {
	var players []db.RoomUser
	if err := db.DB.Find(&players, "room_id = ?", roomID).Error; err != nil {
		return []db.RoomUser{}, err
	}
	return players, nil
}

func _getRoomPlayerListForGrpc(roomID string) ([]*protos.RoomUserData, error) {
	roomPlayers, err := _getRoomPlayerList(roomID)
	if err != nil {
		return []*protos.RoomUserData{}, err
	}
	var players []*protos.RoomUserData
	for _, player := range roomPlayers {
		players = append(players, _roomUserToGrpcRoomUserData(&player))
	}
	return players, nil
}

func _setUserStatus(roomID string, playerName string, status protos.RoomUserStatus) (*db.RoomUser, error) {
	var roomUser db.RoomUser
	if err := db.DB.Find(&roomUser, "roomID = ? AND player_name = ?", roomID, playerName).Error; err != nil {
		return nil, err
	}
	roomUser.Status = status
	return &roomUser, db.DB.Updates(&roomUser).Error
}

func _roomDataToRoomResultData(room *db.Room, fullInfo bool) *protos.RoomData {
	r := &protos.RoomData{
		Id:             room.ID.String(),
		RoomTypeID:     room.RoomTypeID,
		RoomSubTypeIds: room.RoomSubTypeIds,
		MaxPlayer:      room.MaxPlayer,
		Status:         room.Status,
		CurPlayer:      int32(len(room.RoomUsers)),
		Announcement:   "-",
		Avatar:         room.Avatar,
		CreateTime:     room.CreatedAt.UnixMilli(),
		UpdateTime:     room.UpdatedAt.UnixMilli(),
	}
	if fullInfo {
		r.Owner = room.Owner
		r.Announcement = room.Announcement
	} else {
		r.Owner = _maskMiddleChars(room.Owner)
	}
	return r
}
func _roomUserToGrpcRoomUserData(user *db.RoomUser) *protos.RoomUserData {
	return &protos.RoomUserData{
		Id:         user.ID.String(),
		PlayerName: user.PlayerName,
		Avatar:     user.Avatar,
		Status:     user.Status,
	}
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
