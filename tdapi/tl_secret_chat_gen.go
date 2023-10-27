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

// SecretChat represents TL type `secretChat#d7a70bcb`.
type SecretChat struct {
	// Secret chat identifier
	ID int32
	// Identifier of the chat partner
	UserID int64
	// State of the secret chat
	State SecretChatStateClass
	// True, if the chat was created by the current user; false otherwise
	IsOutbound bool
	// Hash of the currently used key for comparison with the hash of the chat partner's key.
	// This is a string of 36 little-endian bytes, which must be split into groups of 2 bits,
	// each denoting a pixel of one of 4 colors FFFFFF, D5E6F3, 2D5775, and 2F99C9.
	KeyHash []byte
	// Secret chat layer; determines features supported by the chat partner's application.
	// Nested text entities and underline and strikethrough entities are supported if the
	// layer >= 101,
	Layer int32
}

// SecretChatTypeID is TL type id of SecretChat.
const SecretChatTypeID = 0xd7a70bcb

// Ensuring interfaces in compile-time for SecretChat.
var (
	_ bin.Encoder     = &SecretChat{}
	_ bin.Decoder     = &SecretChat{}
	_ bin.BareEncoder = &SecretChat{}
	_ bin.BareDecoder = &SecretChat{}
)

func (s *SecretChat) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.ID == 0) {
		return false
	}
	if !(s.UserID == 0) {
		return false
	}
	if !(s.State == nil) {
		return false
	}
	if !(s.IsOutbound == false) {
		return false
	}
	if !(s.KeyHash == nil) {
		return false
	}
	if !(s.Layer == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SecretChat) String() string {
	if s == nil {
		return "SecretChat(nil)"
	}
	type Alias SecretChat
	return fmt.Sprintf("SecretChat%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SecretChat) TypeID() uint32 {
	return SecretChatTypeID
}

// TypeName returns name of type in TL schema.
func (*SecretChat) TypeName() string {
	return "secretChat"
}

// TypeInfo returns info about TL type.
func (s *SecretChat) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "secretChat",
		ID:   SecretChatTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "State",
			SchemaName: "state",
		},
		{
			Name:       "IsOutbound",
			SchemaName: "is_outbound",
		},
		{
			Name:       "KeyHash",
			SchemaName: "key_hash",
		},
		{
			Name:       "Layer",
			SchemaName: "layer",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SecretChat) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode secretChat#d7a70bcb as nil")
	}
	b.PutID(SecretChatTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SecretChat) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode secretChat#d7a70bcb as nil")
	}
	b.PutInt32(s.ID)
	b.PutInt53(s.UserID)
	if s.State == nil {
		return fmt.Errorf("unable to encode secretChat#d7a70bcb: field state is nil")
	}
	if err := s.State.Encode(b); err != nil {
		return fmt.Errorf("unable to encode secretChat#d7a70bcb: field state: %w", err)
	}
	b.PutBool(s.IsOutbound)
	b.PutBytes(s.KeyHash)
	b.PutInt32(s.Layer)
	return nil
}

// Decode implements bin.Decoder.
func (s *SecretChat) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode secretChat#d7a70bcb to nil")
	}
	if err := b.ConsumeID(SecretChatTypeID); err != nil {
		return fmt.Errorf("unable to decode secretChat#d7a70bcb: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SecretChat) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode secretChat#d7a70bcb to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode secretChat#d7a70bcb: field id: %w", err)
		}
		s.ID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode secretChat#d7a70bcb: field user_id: %w", err)
		}
		s.UserID = value
	}
	{
		value, err := DecodeSecretChatState(b)
		if err != nil {
			return fmt.Errorf("unable to decode secretChat#d7a70bcb: field state: %w", err)
		}
		s.State = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode secretChat#d7a70bcb: field is_outbound: %w", err)
		}
		s.IsOutbound = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode secretChat#d7a70bcb: field key_hash: %w", err)
		}
		s.KeyHash = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode secretChat#d7a70bcb: field layer: %w", err)
		}
		s.Layer = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SecretChat) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode secretChat#d7a70bcb as nil")
	}
	b.ObjStart()
	b.PutID("secretChat")
	b.Comma()
	b.FieldStart("id")
	b.PutInt32(s.ID)
	b.Comma()
	b.FieldStart("user_id")
	b.PutInt53(s.UserID)
	b.Comma()
	b.FieldStart("state")
	if s.State == nil {
		return fmt.Errorf("unable to encode secretChat#d7a70bcb: field state is nil")
	}
	if err := s.State.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode secretChat#d7a70bcb: field state: %w", err)
	}
	b.Comma()
	b.FieldStart("is_outbound")
	b.PutBool(s.IsOutbound)
	b.Comma()
	b.FieldStart("key_hash")
	b.PutBytes(s.KeyHash)
	b.Comma()
	b.FieldStart("layer")
	b.PutInt32(s.Layer)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SecretChat) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode secretChat#d7a70bcb to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("secretChat"); err != nil {
				return fmt.Errorf("unable to decode secretChat#d7a70bcb: %w", err)
			}
		case "id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode secretChat#d7a70bcb: field id: %w", err)
			}
			s.ID = value
		case "user_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode secretChat#d7a70bcb: field user_id: %w", err)
			}
			s.UserID = value
		case "state":
			value, err := DecodeTDLibJSONSecretChatState(b)
			if err != nil {
				return fmt.Errorf("unable to decode secretChat#d7a70bcb: field state: %w", err)
			}
			s.State = value
		case "is_outbound":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode secretChat#d7a70bcb: field is_outbound: %w", err)
			}
			s.IsOutbound = value
		case "key_hash":
			value, err := b.Bytes()
			if err != nil {
				return fmt.Errorf("unable to decode secretChat#d7a70bcb: field key_hash: %w", err)
			}
			s.KeyHash = value
		case "layer":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode secretChat#d7a70bcb: field layer: %w", err)
			}
			s.Layer = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetID returns value of ID field.
func (s *SecretChat) GetID() (value int32) {
	if s == nil {
		return
	}
	return s.ID
}

// GetUserID returns value of UserID field.
func (s *SecretChat) GetUserID() (value int64) {
	if s == nil {
		return
	}
	return s.UserID
}

// GetState returns value of State field.
func (s *SecretChat) GetState() (value SecretChatStateClass) {
	if s == nil {
		return
	}
	return s.State
}

// GetIsOutbound returns value of IsOutbound field.
func (s *SecretChat) GetIsOutbound() (value bool) {
	if s == nil {
		return
	}
	return s.IsOutbound
}

// GetKeyHash returns value of KeyHash field.
func (s *SecretChat) GetKeyHash() (value []byte) {
	if s == nil {
		return
	}
	return s.KeyHash
}

// GetLayer returns value of Layer field.
func (s *SecretChat) GetLayer() (value int32) {
	if s == nil {
		return
	}
	return s.Layer
}
