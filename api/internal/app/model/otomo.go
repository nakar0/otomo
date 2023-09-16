package model

import (
	"context"
	"otomo/internal/pkg/errs"
)

// TODO: Add tests
type Converser interface {
	Respond(
		ctx context.Context,
		msg *Message,
		memory *Memory,
		listeningFunc func(ctx context.Context, chunk []byte) error,
	) (*Message, error)
	Message(
		ctx context.Context,
		memory *Memory,
		listeningFunc func(ctx context.Context, chunk []byte) error,
	) (*Message, error)
}
type Summarizer interface {
	Summarize(
		ctx context.Context,
		newMsgs []*Message,
		currentSummary string,
	) (string, error)
}

type Otomo struct {
	UserID UserID `firestore:"user_id"`
	Memory Memory `firestore:"memory"`

	converser  Converser
	summarizer Summarizer
}

func RestoreOtomo(
	userID UserID,
	memory Memory,
) (*Otomo, error) {
	return &Otomo{
		UserID: userID,
		Memory: memory,
	}, nil
}

func (o *Otomo) WithConverser(converser Converser) *Otomo {
	newOtomo := *o
	newOtomo.converser = converser
	return &newOtomo
}

func (o *Otomo) WithSummarizer(summarizer Summarizer) *Otomo {
	newOtomo := *o
	newOtomo.summarizer = summarizer
	return &newOtomo
}

func (o *Otomo) Respond(
	ctx context.Context,
	msg *Message,
) (*Otomo, *Message, error) {
	if err := o.validateAbilities(); err != nil {
		return nil, nil, err
	}

	// TODO: listeningFunc
	respond, err := o.converser.Respond(ctx, msg, &o.Memory, nil)
	if err != nil {
		return nil, nil, err
	}

	summary, err := o.summarizer.Summarize(
		ctx, []*Message{msg, respond}, o.Memory.Summary)
	if err != nil {
		return nil, nil, err
	}

	newOtomo := *o
	newOtomo.Memory.Summary = summary
	return &newOtomo, respond, nil
}

func (o *Otomo) Message(
	ctx context.Context,
) (*Otomo, *Message, error) {
	if err := o.validateAbilities(); err != nil {
		return nil, nil, err
	}

	// TODO: listeningFunc
	newMsg, err := o.converser.Message(ctx, &o.Memory, nil)
	if err != nil {
		return nil, nil, err
	}

	summary, err := o.summarizer.Summarize(
		ctx, []*Message{newMsg}, o.Memory.Summary)
	if err != nil {
		return nil, nil, err
	}

	newOtomo := *o
	newOtomo.Memory.Summary = summary

	return &newOtomo, newMsg, nil
}

func (o *Otomo) validateAbilities() error {
	if o.converser == nil {
		return &errs.Error{
			Message: "converser is not set",
			Cause:   errs.CauseUnavailable,
			Domain:  errs.DomainOtomo,
			Field:   errs.FieldAbility,
		}
	}
	if o.summarizer == nil {
		return &errs.Error{
			Message: "summarizer is not set",
			Cause:   errs.CauseUnavailable,
			Domain:  errs.DomainOtomo,
			Field:   errs.FieldAbility,
		}

	}
	return nil
}