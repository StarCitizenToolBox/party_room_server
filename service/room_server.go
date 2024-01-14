package service

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"party_room_server/db"
	"party_room_server/protos"
	"sync"
)

type _usersStream struct {
	PlayerName       string
	DeviceUUID       uuid.UUID
	Stream           protos.IndexService_JoinRoomServer
	StreamCancelFunc context.CancelFunc
}

type _roomUsers struct {
	ID    uuid.UUID
	Users map[string]_usersStream
}

type _roomServer struct {
	mu    sync.RWMutex
	Rooms map[uuid.UUID]_roomUsers
}

var RoomServer = _roomServer{}

func (r *_roomServer) CreateRoom(roomID uuid.UUID) _roomUsers {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.Rooms[roomID]; !ok {
		r.Rooms[roomID] = _roomUsers{
			ID:    roomID,
			Users: map[string]_usersStream{},
		}
	}
	return r.Rooms[roomID]
}

func (r *_roomServer) AddUser(roomID uuid.UUID, playerName string, deviceUUID uuid.UUID, stream protos.IndexService_JoinRoomServer, streamCancelFunc context.CancelFunc) error {
	room := r.CreateRoom(roomID)
	r.mu.Lock()
	defer r.mu.Unlock()

	if v, ok := room.Users[playerName]; ok {
		defer func() {
			v.StreamCancelFunc()
			delete(room.Users, playerName)
		}()
		if err := v.Stream.Send(&protos.RoomUpdateMessage{
			RoomUpdateType: protos.RoomUpdateType_RoomClose,
		}); err != nil {
			return err
		}
	}
	room.Users[playerName] = _usersStream{
		PlayerName:       playerName,
		DeviceUUID:       deviceUUID,
		StreamCancelFunc: streamCancelFunc,
		Stream:           stream,
	}
	return nil
}

func (r *_roomServer) DelUser(roomID uuid.UUID, playerName string) error {
	room := r.CreateRoom(roomID)
	r.mu.Lock()
	defer r.mu.Unlock()
	if v, ok := room.Users[playerName]; ok {
		defer func() {
			v.StreamCancelFunc()
			delete(room.Users, playerName)
		}()
		if err := v.Stream.Send(&protos.RoomUpdateMessage{
			RoomUpdateType: protos.RoomUpdateType_RoomClose,
		}); err != nil {
			return err
		}
	}
	return nil
}

func (r *_roomServer) SendBroadcastMessage(roomID uuid.UUID, msg *protos.RoomUpdateMessage) {
	room := r.CreateRoom(roomID)
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, usersStream := range room.Users {
		_ = usersStream.Stream.Send(msg)
	}
}

func (r *_roomServer) sendMessageToRoomUser(roomID uuid.UUID, userName string, msg *protos.RoomUpdateMessage) error {
	room := r.CreateRoom(roomID)
	r.mu.RLock()
	defer r.mu.RUnlock()
	if user, ok := room.Users[userName]; ok {
		return user.Stream.Send(msg)
	}
	return errors.New("USER_NOT_FOUNT")
}

func _onUserJoinRoom(user *db.RoomUser, stream protos.IndexService_JoinRoomServer, streamCancelFunc context.CancelFunc) error {
	deviceUUID, err := uuid.Parse(user.DeviceUUID)
	if err != nil {
		return err
	}
	err = RoomServer.AddUser(user.RoomID, user.PlayerName, deviceUUID, stream, streamCancelFunc)
	if err != nil {
		return err
	}
	// 向刚加入的用户发送全部房间信息
	roomUsersList, err := _getRoomPlayerListForGrpc(user.RoomID.String())
	if err != nil {
		return err
	}
	if err := RoomServer.sendMessageToRoomUser(user.RoomID, user.PlayerName, &protos.RoomUpdateMessage{
		RoomUpdateType: protos.RoomUpdateType_RoomUpdateData,
		RoomData:       _roomDataToRoomResultData(user.Room, false),
		UsersData:      roomUsersList,
	}); err != nil {
		return err
	}
	// 向房间内所有用户发送用户加入提醒
	RoomServer.SendBroadcastMessage(user.RoomID, &protos.RoomUpdateMessage{
		RoomUpdateType: protos.RoomUpdateType_RoomUpdateData,
		UsersData: []*protos.RoomUserData{
			_roomUserToGrpcRoomUserData(user),
		},
	})
	return nil
}

func _onUserOfflineRoom(roomID string, playerName string, _ string) {
	user, err := _setUserStatus(roomID, playerName, protos.RoomUserStatus_RoomUserStatusLostOffline)
	if err != nil {
		return
	}
	// 向房间内所有用户发送用户状态更新
	RoomServer.SendBroadcastMessage(user.RoomID, &protos.RoomUpdateMessage{
		RoomUpdateType: protos.RoomUpdateType_RoomUpdateData,
		UsersData: []*protos.RoomUserData{
			_roomUserToGrpcRoomUserData(user),
		},
	})
	// 删除用户
	_ = RoomServer.DelUser(user.RoomID, playerName)
}
