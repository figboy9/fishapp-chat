// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: chat.proto

package chat_grpc

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

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _chat_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Room with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Room) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for PostId

	if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RoomValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RoomValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// RoomValidationError is the validation error returned by Room.Validate if the
// designated constraints aren't met.
type RoomValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RoomValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RoomValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RoomValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RoomValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RoomValidationError) ErrorName() string { return "RoomValidationError" }

// Error satisfies the builtin error interface
func (e RoomValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRoom.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RoomValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RoomValidationError{}

// Validate checks the field values on Member with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Member) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for RoomId

	// no validation rules for UserId

	if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MemberValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MemberValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// MemberValidationError is the validation error returned by Member.Validate if
// the designated constraints aren't met.
type MemberValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MemberValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MemberValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MemberValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MemberValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MemberValidationError) ErrorName() string { return "MemberValidationError" }

// Error satisfies the builtin error interface
func (e MemberValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMember.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MemberValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MemberValidationError{}

// Validate checks the field values on Message with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Message) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Body

	// no validation rules for RoomId

	// no validation rules for UserId

	if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MessageValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MessageValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

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

// Validate checks the field values on CreateRoomReq with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *CreateRoomReq) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetPostId() < 1 {
		return CreateRoomReqValidationError{
			field:  "PostId",
			reason: "value must be greater than or equal to 1",
		}
	}

	if m.GetUserId() < 1 {
		return CreateRoomReqValidationError{
			field:  "UserId",
			reason: "value must be greater than or equal to 1",
		}
	}

	return nil
}

// CreateRoomReqValidationError is the validation error returned by
// CreateRoomReq.Validate if the designated constraints aren't met.
type CreateRoomReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateRoomReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateRoomReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateRoomReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateRoomReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateRoomReqValidationError) ErrorName() string { return "CreateRoomReqValidationError" }

// Error satisfies the builtin error interface
func (e CreateRoomReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateRoomReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateRoomReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateRoomReqValidationError{}

// Validate checks the field values on GetRoomReq with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *GetRoomReq) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetPostId() < 1 {
		return GetRoomReqValidationError{
			field:  "PostId",
			reason: "value must be greater than or equal to 1",
		}
	}

	if m.GetUserId() < 1 {
		return GetRoomReqValidationError{
			field:  "UserId",
			reason: "value must be greater than or equal to 1",
		}
	}

	return nil
}

// GetRoomReqValidationError is the validation error returned by
// GetRoomReq.Validate if the designated constraints aren't met.
type GetRoomReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetRoomReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetRoomReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetRoomReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetRoomReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetRoomReqValidationError) ErrorName() string { return "GetRoomReqValidationError" }

// Error satisfies the builtin error interface
func (e GetRoomReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetRoomReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetRoomReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetRoomReqValidationError{}

// Validate checks the field values on ListMembersReq with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ListMembersReq) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetRoomId() < 1 {
		return ListMembersReqValidationError{
			field:  "RoomId",
			reason: "value must be greater than or equal to 1",
		}
	}

	if m.GetUserId() < 1 {
		return ListMembersReqValidationError{
			field:  "UserId",
			reason: "value must be greater than or equal to 1",
		}
	}

	return nil
}

// ListMembersReqValidationError is the validation error returned by
// ListMembersReq.Validate if the designated constraints aren't met.
type ListMembersReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListMembersReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListMembersReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListMembersReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListMembersReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListMembersReqValidationError) ErrorName() string { return "ListMembersReqValidationError" }

// Error satisfies the builtin error interface
func (e ListMembersReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListMembersReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListMembersReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListMembersReqValidationError{}

// Validate checks the field values on ListMembersRes with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ListMembersRes) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetMembers() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListMembersResValidationError{
					field:  fmt.Sprintf("Members[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListMembersResValidationError is the validation error returned by
// ListMembersRes.Validate if the designated constraints aren't met.
type ListMembersResValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListMembersResValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListMembersResValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListMembersResValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListMembersResValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListMembersResValidationError) ErrorName() string { return "ListMembersResValidationError" }

// Error satisfies the builtin error interface
func (e ListMembersResValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListMembersRes.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListMembersResValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListMembersResValidationError{}

// Validate checks the field values on CreateMemberReq with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *CreateMemberReq) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetRoomId() < 1 {
		return CreateMemberReqValidationError{
			field:  "RoomId",
			reason: "value must be greater than or equal to 1",
		}
	}

	if m.GetUserId() < 1 {
		return CreateMemberReqValidationError{
			field:  "UserId",
			reason: "value must be greater than or equal to 1",
		}
	}

	return nil
}

// CreateMemberReqValidationError is the validation error returned by
// CreateMemberReq.Validate if the designated constraints aren't met.
type CreateMemberReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateMemberReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateMemberReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateMemberReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateMemberReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateMemberReqValidationError) ErrorName() string { return "CreateMemberReqValidationError" }

// Error satisfies the builtin error interface
func (e CreateMemberReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateMemberReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateMemberReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateMemberReqValidationError{}

// Validate checks the field values on DeleteMemberReq with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *DeleteMemberReq) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetRoomId() < 1 {
		return DeleteMemberReqValidationError{
			field:  "RoomId",
			reason: "value must be greater than or equal to 1",
		}
	}

	if m.GetUserId() < 1 {
		return DeleteMemberReqValidationError{
			field:  "UserId",
			reason: "value must be greater than or equal to 1",
		}
	}

	return nil
}

// DeleteMemberReqValidationError is the validation error returned by
// DeleteMemberReq.Validate if the designated constraints aren't met.
type DeleteMemberReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteMemberReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteMemberReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteMemberReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteMemberReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteMemberReqValidationError) ErrorName() string { return "DeleteMemberReqValidationError" }

// Error satisfies the builtin error interface
func (e DeleteMemberReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteMemberReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteMemberReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteMemberReqValidationError{}

// Validate checks the field values on CreateMessageReq with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *CreateMessageReq) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Body

	if m.GetRoomId() < 1 {
		return CreateMessageReqValidationError{
			field:  "RoomId",
			reason: "value must be greater than or equal to 1",
		}
	}

	if m.GetUserId() < 1 {
		return CreateMessageReqValidationError{
			field:  "UserId",
			reason: "value must be greater than or equal to 1",
		}
	}

	return nil
}

// CreateMessageReqValidationError is the validation error returned by
// CreateMessageReq.Validate if the designated constraints aren't met.
type CreateMessageReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateMessageReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateMessageReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateMessageReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateMessageReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateMessageReqValidationError) ErrorName() string { return "CreateMessageReqValidationError" }

// Error satisfies the builtin error interface
func (e CreateMessageReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateMessageReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateMessageReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateMessageReqValidationError{}

// Validate checks the field values on ListMessagesReq with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ListMessagesReq) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetRoomId() < 1 {
		return ListMessagesReqValidationError{
			field:  "RoomId",
			reason: "value must be greater than or equal to 1",
		}
	}

	if m.GetUserId() < 1 {
		return ListMessagesReqValidationError{
			field:  "UserId",
			reason: "value must be greater than or equal to 1",
		}
	}

	return nil
}

// ListMessagesReqValidationError is the validation error returned by
// ListMessagesReq.Validate if the designated constraints aren't met.
type ListMessagesReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListMessagesReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListMessagesReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListMessagesReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListMessagesReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListMessagesReqValidationError) ErrorName() string { return "ListMessagesReqValidationError" }

// Error satisfies the builtin error interface
func (e ListMessagesReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListMessagesReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListMessagesReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListMessagesReqValidationError{}

// Validate checks the field values on ListMessagesRes with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ListMessagesRes) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetMessages() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListMessagesResValidationError{
					field:  fmt.Sprintf("Messages[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListMessagesResValidationError is the validation error returned by
// ListMessagesRes.Validate if the designated constraints aren't met.
type ListMessagesResValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListMessagesResValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListMessagesResValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListMessagesResValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListMessagesResValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListMessagesResValidationError) ErrorName() string { return "ListMessagesResValidationError" }

// Error satisfies the builtin error interface
func (e ListMessagesResValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListMessagesRes.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListMessagesResValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListMessagesResValidationError{}

// Validate checks the field values on StreamMessageReq with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *StreamMessageReq) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetRoomId() < 1 {
		return StreamMessageReqValidationError{
			field:  "RoomId",
			reason: "value must be greater than or equal to 1",
		}
	}

	if m.GetUserId() < 1 {
		return StreamMessageReqValidationError{
			field:  "UserId",
			reason: "value must be greater than or equal to 1",
		}
	}

	return nil
}

// StreamMessageReqValidationError is the validation error returned by
// StreamMessageReq.Validate if the designated constraints aren't met.
type StreamMessageReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StreamMessageReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StreamMessageReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StreamMessageReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StreamMessageReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StreamMessageReqValidationError) ErrorName() string { return "StreamMessageReqValidationError" }

// Error satisfies the builtin error interface
func (e StreamMessageReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStreamMessageReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StreamMessageReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StreamMessageReqValidationError{}
