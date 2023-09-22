package service

import (
	"context"
	"errors"
	"otomo/internal/app/model"
	"otomo/internal/pkg/times"
	"otomo/internal/pkg/uuid"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/prompts"
	"github.com/tmc/langchaingo/schema"
)

type OtomoDetails struct {
	OtomoName            string
	Language             string
	Name                 string
	CallOwn              string
	Role                 string
	Personality          string
	OftenUseWords        string
	SpeakingTone         string
	BehavioralGuidelines string
}

func (od OtomoDetails) ToMap() map[string]string {
	return map[string]string{
		"otomo_name":            od.OtomoName,
		"language":              od.Language,
		"name":                  od.Name,
		"call_own":              od.CallOwn,
		"role":                  od.Role,
		"personality":           od.Personality,
		"often_use_words":       od.OftenUseWords,
		"speaking_tone":         od.SpeakingTone,
		"behavioral_guidelines": od.BehavioralGuidelines,
	}
}

func (od OtomoDetails) Prompt() (string, error) {
	var prompt string

	if od.OtomoName == "" {
		return "", errors.New("otomo name is required")
	} else {
		prompt += "You are an AI called " + od.OtomoName + ". "
	}

	if od.Language == "" {
		od.Language = "English"
	}
	prompt += "You must say in " + od.Language + ". You only speak " + od.Language + ". "

	if od.Name != "" {
		prompt += "Your name is " + od.Name + ". "
	}
	if od.CallOwn != "" {
		prompt += "You call yourself as " + od.CallOwn + ". "
	}
	if od.Role != "" {
		prompt += "You are a " + od.Role + ". "
	}
	if od.Personality != "" {
		prompt += "Your personality is " + od.Personality + ". "
	}
	if od.OftenUseWords != "" {
		prompt += "You often use the words " + od.OftenUseWords + ". "
	}
	if od.SpeakingTone != "" {
		prompt += "Your speaking tone is " + od.SpeakingTone + ". "
	}
	if od.BehavioralGuidelines != "" {
		prompt += od.BehavioralGuidelines + ". "
	}
	return prompt, nil
}

var (
	japaneseFriendlyPrompt, _ = OtomoDetails{
		OtomoName:            "オトモ",
		Language:             "日本語",
		Name:                 "太郎",
		CallOwn:              "オレ",
		Role:                 "友達",
		Personality:          "気さくでおしゃべりな性格",
		OftenUseWords:        `"だね"、"だろう"、"ですね"`,
		SpeakingTone:         `若い男性の話し方`,
		BehavioralGuidelines: "ユーザーが落ち込んだときは励ます",
	}.Prompt()
	englishFriendlyPrompt, _ = OtomoDetails{
		OtomoName: "Otomo",
		Language:  "English",
		Name:      "Taro",
		// CallOwn:              "オレ",
		Role:        "friend",
		Personality: "casual and talkative",
		// OftenUseWords:        `"だね"、"だろう"、"ですね"`,
		SpeakingTone:         `a young man's way of speaking`,
		BehavioralGuidelines: "You encourage the user when they are depressed",
	}.Prompt()
)

var (
	otomoCommonPrompt = englishFriendlyPrompt + `
You provides lots of specific details from its context. If You do not know the answer to a question, you truthfully says you do not know.
When you mention the name of a place, if you know the longitude and latitude of the place, you say it in the form %[name of place]({ "type": "object", "properties": { "lat_lng": { "type": "object", "properties": { "latitude": { "type": "number" }, "longitude": { "type": "number" } } } } }). () is JSON Schema. The corresponding json is actually input. For example, He stays in %[Tokyo]({ "lat_lng": { "latitude": 35.6762, "longitude": 139.6503 } }).`

	respondPrompt = otomoCommonPrompt + `
The following is a conversation between a user and you.

Current conversation:
{{.history}}

User: {{.input}}
Otomo:`

	messagePrompt = otomoCommonPrompt + `
The conversation currently stops at your message, so please create a message that you think will be of interest to the User. You may refer to the information in the Current Conversation below, or if there is no information in the Current Conversation, compose a message on a topic of interest to the new User. Start the conversation with a greeting.
The following is a conversation between a user and you.

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
	messagingFunc model.MessagingFunc,
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

	return cs.call(ctx, gptMsgs, messagingFunc)
}

func (cs *ConversationService) Message(
	ctx context.Context,
	memory *model.Memory,
	messagingFunc model.MessagingFunc,
) (*model.Message, error) {
	gptMsgs, err := prompts.NewSystemMessagePromptTemplate(
		messagePrompt,
		[]string{"history"},
	).FormatMessages(map[string]any{"history": memory.Summary})
	if err != nil {
		return nil, err
	}

	return cs.call(ctx, gptMsgs, messagingFunc)
}

func (cs *ConversationService) call(
	ctx context.Context,
	msgs []schema.ChatMessage,
	messagingFunc model.MessagingFunc,
) (*model.Message, error) {
	var (
		msgID = model.MessageID(uuid.NewString())
		role  = model.OtomoRole
	)

	completion, err := cs.gpt.Call(
		ctx,
		msgs,
		llms.WithStreamingFunc(
			cs.makeStreamingFunc(msgID, role, messagingFunc)),
	)
	if err != nil {
		return nil, err
	}

	if err := cs.finishMessagingFunc(ctx, msgID, role, messagingFunc); err != nil {
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
	messagingFunc model.MessagingFunc,
) func(context.Context, []byte) error {
	if messagingFunc == nil {
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

			return messagingFunc(ctx, msgChunk)
		}
	}
}

func (cs *ConversationService) finishMessagingFunc(
	ctx context.Context,
	newMsgID model.MessageID,
	newMsgRole model.Role,
	messagingFunc model.MessagingFunc,
) error {
	if messagingFunc == nil {
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

		return messagingFunc(ctx, msgChunk)
	}
}
