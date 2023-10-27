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

// ContactsGetBlockedRequest represents TL type `contacts.getBlocked#9a868f80`.
type ContactsGetBlockedRequest struct {
	// Flags field of ContactsGetBlockedRequest.
	Flags bin.Fields
	// MyStoriesFrom field of ContactsGetBlockedRequest.
	MyStoriesFrom bool
	// Offset field of ContactsGetBlockedRequest.
	Offset int
	// Limit field of ContactsGetBlockedRequest.
	Limit int
}

// ContactsGetBlockedRequestTypeID is TL type id of ContactsGetBlockedRequest.
const ContactsGetBlockedRequestTypeID = 0x9a868f80

// Ensuring interfaces in compile-time for ContactsGetBlockedRequest.
var (
	_ bin.Encoder     = &ContactsGetBlockedRequest{}
	_ bin.Decoder     = &ContactsGetBlockedRequest{}
	_ bin.BareEncoder = &ContactsGetBlockedRequest{}
	_ bin.BareDecoder = &ContactsGetBlockedRequest{}
)

func (g *ContactsGetBlockedRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Flags.Zero()) {
		return false
	}
	if !(g.MyStoriesFrom == false) {
		return false
	}
	if !(g.Offset == 0) {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *ContactsGetBlockedRequest) String() string {
	if g == nil {
		return "ContactsGetBlockedRequest(nil)"
	}
	type Alias ContactsGetBlockedRequest
	return fmt.Sprintf("ContactsGetBlockedRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ContactsGetBlockedRequest) TypeID() uint32 {
	return ContactsGetBlockedRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ContactsGetBlockedRequest) TypeName() string {
	return "contacts.getBlocked"
}

// TypeInfo returns info about TL type.
func (g *ContactsGetBlockedRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "contacts.getBlocked",
		ID:   ContactsGetBlockedRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "MyStoriesFrom",
			SchemaName: "my_stories_from",
			Null:       !g.Flags.Has(0),
		},
		{
			Name:       "Offset",
			SchemaName: "offset",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (g *ContactsGetBlockedRequest) SetFlags() {
	if !(g.MyStoriesFrom == false) {
		g.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (g *ContactsGetBlockedRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode contacts.getBlocked#9a868f80 as nil")
	}
	b.PutID(ContactsGetBlockedRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *ContactsGetBlockedRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode contacts.getBlocked#9a868f80 as nil")
	}
	g.SetFlags()
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode contacts.getBlocked#9a868f80: field flags: %w", err)
	}
	b.PutInt(g.Offset)
	b.PutInt(g.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (g *ContactsGetBlockedRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode contacts.getBlocked#9a868f80 to nil")
	}
	if err := b.ConsumeID(ContactsGetBlockedRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.getBlocked#9a868f80: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *ContactsGetBlockedRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode contacts.getBlocked#9a868f80 to nil")
	}
	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode contacts.getBlocked#9a868f80: field flags: %w", err)
		}
	}
	g.MyStoriesFrom = g.Flags.Has(0)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.getBlocked#9a868f80: field offset: %w", err)
		}
		g.Offset = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.getBlocked#9a868f80: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// SetMyStoriesFrom sets value of MyStoriesFrom conditional field.
func (g *ContactsGetBlockedRequest) SetMyStoriesFrom(value bool) {
	if value {
		g.Flags.Set(0)
		g.MyStoriesFrom = true
	} else {
		g.Flags.Unset(0)
		g.MyStoriesFrom = false
	}
}

// GetMyStoriesFrom returns value of MyStoriesFrom conditional field.
func (g *ContactsGetBlockedRequest) GetMyStoriesFrom() (value bool) {
	if g == nil {
		return
	}
	return g.Flags.Has(0)
}

// GetOffset returns value of Offset field.
func (g *ContactsGetBlockedRequest) GetOffset() (value int) {
	if g == nil {
		return
	}
	return g.Offset
}

// GetLimit returns value of Limit field.
func (g *ContactsGetBlockedRequest) GetLimit() (value int) {
	if g == nil {
		return
	}
	return g.Limit
}

// ContactsGetBlocked invokes method contacts.getBlocked#9a868f80 returning error if any.
func (c *Client) ContactsGetBlocked(ctx context.Context, request *ContactsGetBlockedRequest) (ContactsBlockedClass, error) {
	var result ContactsBlockedBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Blocked, nil
}
