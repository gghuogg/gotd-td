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

// MessagesForumTopics represents TL type `messages.forumTopics#367617d3`.
type MessagesForumTopics struct {
	// Flags field of MessagesForumTopics.
	Flags bin.Fields
	// OrderByCreateDate field of MessagesForumTopics.
	OrderByCreateDate bool
	// Count field of MessagesForumTopics.
	Count int
	// Topics field of MessagesForumTopics.
	Topics []ForumTopicClass
	// Messages field of MessagesForumTopics.
	Messages []MessageClass
	// Chats field of MessagesForumTopics.
	Chats []ChatClass
	// Users field of MessagesForumTopics.
	Users []UserClass
	// Pts field of MessagesForumTopics.
	Pts int
}

// MessagesForumTopicsTypeID is TL type id of MessagesForumTopics.
const MessagesForumTopicsTypeID = 0x367617d3

// Ensuring interfaces in compile-time for MessagesForumTopics.
var (
	_ bin.Encoder     = &MessagesForumTopics{}
	_ bin.Decoder     = &MessagesForumTopics{}
	_ bin.BareEncoder = &MessagesForumTopics{}
	_ bin.BareDecoder = &MessagesForumTopics{}
)

func (f *MessagesForumTopics) Zero() bool {
	if f == nil {
		return true
	}
	if !(f.Flags.Zero()) {
		return false
	}
	if !(f.OrderByCreateDate == false) {
		return false
	}
	if !(f.Count == 0) {
		return false
	}
	if !(f.Topics == nil) {
		return false
	}
	if !(f.Messages == nil) {
		return false
	}
	if !(f.Chats == nil) {
		return false
	}
	if !(f.Users == nil) {
		return false
	}
	if !(f.Pts == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (f *MessagesForumTopics) String() string {
	if f == nil {
		return "MessagesForumTopics(nil)"
	}
	type Alias MessagesForumTopics
	return fmt.Sprintf("MessagesForumTopics%+v", Alias(*f))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesForumTopics) TypeID() uint32 {
	return MessagesForumTopicsTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesForumTopics) TypeName() string {
	return "messages.forumTopics"
}

// TypeInfo returns info about TL type.
func (f *MessagesForumTopics) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.forumTopics",
		ID:   MessagesForumTopicsTypeID,
	}
	if f == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "OrderByCreateDate",
			SchemaName: "order_by_create_date",
			Null:       !f.Flags.Has(0),
		},
		{
			Name:       "Count",
			SchemaName: "count",
		},
		{
			Name:       "Topics",
			SchemaName: "topics",
		},
		{
			Name:       "Messages",
			SchemaName: "messages",
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
			Name:       "Pts",
			SchemaName: "pts",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (f *MessagesForumTopics) SetFlags() {
	if !(f.OrderByCreateDate == false) {
		f.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (f *MessagesForumTopics) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode messages.forumTopics#367617d3 as nil")
	}
	b.PutID(MessagesForumTopicsTypeID)
	return f.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (f *MessagesForumTopics) EncodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode messages.forumTopics#367617d3 as nil")
	}
	f.SetFlags()
	if err := f.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.forumTopics#367617d3: field flags: %w", err)
	}
	b.PutInt(f.Count)
	b.PutVectorHeader(len(f.Topics))
	for idx, v := range f.Topics {
		if v == nil {
			return fmt.Errorf("unable to encode messages.forumTopics#367617d3: field topics element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.forumTopics#367617d3: field topics element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(f.Messages))
	for idx, v := range f.Messages {
		if v == nil {
			return fmt.Errorf("unable to encode messages.forumTopics#367617d3: field messages element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.forumTopics#367617d3: field messages element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(f.Chats))
	for idx, v := range f.Chats {
		if v == nil {
			return fmt.Errorf("unable to encode messages.forumTopics#367617d3: field chats element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.forumTopics#367617d3: field chats element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(f.Users))
	for idx, v := range f.Users {
		if v == nil {
			return fmt.Errorf("unable to encode messages.forumTopics#367617d3: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.forumTopics#367617d3: field users element with index %d: %w", idx, err)
		}
	}
	b.PutInt(f.Pts)
	return nil
}

// Decode implements bin.Decoder.
func (f *MessagesForumTopics) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode messages.forumTopics#367617d3 to nil")
	}
	if err := b.ConsumeID(MessagesForumTopicsTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.forumTopics#367617d3: %w", err)
	}
	return f.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (f *MessagesForumTopics) DecodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode messages.forumTopics#367617d3 to nil")
	}
	{
		if err := f.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.forumTopics#367617d3: field flags: %w", err)
		}
	}
	f.OrderByCreateDate = f.Flags.Has(0)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.forumTopics#367617d3: field count: %w", err)
		}
		f.Count = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.forumTopics#367617d3: field topics: %w", err)
		}

		if headerLen > 0 {
			f.Topics = make([]ForumTopicClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeForumTopic(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.forumTopics#367617d3: field topics: %w", err)
			}
			f.Topics = append(f.Topics, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.forumTopics#367617d3: field messages: %w", err)
		}

		if headerLen > 0 {
			f.Messages = make([]MessageClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessage(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.forumTopics#367617d3: field messages: %w", err)
			}
			f.Messages = append(f.Messages, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.forumTopics#367617d3: field chats: %w", err)
		}

		if headerLen > 0 {
			f.Chats = make([]ChatClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChat(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.forumTopics#367617d3: field chats: %w", err)
			}
			f.Chats = append(f.Chats, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.forumTopics#367617d3: field users: %w", err)
		}

		if headerLen > 0 {
			f.Users = make([]UserClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.forumTopics#367617d3: field users: %w", err)
			}
			f.Users = append(f.Users, value)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.forumTopics#367617d3: field pts: %w", err)
		}
		f.Pts = value
	}
	return nil
}

// SetOrderByCreateDate sets value of OrderByCreateDate conditional field.
func (f *MessagesForumTopics) SetOrderByCreateDate(value bool) {
	if value {
		f.Flags.Set(0)
		f.OrderByCreateDate = true
	} else {
		f.Flags.Unset(0)
		f.OrderByCreateDate = false
	}
}

// GetOrderByCreateDate returns value of OrderByCreateDate conditional field.
func (f *MessagesForumTopics) GetOrderByCreateDate() (value bool) {
	if f == nil {
		return
	}
	return f.Flags.Has(0)
}

// GetCount returns value of Count field.
func (f *MessagesForumTopics) GetCount() (value int) {
	if f == nil {
		return
	}
	return f.Count
}

// GetTopics returns value of Topics field.
func (f *MessagesForumTopics) GetTopics() (value []ForumTopicClass) {
	if f == nil {
		return
	}
	return f.Topics
}

// GetMessages returns value of Messages field.
func (f *MessagesForumTopics) GetMessages() (value []MessageClass) {
	if f == nil {
		return
	}
	return f.Messages
}

// GetChats returns value of Chats field.
func (f *MessagesForumTopics) GetChats() (value []ChatClass) {
	if f == nil {
		return
	}
	return f.Chats
}

// GetUsers returns value of Users field.
func (f *MessagesForumTopics) GetUsers() (value []UserClass) {
	if f == nil {
		return
	}
	return f.Users
}

// GetPts returns value of Pts field.
func (f *MessagesForumTopics) GetPts() (value int) {
	if f == nil {
		return
	}
	return f.Pts
}
