// Code generated by gotdgen, DO NOT EDIT.

package tdapi

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// SendMessageRequest represents TL type `sendMessage#171a006f`.
type SendMessageRequest struct {
	// Target chat
	ChatID int64
	// If not 0, a message thread identifier in which the message will be sent
	MessageThreadID int64
	// Identifier of the replied message or story; pass null if none
	ReplyTo MessageReplyToClass
	// Options to be used to send the message; pass null to use default options
	Options MessageSendOptions
	// Markup for replying to the message; pass null if none; for bots only
	ReplyMarkup ReplyMarkupClass
	// The content of the message to be sent
	InputMessageContent InputMessageContentClass
}

// SendMessageRequestTypeID is TL type id of SendMessageRequest.
const SendMessageRequestTypeID = 0x171a006f

// Ensuring interfaces in compile-time for SendMessageRequest.
var (
	_ bin.Encoder     = &SendMessageRequest{}
	_ bin.Decoder     = &SendMessageRequest{}
	_ bin.BareEncoder = &SendMessageRequest{}
	_ bin.BareDecoder = &SendMessageRequest{}
)

func (s *SendMessageRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.ChatID == 0) {
		return false
	}
	if !(s.MessageThreadID == 0) {
		return false
	}
	if !(s.ReplyTo == nil) {
		return false
	}
	if !(s.Options.Zero()) {
		return false
	}
	if !(s.ReplyMarkup == nil) {
		return false
	}
	if !(s.InputMessageContent == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SendMessageRequest) String() string {
	if s == nil {
		return "SendMessageRequest(nil)"
	}
	type Alias SendMessageRequest
	return fmt.Sprintf("SendMessageRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SendMessageRequest) TypeID() uint32 {
	return SendMessageRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SendMessageRequest) TypeName() string {
	return "sendMessage"
}

// TypeInfo returns info about TL type.
func (s *SendMessageRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "sendMessage",
		ID:   SendMessageRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "MessageThreadID",
			SchemaName: "message_thread_id",
		},
		{
			Name:       "ReplyTo",
			SchemaName: "reply_to",
		},
		{
			Name:       "Options",
			SchemaName: "options",
		},
		{
			Name:       "ReplyMarkup",
			SchemaName: "reply_markup",
		},
		{
			Name:       "InputMessageContent",
			SchemaName: "input_message_content",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SendMessageRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sendMessage#171a006f as nil")
	}
	b.PutID(SendMessageRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SendMessageRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sendMessage#171a006f as nil")
	}
	b.PutInt53(s.ChatID)
	b.PutInt53(s.MessageThreadID)
	if s.ReplyTo == nil {
		return fmt.Errorf("unable to encode sendMessage#171a006f: field reply_to is nil")
	}
	if err := s.ReplyTo.Encode(b); err != nil {
		return fmt.Errorf("unable to encode sendMessage#171a006f: field reply_to: %w", err)
	}
	if err := s.Options.Encode(b); err != nil {
		return fmt.Errorf("unable to encode sendMessage#171a006f: field options: %w", err)
	}
	if s.ReplyMarkup == nil {
		return fmt.Errorf("unable to encode sendMessage#171a006f: field reply_markup is nil")
	}
	if err := s.ReplyMarkup.Encode(b); err != nil {
		return fmt.Errorf("unable to encode sendMessage#171a006f: field reply_markup: %w", err)
	}
	if s.InputMessageContent == nil {
		return fmt.Errorf("unable to encode sendMessage#171a006f: field input_message_content is nil")
	}
	if err := s.InputMessageContent.Encode(b); err != nil {
		return fmt.Errorf("unable to encode sendMessage#171a006f: field input_message_content: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *SendMessageRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sendMessage#171a006f to nil")
	}
	if err := b.ConsumeID(SendMessageRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode sendMessage#171a006f: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SendMessageRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sendMessage#171a006f to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode sendMessage#171a006f: field chat_id: %w", err)
		}
		s.ChatID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode sendMessage#171a006f: field message_thread_id: %w", err)
		}
		s.MessageThreadID = value
	}
	{
		value, err := DecodeMessageReplyTo(b)
		if err != nil {
			return fmt.Errorf("unable to decode sendMessage#171a006f: field reply_to: %w", err)
		}
		s.ReplyTo = value
	}
	{
		if err := s.Options.Decode(b); err != nil {
			return fmt.Errorf("unable to decode sendMessage#171a006f: field options: %w", err)
		}
	}
	{
		value, err := DecodeReplyMarkup(b)
		if err != nil {
			return fmt.Errorf("unable to decode sendMessage#171a006f: field reply_markup: %w", err)
		}
		s.ReplyMarkup = value
	}
	{
		value, err := DecodeInputMessageContent(b)
		if err != nil {
			return fmt.Errorf("unable to decode sendMessage#171a006f: field input_message_content: %w", err)
		}
		s.InputMessageContent = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SendMessageRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode sendMessage#171a006f as nil")
	}
	b.ObjStart()
	b.PutID("sendMessage")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(s.ChatID)
	b.Comma()
	b.FieldStart("message_thread_id")
	b.PutInt53(s.MessageThreadID)
	b.Comma()
	b.FieldStart("reply_to")
	if s.ReplyTo == nil {
		return fmt.Errorf("unable to encode sendMessage#171a006f: field reply_to is nil")
	}
	if err := s.ReplyTo.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode sendMessage#171a006f: field reply_to: %w", err)
	}
	b.Comma()
	b.FieldStart("options")
	if err := s.Options.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode sendMessage#171a006f: field options: %w", err)
	}
	b.Comma()
	b.FieldStart("reply_markup")
	if s.ReplyMarkup == nil {
		return fmt.Errorf("unable to encode sendMessage#171a006f: field reply_markup is nil")
	}
	if err := s.ReplyMarkup.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode sendMessage#171a006f: field reply_markup: %w", err)
	}
	b.Comma()
	b.FieldStart("input_message_content")
	if s.InputMessageContent == nil {
		return fmt.Errorf("unable to encode sendMessage#171a006f: field input_message_content is nil")
	}
	if err := s.InputMessageContent.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode sendMessage#171a006f: field input_message_content: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SendMessageRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode sendMessage#171a006f to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("sendMessage"); err != nil {
				return fmt.Errorf("unable to decode sendMessage#171a006f: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode sendMessage#171a006f: field chat_id: %w", err)
			}
			s.ChatID = value
		case "message_thread_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode sendMessage#171a006f: field message_thread_id: %w", err)
			}
			s.MessageThreadID = value
		case "reply_to":
			value, err := DecodeTDLibJSONMessageReplyTo(b)
			if err != nil {
				return fmt.Errorf("unable to decode sendMessage#171a006f: field reply_to: %w", err)
			}
			s.ReplyTo = value
		case "options":
			if err := s.Options.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode sendMessage#171a006f: field options: %w", err)
			}
		case "reply_markup":
			value, err := DecodeTDLibJSONReplyMarkup(b)
			if err != nil {
				return fmt.Errorf("unable to decode sendMessage#171a006f: field reply_markup: %w", err)
			}
			s.ReplyMarkup = value
		case "input_message_content":
			value, err := DecodeTDLibJSONInputMessageContent(b)
			if err != nil {
				return fmt.Errorf("unable to decode sendMessage#171a006f: field input_message_content: %w", err)
			}
			s.InputMessageContent = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (s *SendMessageRequest) GetChatID() (value int64) {
	if s == nil {
		return
	}
	return s.ChatID
}

// GetMessageThreadID returns value of MessageThreadID field.
func (s *SendMessageRequest) GetMessageThreadID() (value int64) {
	if s == nil {
		return
	}
	return s.MessageThreadID
}

// GetReplyTo returns value of ReplyTo field.
func (s *SendMessageRequest) GetReplyTo() (value MessageReplyToClass) {
	if s == nil {
		return
	}
	return s.ReplyTo
}

// GetOptions returns value of Options field.
func (s *SendMessageRequest) GetOptions() (value MessageSendOptions) {
	if s == nil {
		return
	}
	return s.Options
}

// GetReplyMarkup returns value of ReplyMarkup field.
func (s *SendMessageRequest) GetReplyMarkup() (value ReplyMarkupClass) {
	if s == nil {
		return
	}
	return s.ReplyMarkup
}

// GetInputMessageContent returns value of InputMessageContent field.
func (s *SendMessageRequest) GetInputMessageContent() (value InputMessageContentClass) {
	if s == nil {
		return
	}
	return s.InputMessageContent
}

// SendMessage invokes method sendMessage#171a006f returning error if any.
func (c *Client) SendMessage(ctx context.Context, request *SendMessageRequest) (*Message, error) {
	var result Message

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
