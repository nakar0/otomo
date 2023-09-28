// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: message.proto

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

// Validate checks the field values on Message with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Message) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Message with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in MessageMultiError, or nil if none found.
func (m *Message) ValidateAll() error {
	return m.validate(true)
}

func (m *Message) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Text

	// no validation rules for Role

	if all {
		switch v := interface{}(m.GetSentAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, MessageValidationError{
					field:  "SentAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, MessageValidationError{
					field:  "SentAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSentAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MessageValidationError{
				field:  "SentAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetClientId()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, MessageValidationError{
					field:  "ClientId",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, MessageValidationError{
					field:  "ClientId",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetClientId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MessageValidationError{
				field:  "ClientId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetLocationAnalysis()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, MessageValidationError{
					field:  "LocationAnalysis",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, MessageValidationError{
					field:  "LocationAnalysis",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetLocationAnalysis()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MessageValidationError{
				field:  "LocationAnalysis",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return MessageMultiError(errors)
	}
	return nil
}

// MessageMultiError is an error wrapping multiple validation errors returned
// by Message.ValidateAll() if the designated constraints aren't met.
type MessageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MessageMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MessageMultiError) AllErrors() []error { return m }

// MessageValidationError is the validation error returned by Message.Validate
// if the designated constraints aren't met.
type MessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MessageValidationError) ErrorName() string { return "MessageValidationError" }

// Error satisfies the builtin error interface
func (e MessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MessageValidationError{}

// Validate checks the field values on MessageChunk with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *MessageChunk) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MessageChunk with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in MessageChunkMultiError, or
// nil if none found.
func (m *MessageChunk) ValidateAll() error {
	return m.validate(true)
}

func (m *MessageChunk) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for MessageId

	// no validation rules for Text

	// no validation rules for Role

	if all {
		switch v := interface{}(m.GetSentAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, MessageChunkValidationError{
					field:  "SentAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, MessageChunkValidationError{
					field:  "SentAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSentAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MessageChunkValidationError{
				field:  "SentAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetClientId()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, MessageChunkValidationError{
					field:  "ClientId",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, MessageChunkValidationError{
					field:  "ClientId",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetClientId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MessageChunkValidationError{
				field:  "ClientId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for IsLast

	if len(errors) > 0 {
		return MessageChunkMultiError(errors)
	}
	return nil
}

// MessageChunkMultiError is an error wrapping multiple validation errors
// returned by MessageChunk.ValidateAll() if the designated constraints aren't met.
type MessageChunkMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MessageChunkMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MessageChunkMultiError) AllErrors() []error { return m }

// MessageChunkValidationError is the validation error returned by
// MessageChunk.Validate if the designated constraints aren't met.
type MessageChunkValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MessageChunkValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MessageChunkValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MessageChunkValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MessageChunkValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MessageChunkValidationError) ErrorName() string { return "MessageChunkValidationError" }

// Error satisfies the builtin error interface
func (e MessageChunkValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMessageChunk.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MessageChunkValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MessageChunkValidationError{}

// Validate checks the field values on LocationAnalysis with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *LocationAnalysis) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LocationAnalysis with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// LocationAnalysisMultiError, or nil if none found.
func (m *LocationAnalysis) ValidateAll() error {
	return m.validate(true)
}

func (m *LocationAnalysis) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetLocations() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, LocationAnalysisValidationError{
						field:  fmt.Sprintf("Locations[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, LocationAnalysisValidationError{
						field:  fmt.Sprintf("Locations[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LocationAnalysisValidationError{
					field:  fmt.Sprintf("Locations[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if all {
		switch v := interface{}(m.GetAnalyzedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, LocationAnalysisValidationError{
					field:  "AnalyzedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, LocationAnalysisValidationError{
					field:  "AnalyzedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAnalyzedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LocationAnalysisValidationError{
				field:  "AnalyzedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetError()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, LocationAnalysisValidationError{
					field:  "Error",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, LocationAnalysisValidationError{
					field:  "Error",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetError()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LocationAnalysisValidationError{
				field:  "Error",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return LocationAnalysisMultiError(errors)
	}
	return nil
}

// LocationAnalysisMultiError is an error wrapping multiple validation errors
// returned by LocationAnalysis.ValidateAll() if the designated constraints
// aren't met.
type LocationAnalysisMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LocationAnalysisMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LocationAnalysisMultiError) AllErrors() []error { return m }

// LocationAnalysisValidationError is the validation error returned by
// LocationAnalysis.Validate if the designated constraints aren't met.
type LocationAnalysisValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LocationAnalysisValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LocationAnalysisValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LocationAnalysisValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LocationAnalysisValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LocationAnalysisValidationError) ErrorName() string { return "LocationAnalysisValidationError" }

// Error satisfies the builtin error interface
func (e LocationAnalysisValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLocationAnalysis.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LocationAnalysisValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LocationAnalysisValidationError{}

// Validate checks the field values on AnalyzedLocation with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *AnalyzedLocation) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AnalyzedLocation with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AnalyzedLocationMultiError, or nil if none found.
func (m *AnalyzedLocation) ValidateAll() error {
	return m.validate(true)
}

func (m *AnalyzedLocation) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Text

	if all {
		switch v := interface{}(m.GetLocation()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AnalyzedLocationValidationError{
					field:  "Location",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AnalyzedLocationValidationError{
					field:  "Location",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetLocation()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AnalyzedLocationValidationError{
				field:  "Location",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return AnalyzedLocationMultiError(errors)
	}
	return nil
}

// AnalyzedLocationMultiError is an error wrapping multiple validation errors
// returned by AnalyzedLocation.ValidateAll() if the designated constraints
// aren't met.
type AnalyzedLocationMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AnalyzedLocationMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AnalyzedLocationMultiError) AllErrors() []error { return m }

// AnalyzedLocationValidationError is the validation error returned by
// AnalyzedLocation.Validate if the designated constraints aren't met.
type AnalyzedLocationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AnalyzedLocationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AnalyzedLocationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AnalyzedLocationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AnalyzedLocationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AnalyzedLocationValidationError) ErrorName() string { return "AnalyzedLocationValidationError" }

// Error satisfies the builtin error interface
func (e AnalyzedLocationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAnalyzedLocation.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AnalyzedLocationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AnalyzedLocationValidationError{}

// Validate checks the field values on MonthlySentCount with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *MonthlySentCount) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MonthlySentCount with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// MonthlySentCountMultiError, or nil if none found.
func (m *MonthlySentCount) ValidateAll() error {
	return m.validate(true)
}

func (m *MonthlySentCount) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetYearMonth()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, MonthlySentCountValidationError{
					field:  "YearMonth",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, MonthlySentCountValidationError{
					field:  "YearMonth",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetYearMonth()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MonthlySentCountValidationError{
				field:  "YearMonth",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Count

	for idx, item := range m.GetDailySentCount() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, MonthlySentCountValidationError{
						field:  fmt.Sprintf("DailySentCount[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, MonthlySentCountValidationError{
						field:  fmt.Sprintf("DailySentCount[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MonthlySentCountValidationError{
					field:  fmt.Sprintf("DailySentCount[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return MonthlySentCountMultiError(errors)
	}
	return nil
}

// MonthlySentCountMultiError is an error wrapping multiple validation errors
// returned by MonthlySentCount.ValidateAll() if the designated constraints
// aren't met.
type MonthlySentCountMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MonthlySentCountMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MonthlySentCountMultiError) AllErrors() []error { return m }

// MonthlySentCountValidationError is the validation error returned by
// MonthlySentCount.Validate if the designated constraints aren't met.
type MonthlySentCountValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MonthlySentCountValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MonthlySentCountValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MonthlySentCountValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MonthlySentCountValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MonthlySentCountValidationError) ErrorName() string { return "MonthlySentCountValidationError" }

// Error satisfies the builtin error interface
func (e MonthlySentCountValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMonthlySentCount.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MonthlySentCountValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MonthlySentCountValidationError{}

// Validate checks the field values on DailySentCount with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DailySentCount) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DailySentCount with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DailySentCountMultiError,
// or nil if none found.
func (m *DailySentCount) ValidateAll() error {
	return m.validate(true)
}

func (m *DailySentCount) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetDate()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DailySentCountValidationError{
					field:  "Date",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DailySentCountValidationError{
					field:  "Date",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DailySentCountValidationError{
				field:  "Date",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Count

	if len(errors) > 0 {
		return DailySentCountMultiError(errors)
	}
	return nil
}

// DailySentCountMultiError is an error wrapping multiple validation errors
// returned by DailySentCount.ValidateAll() if the designated constraints
// aren't met.
type DailySentCountMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DailySentCountMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DailySentCountMultiError) AllErrors() []error { return m }

// DailySentCountValidationError is the validation error returned by
// DailySentCount.Validate if the designated constraints aren't met.
type DailySentCountValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DailySentCountValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DailySentCountValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DailySentCountValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DailySentCountValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DailySentCountValidationError) ErrorName() string { return "DailySentCountValidationError" }

// Error satisfies the builtin error interface
func (e DailySentCountValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDailySentCount.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DailySentCountValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DailySentCountValidationError{}

// Validate checks the field values on RemainingSendCount with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RemainingSendCount) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RemainingSendCount with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RemainingSendCountMultiError, or nil if none found.
func (m *RemainingSendCount) ValidateAll() error {
	return m.validate(true)
}

func (m *RemainingSendCount) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetMonthly()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RemainingSendCountValidationError{
					field:  "Monthly",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RemainingSendCountValidationError{
					field:  "Monthly",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMonthly()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RemainingSendCountValidationError{
				field:  "Monthly",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetDaily()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RemainingSendCountValidationError{
					field:  "Daily",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RemainingSendCountValidationError{
					field:  "Daily",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDaily()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RemainingSendCountValidationError{
				field:  "Daily",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return RemainingSendCountMultiError(errors)
	}
	return nil
}

// RemainingSendCountMultiError is an error wrapping multiple validation errors
// returned by RemainingSendCount.ValidateAll() if the designated constraints
// aren't met.
type RemainingSendCountMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RemainingSendCountMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RemainingSendCountMultiError) AllErrors() []error { return m }

// RemainingSendCountValidationError is the validation error returned by
// RemainingSendCount.Validate if the designated constraints aren't met.
type RemainingSendCountValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemainingSendCountValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemainingSendCountValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemainingSendCountValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemainingSendCountValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemainingSendCountValidationError) ErrorName() string {
	return "RemainingSendCountValidationError"
}

// Error satisfies the builtin error interface
func (e RemainingSendCountValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemainingSendCount.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemainingSendCountValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemainingSendCountValidationError{}

// Validate checks the field values on RemainingMonthlySendCount with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RemainingMonthlySendCount) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RemainingMonthlySendCount with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RemainingMonthlySendCountMultiError, or nil if none found.
func (m *RemainingMonthlySendCount) ValidateAll() error {
	return m.validate(true)
}

func (m *RemainingMonthlySendCount) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetYearMonth()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RemainingMonthlySendCountValidationError{
					field:  "YearMonth",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RemainingMonthlySendCountValidationError{
					field:  "YearMonth",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetYearMonth()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RemainingMonthlySendCountValidationError{
				field:  "YearMonth",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Count

	if len(errors) > 0 {
		return RemainingMonthlySendCountMultiError(errors)
	}
	return nil
}

// RemainingMonthlySendCountMultiError is an error wrapping multiple validation
// errors returned by RemainingMonthlySendCount.ValidateAll() if the
// designated constraints aren't met.
type RemainingMonthlySendCountMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RemainingMonthlySendCountMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RemainingMonthlySendCountMultiError) AllErrors() []error { return m }

// RemainingMonthlySendCountValidationError is the validation error returned by
// RemainingMonthlySendCount.Validate if the designated constraints aren't met.
type RemainingMonthlySendCountValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemainingMonthlySendCountValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemainingMonthlySendCountValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemainingMonthlySendCountValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemainingMonthlySendCountValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemainingMonthlySendCountValidationError) ErrorName() string {
	return "RemainingMonthlySendCountValidationError"
}

// Error satisfies the builtin error interface
func (e RemainingMonthlySendCountValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemainingMonthlySendCount.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemainingMonthlySendCountValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemainingMonthlySendCountValidationError{}

// Validate checks the field values on RemainingDailySendCount with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RemainingDailySendCount) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RemainingDailySendCount with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RemainingDailySendCountMultiError, or nil if none found.
func (m *RemainingDailySendCount) ValidateAll() error {
	return m.validate(true)
}

func (m *RemainingDailySendCount) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetDate()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RemainingDailySendCountValidationError{
					field:  "Date",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RemainingDailySendCountValidationError{
					field:  "Date",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RemainingDailySendCountValidationError{
				field:  "Date",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Count

	if len(errors) > 0 {
		return RemainingDailySendCountMultiError(errors)
	}
	return nil
}

// RemainingDailySendCountMultiError is an error wrapping multiple validation
// errors returned by RemainingDailySendCount.ValidateAll() if the designated
// constraints aren't met.
type RemainingDailySendCountMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RemainingDailySendCountMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RemainingDailySendCountMultiError) AllErrors() []error { return m }

// RemainingDailySendCountValidationError is the validation error returned by
// RemainingDailySendCount.Validate if the designated constraints aren't met.
type RemainingDailySendCountValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemainingDailySendCountValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemainingDailySendCountValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemainingDailySendCountValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemainingDailySendCountValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemainingDailySendCountValidationError) ErrorName() string {
	return "RemainingDailySendCountValidationError"
}

// Error satisfies the builtin error interface
func (e RemainingDailySendCountValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemainingDailySendCount.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemainingDailySendCountValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemainingDailySendCountValidationError{}
