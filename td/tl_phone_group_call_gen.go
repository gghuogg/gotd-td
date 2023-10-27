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

// PhoneGroupCall represents TL type `phone.groupCall#9e727aad`.
type PhoneGroupCall struct {
	// Call field of PhoneGroupCall.
	Call GroupCallClass
	// Participants field of PhoneGroupCall.
	Participants []GroupCallParticipant
	// ParticipantsNextOffset field of PhoneGroupCall.
	ParticipantsNextOffset string
	// Chats field of PhoneGroupCall.
	Chats []ChatClass
	// Users field of PhoneGroupCall.
	Users []UserClass
}

// PhoneGroupCallTypeID is TL type id of PhoneGroupCall.
const PhoneGroupCallTypeID = 0x9e727aad

// Ensuring interfaces in compile-time for PhoneGroupCall.
var (
	_ bin.Encoder     = &PhoneGroupCall{}
	_ bin.Decoder     = &PhoneGroupCall{}
	_ bin.BareEncoder = &PhoneGroupCall{}
	_ bin.BareDecoder = &PhoneGroupCall{}
)

func (g *PhoneGroupCall) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Call == nil) {
		return false
	}
	if !(g.Participants == nil) {
		return false
	}
	if !(g.ParticipantsNextOffset == "") {
		return false
	}
	if !(g.Chats == nil) {
		return false
	}
	if !(g.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *PhoneGroupCall) String() string {
	if g == nil {
		return "PhoneGroupCall(nil)"
	}
	type Alias PhoneGroupCall
	return fmt.Sprintf("PhoneGroupCall%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PhoneGroupCall) TypeID() uint32 {
	return PhoneGroupCallTypeID
}

// TypeName returns name of type in TL schema.
func (*PhoneGroupCall) TypeName() string {
	return "phone.groupCall"
}

// TypeInfo returns info about TL type.
func (g *PhoneGroupCall) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "phone.groupCall",
		ID:   PhoneGroupCallTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Call",
			SchemaName: "call",
		},
		{
			Name:       "Participants",
			SchemaName: "participants",
		},
		{
			Name:       "ParticipantsNextOffset",
			SchemaName: "participants_next_offset",
		},
		{
			Name:       "Chats",
			SchemaName: "chats",
		},
		{
			Name:       "Users",
			SchemaName: "users",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *PhoneGroupCall) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode phone.groupCall#9e727aad as nil")
	}
	b.PutID(PhoneGroupCallTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *PhoneGroupCall) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode phone.groupCall#9e727aad as nil")
	}
	if g.Call == nil {
		return fmt.Errorf("unable to encode phone.groupCall#9e727aad: field call is nil")
	}
	if err := g.Call.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.groupCall#9e727aad: field call: %w", err)
	}
	b.PutVectorHeader(len(g.Participants))
	for idx, v := range g.Participants {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode phone.groupCall#9e727aad: field participants element with index %d: %w", idx, err)
		}
	}
	b.PutString(g.ParticipantsNextOffset)
	b.PutVectorHeader(len(g.Chats))
	for idx, v := range g.Chats {
		if v == nil {
			return fmt.Errorf("unable to encode phone.groupCall#9e727aad: field chats element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode phone.groupCall#9e727aad: field chats element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(g.Users))
	for idx, v := range g.Users {
		if v == nil {
			return fmt.Errorf("unable to encode phone.groupCall#9e727aad: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode phone.groupCall#9e727aad: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *PhoneGroupCall) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode phone.groupCall#9e727aad to nil")
	}
	if err := b.ConsumeID(PhoneGroupCallTypeID); err != nil {
		return fmt.Errorf("unable to decode phone.groupCall#9e727aad: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *PhoneGroupCall) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode phone.groupCall#9e727aad to nil")
	}
	{
		value, err := DecodeGroupCall(b)
		if err != nil {
			return fmt.Errorf("unable to decode phone.groupCall#9e727aad: field call: %w", err)
		}
		g.Call = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode phone.groupCall#9e727aad: field participants: %w", err)
		}

		if headerLen > 0 {
			g.Participants = make([]GroupCallParticipant, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value GroupCallParticipant
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode phone.groupCall#9e727aad: field participants: %w", err)
			}
			g.Participants = append(g.Participants, value)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode phone.groupCall#9e727aad: field participants_next_offset: %w", err)
		}
		g.ParticipantsNextOffset = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode phone.groupCall#9e727aad: field chats: %w", err)
		}

		if headerLen > 0 {
			g.Chats = make([]ChatClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChat(b)
			if err != nil {
				return fmt.Errorf("unable to decode phone.groupCall#9e727aad: field chats: %w", err)
			}
			g.Chats = append(g.Chats, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode phone.groupCall#9e727aad: field users: %w", err)
		}

		if headerLen > 0 {
			g.Users = make([]UserClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode phone.groupCall#9e727aad: field users: %w", err)
			}
			g.Users = append(g.Users, value)
		}
	}
	return nil
}

// GetCall returns value of Call field.
func (g *PhoneGroupCall) GetCall() (value GroupCallClass) {
	if g == nil {
		return
	}
	return g.Call
}

// GetParticipants returns value of Participants field.
func (g *PhoneGroupCall) GetParticipants() (value []GroupCallParticipant) {
	if g == nil {
		return
	}
	return g.Participants
}

// GetParticipantsNextOffset returns value of ParticipantsNextOffset field.
func (g *PhoneGroupCall) GetParticipantsNextOffset() (value string) {
	if g == nil {
		return
	}
	return g.ParticipantsNextOffset
}

// GetChats returns value of Chats field.
func (g *PhoneGroupCall) GetChats() (value []ChatClass) {
	if g == nil {
		return
	}
	return g.Chats
}

// GetUsers returns value of Users field.
func (g *PhoneGroupCall) GetUsers() (value []UserClass) {
	if g == nil {
		return
	}
	return g.Users
}
