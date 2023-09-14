package controller

import (
	"context"
	"otomo/internal/app/controller/repo"
	"otomo/internal/app/controller/svc"
	"otomo/internal/app/grpcgen"
	"otomo/internal/app/model"
	"otomo/internal/pkg/ctxs"
	"otomo/internal/pkg/errs"
	"otomo/internal/pkg/times"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// TODO: Add tests
// TODO: Add presenter

var _ grpcgen.ChatServiceServer = (*ChatController)(nil)

var (
	ErrDomainMessage = "message"
)

type ChatController struct {
	errorPresenter errorPresenter
	chatFactory    *model.ChatFactory
	msgFactory     *model.MessageFactory
	chatRepo       repo.ChatRepository
	msgRepo        repo.MessageRepository
	chatService    svc.ChatService
	summaryService svc.SummaryService
}

func NewChatController(
	errorPresenter errorPresenter,
	chatFactory *model.ChatFactory,
	msgFactory *model.MessageFactory,
	chatRepo repo.ChatRepository,
	msgRepo repo.MessageRepository,
	chatService svc.ChatService,
	summaryService svc.SummaryService,
) *ChatController {
	return &ChatController{
		errorPresenter: errorPresenter,
		chatFactory:    chatFactory,
		msgFactory:     msgFactory,
		chatRepo:       chatRepo,
		msgRepo:        msgRepo,
		chatService:    chatService,
		summaryService: summaryService,
	}
}

// TODO: Implement transaction
func (c *ChatController) SendMessage(
	req *grpcgen.ChatService_SendMessageRequest,
	stream grpcgen.ChatService_SendMessageServer,
) error {
	if err := c.sendMessage(req, stream); err != nil {
		return c.toGrpcError(stream.Context(), err)
	}
	return nil
}

func (c *ChatController) sendMessage(
	req *grpcgen.ChatService_SendMessageRequest,
	stream grpcgen.ChatService_SendMessageServer,
) error {
	var ctx = stream.Context()

	if err := req.ValidateAll(); err != nil {
		return err
	}

	userID, err := ctxs.UserIDFromContext(ctx)
	if err != nil {
		return err
	}

	msg, err := c.msgFactory.New(req.GetText(), model.UserRole)
	if err != nil {
		return err
	}

	if err := c.msgRepo.Add(ctx, userID, msg); err != nil {
		return err
	}

	chat, err := c.chatRepo.Get(ctx, userID)
	if err != nil {
		return err
	}

	reply, err := c.chatService.Send(
		ctx,
		userID,
		msg,
		chat.Summary,
		func(ctx context.Context, chunk []byte) error {
			return stream.Send(&grpcgen.ChatService_SendMessageStreamResponse{
				Text:   string(chunk),
				SentAt: timestamppb.New(times.C.Now()),
			})
		},
	)
	if err != nil {
		return err
	}

	if err := c.msgRepo.Add(ctx, userID, reply); err != nil {
		return err
	}

	newSummary, err := c.summaryService.Summarize(
		ctx, []*model.Message{msg, reply}, chat.Summary)
	if err != nil {
		return err
	}

	return c.saveChat(ctx, userID, newSummary)
}

func (c *ChatController) saveChat(
	ctx context.Context,
	userID model.UserID,
	summary string,
) error {
	newMsgChat, err := c.chatFactory.New(summary)
	if err != nil {
		return err
	}

	return c.chatRepo.Save(ctx, userID, newMsgChat)
}

func (c *ChatController) ListMessages(
	ctx context.Context,
	req *grpcgen.ChatService_ListMessagesRequest,
) (*grpcgen.ChatService_ListMessagesResponse, error) {
	res, err := c.listMessages(ctx, req)
	if err != nil {
		return nil, c.toGrpcError(ctx, err)
	}

	return res, nil
}

func (c *ChatController) listMessages(
	ctx context.Context,
	req *grpcgen.ChatService_ListMessagesRequest,
) (*grpcgen.ChatService_ListMessagesResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, err
	}

	if !ctxs.UserIs(ctx, model.UserID(req.UserId)) {
		return nil, status.New(codes.PermissionDenied, "can only get own list").Err()
	}

	msgs, err := c.msgRepo.List(
		ctx,
		model.UserID(req.GetUserId()),
		&repo.MessagePage{
			Size: int(req.GetPageSize()),
			StartAfterMessageID: model.MessageID(
				req.GetPageStartAfterMessageId()),
		})
	if err != nil {
		return nil, err
	}

	resMsgs, err := c.toGrpcMessages(msgs)
	if err != nil {
		return nil, err
	}

	return &grpcgen.ChatService_ListMessagesResponse{
		PageSize: uint32(len(resMsgs)),
		Messages: resMsgs,
	}, nil
}

func (c *ChatController) toGrpcMessages(
	msgs []*model.Message,
) ([]*grpcgen.Message, error) {
	grpcMsgs := make([]*grpcgen.Message, len(msgs))
	for i, msg := range msgs {
		grpcMsg, err := c.toGrpcMessage(msg)
		if err != nil {
			return nil, err
		}
		grpcMsgs[i] = grpcMsg
	}
	return grpcMsgs, nil
}

func (c *ChatController) toGrpcMessage(
	msg *model.Message,
) (*grpcgen.Message, error) {
	role, err := toGrpcRole(msg.Role)
	if err != nil {
		return nil, err
	}

	return &grpcgen.Message{
		Id:     string(msg.ID),
		Text:   msg.Text,
		Role:   role,
		SentAt: timestamppb.New(msg.SentAt),
	}, nil
}

func toGrpcRole(r model.Role) (role grpcgen.Role, err error) {
	switch r {
	case model.UserRole:
		role = grpcgen.Role_USER
	case model.OtomoRole:
		role = grpcgen.Role_OTOMO
	default:
		err = &errs.Error{
			Domain: errs.DomainMessage,
			Cause:  errs.CauseNotExist,
			Field:  errs.FieldRole,
		}
	}
	return
}

func (c *ChatController) toGrpcError(ctx context.Context, err error) error {
	return c.errorPresenter.ErrorOutput(ctx, err).Err()
}

func (c *ChatController) AskToMessage(
	context context.Context,
	req *grpcgen.ChatService_AskToMessageRequest,
) (*grpcgen.ChatService_AskToMessageResponse, error) {
	panic("not implemented") // TODO: Implement
}
