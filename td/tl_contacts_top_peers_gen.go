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

// ContactsTopPeersNotModified represents TL type `contacts.topPeersNotModified#de266ef5`.
type ContactsTopPeersNotModified struct {
}

// ContactsTopPeersNotModifiedTypeID is TL type id of ContactsTopPeersNotModified.
const ContactsTopPeersNotModifiedTypeID = 0xde266ef5

// construct implements constructor of ContactsTopPeersClass.
func (t ContactsTopPeersNotModified) construct() ContactsTopPeersClass { return &t }

// Ensuring interfaces in compile-time for ContactsTopPeersNotModified.
var (
	_ bin.Encoder     = &ContactsTopPeersNotModified{}
	_ bin.Decoder     = &ContactsTopPeersNotModified{}
	_ bin.BareEncoder = &ContactsTopPeersNotModified{}
	_ bin.BareDecoder = &ContactsTopPeersNotModified{}

	_ ContactsTopPeersClass = &ContactsTopPeersNotModified{}
)

func (t *ContactsTopPeersNotModified) Zero() bool {
	if t == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (t *ContactsTopPeersNotModified) String() string {
	if t == nil {
		return "ContactsTopPeersNotModified(nil)"
	}
	type Alias ContactsTopPeersNotModified
	return fmt.Sprintf("ContactsTopPeersNotModified%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ContactsTopPeersNotModified) TypeID() uint32 {
	return ContactsTopPeersNotModifiedTypeID
}

// TypeName returns name of type in TL schema.
func (*ContactsTopPeersNotModified) TypeName() string {
	return "contacts.topPeersNotModified"
}

// TypeInfo returns info about TL type.
func (t *ContactsTopPeersNotModified) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "contacts.topPeersNotModified",
		ID:   ContactsTopPeersNotModifiedTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (t *ContactsTopPeersNotModified) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode contacts.topPeersNotModified#de266ef5 as nil")
	}
	b.PutID(ContactsTopPeersNotModifiedTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *ContactsTopPeersNotModified) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode contacts.topPeersNotModified#de266ef5 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *ContactsTopPeersNotModified) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode contacts.topPeersNotModified#de266ef5 to nil")
	}
	if err := b.ConsumeID(ContactsTopPeersNotModifiedTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.topPeersNotModified#de266ef5: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *ContactsTopPeersNotModified) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode contacts.topPeersNotModified#de266ef5 to nil")
	}
	return nil
}

// ContactsTopPeers represents TL type `contacts.topPeers#70b772a8`.
type ContactsTopPeers struct {
	// Categories field of ContactsTopPeers.
	Categories []TopPeerCategoryPeers
	// Chats field of ContactsTopPeers.
	Chats []ChatClass
	// Users field of ContactsTopPeers.
	Users []UserClass
}

// ContactsTopPeersTypeID is TL type id of ContactsTopPeers.
const ContactsTopPeersTypeID = 0x70b772a8

// construct implements constructor of ContactsTopPeersClass.
func (t ContactsTopPeers) construct() ContactsTopPeersClass { return &t }

// Ensuring interfaces in compile-time for ContactsTopPeers.
var (
	_ bin.Encoder     = &ContactsTopPeers{}
	_ bin.Decoder     = &ContactsTopPeers{}
	_ bin.BareEncoder = &ContactsTopPeers{}
	_ bin.BareDecoder = &ContactsTopPeers{}

	_ ContactsTopPeersClass = &ContactsTopPeers{}
)

func (t *ContactsTopPeers) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Categories == nil) {
		return false
	}
	if !(t.Chats == nil) {
		return false
	}
	if !(t.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *ContactsTopPeers) String() string {
	if t == nil {
		return "ContactsTopPeers(nil)"
	}
	type Alias ContactsTopPeers
	return fmt.Sprintf("ContactsTopPeers%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ContactsTopPeers) TypeID() uint32 {
	return ContactsTopPeersTypeID
}

// TypeName returns name of type in TL schema.
func (*ContactsTopPeers) TypeName() string {
	return "contacts.topPeers"
}

// TypeInfo returns info about TL type.
func (t *ContactsTopPeers) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "contacts.topPeers",
		ID:   ContactsTopPeersTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Categories",
			SchemaName: "categories",
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
func (t *ContactsTopPeers) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode contacts.topPeers#70b772a8 as nil")
	}
	b.PutID(ContactsTopPeersTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *ContactsTopPeers) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode contacts.topPeers#70b772a8 as nil")
	}
	b.PutVectorHeader(len(t.Categories))
	for idx, v := range t.Categories {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode contacts.topPeers#70b772a8: field categories element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(t.Chats))
	for idx, v := range t.Chats {
		if v == nil {
			return fmt.Errorf("unable to encode contacts.topPeers#70b772a8: field chats element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode contacts.topPeers#70b772a8: field chats element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(t.Users))
	for idx, v := range t.Users {
		if v == nil {
			return fmt.Errorf("unable to encode contacts.topPeers#70b772a8: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode contacts.topPeers#70b772a8: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *ContactsTopPeers) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode contacts.topPeers#70b772a8 to nil")
	}
	if err := b.ConsumeID(ContactsTopPeersTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.topPeers#70b772a8: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *ContactsTopPeers) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode contacts.topPeers#70b772a8 to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.topPeers#70b772a8: field categories: %w", err)
		}

		if headerLen > 0 {
			t.Categories = make([]TopPeerCategoryPeers, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value TopPeerCategoryPeers
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode contacts.topPeers#70b772a8: field categories: %w", err)
			}
			t.Categories = append(t.Categories, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.topPeers#70b772a8: field chats: %w", err)
		}

		if headerLen > 0 {
			t.Chats = make([]ChatClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChat(b)
			if err != nil {
				return fmt.Errorf("unable to decode contacts.topPeers#70b772a8: field chats: %w", err)
			}
			t.Chats = append(t.Chats, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.topPeers#70b772a8: field users: %w", err)
		}

		if headerLen > 0 {
			t.Users = make([]UserClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode contacts.topPeers#70b772a8: field users: %w", err)
			}
			t.Users = append(t.Users, value)
		}
	}
	return nil
}

// GetCategories returns value of Categories field.
func (t *ContactsTopPeers) GetCategories() (value []TopPeerCategoryPeers) {
	if t == nil {
		return
	}
	return t.Categories
}

// GetChats returns value of Chats field.
func (t *ContactsTopPeers) GetChats() (value []ChatClass) {
	if t == nil {
		return
	}
	return t.Chats
}

// GetUsers returns value of Users field.
func (t *ContactsTopPeers) GetUsers() (value []UserClass) {
	if t == nil {
		return
	}
	return t.Users
}

// ContactsTopPeersDisabled represents TL type `contacts.topPeersDisabled#b52c939d`.
type ContactsTopPeersDisabled struct {
}

// ContactsTopPeersDisabledTypeID is TL type id of ContactsTopPeersDisabled.
const ContactsTopPeersDisabledTypeID = 0xb52c939d

// construct implements constructor of ContactsTopPeersClass.
func (t ContactsTopPeersDisabled) construct() ContactsTopPeersClass { return &t }

// Ensuring interfaces in compile-time for ContactsTopPeersDisabled.
var (
	_ bin.Encoder     = &ContactsTopPeersDisabled{}
	_ bin.Decoder     = &ContactsTopPeersDisabled{}
	_ bin.BareEncoder = &ContactsTopPeersDisabled{}
	_ bin.BareDecoder = &ContactsTopPeersDisabled{}

	_ ContactsTopPeersClass = &ContactsTopPeersDisabled{}
)

func (t *ContactsTopPeersDisabled) Zero() bool {
	if t == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (t *ContactsTopPeersDisabled) String() string {
	if t == nil {
		return "ContactsTopPeersDisabled(nil)"
	}
	type Alias ContactsTopPeersDisabled
	return fmt.Sprintf("ContactsTopPeersDisabled%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ContactsTopPeersDisabled) TypeID() uint32 {
	return ContactsTopPeersDisabledTypeID
}

// TypeName returns name of type in TL schema.
func (*ContactsTopPeersDisabled) TypeName() string {
	return "contacts.topPeersDisabled"
}

// TypeInfo returns info about TL type.
func (t *ContactsTopPeersDisabled) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "contacts.topPeersDisabled",
		ID:   ContactsTopPeersDisabledTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (t *ContactsTopPeersDisabled) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode contacts.topPeersDisabled#b52c939d as nil")
	}
	b.PutID(ContactsTopPeersDisabledTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *ContactsTopPeersDisabled) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode contacts.topPeersDisabled#b52c939d as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *ContactsTopPeersDisabled) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode contacts.topPeersDisabled#b52c939d to nil")
	}
	if err := b.ConsumeID(ContactsTopPeersDisabledTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.topPeersDisabled#b52c939d: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *ContactsTopPeersDisabled) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode contacts.topPeersDisabled#b52c939d to nil")
	}
	return nil
}

// ContactsTopPeersClassName is schema name of ContactsTopPeersClass.
const ContactsTopPeersClassName = "contacts.TopPeers"

// ContactsTopPeersClass represents contacts.TopPeers generic type.
//
// Example:
//
//	g, err := td.DecodeContactsTopPeers(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *td.ContactsTopPeersNotModified: // contacts.topPeersNotModified#de266ef5
//	case *td.ContactsTopPeers: // contacts.topPeers#70b772a8
//	case *td.ContactsTopPeersDisabled: // contacts.topPeersDisabled#b52c939d
//	default: panic(v)
//	}
type ContactsTopPeersClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() ContactsTopPeersClass

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
}

// DecodeContactsTopPeers implements binary de-serialization for ContactsTopPeersClass.
func DecodeContactsTopPeers(buf *bin.Buffer) (ContactsTopPeersClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case ContactsTopPeersNotModifiedTypeID:
		// Decoding contacts.topPeersNotModified#de266ef5.
		v := ContactsTopPeersNotModified{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ContactsTopPeersClass: %w", err)
		}
		return &v, nil
	case ContactsTopPeersTypeID:
		// Decoding contacts.topPeers#70b772a8.
		v := ContactsTopPeers{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ContactsTopPeersClass: %w", err)
		}
		return &v, nil
	case ContactsTopPeersDisabledTypeID:
		// Decoding contacts.topPeersDisabled#b52c939d.
		v := ContactsTopPeersDisabled{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ContactsTopPeersClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode ContactsTopPeersClass: %w", bin.NewUnexpectedID(id))
	}
}

// ContactsTopPeers boxes the ContactsTopPeersClass providing a helper.
type ContactsTopPeersBox struct {
	TopPeers ContactsTopPeersClass
}

// Decode implements bin.Decoder for ContactsTopPeersBox.
func (b *ContactsTopPeersBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode ContactsTopPeersBox to nil")
	}
	v, err := DecodeContactsTopPeers(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.TopPeers = v
	return nil
}

// Encode implements bin.Encode for ContactsTopPeersBox.
func (b *ContactsTopPeersBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.TopPeers == nil {
		return fmt.Errorf("unable to encode ContactsTopPeersClass as nil")
	}
	return b.TopPeers.Encode(buf)
}
