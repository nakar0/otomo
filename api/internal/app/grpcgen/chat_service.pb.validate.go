// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: chat_service.proto

package grpcgen

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
)

// Validate checks the field values on ChatService_SendMessageRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ChatService_SendMessageRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ChatService_SendMessageRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// ChatService_SendMessageRequestMultiError, or nil if none found.
func (m *ChatService_SendMessageRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ChatService_SendMessageRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetUserId()) < 1 {
		err := ChatService_SendMessageRequestValidationError{
			field:  "UserId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetText()); l < 1 || l > 10000 {
		err := ChatService_SendMessageRequestValidationError{
			field:  "Text",
			reason: "value length must be between 1 and 10000 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetClientId()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ChatService_SendMessageRequestValidationError{
					field:  "ClientId",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ChatService_SendMessageRequestValidationError{
					field:  "ClientId",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetClientId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ChatService_SendMessageRequestValidationError{
				field:  "ClientId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ChatService_SendMessageRequestMultiError(errors)
	}
	return nil
}

// ChatService_SendMessageRequestMultiError is an error wrapping multiple
// validation errors returned by ChatService_SendMessageRequest.ValidateAll()
// if the designated constraints aren't met.
type ChatService_SendMessageRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ChatService_SendMessageRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ChatService_SendMessageRequestMultiError) AllErrors() []error { return m }

// ChatService_SendMessageRequestValidationError is the validation error
// returned by ChatService_SendMessageRequest.Validate if the designated
// constraints aren't met.
type ChatService_SendMessageRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ChatService_SendMessageRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ChatService_SendMessageRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ChatService_SendMessageRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ChatService_SendMessageRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ChatService_SendMessageRequestValidationError) ErrorName() string {
	return "ChatService_SendMessageRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ChatService_SendMessageRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sChatService_SendMessageRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ChatService_SendMessageRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ChatService_SendMessageRequestValidationError{}

// Validate checks the field values on ChatService_SendMessageResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ChatService_SendMessageResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ChatService_SendMessageResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// ChatService_SendMessageResponseMultiError, or nil if none found.
func (m *ChatService_SendMessageResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ChatService_SendMessageResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetMessage()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ChatService_SendMessageResponseValidationError{
					field:  "Message",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ChatService_SendMessageResponseValidationError{
					field:  "Message",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMessage()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ChatService_SendMessageResponseValidationError{
				field:  "Message",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ChatService_SendMessageResponseMultiError(errors)
	}
	return nil
}

// ChatService_SendMessageResponseMultiError is an error wrapping multiple
// validation errors returned by ChatService_SendMessageResponse.ValidateAll()
// if the designated constraints aren't met.
type ChatService_SendMessageResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ChatService_SendMessageResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ChatService_SendMessageResponseMultiError) AllErrors() []error { return m }

// ChatService_SendMessageResponseValidationError is the validation error
// returned by ChatService_SendMessageResponse.Validate if the designated
// constraints aren't met.
type ChatService_SendMessageResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ChatService_SendMessageResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ChatService_SendMessageResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ChatService_SendMessageResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ChatService_SendMessageResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ChatService_SendMessageResponseValidationError) ErrorName() string {
	return "ChatService_SendMessageResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ChatService_SendMessageResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sChatService_SendMessageResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ChatService_SendMessageResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ChatService_SendMessageResponseValidationError{}

// Validate checks the field values on ChatService_ListMessagesRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ChatService_ListMessagesRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ChatService_ListMessagesRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// ChatService_ListMessagesRequestMultiError, or nil if none found.
func (m *ChatService_ListMessagesRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ChatService_ListMessagesRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for PageSize

	// no validation rules for PageStartAfterMessageId

	if utf8.RuneCountInString(m.GetUserId()) < 1 {
		err := ChatService_ListMessagesRequestValidationError{
			field:  "UserId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ChatService_ListMessagesRequestMultiError(errors)
	}
	return nil
}

// ChatService_ListMessagesRequestMultiError is an error wrapping multiple
// validation errors returned by ChatService_ListMessagesRequest.ValidateAll()
// if the designated constraints aren't met.
type ChatService_ListMessagesRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ChatService_ListMessagesRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ChatService_ListMessagesRequestMultiError) AllErrors() []error { return m }

// ChatService_ListMessagesRequestValidationError is the validation error
// returned by ChatService_ListMessagesRequest.Validate if the designated
// constraints aren't met.
type ChatService_ListMessagesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ChatService_ListMessagesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ChatService_ListMessagesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ChatService_ListMessagesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ChatService_ListMessagesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ChatService_ListMessagesRequestValidationError) ErrorName() string {
	return "ChatService_ListMessagesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ChatService_ListMessagesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sChatService_ListMessagesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ChatService_ListMessagesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ChatService_ListMessagesRequestValidationError{}

// Validate checks the field values on ChatService_ListMessagesResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *ChatService_ListMessagesResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ChatService_ListMessagesResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// ChatService_ListMessagesResponseMultiError, or nil if none found.
func (m *ChatService_ListMessagesResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ChatService_ListMessagesResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for PageSize

	// no validation rules for HasMore

	for idx, item := range m.GetMessages() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ChatService_ListMessagesResponseValidationError{
						field:  fmt.Sprintf("Messages[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ChatService_ListMessagesResponseValidationError{
						field:  fmt.Sprintf("Messages[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ChatService_ListMessagesResponseValidationError{
					field:  fmt.Sprintf("Messages[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ChatService_ListMessagesResponseMultiError(errors)
	}
	return nil
}

// ChatService_ListMessagesResponseMultiError is an error wrapping multiple
// validation errors returned by
// ChatService_ListMessagesResponse.ValidateAll() if the designated
// constraints aren't met.
type ChatService_ListMessagesResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ChatService_ListMessagesResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ChatService_ListMessagesResponseMultiError) AllErrors() []error { return m }

// ChatService_ListMessagesResponseValidationError is the validation error
// returned by ChatService_ListMessagesResponse.Validate if the designated
// constraints aren't met.
type ChatService_ListMessagesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ChatService_ListMessagesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ChatService_ListMessagesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ChatService_ListMessagesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ChatService_ListMessagesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ChatService_ListMessagesResponseValidationError) ErrorName() string {
	return "ChatService_ListMessagesResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ChatService_ListMessagesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sChatService_ListMessagesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ChatService_ListMessagesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ChatService_ListMessagesResponseValidationError{}

// Validate checks the field values on ChatService_AskToMessageRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ChatService_AskToMessageRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ChatService_AskToMessageRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// ChatService_AskToMessageRequestMultiError, or nil if none found.
func (m *ChatService_AskToMessageRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ChatService_AskToMessageRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetUserId()) < 1 {
		err := ChatService_AskToMessageRequestValidationError{
			field:  "UserId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ChatService_AskToMessageRequestMultiError(errors)
	}
	return nil
}

// ChatService_AskToMessageRequestMultiError is an error wrapping multiple
// validation errors returned by ChatService_AskToMessageRequest.ValidateAll()
// if the designated constraints aren't met.
type ChatService_AskToMessageRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ChatService_AskToMessageRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ChatService_AskToMessageRequestMultiError) AllErrors() []error { return m }

// ChatService_AskToMessageRequestValidationError is the validation error
// returned by ChatService_AskToMessageRequest.Validate if the designated
// constraints aren't met.
type ChatService_AskToMessageRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ChatService_AskToMessageRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ChatService_AskToMessageRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ChatService_AskToMessageRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ChatService_AskToMessageRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ChatService_AskToMessageRequestValidationError) ErrorName() string {
	return "ChatService_AskToMessageRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ChatService_AskToMessageRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sChatService_AskToMessageRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ChatService_AskToMessageRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ChatService_AskToMessageRequestValidationError{}

// Validate checks the field values on ChatService_AskToMessageResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *ChatService_AskToMessageResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ChatService_AskToMessageResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// ChatService_AskToMessageResponseMultiError, or nil if none found.
func (m *ChatService_AskToMessageResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ChatService_AskToMessageResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetMessage()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ChatService_AskToMessageResponseValidationError{
					field:  "Message",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ChatService_AskToMessageResponseValidationError{
					field:  "Message",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMessage()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ChatService_AskToMessageResponseValidationError{
				field:  "Message",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ChatService_AskToMessageResponseMultiError(errors)
	}
	return nil
}

// ChatService_AskToMessageResponseMultiError is an error wrapping multiple
// validation errors returned by
// ChatService_AskToMessageResponse.ValidateAll() if the designated
// constraints aren't met.
type ChatService_AskToMessageResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ChatService_AskToMessageResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ChatService_AskToMessageResponseMultiError) AllErrors() []error { return m }

// ChatService_AskToMessageResponseValidationError is the validation error
// returned by ChatService_AskToMessageResponse.Validate if the designated
// constraints aren't met.
type ChatService_AskToMessageResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ChatService_AskToMessageResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ChatService_AskToMessageResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ChatService_AskToMessageResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ChatService_AskToMessageResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ChatService_AskToMessageResponseValidationError) ErrorName() string {
	return "ChatService_AskToMessageResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ChatService_AskToMessageResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sChatService_AskToMessageResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ChatService_AskToMessageResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ChatService_AskToMessageResponseValidationError{}

// Validate checks the field values on ChatService_MessagingStreamRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *ChatService_MessagingStreamRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ChatService_MessagingStreamRequest
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// ChatService_MessagingStreamRequestMultiError, or nil if none found.
func (m *ChatService_MessagingStreamRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ChatService_MessagingStreamRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetUserId()) < 1 {
		err := ChatService_MessagingStreamRequestValidationError{
			field:  "UserId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ChatService_MessagingStreamRequestMultiError(errors)
	}
	return nil
}

// ChatService_MessagingStreamRequestMultiError is an error wrapping multiple
// validation errors returned by
// ChatService_MessagingStreamRequest.ValidateAll() if the designated
// constraints aren't met.
type ChatService_MessagingStreamRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ChatService_MessagingStreamRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ChatService_MessagingStreamRequestMultiError) AllErrors() []error { return m }

// ChatService_MessagingStreamRequestValidationError is the validation error
// returned by ChatService_MessagingStreamRequest.Validate if the designated
// constraints aren't met.
type ChatService_MessagingStreamRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ChatService_MessagingStreamRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ChatService_MessagingStreamRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ChatService_MessagingStreamRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ChatService_MessagingStreamRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ChatService_MessagingStreamRequestValidationError) ErrorName() string {
	return "ChatService_MessagingStreamRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ChatService_MessagingStreamRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sChatService_MessagingStreamRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ChatService_MessagingStreamRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ChatService_MessagingStreamRequestValidationError{}

// Validate checks the field values on ChatService_MessagingStreamResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *ChatService_MessagingStreamResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ChatService_MessagingStreamResponse
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// ChatService_MessagingStreamResponseMultiError, or nil if none found.
func (m *ChatService_MessagingStreamResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ChatService_MessagingStreamResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetChunk()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ChatService_MessagingStreamResponseValidationError{
					field:  "Chunk",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ChatService_MessagingStreamResponseValidationError{
					field:  "Chunk",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetChunk()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ChatService_MessagingStreamResponseValidationError{
				field:  "Chunk",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ChatService_MessagingStreamResponseMultiError(errors)
	}
	return nil
}

// ChatService_MessagingStreamResponseMultiError is an error wrapping multiple
// validation errors returned by
// ChatService_MessagingStreamResponse.ValidateAll() if the designated
// constraints aren't met.
type ChatService_MessagingStreamResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ChatService_MessagingStreamResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ChatService_MessagingStreamResponseMultiError) AllErrors() []error { return m }

// ChatService_MessagingStreamResponseValidationError is the validation error
// returned by ChatService_MessagingStreamResponse.Validate if the designated
// constraints aren't met.
type ChatService_MessagingStreamResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ChatService_MessagingStreamResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ChatService_MessagingStreamResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ChatService_MessagingStreamResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ChatService_MessagingStreamResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ChatService_MessagingStreamResponseValidationError) ErrorName() string {
	return "ChatService_MessagingStreamResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ChatService_MessagingStreamResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sChatService_MessagingStreamResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ChatService_MessagingStreamResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ChatService_MessagingStreamResponseValidationError{}
