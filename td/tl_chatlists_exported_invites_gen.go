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

// ChatlistsExportedInvites represents TL type `chatlists.exportedInvites#10ab6dc7`.
type ChatlistsExportedInvites struct {
	// Invites field of ChatlistsExportedInvites.
	Invites []ExportedChatlistInvite
	// Chats field of ChatlistsExportedInvites.
	Chats []ChatClass
	// Users field of ChatlistsExportedInvites.
	Users []UserClass
}

// ChatlistsExportedInvitesTypeID is TL type id of ChatlistsExportedInvites.
const ChatlistsExportedInvitesTypeID = 0x10ab6dc7

// Ensuring interfaces in compile-time for ChatlistsExportedInvites.
var (
	_ bin.Encoder     = &ChatlistsExportedInvites{}
	_ bin.Decoder     = &ChatlistsExportedInvites{}
	_ bin.BareEncoder = &ChatlistsExportedInvites{}
	_ bin.BareDecoder = &ChatlistsExportedInvites{}
)

func (e *ChatlistsExportedInvites) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.Invites == nil) {
		return false
	}
	if !(e.Chats == nil) {
		return false
	}
	if !(e.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *ChatlistsExportedInvites) String() string {
	if e == nil {
		return "ChatlistsExportedInvites(nil)"
	}
	type Alias ChatlistsExportedInvites
	return fmt.Sprintf("ChatlistsExportedInvites%+v", Alias(*e))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatlistsExportedInvites) TypeID() uint32 {
	return ChatlistsExportedInvitesTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatlistsExportedInvites) TypeName() string {
	return "chatlists.exportedInvites"
}

// TypeInfo returns info about TL type.
func (e *ChatlistsExportedInvites) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatlists.exportedInvites",
		ID:   ChatlistsExportedInvitesTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Invites",
			SchemaName: "invites",
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
func (e *ChatlistsExportedInvites) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode chatlists.exportedInvites#10ab6dc7 as nil")
	}
	b.PutID(ChatlistsExportedInvitesTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *ChatlistsExportedInvites) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode chatlists.exportedInvites#10ab6dc7 as nil")
	}
	b.PutVectorHeader(len(e.Invites))
	for idx, v := range e.Invites {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode chatlists.exportedInvites#10ab6dc7: field invites element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(e.Chats))
	for idx, v := range e.Chats {
		if v == nil {
			return fmt.Errorf("unable to encode chatlists.exportedInvites#10ab6dc7: field chats element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode chatlists.exportedInvites#10ab6dc7: field chats element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(e.Users))
	for idx, v := range e.Users {
		if v == nil {
			return fmt.Errorf("unable to encode chatlists.exportedInvites#10ab6dc7: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode chatlists.exportedInvites#10ab6dc7: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (e *ChatlistsExportedInvites) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode chatlists.exportedInvites#10ab6dc7 to nil")
	}
	if err := b.ConsumeID(ChatlistsExportedInvitesTypeID); err != nil {
		return fmt.Errorf("unable to decode chatlists.exportedInvites#10ab6dc7: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *ChatlistsExportedInvites) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode chatlists.exportedInvites#10ab6dc7 to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode chatlists.exportedInvites#10ab6dc7: field invites: %w", err)
		}

		if headerLen > 0 {
			e.Invites = make([]ExportedChatlistInvite, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value ExportedChatlistInvite
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode chatlists.exportedInvites#10ab6dc7: field invites: %w", err)
			}
			e.Invites = append(e.Invites, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode chatlists.exportedInvites#10ab6dc7: field chats: %w", err)
		}

		if headerLen > 0 {
			e.Chats = make([]ChatClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChat(b)
			if err != nil {
				return fmt.Errorf("unable to decode chatlists.exportedInvites#10ab6dc7: field chats: %w", err)
			}
			e.Chats = append(e.Chats, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode chatlists.exportedInvites#10ab6dc7: field users: %w", err)
		}

		if headerLen > 0 {
			e.Users = make([]UserClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode chatlists.exportedInvites#10ab6dc7: field users: %w", err)
			}
			e.Users = append(e.Users, value)
		}
	}
	return nil
}

// GetInvites returns value of Invites field.
func (e *ChatlistsExportedInvites) GetInvites() (value []ExportedChatlistInvite) {
	if e == nil {
		return
	}
	return e.Invites
}

// GetChats returns value of Chats field.
func (e *ChatlistsExportedInvites) GetChats() (value []ChatClass) {
	if e == nil {
		return
	}
	return e.Chats
}

// GetUsers returns value of Users field.
func (e *ChatlistsExportedInvites) GetUsers() (value []UserClass) {
	if e == nil {
		return
	}
	return e.Users
}
