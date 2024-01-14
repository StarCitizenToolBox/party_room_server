package service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"party_room_server/protos"
)

type ChatServiceImpl struct {
	protos.UnimplementedChatServiceServer
}

func (ChatServiceImpl) ListenMessage(*protos.PreUser, protos.ChatService_ListenMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method ListenMessage not implemented")
}
func (ChatServiceImpl) SendMessage(context.Context, *protos.ChatMessage) (*protos.BaseRespData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
