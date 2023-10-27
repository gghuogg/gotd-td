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

// ConnectedWebsites represents TL type `connectedWebsites#f0c8b5ea`.
type ConnectedWebsites struct {
	// List of connected websites
	Websites []ConnectedWebsite
}

// ConnectedWebsitesTypeID is TL type id of ConnectedWebsites.
const ConnectedWebsitesTypeID = 0xf0c8b5ea

// Ensuring interfaces in compile-time for ConnectedWebsites.
var (
	_ bin.Encoder     = &ConnectedWebsites{}
	_ bin.Decoder     = &ConnectedWebsites{}
	_ bin.BareEncoder = &ConnectedWebsites{}
	_ bin.BareDecoder = &ConnectedWebsites{}
)

func (c *ConnectedWebsites) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Websites == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ConnectedWebsites) String() string {
	if c == nil {
		return "ConnectedWebsites(nil)"
	}
	type Alias ConnectedWebsites
	return fmt.Sprintf("ConnectedWebsites%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ConnectedWebsites) TypeID() uint32 {
	return ConnectedWebsitesTypeID
}

// TypeName returns name of type in TL schema.
func (*ConnectedWebsites) TypeName() string {
	return "connectedWebsites"
}

// TypeInfo returns info about TL type.
func (c *ConnectedWebsites) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "connectedWebsites",
		ID:   ConnectedWebsitesTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Websites",
			SchemaName: "websites",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ConnectedWebsites) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode connectedWebsites#f0c8b5ea as nil")
	}
	b.PutID(ConnectedWebsitesTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ConnectedWebsites) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode connectedWebsites#f0c8b5ea as nil")
	}
	b.PutInt(len(c.Websites))
	for idx, v := range c.Websites {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare connectedWebsites#f0c8b5ea: field websites element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ConnectedWebsites) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode connectedWebsites#f0c8b5ea to nil")
	}
	if err := b.ConsumeID(ConnectedWebsitesTypeID); err != nil {
		return fmt.Errorf("unable to decode connectedWebsites#f0c8b5ea: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ConnectedWebsites) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode connectedWebsites#f0c8b5ea to nil")
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode connectedWebsites#f0c8b5ea: field websites: %w", err)
		}

		if headerLen > 0 {
			c.Websites = make([]ConnectedWebsite, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value ConnectedWebsite
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare connectedWebsites#f0c8b5ea: field websites: %w", err)
			}
			c.Websites = append(c.Websites, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *ConnectedWebsites) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode connectedWebsites#f0c8b5ea as nil")
	}
	b.ObjStart()
	b.PutID("connectedWebsites")
	b.Comma()
	b.FieldStart("websites")
	b.ArrStart()
	for idx, v := range c.Websites {
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode connectedWebsites#f0c8b5ea: field websites element with index %d: %w", idx, err)
		}
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *ConnectedWebsites) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode connectedWebsites#f0c8b5ea to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("connectedWebsites"); err != nil {
				return fmt.Errorf("unable to decode connectedWebsites#f0c8b5ea: %w", err)
			}
		case "websites":
			if err := b.Arr(func(b tdjson.Decoder) error {
				var value ConnectedWebsite
				if err := value.DecodeTDLibJSON(b); err != nil {
					return fmt.Errorf("unable to decode connectedWebsites#f0c8b5ea: field websites: %w", err)
				}
				c.Websites = append(c.Websites, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode connectedWebsites#f0c8b5ea: field websites: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetWebsites returns value of Websites field.
func (c *ConnectedWebsites) GetWebsites() (value []ConnectedWebsite) {
	if c == nil {
		return
	}
	return c.Websites
}
