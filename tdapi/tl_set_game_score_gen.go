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

// SetGameScoreRequest represents TL type `setGameScore#7eccedc6`.
type SetGameScoreRequest struct {
	// The chat to which the message with the game belongs
	ChatID int64
	// Identifier of the message
	MessageID int64
	// Pass true to edit the game message to include the current scoreboard
	EditMessage bool
	// User identifier
	UserID int64
	// The new score
	Score int32
	// Pass true to update the score even if it decreases. If the score is 0, the user will
	// be deleted from the high score table
	Force bool
}

// SetGameScoreRequestTypeID is TL type id of SetGameScoreRequest.
const SetGameScoreRequestTypeID = 0x7eccedc6

// Ensuring interfaces in compile-time for SetGameScoreRequest.
var (
	_ bin.Encoder     = &SetGameScoreRequest{}
	_ bin.Decoder     = &SetGameScoreRequest{}
	_ bin.BareEncoder = &SetGameScoreRequest{}
	_ bin.BareDecoder = &SetGameScoreRequest{}
)

func (s *SetGameScoreRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.ChatID == 0) {
		return false
	}
	if !(s.MessageID == 0) {
		return false
	}
	if !(s.EditMessage == false) {
		return false
	}
	if !(s.UserID == 0) {
		return false
	}
	if !(s.Score == 0) {
		return false
	}
	if !(s.Force == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SetGameScoreRequest) String() string {
	if s == nil {
		return "SetGameScoreRequest(nil)"
	}
	type Alias SetGameScoreRequest
	return fmt.Sprintf("SetGameScoreRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SetGameScoreRequest) TypeID() uint32 {
	return SetGameScoreRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SetGameScoreRequest) TypeName() string {
	return "setGameScore"
}

// TypeInfo returns info about TL type.
func (s *SetGameScoreRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "setGameScore",
		ID:   SetGameScoreRequestTypeID,
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
			Name:       "MessageID",
			SchemaName: "message_id",
		},
		{
			Name:       "EditMessage",
			SchemaName: "edit_message",
		},
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "Score",
			SchemaName: "score",
		},
		{
			Name:       "Force",
			SchemaName: "force",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SetGameScoreRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setGameScore#7eccedc6 as nil")
	}
	b.PutID(SetGameScoreRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SetGameScoreRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setGameScore#7eccedc6 as nil")
	}
	b.PutInt53(s.ChatID)
	b.PutInt53(s.MessageID)
	b.PutBool(s.EditMessage)
	b.PutInt53(s.UserID)
	b.PutInt32(s.Score)
	b.PutBool(s.Force)
	return nil
}

// Decode implements bin.Decoder.
func (s *SetGameScoreRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setGameScore#7eccedc6 to nil")
	}
	if err := b.ConsumeID(SetGameScoreRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode setGameScore#7eccedc6: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SetGameScoreRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setGameScore#7eccedc6 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode setGameScore#7eccedc6: field chat_id: %w", err)
		}
		s.ChatID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode setGameScore#7eccedc6: field message_id: %w", err)
		}
		s.MessageID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode setGameScore#7eccedc6: field edit_message: %w", err)
		}
		s.EditMessage = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode setGameScore#7eccedc6: field user_id: %w", err)
		}
		s.UserID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode setGameScore#7eccedc6: field score: %w", err)
		}
		s.Score = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode setGameScore#7eccedc6: field force: %w", err)
		}
		s.Force = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SetGameScoreRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode setGameScore#7eccedc6 as nil")
	}
	b.ObjStart()
	b.PutID("setGameScore")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(s.ChatID)
	b.Comma()
	b.FieldStart("message_id")
	b.PutInt53(s.MessageID)
	b.Comma()
	b.FieldStart("edit_message")
	b.PutBool(s.EditMessage)
	b.Comma()
	b.FieldStart("user_id")
	b.PutInt53(s.UserID)
	b.Comma()
	b.FieldStart("score")
	b.PutInt32(s.Score)
	b.Comma()
	b.FieldStart("force")
	b.PutBool(s.Force)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SetGameScoreRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode setGameScore#7eccedc6 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("setGameScore"); err != nil {
				return fmt.Errorf("unable to decode setGameScore#7eccedc6: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode setGameScore#7eccedc6: field chat_id: %w", err)
			}
			s.ChatID = value
		case "message_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode setGameScore#7eccedc6: field message_id: %w", err)
			}
			s.MessageID = value
		case "edit_message":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode setGameScore#7eccedc6: field edit_message: %w", err)
			}
			s.EditMessage = value
		case "user_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode setGameScore#7eccedc6: field user_id: %w", err)
			}
			s.UserID = value
		case "score":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode setGameScore#7eccedc6: field score: %w", err)
			}
			s.Score = value
		case "force":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode setGameScore#7eccedc6: field force: %w", err)
			}
			s.Force = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (s *SetGameScoreRequest) GetChatID() (value int64) {
	if s == nil {
		return
	}
	return s.ChatID
}

// GetMessageID returns value of MessageID field.
func (s *SetGameScoreRequest) GetMessageID() (value int64) {
	if s == nil {
		return
	}
	return s.MessageID
}

// GetEditMessage returns value of EditMessage field.
func (s *SetGameScoreRequest) GetEditMessage() (value bool) {
	if s == nil {
		return
	}
	return s.EditMessage
}

// GetUserID returns value of UserID field.
func (s *SetGameScoreRequest) GetUserID() (value int64) {
	if s == nil {
		return
	}
	return s.UserID
}

// GetScore returns value of Score field.
func (s *SetGameScoreRequest) GetScore() (value int32) {
	if s == nil {
		return
	}
	return s.Score
}

// GetForce returns value of Force field.
func (s *SetGameScoreRequest) GetForce() (value bool) {
	if s == nil {
		return
	}
	return s.Force
}

// SetGameScore invokes method setGameScore#7eccedc6 returning error if any.
func (c *Client) SetGameScore(ctx context.Context, request *SetGameScoreRequest) (*Message, error) {
	var result Message

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
