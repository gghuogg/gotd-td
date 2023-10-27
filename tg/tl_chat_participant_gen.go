// Code generated by gotdgen, DO NOT EDIT.

package tg

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

// ChatParticipant represents TL type `chatParticipant#c02d4007`.
// Group member.
//
// See https://core.telegram.org/constructor/chatParticipant for reference.
type ChatParticipant struct {
	// Member user ID
	UserID int64
	// ID of the user that added the member to the group
	InviterID int64
	// Date added to the group
	Date int
}

// ChatParticipantTypeID is TL type id of ChatParticipant.
const ChatParticipantTypeID = 0xc02d4007

// construct implements constructor of ChatParticipantClass.
func (c ChatParticipant) construct() ChatParticipantClass { return &c }

// Ensuring interfaces in compile-time for ChatParticipant.
var (
	_ bin.Encoder     = &ChatParticipant{}
	_ bin.Decoder     = &ChatParticipant{}
	_ bin.BareEncoder = &ChatParticipant{}
	_ bin.BareDecoder = &ChatParticipant{}

	_ ChatParticipantClass = &ChatParticipant{}
)

func (c *ChatParticipant) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.UserID == 0) {
		return false
	}
	if !(c.InviterID == 0) {
		return false
	}
	if !(c.Date == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatParticipant) String() string {
	if c == nil {
		return "ChatParticipant(nil)"
	}
	type Alias ChatParticipant
	return fmt.Sprintf("ChatParticipant%+v", Alias(*c))
}

// FillFrom fills ChatParticipant from given interface.
func (c *ChatParticipant) FillFrom(from interface {
	GetUserID() (value int64)
	GetInviterID() (value int64)
	GetDate() (value int)
}) {
	c.UserID = from.GetUserID()
	c.InviterID = from.GetInviterID()
	c.Date = from.GetDate()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatParticipant) TypeID() uint32 {
	return ChatParticipantTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatParticipant) TypeName() string {
	return "chatParticipant"
}

// TypeInfo returns info about TL type.
func (c *ChatParticipant) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatParticipant",
		ID:   ChatParticipantTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "InviterID",
			SchemaName: "inviter_id",
		},
		{
			Name:       "Date",
			SchemaName: "date",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatParticipant) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatParticipant#c02d4007 as nil")
	}
	b.PutID(ChatParticipantTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatParticipant) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatParticipant#c02d4007 as nil")
	}
	b.PutLong(c.UserID)
	b.PutLong(c.InviterID)
	b.PutInt(c.Date)
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatParticipant) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatParticipant#c02d4007 to nil")
	}
	if err := b.ConsumeID(ChatParticipantTypeID); err != nil {
		return fmt.Errorf("unable to decode chatParticipant#c02d4007: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatParticipant) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatParticipant#c02d4007 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode chatParticipant#c02d4007: field user_id: %w", err)
		}
		c.UserID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode chatParticipant#c02d4007: field inviter_id: %w", err)
		}
		c.InviterID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatParticipant#c02d4007: field date: %w", err)
		}
		c.Date = value
	}
	return nil
}

// GetUserID returns value of UserID field.
func (c *ChatParticipant) GetUserID() (value int64) {
	if c == nil {
		return
	}
	return c.UserID
}

// GetInviterID returns value of InviterID field.
func (c *ChatParticipant) GetInviterID() (value int64) {
	if c == nil {
		return
	}
	return c.InviterID
}

// GetDate returns value of Date field.
func (c *ChatParticipant) GetDate() (value int) {
	if c == nil {
		return
	}
	return c.Date
}

// ChatParticipantCreator represents TL type `chatParticipantCreator#e46bcee4`.
// Represents the creator of the group
//
// See https://core.telegram.org/constructor/chatParticipantCreator for reference.
type ChatParticipantCreator struct {
	// ID of the user that created the group
	UserID int64
}

// ChatParticipantCreatorTypeID is TL type id of ChatParticipantCreator.
const ChatParticipantCreatorTypeID = 0xe46bcee4

// construct implements constructor of ChatParticipantClass.
func (c ChatParticipantCreator) construct() ChatParticipantClass { return &c }

// Ensuring interfaces in compile-time for ChatParticipantCreator.
var (
	_ bin.Encoder     = &ChatParticipantCreator{}
	_ bin.Decoder     = &ChatParticipantCreator{}
	_ bin.BareEncoder = &ChatParticipantCreator{}
	_ bin.BareDecoder = &ChatParticipantCreator{}

	_ ChatParticipantClass = &ChatParticipantCreator{}
)

func (c *ChatParticipantCreator) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.UserID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatParticipantCreator) String() string {
	if c == nil {
		return "ChatParticipantCreator(nil)"
	}
	type Alias ChatParticipantCreator
	return fmt.Sprintf("ChatParticipantCreator%+v", Alias(*c))
}

// FillFrom fills ChatParticipantCreator from given interface.
func (c *ChatParticipantCreator) FillFrom(from interface {
	GetUserID() (value int64)
}) {
	c.UserID = from.GetUserID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatParticipantCreator) TypeID() uint32 {
	return ChatParticipantCreatorTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatParticipantCreator) TypeName() string {
	return "chatParticipantCreator"
}

// TypeInfo returns info about TL type.
func (c *ChatParticipantCreator) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatParticipantCreator",
		ID:   ChatParticipantCreatorTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatParticipantCreator) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatParticipantCreator#e46bcee4 as nil")
	}
	b.PutID(ChatParticipantCreatorTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatParticipantCreator) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatParticipantCreator#e46bcee4 as nil")
	}
	b.PutLong(c.UserID)
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatParticipantCreator) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatParticipantCreator#e46bcee4 to nil")
	}
	if err := b.ConsumeID(ChatParticipantCreatorTypeID); err != nil {
		return fmt.Errorf("unable to decode chatParticipantCreator#e46bcee4: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatParticipantCreator) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatParticipantCreator#e46bcee4 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode chatParticipantCreator#e46bcee4: field user_id: %w", err)
		}
		c.UserID = value
	}
	return nil
}

// GetUserID returns value of UserID field.
func (c *ChatParticipantCreator) GetUserID() (value int64) {
	if c == nil {
		return
	}
	return c.UserID
}

// ChatParticipantAdmin represents TL type `chatParticipantAdmin#a0933f5b`.
// Chat admin
//
// See https://core.telegram.org/constructor/chatParticipantAdmin for reference.
type ChatParticipantAdmin struct {
	// ID of a group member that is admin
	UserID int64
	// ID of the user that added the member to the group
	InviterID int64
	// Date when the user was added
	Date int
}

// ChatParticipantAdminTypeID is TL type id of ChatParticipantAdmin.
const ChatParticipantAdminTypeID = 0xa0933f5b

// construct implements constructor of ChatParticipantClass.
func (c ChatParticipantAdmin) construct() ChatParticipantClass { return &c }

// Ensuring interfaces in compile-time for ChatParticipantAdmin.
var (
	_ bin.Encoder     = &ChatParticipantAdmin{}
	_ bin.Decoder     = &ChatParticipantAdmin{}
	_ bin.BareEncoder = &ChatParticipantAdmin{}
	_ bin.BareDecoder = &ChatParticipantAdmin{}

	_ ChatParticipantClass = &ChatParticipantAdmin{}
)

func (c *ChatParticipantAdmin) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.UserID == 0) {
		return false
	}
	if !(c.InviterID == 0) {
		return false
	}
	if !(c.Date == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatParticipantAdmin) String() string {
	if c == nil {
		return "ChatParticipantAdmin(nil)"
	}
	type Alias ChatParticipantAdmin
	return fmt.Sprintf("ChatParticipantAdmin%+v", Alias(*c))
}

// FillFrom fills ChatParticipantAdmin from given interface.
func (c *ChatParticipantAdmin) FillFrom(from interface {
	GetUserID() (value int64)
	GetInviterID() (value int64)
	GetDate() (value int)
}) {
	c.UserID = from.GetUserID()
	c.InviterID = from.GetInviterID()
	c.Date = from.GetDate()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatParticipantAdmin) TypeID() uint32 {
	return ChatParticipantAdminTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatParticipantAdmin) TypeName() string {
	return "chatParticipantAdmin"
}

// TypeInfo returns info about TL type.
func (c *ChatParticipantAdmin) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatParticipantAdmin",
		ID:   ChatParticipantAdminTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "InviterID",
			SchemaName: "inviter_id",
		},
		{
			Name:       "Date",
			SchemaName: "date",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatParticipantAdmin) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatParticipantAdmin#a0933f5b as nil")
	}
	b.PutID(ChatParticipantAdminTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatParticipantAdmin) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatParticipantAdmin#a0933f5b as nil")
	}
	b.PutLong(c.UserID)
	b.PutLong(c.InviterID)
	b.PutInt(c.Date)
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatParticipantAdmin) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatParticipantAdmin#a0933f5b to nil")
	}
	if err := b.ConsumeID(ChatParticipantAdminTypeID); err != nil {
		return fmt.Errorf("unable to decode chatParticipantAdmin#a0933f5b: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatParticipantAdmin) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatParticipantAdmin#a0933f5b to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode chatParticipantAdmin#a0933f5b: field user_id: %w", err)
		}
		c.UserID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode chatParticipantAdmin#a0933f5b: field inviter_id: %w", err)
		}
		c.InviterID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatParticipantAdmin#a0933f5b: field date: %w", err)
		}
		c.Date = value
	}
	return nil
}

// GetUserID returns value of UserID field.
func (c *ChatParticipantAdmin) GetUserID() (value int64) {
	if c == nil {
		return
	}
	return c.UserID
}

// GetInviterID returns value of InviterID field.
func (c *ChatParticipantAdmin) GetInviterID() (value int64) {
	if c == nil {
		return
	}
	return c.InviterID
}

// GetDate returns value of Date field.
func (c *ChatParticipantAdmin) GetDate() (value int) {
	if c == nil {
		return
	}
	return c.Date
}

// ChatParticipantClassName is schema name of ChatParticipantClass.
const ChatParticipantClassName = "ChatParticipant"

// ChatParticipantClass represents ChatParticipant generic type.
//
// See https://core.telegram.org/type/ChatParticipant for reference.
//
// Example:
//
//	g, err := tg.DecodeChatParticipant(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.ChatParticipant: // chatParticipant#c02d4007
//	case *tg.ChatParticipantCreator: // chatParticipantCreator#e46bcee4
//	case *tg.ChatParticipantAdmin: // chatParticipantAdmin#a0933f5b
//	default: panic(v)
//	}
type ChatParticipantClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() ChatParticipantClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	// Member user ID
	GetUserID() (value int64)
}

// DecodeChatParticipant implements binary de-serialization for ChatParticipantClass.
func DecodeChatParticipant(buf *bin.Buffer) (ChatParticipantClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case ChatParticipantTypeID:
		// Decoding chatParticipant#c02d4007.
		v := ChatParticipant{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatParticipantClass: %w", err)
		}
		return &v, nil
	case ChatParticipantCreatorTypeID:
		// Decoding chatParticipantCreator#e46bcee4.
		v := ChatParticipantCreator{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatParticipantClass: %w", err)
		}
		return &v, nil
	case ChatParticipantAdminTypeID:
		// Decoding chatParticipantAdmin#a0933f5b.
		v := ChatParticipantAdmin{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatParticipantClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode ChatParticipantClass: %w", bin.NewUnexpectedID(id))
	}
}

// ChatParticipant boxes the ChatParticipantClass providing a helper.
type ChatParticipantBox struct {
	ChatParticipant ChatParticipantClass
}

// Decode implements bin.Decoder for ChatParticipantBox.
func (b *ChatParticipantBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode ChatParticipantBox to nil")
	}
	v, err := DecodeChatParticipant(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.ChatParticipant = v
	return nil
}

// Encode implements bin.Encode for ChatParticipantBox.
func (b *ChatParticipantBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.ChatParticipant == nil {
		return fmt.Errorf("unable to encode ChatParticipantClass as nil")
	}
	return b.ChatParticipant.Encode(buf)
}
