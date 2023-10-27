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

// PhoneGroupParticipants represents TL type `phone.groupParticipants#f47751b6`.
type PhoneGroupParticipants struct {
	// Count field of PhoneGroupParticipants.
	Count int
	// Participants field of PhoneGroupParticipants.
	Participants []GroupCallParticipant
	// NextOffset field of PhoneGroupParticipants.
	NextOffset string
	// Chats field of PhoneGroupParticipants.
	Chats []ChatClass
	// Users field of PhoneGroupParticipants.
	Users []UserClass
	// Version field of PhoneGroupParticipants.
	Version int
}

// PhoneGroupParticipantsTypeID is TL type id of PhoneGroupParticipants.
const PhoneGroupParticipantsTypeID = 0xf47751b6

// Ensuring interfaces in compile-time for PhoneGroupParticipants.
var (
	_ bin.Encoder     = &PhoneGroupParticipants{}
	_ bin.Decoder     = &PhoneGroupParticipants{}
	_ bin.BareEncoder = &PhoneGroupParticipants{}
	_ bin.BareDecoder = &PhoneGroupParticipants{}
)

func (g *PhoneGroupParticipants) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Count == 0) {
		return false
	}
	if !(g.Participants == nil) {
		return false
	}
	if !(g.NextOffset == "") {
		return false
	}
	if !(g.Chats == nil) {
		return false
	}
	if !(g.Users == nil) {
		return false
	}
	if !(g.Version == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *PhoneGroupParticipants) String() string {
	if g == nil {
		return "PhoneGroupParticipants(nil)"
	}
	type Alias PhoneGroupParticipants
	return fmt.Sprintf("PhoneGroupParticipants%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PhoneGroupParticipants) TypeID() uint32 {
	return PhoneGroupParticipantsTypeID
}

// TypeName returns name of type in TL schema.
func (*PhoneGroupParticipants) TypeName() string {
	return "phone.groupParticipants"
}

// TypeInfo returns info about TL type.
func (g *PhoneGroupParticipants) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "phone.groupParticipants",
		ID:   PhoneGroupParticipantsTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Count",
			SchemaName: "count",
		},
		{
			Name:       "Participants",
			SchemaName: "participants",
		},
		{
			Name:       "NextOffset",
			SchemaName: "next_offset",
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
			Name:       "Version",
			SchemaName: "version",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *PhoneGroupParticipants) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode phone.groupParticipants#f47751b6 as nil")
	}
	b.PutID(PhoneGroupParticipantsTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *PhoneGroupParticipants) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode phone.groupParticipants#f47751b6 as nil")
	}
	b.PutInt(g.Count)
	b.PutVectorHeader(len(g.Participants))
	for idx, v := range g.Participants {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode phone.groupParticipants#f47751b6: field participants element with index %d: %w", idx, err)
		}
	}
	b.PutString(g.NextOffset)
	b.PutVectorHeader(len(g.Chats))
	for idx, v := range g.Chats {
		if v == nil {
			return fmt.Errorf("unable to encode phone.groupParticipants#f47751b6: field chats element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode phone.groupParticipants#f47751b6: field chats element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(g.Users))
	for idx, v := range g.Users {
		if v == nil {
			return fmt.Errorf("unable to encode phone.groupParticipants#f47751b6: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode phone.groupParticipants#f47751b6: field users element with index %d: %w", idx, err)
		}
	}
	b.PutInt(g.Version)
	return nil
}

// Decode implements bin.Decoder.
func (g *PhoneGroupParticipants) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode phone.groupParticipants#f47751b6 to nil")
	}
	if err := b.ConsumeID(PhoneGroupParticipantsTypeID); err != nil {
		return fmt.Errorf("unable to decode phone.groupParticipants#f47751b6: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *PhoneGroupParticipants) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode phone.groupParticipants#f47751b6 to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode phone.groupParticipants#f47751b6: field count: %w", err)
		}
		g.Count = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode phone.groupParticipants#f47751b6: field participants: %w", err)
		}

		if headerLen > 0 {
			g.Participants = make([]GroupCallParticipant, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value GroupCallParticipant
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode phone.groupParticipants#f47751b6: field participants: %w", err)
			}
			g.Participants = append(g.Participants, value)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode phone.groupParticipants#f47751b6: field next_offset: %w", err)
		}
		g.NextOffset = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode phone.groupParticipants#f47751b6: field chats: %w", err)
		}

		if headerLen > 0 {
			g.Chats = make([]ChatClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChat(b)
			if err != nil {
				return fmt.Errorf("unable to decode phone.groupParticipants#f47751b6: field chats: %w", err)
			}
			g.Chats = append(g.Chats, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode phone.groupParticipants#f47751b6: field users: %w", err)
		}

		if headerLen > 0 {
			g.Users = make([]UserClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode phone.groupParticipants#f47751b6: field users: %w", err)
			}
			g.Users = append(g.Users, value)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode phone.groupParticipants#f47751b6: field version: %w", err)
		}
		g.Version = value
	}
	return nil
}

// GetCount returns value of Count field.
func (g *PhoneGroupParticipants) GetCount() (value int) {
	if g == nil {
		return
	}
	return g.Count
}

// GetParticipants returns value of Participants field.
func (g *PhoneGroupParticipants) GetParticipants() (value []GroupCallParticipant) {
	if g == nil {
		return
	}
	return g.Participants
}

// GetNextOffset returns value of NextOffset field.
func (g *PhoneGroupParticipants) GetNextOffset() (value string) {
	if g == nil {
		return
	}
	return g.NextOffset
}

// GetChats returns value of Chats field.
func (g *PhoneGroupParticipants) GetChats() (value []ChatClass) {
	if g == nil {
		return
	}
	return g.Chats
}

// GetUsers returns value of Users field.
func (g *PhoneGroupParticipants) GetUsers() (value []UserClass) {
	if g == nil {
		return
	}
	return g.Users
}

// GetVersion returns value of Version field.
func (g *PhoneGroupParticipants) GetVersion() (value int) {
	if g == nil {
		return
	}
	return g.Version
}
