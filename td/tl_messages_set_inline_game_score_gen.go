// Code generated by gotdgen, DO NOT EDIT.

package td

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

// MessagesSetInlineGameScoreRequest represents TL type `messages.setInlineGameScore#15ad9f64`.
type MessagesSetInlineGameScoreRequest struct {
	// Flags field of MessagesSetInlineGameScoreRequest.
	Flags bin.Fields
	// EditMessage field of MessagesSetInlineGameScoreRequest.
	EditMessage bool
	// Force field of MessagesSetInlineGameScoreRequest.
	Force bool
	// ID field of MessagesSetInlineGameScoreRequest.
	ID InputBotInlineMessageIDClass
	// UserID field of MessagesSetInlineGameScoreRequest.
	UserID InputUserClass
	// Score field of MessagesSetInlineGameScoreRequest.
	Score int
}

// MessagesSetInlineGameScoreRequestTypeID is TL type id of MessagesSetInlineGameScoreRequest.
const MessagesSetInlineGameScoreRequestTypeID = 0x15ad9f64

// Ensuring interfaces in compile-time for MessagesSetInlineGameScoreRequest.
var (
	_ bin.Encoder     = &MessagesSetInlineGameScoreRequest{}
	_ bin.Decoder     = &MessagesSetInlineGameScoreRequest{}
	_ bin.BareEncoder = &MessagesSetInlineGameScoreRequest{}
	_ bin.BareDecoder = &MessagesSetInlineGameScoreRequest{}
)

func (s *MessagesSetInlineGameScoreRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.EditMessage == false) {
		return false
	}
	if !(s.Force == false) {
		return false
	}
	if !(s.ID == nil) {
		return false
	}
	if !(s.UserID == nil) {
		return false
	}
	if !(s.Score == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesSetInlineGameScoreRequest) String() string {
	if s == nil {
		return "MessagesSetInlineGameScoreRequest(nil)"
	}
	type Alias MessagesSetInlineGameScoreRequest
	return fmt.Sprintf("MessagesSetInlineGameScoreRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesSetInlineGameScoreRequest) TypeID() uint32 {
	return MessagesSetInlineGameScoreRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesSetInlineGameScoreRequest) TypeName() string {
	return "messages.setInlineGameScore"
}

// TypeInfo returns info about TL type.
func (s *MessagesSetInlineGameScoreRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.setInlineGameScore",
		ID:   MessagesSetInlineGameScoreRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "EditMessage",
			SchemaName: "edit_message",
			Null:       !s.Flags.Has(0),
		},
		{
			Name:       "Force",
			SchemaName: "force",
			Null:       !s.Flags.Has(1),
		},
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "Score",
			SchemaName: "score",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (s *MessagesSetInlineGameScoreRequest) SetFlags() {
	if !(s.EditMessage == false) {
		s.Flags.Set(0)
	}
	if !(s.Force == false) {
		s.Flags.Set(1)
	}
}

// Encode implements bin.Encoder.
func (s *MessagesSetInlineGameScoreRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.setInlineGameScore#15ad9f64 as nil")
	}
	b.PutID(MessagesSetInlineGameScoreRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *MessagesSetInlineGameScoreRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.setInlineGameScore#15ad9f64 as nil")
	}
	s.SetFlags()
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.setInlineGameScore#15ad9f64: field flags: %w", err)
	}
	if s.ID == nil {
		return fmt.Errorf("unable to encode messages.setInlineGameScore#15ad9f64: field id is nil")
	}
	if err := s.ID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.setInlineGameScore#15ad9f64: field id: %w", err)
	}
	if s.UserID == nil {
		return fmt.Errorf("unable to encode messages.setInlineGameScore#15ad9f64: field user_id is nil")
	}
	if err := s.UserID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.setInlineGameScore#15ad9f64: field user_id: %w", err)
	}
	b.PutInt(s.Score)
	return nil
}

// Decode implements bin.Decoder.
func (s *MessagesSetInlineGameScoreRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.setInlineGameScore#15ad9f64 to nil")
	}
	if err := b.ConsumeID(MessagesSetInlineGameScoreRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.setInlineGameScore#15ad9f64: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *MessagesSetInlineGameScoreRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.setInlineGameScore#15ad9f64 to nil")
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.setInlineGameScore#15ad9f64: field flags: %w", err)
		}
	}
	s.EditMessage = s.Flags.Has(0)
	s.Force = s.Flags.Has(1)
	{
		value, err := DecodeInputBotInlineMessageID(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.setInlineGameScore#15ad9f64: field id: %w", err)
		}
		s.ID = value
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.setInlineGameScore#15ad9f64: field user_id: %w", err)
		}
		s.UserID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.setInlineGameScore#15ad9f64: field score: %w", err)
		}
		s.Score = value
	}
	return nil
}

// SetEditMessage sets value of EditMessage conditional field.
func (s *MessagesSetInlineGameScoreRequest) SetEditMessage(value bool) {
	if value {
		s.Flags.Set(0)
		s.EditMessage = true
	} else {
		s.Flags.Unset(0)
		s.EditMessage = false
	}
}

// GetEditMessage returns value of EditMessage conditional field.
func (s *MessagesSetInlineGameScoreRequest) GetEditMessage() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(0)
}

// SetForce sets value of Force conditional field.
func (s *MessagesSetInlineGameScoreRequest) SetForce(value bool) {
	if value {
		s.Flags.Set(1)
		s.Force = true
	} else {
		s.Flags.Unset(1)
		s.Force = false
	}
}

// GetForce returns value of Force conditional field.
func (s *MessagesSetInlineGameScoreRequest) GetForce() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(1)
}

// GetID returns value of ID field.
func (s *MessagesSetInlineGameScoreRequest) GetID() (value InputBotInlineMessageIDClass) {
	if s == nil {
		return
	}
	return s.ID
}

// GetUserID returns value of UserID field.
func (s *MessagesSetInlineGameScoreRequest) GetUserID() (value InputUserClass) {
	if s == nil {
		return
	}
	return s.UserID
}

// GetScore returns value of Score field.
func (s *MessagesSetInlineGameScoreRequest) GetScore() (value int) {
	if s == nil {
		return
	}
	return s.Score
}

// MessagesSetInlineGameScore invokes method messages.setInlineGameScore#15ad9f64 returning error if any.
func (c *Client) MessagesSetInlineGameScore(ctx context.Context, request *MessagesSetInlineGameScoreRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
