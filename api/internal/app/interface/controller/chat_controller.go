package controller

import "otomo/internal/app/interface/controller/grpc/grpcgen"

var _ grpcgen.ChatServiceServer = (*ChatController)(nil)

type ChatController struct{}

func (c *ChatController) SendMessage(
	_ *grpcgen.ChatService_SendMessageRequest,
	_ grpcgen.ChatService_SendMessageServer,
) error {
	// STEP: Add a message to the repo
	// STEP: make a reply
	// STEP: response
	// STEP: Add a message to the repo
	panic("not implemented") // TODO: Implement
}
