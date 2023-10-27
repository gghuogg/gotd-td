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

// MessagesMessageReactionsList represents TL type `messages.messageReactionsList#31bd492d`.
type MessagesMessageReactionsList struct {
	// Flags field of MessagesMessageReactionsList.
	Flags bin.Fields
	// Count field of MessagesMessageReactionsList.
	Count int
	// Reactions field of MessagesMessageReactionsList.
	Reactions []MessagePeerReaction
	// Chats field of MessagesMessageReactionsList.
	Chats []ChatClass
	// Users field of MessagesMessageReactionsList.
	Users []UserClass
	// NextOffset field of MessagesMessageReactionsList.
	//
	// Use SetNextOffset and GetNextOffset helpers.
	NextOffset string
}

// MessagesMessageReactionsListTypeID is TL type id of MessagesMessageReactionsList.
const MessagesMessageReactionsListTypeID = 0x31bd492d

// Ensuring interfaces in compile-time for MessagesMessageReactionsList.
var (
	_ bin.Encoder     = &MessagesMessageReactionsList{}
	_ bin.Decoder     = &MessagesMessageReactionsList{}
	_ bin.BareEncoder = &MessagesMessageReactionsList{}
	_ bin.BareDecoder = &MessagesMessageReactionsList{}
)

func (m *MessagesMessageReactionsList) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.Flags.Zero()) {
		return false
	}
	if !(m.Count == 0) {
		return false
	}
	if !(m.Reactions == nil) {
		return false
	}
	if !(m.Chats == nil) {
		return false
	}
	if !(m.Users == nil) {
		return false
	}
	if !(m.NextOffset == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessagesMessageReactionsList) String() string {
	if m == nil {
		return "MessagesMessageReactionsList(nil)"
	}
	type Alias MessagesMessageReactionsList
	return fmt.Sprintf("MessagesMessageReactionsList%+v", Alias(*m))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesMessageReactionsList) TypeID() uint32 {
	return MessagesMessageReactionsListTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesMessageReactionsList) TypeName() string {
	return "messages.messageReactionsList"
}

// TypeInfo returns info about TL type.
func (m *MessagesMessageReactionsList) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.messageReactionsList",
		ID:   MessagesMessageReactionsListTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Count",
			SchemaName: "count",
		},
		{
			Name:       "Reactions",
			SchemaName: "reactions",
		},
		{
			Name:       "Chats",
			SchemaName: "chats",
		},
		{
			Name:       "Users",
			SchemaName: "users",
		},
		{
			Name:       "NextOffset",
			SchemaName: "next_offset",
			Null:       !m.Flags.Has(0),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (m *MessagesMessageReactionsList) SetFlags() {
	if !(m.NextOffset == "") {
		m.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (m *MessagesMessageReactionsList) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messages.messageReactionsList#31bd492d as nil")
	}
	b.PutID(MessagesMessageReactionsListTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MessagesMessageReactionsList) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messages.messageReactionsList#31bd492d as nil")
	}
	m.SetFlags()
	if err := m.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.messageReactionsList#31bd492d: field flags: %w", err)
	}
	b.PutInt(m.Count)
	b.PutVectorHeader(len(m.Reactions))
	for idx, v := range m.Reactions {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.messageReactionsList#31bd492d: field reactions element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(m.Chats))
	for idx, v := range m.Chats {
		if v == nil {
			return fmt.Errorf("unable to encode messages.messageReactionsList#31bd492d: field chats element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.messageReactionsList#31bd492d: field chats element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(m.Users))
	for idx, v := range m.Users {
		if v == nil {
			return fmt.Errorf("unable to encode messages.messageReactionsList#31bd492d: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.messageReactionsList#31bd492d: field users element with index %d: %w", idx, err)
		}
	}
	if m.Flags.Has(0) {
		b.PutString(m.NextOffset)
	}
	return nil
}

// Decode implements bin.Decoder.
func (m *MessagesMessageReactionsList) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messages.messageReactionsList#31bd492d to nil")
	}
	if err := b.ConsumeID(MessagesMessageReactionsListTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.messageReactionsList#31bd492d: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MessagesMessageReactionsList) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messages.messageReactionsList#31bd492d to nil")
	}
	{
		if err := m.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.messageReactionsList#31bd492d: field flags: %w", err)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.messageReactionsList#31bd492d: field count: %w", err)
		}
		m.Count = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.messageReactionsList#31bd492d: field reactions: %w", err)
		}

		if headerLen > 0 {
			m.Reactions = make([]MessagePeerReaction, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value MessagePeerReaction
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode messages.messageReactionsList#31bd492d: field reactions: %w", err)
			}
			m.Reactions = append(m.Reactions, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.messageReactionsList#31bd492d: field chats: %w", err)
		}

		if headerLen > 0 {
			m.Chats = make([]ChatClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChat(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.messageReactionsList#31bd492d: field chats: %w", err)
			}
			m.Chats = append(m.Chats, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.messageReactionsList#31bd492d: field users: %w", err)
		}

		if headerLen > 0 {
			m.Users = make([]UserClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.messageReactionsList#31bd492d: field users: %w", err)
			}
			m.Users = append(m.Users, value)
		}
	}
	if m.Flags.Has(0) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.messageReactionsList#31bd492d: field next_offset: %w", err)
		}
		m.NextOffset = value
	}
	return nil
}

// GetCount returns value of Count field.
func (m *MessagesMessageReactionsList) GetCount() (value int) {
	if m == nil {
		return
	}
	return m.Count
}

// GetReactions returns value of Reactions field.
func (m *MessagesMessageReactionsList) GetReactions() (value []MessagePeerReaction) {
	if m == nil {
		return
	}
	return m.Reactions
}

// GetChats returns value of Chats field.
func (m *MessagesMessageReactionsList) GetChats() (value []ChatClass) {
	if m == nil {
		return
	}
	return m.Chats
}

// GetUsers returns value of Users field.
func (m *MessagesMessageReactionsList) GetUsers() (value []UserClass) {
	if m == nil {
		return
	}
	return m.Users
}

// SetNextOffset sets value of NextOffset conditional field.
func (m *MessagesMessageReactionsList) SetNextOffset(value string) {
	m.Flags.Set(0)
	m.NextOffset = value
}

// GetNextOffset returns value of NextOffset conditional field and
// boolean which is true if field was set.
func (m *MessagesMessageReactionsList) GetNextOffset() (value string, ok bool) {
	if m == nil {
		return
	}
	if !m.Flags.Has(0) {
		return value, false
	}
	return m.NextOffset, true
}
