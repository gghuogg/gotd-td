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

// IPPort represents TL type `ipPort#d433ad73`.
type IPPort struct {
	// Ipv4 field of IPPort.
	Ipv4 int
	// Port field of IPPort.
	Port int
}

// IPPortTypeID is TL type id of IPPort.
const IPPortTypeID = 0xd433ad73

// construct implements constructor of IPPortClass.
func (i IPPort) construct() IPPortClass { return &i }

// Ensuring interfaces in compile-time for IPPort.
var (
	_ bin.Encoder     = &IPPort{}
	_ bin.Decoder     = &IPPort{}
	_ bin.BareEncoder = &IPPort{}
	_ bin.BareDecoder = &IPPort{}

	_ IPPortClass = &IPPort{}
)

func (i *IPPort) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Ipv4 == 0) {
		return false
	}
	if !(i.Port == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *IPPort) String() string {
	if i == nil {
		return "IPPort(nil)"
	}
	type Alias IPPort
	return fmt.Sprintf("IPPort%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*IPPort) TypeID() uint32 {
	return IPPortTypeID
}

// TypeName returns name of type in TL schema.
func (*IPPort) TypeName() string {
	return "ipPort"
}

// TypeInfo returns info about TL type.
func (i *IPPort) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "ipPort",
		ID:   IPPortTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Ipv4",
			SchemaName: "ipv4",
		},
		{
			Name:       "Port",
			SchemaName: "port",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *IPPort) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode ipPort#d433ad73 as nil")
	}
	b.PutID(IPPortTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *IPPort) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode ipPort#d433ad73 as nil")
	}
	b.PutInt(i.Ipv4)
	b.PutInt(i.Port)
	return nil
}

// Decode implements bin.Decoder.
func (i *IPPort) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode ipPort#d433ad73 to nil")
	}
	if err := b.ConsumeID(IPPortTypeID); err != nil {
		return fmt.Errorf("unable to decode ipPort#d433ad73: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *IPPort) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode ipPort#d433ad73 to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode ipPort#d433ad73: field ipv4: %w", err)
		}
		i.Ipv4 = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode ipPort#d433ad73: field port: %w", err)
		}
		i.Port = value
	}
	return nil
}

// GetIpv4 returns value of Ipv4 field.
func (i *IPPort) GetIpv4() (value int) {
	if i == nil {
		return
	}
	return i.Ipv4
}

// GetPort returns value of Port field.
func (i *IPPort) GetPort() (value int) {
	if i == nil {
		return
	}
	return i.Port
}

// IPPortSecret represents TL type `ipPortSecret#37982646`.
type IPPortSecret struct {
	// Ipv4 field of IPPortSecret.
	Ipv4 int
	// Port field of IPPortSecret.
	Port int
	// Secret field of IPPortSecret.
	Secret []byte
}

// IPPortSecretTypeID is TL type id of IPPortSecret.
const IPPortSecretTypeID = 0x37982646

// construct implements constructor of IPPortClass.
func (i IPPortSecret) construct() IPPortClass { return &i }

// Ensuring interfaces in compile-time for IPPortSecret.
var (
	_ bin.Encoder     = &IPPortSecret{}
	_ bin.Decoder     = &IPPortSecret{}
	_ bin.BareEncoder = &IPPortSecret{}
	_ bin.BareDecoder = &IPPortSecret{}

	_ IPPortClass = &IPPortSecret{}
)

func (i *IPPortSecret) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Ipv4 == 0) {
		return false
	}
	if !(i.Port == 0) {
		return false
	}
	if !(i.Secret == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *IPPortSecret) String() string {
	if i == nil {
		return "IPPortSecret(nil)"
	}
	type Alias IPPortSecret
	return fmt.Sprintf("IPPortSecret%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*IPPortSecret) TypeID() uint32 {
	return IPPortSecretTypeID
}

// TypeName returns name of type in TL schema.
func (*IPPortSecret) TypeName() string {
	return "ipPortSecret"
}

// TypeInfo returns info about TL type.
func (i *IPPortSecret) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "ipPortSecret",
		ID:   IPPortSecretTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Ipv4",
			SchemaName: "ipv4",
		},
		{
			Name:       "Port",
			SchemaName: "port",
		},
		{
			Name:       "Secret",
			SchemaName: "secret",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *IPPortSecret) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode ipPortSecret#37982646 as nil")
	}
	b.PutID(IPPortSecretTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *IPPortSecret) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode ipPortSecret#37982646 as nil")
	}
	b.PutInt(i.Ipv4)
	b.PutInt(i.Port)
	b.PutBytes(i.Secret)
	return nil
}

// Decode implements bin.Decoder.
func (i *IPPortSecret) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode ipPortSecret#37982646 to nil")
	}
	if err := b.ConsumeID(IPPortSecretTypeID); err != nil {
		return fmt.Errorf("unable to decode ipPortSecret#37982646: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *IPPortSecret) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode ipPortSecret#37982646 to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode ipPortSecret#37982646: field ipv4: %w", err)
		}
		i.Ipv4 = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode ipPortSecret#37982646: field port: %w", err)
		}
		i.Port = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode ipPortSecret#37982646: field secret: %w", err)
		}
		i.Secret = value
	}
	return nil
}

// GetIpv4 returns value of Ipv4 field.
func (i *IPPortSecret) GetIpv4() (value int) {
	if i == nil {
		return
	}
	return i.Ipv4
}

// GetPort returns value of Port field.
func (i *IPPortSecret) GetPort() (value int) {
	if i == nil {
		return
	}
	return i.Port
}

// GetSecret returns value of Secret field.
func (i *IPPortSecret) GetSecret() (value []byte) {
	if i == nil {
		return
	}
	return i.Secret
}

// IPPortClassName is schema name of IPPortClass.
const IPPortClassName = "IpPort"

// IPPortClass represents IpPort generic type.
//
// Example:
//
//	g, err := td.DecodeIPPort(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *td.IPPort: // ipPort#d433ad73
//	case *td.IPPortSecret: // ipPortSecret#37982646
//	default: panic(v)
//	}
type IPPortClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() IPPortClass

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

	// Ipv4 field of IPPort.
	GetIpv4() (value int)
	// Port field of IPPort.
	GetPort() (value int)
}

// DecodeIPPort implements binary de-serialization for IPPortClass.
func DecodeIPPort(buf *bin.Buffer) (IPPortClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case IPPortTypeID:
		// Decoding ipPort#d433ad73.
		v := IPPort{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode IPPortClass: %w", err)
		}
		return &v, nil
	case IPPortSecretTypeID:
		// Decoding ipPortSecret#37982646.
		v := IPPortSecret{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode IPPortClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode IPPortClass: %w", bin.NewUnexpectedID(id))
	}
}

// IPPort boxes the IPPortClass providing a helper.
type IPPortBox struct {
	IpPort IPPortClass
}

// Decode implements bin.Decoder for IPPortBox.
func (b *IPPortBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode IPPortBox to nil")
	}
	v, err := DecodeIPPort(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.IpPort = v
	return nil
}

// Encode implements bin.Encode for IPPortBox.
func (b *IPPortBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.IpPort == nil {
		return fmt.Errorf("unable to encode IPPortClass as nil")
	}
	return b.IpPort.Encode(buf)
}
