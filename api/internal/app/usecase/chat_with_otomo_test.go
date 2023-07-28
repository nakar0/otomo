package usecase

import (
	"context"
	"otomo/internal/app/domain/entity/message"
	"otomo/internal/app/domain/entity/user"
	"otomo/internal/app/domain/gateway/infra/mock_infra"
	"otomo/internal/app/domain/gateway/repo/mock_repo"
	"otomo/pkg/uuid"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type chatWithOtomoUseCaseFields struct {
	msgMaker *mock_infra.MockMessageMaker
	msgRepo  *mock_repo.MockMessageWithOtomoRepository
}

func newChatWithOtomoUseCaseFields(t *testing.T) *chatWithOtomoUseCaseFields {
	ctrl := gomock.NewController(t)
	msgMaker := mock_infra.NewMockMessageMaker(ctrl)
	msgRepo := mock_repo.NewMockMessageWithOtomoRepository(ctrl)
	return &chatWithOtomoUseCaseFields{
		msgMaker: msgMaker,
		msgRepo:  msgRepo,
	}
}

func TestChatWithOtomoUseCase_MessageToOtomo(t *testing.T) {
	var (
		giveCtx            = context.TODO()
		giveUserID         = user.ID(uuid.NewString())
		giveText           = "Hello, Otomo!"
		truncatedSecondNow = time.Now().Truncate(time.Second)
		reply              = message.RestoreMessageWithOtomo(
			message.MessageWithOtomoID(uuid.NewString()),
			giveUserID,
			message.OtomoRole,
			message.UserRole,
			"Hello, User!",
			time.Now(),
		)
	)
	type args struct {
		ctx    context.Context
		userID user.ID
		text   string
	}
	tests := []struct {
		name      string
		fields    *chatWithOtomoUseCaseFields
		args      args
		want      *message.MessageWithOtomo
		wantIsErr bool
	}{
		// TODO: Add a test for success
		{
			name: strings.Join([]string{
				"should return *message.MessageWithOtomo",
				"when give a valid args",
			}, " "),
			fields: func() *chatWithOtomoUseCaseFields {
				var (
					fields            = newChatWithOtomoUseCaseFields(t)
					assertUserMessage = func(msg *message.MessageWithOtomo) {
						assert.Len(t, msg.ID(), 36)
						assert.Equal(t, giveUserID, msg.UserID())
						assert.Equal(t, message.UserRole, msg.Sender())
						assert.Equal(t, message.OtomoRole, msg.Receiver())
						assert.Equal(t, giveText, msg.Text())
						assert.Equal(
							t,
							truncatedSecondNow,
							msg.SentAt().Truncate(time.Second),
						)
					}
				)
				fields.msgRepo.EXPECT().Add(
					giveCtx,
					gomock.Any(),
				).DoAndReturn(
					func(ctx context.Context, msg *message.MessageWithOtomo) error {
						assertUserMessage(msg)
						return nil
					},
				)
				fields.msgMaker.EXPECT().MakeFromMessageWithOtomo(
					giveCtx,
					gomock.Any(),
				).DoAndReturn(
					func(ctx context.Context, msg *message.MessageWithOtomo) (
						*message.MessageWithOtomo,
						error,
					) {
						assertUserMessage(msg)
						return reply, nil
					},
				)
				fields.msgRepo.EXPECT().Add(
					giveCtx,
					reply,
				).Return(nil)
				return fields
			}(),
			args:      args{ctx: giveCtx, userID: giveUserID, text: giveText},
			want:      reply,
			wantIsErr: false,
		},
		// TODO: Add test for occurs error
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			u := &ChatWithOtomoUseCase{
				msgMaker: tt.fields.msgMaker,
				msgRepo:  tt.fields.msgRepo,
			}
			got, err := u.MessageToOtomo(tt.args.ctx, tt.args.userID, tt.args.text)
			if (err != nil) != tt.wantIsErr {
				t.Errorf("ChatWithOtomoUseCase.MessageToOtomo() error = %v, wantIsErr %v", err, tt.wantIsErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
