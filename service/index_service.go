package service

import (
	"context"
	"errors"
	"party_room_server/build_conf"
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
