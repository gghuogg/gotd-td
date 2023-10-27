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

// ExportedMessageLink represents TL type `exportedMessageLink#5dab1af4`.
type ExportedMessageLink struct {
	// Link field of ExportedMessageLink.
	Link string
	// HTML field of ExportedMessageLink.
	HTML string
}

// ExportedMessageLinkTypeID is TL type id of ExportedMessageLink.
const ExportedMessageLinkTypeID = 0x5dab1af4

// Ensuring interfaces in compile-time for ExportedMessageLink.
var (
	_ bin.Encoder     = &ExportedMessageLink{}
	_ bin.Decoder     = &ExportedMessageLink{}
	_ bin.BareEncoder = &ExportedMessageLink{}
	_ bin.BareDecoder = &ExportedMessageLink{}
)

func (e *ExportedMessageLink) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.Link == "") {
		return false
	}
	if !(e.HTML == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *ExportedMessageLink) String() string {
	if e == nil {
		return "ExportedMessageLink(nil)"
	}
	type Alias ExportedMessageLink
	return fmt.Sprintf("ExportedMessageLink%+v", Alias(*e))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ExportedMessageLink) TypeID() uint32 {
	return ExportedMessageLinkTypeID
}

// TypeName returns name of type in TL schema.
func (*ExportedMessageLink) TypeName() string {
	return "exportedMessageLink"
}

// TypeInfo returns info about TL type.
func (e *ExportedMessageLink) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "exportedMessageLink",
		ID:   ExportedMessageLinkTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Link",
			SchemaName: "link",
		},
		{
			Name:       "HTML",
			SchemaName: "html",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *ExportedMessageLink) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode exportedMessageLink#5dab1af4 as nil")
	}
	b.PutID(ExportedMessageLinkTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *ExportedMessageLink) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode exportedMessageLink#5dab1af4 as nil")
	}
	b.PutString(e.Link)
	b.PutString(e.HTML)
	return nil
}

// Decode implements bin.Decoder.
func (e *ExportedMessageLink) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode exportedMessageLink#5dab1af4 to nil")
	}
	if err := b.ConsumeID(ExportedMessageLinkTypeID); err != nil {
		return fmt.Errorf("unable to decode exportedMessageLink#5dab1af4: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *ExportedMessageLink) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode exportedMessageLink#5dab1af4 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode exportedMessageLink#5dab1af4: field link: %w", err)
		}
		e.Link = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode exportedMessageLink#5dab1af4: field html: %w", err)
		}
		e.HTML = value
	}
	return nil
}

// GetLink returns value of Link field.
func (e *ExportedMessageLink) GetLink() (value string) {
	if e == nil {
		return
	}
	return e.Link
}

// GetHTML returns value of HTML field.
func (e *ExportedMessageLink) GetHTML() (value string) {
	if e == nil {
		return
	}
	return e.HTML
}
