package service

import (
	"context"
	"otomo/internal/app/model"
	"otomo/internal/pkg/times"
	"otomo/internal/pkg/uuid"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/prompts"
	"github.com/tmc/langchaingo/schema"
)

const (
	otomoCommonPrompt = `The following is a friendly conversation between a user and an AI called Otomo. Otomo is talkative and provides lots of specific details from its context. If Otomo does not know the answer to a question, it truthfully says it does not know.
When Otomo mentions the name of a place, if it knows the longitude and latitude of the place, it says it in the form %[name of place]({ "type": "object", "properties": { "lat_lng": { "type": "object", "properties": { "latitude": { "type": "number" }, "longitude": { "type": "number" } } } } }). () is JSON Schema. The corresponding json is actually input. For example, Tokyo is %[Tokyo]({ "lat_lng": { "latitude": 35.6762, "longitude": 139.6503 } }).`

	respondPrompt = otomoCommonPrompt + `

Current conversation:
{{.history}}

User: {{.input}}
Otomo:`

	messagePrompt = otomoCommonPrompt + `
The conversation currently stops at your message, so please create a message that you think will be of interest to the User. You may refer to the information in the Current Conversation below, or if there is no information in the Current Conversation, compose a message on a topic of interest to the new User. Start the conversation with a greeting.

Current conversation:
{{.history}}

Otomo:`
)

var _ model.Converser = (*ConversationService)(nil)

// TODO: Add tests
type ConversationService struct {
	gpt        *openai.Chat
	msgFactory *model.MessageFactory
}

func NewConversationService(gpt *openai.Chat) *ConversationService {
	return &ConversationService{
		gpt: gpt,
	}
}

func (cs *ConversationService) Respond(
	ctx context.Context,
	msg *model.Message,
	memory *model.Memory,
	listeningFunc model.ListeningFunc,
) (*model.Message, error) {
	gptMsgs, err := prompts.NewHumanMessagePromptTemplate(
		respondPrompt,
		[]string{"history", "input"},
	).FormatMessages(
		map[string]any{"history": memory.Summary, "input": msg.Text},
	)
	if err != nil {
		return nil, err
	}

	return cs.call(ctx, gptMsgs, listeningFunc)
}

func (cs *ConversationService) Message(
	ctx context.Context,
	memory *model.Memory,
	listeningFunc model.ListeningFunc,
) (*model.Message, error) {
	gptMsgs, err := prompts.NewSystemMessagePromptTemplate(
		messagePrompt,
		[]string{"history"},
	).FormatMessages(map[string]any{"history": memory.Summary})
	if err != nil {
		return nil, err
	}

	return cs.call(ctx, gptMsgs, listeningFunc)
}

func (cs *ConversationService) call(
	ctx context.Context,
	msgs []schema.ChatMessage,
	listeningFunc model.ListeningFunc,
) (*model.Message, error) {
	var (
		msgID = model.MessageID(uuid.NewString())
		role  = model.OtomoRole
	)

	completion, err := cs.gpt.Call(
		ctx,
		msgs,
		llms.WithStreamingFunc(
			cs.makeStreamingFunc(msgID, role, listeningFunc)),
	)
	if err != nil {
		return nil, err
	}

	if err := cs.finishListening(ctx, msgID, role, listeningFunc); err != nil {
		return nil, err
	}

	return cs.msgFactory.Restore(
		msgID,
		completion.GetContent(),
		role,
		times.C.Now(),
		nil,
	), nil
}

func (cs *ConversationService) makeStreamingFunc(
	newMsgID model.MessageID,
	newMsgRole model.Role,
	listeningFunc model.ListeningFunc,
) func(context.Context, []byte) error {
	if listeningFunc == nil {
		return func(_ context.Context, _ []byte) error {
			return nil
		}
	} else {
		return func(ctx context.Context, chunk []byte) error {
			text := string(chunk)
			msgChunk, err := model.NewMessageChunk(
				newMsgID,
				text,
				newMsgRole,
				times.C.Now(),
				nil,
				false,
			)
			if err != nil {
				return err
			}

			return listeningFunc(ctx, msgChunk)
		}
	}
}

func (cs *ConversationService) finishListening(
	ctx context.Context,
	newMsgID model.MessageID,
	newMsgRole model.Role,
	listeningFunc model.ListeningFunc,
) error {
	if listeningFunc == nil {
		return nil
	} else {
		msgChunk, err := model.NewMessageChunk(
			newMsgID,
			"",
			newMsgRole,
			times.C.Now(),
			nil,
			true,
		)
		if err != nil {
			return err
		}

		return listeningFunc(ctx, msgChunk)
	}
}
