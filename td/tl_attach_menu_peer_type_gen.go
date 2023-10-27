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

// AttachMenuPeerTypeSameBotPM represents TL type `attachMenuPeerTypeSameBotPM#7d6be90e`.
type AttachMenuPeerTypeSameBotPM struct {
}

// AttachMenuPeerTypeSameBotPMTypeID is TL type id of AttachMenuPeerTypeSameBotPM.
const AttachMenuPeerTypeSameBotPMTypeID = 0x7d6be90e

// construct implements constructor of AttachMenuPeerTypeClass.
func (a AttachMenuPeerTypeSameBotPM) construct() AttachMenuPeerTypeClass { return &a }

// Ensuring interfaces in compile-time for AttachMenuPeerTypeSameBotPM.
var (
	_ bin.Encoder     = &AttachMenuPeerTypeSameBotPM{}
	_ bin.Decoder     = &AttachMenuPeerTypeSameBotPM{}
	_ bin.BareEncoder = &AttachMenuPeerTypeSameBotPM{}
	_ bin.BareDecoder = &AttachMenuPeerTypeSameBotPM{}

	_ AttachMenuPeerTypeClass = &AttachMenuPeerTypeSameBotPM{}
)

func (a *AttachMenuPeerTypeSameBotPM) Zero() bool {
	if a == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (a *AttachMenuPeerTypeSameBotPM) String() string {
	if a == nil {
		return "AttachMenuPeerTypeSameBotPM(nil)"
	}
	type Alias AttachMenuPeerTypeSameBotPM
	return fmt.Sprintf("AttachMenuPeerTypeSameBotPM%+v", Alias(*a))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AttachMenuPeerTypeSameBotPM) TypeID() uint32 {
	return AttachMenuPeerTypeSameBotPMTypeID
}

// TypeName returns name of type in TL schema.
func (*AttachMenuPeerTypeSameBotPM) TypeName() string {
	return "attachMenuPeerTypeSameBotPM"
}

// TypeInfo returns info about TL type.
func (a *AttachMenuPeerTypeSameBotPM) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "attachMenuPeerTypeSameBotPM",
		ID:   AttachMenuPeerTypeSameBotPMTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (a *AttachMenuPeerTypeSameBotPM) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode attachMenuPeerTypeSameBotPM#7d6be90e as nil")
	}
	b.PutID(AttachMenuPeerTypeSameBotPMTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *AttachMenuPeerTypeSameBotPM) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode attachMenuPeerTypeSameBotPM#7d6be90e as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (a *AttachMenuPeerTypeSameBotPM) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode attachMenuPeerTypeSameBotPM#7d6be90e to nil")
	}
	if err := b.ConsumeID(AttachMenuPeerTypeSameBotPMTypeID); err != nil {
		return fmt.Errorf("unable to decode attachMenuPeerTypeSameBotPM#7d6be90e: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *AttachMenuPeerTypeSameBotPM) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode attachMenuPeerTypeSameBotPM#7d6be90e to nil")
	}
	return nil
}

// AttachMenuPeerTypeBotPM represents TL type `attachMenuPeerTypeBotPM#c32bfa1a`.
type AttachMenuPeerTypeBotPM struct {
}

// AttachMenuPeerTypeBotPMTypeID is TL type id of AttachMenuPeerTypeBotPM.
const AttachMenuPeerTypeBotPMTypeID = 0xc32bfa1a

// construct implements constructor of AttachMenuPeerTypeClass.
func (a AttachMenuPeerTypeBotPM) construct() AttachMenuPeerTypeClass { return &a }

// Ensuring interfaces in compile-time for AttachMenuPeerTypeBotPM.
var (
	_ bin.Encoder     = &AttachMenuPeerTypeBotPM{}
	_ bin.Decoder     = &AttachMenuPeerTypeBotPM{}
	_ bin.BareEncoder = &AttachMenuPeerTypeBotPM{}
	_ bin.BareDecoder = &AttachMenuPeerTypeBotPM{}

	_ AttachMenuPeerTypeClass = &AttachMenuPeerTypeBotPM{}
)

func (a *AttachMenuPeerTypeBotPM) Zero() bool {
	if a == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (a *AttachMenuPeerTypeBotPM) String() string {
	if a == nil {
		return "AttachMenuPeerTypeBotPM(nil)"
	}
	type Alias AttachMenuPeerTypeBotPM
	return fmt.Sprintf("AttachMenuPeerTypeBotPM%+v", Alias(*a))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AttachMenuPeerTypeBotPM) TypeID() uint32 {
	return AttachMenuPeerTypeBotPMTypeID
}

// TypeName returns name of type in TL schema.
func (*AttachMenuPeerTypeBotPM) TypeName() string {
	return "attachMenuPeerTypeBotPM"
}

// TypeInfo returns info about TL type.
func (a *AttachMenuPeerTypeBotPM) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "attachMenuPeerTypeBotPM",
		ID:   AttachMenuPeerTypeBotPMTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (a *AttachMenuPeerTypeBotPM) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode attachMenuPeerTypeBotPM#c32bfa1a as nil")
	}
	b.PutID(AttachMenuPeerTypeBotPMTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *AttachMenuPeerTypeBotPM) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode attachMenuPeerTypeBotPM#c32bfa1a as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (a *AttachMenuPeerTypeBotPM) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode attachMenuPeerTypeBotPM#c32bfa1a to nil")
	}
	if err := b.ConsumeID(AttachMenuPeerTypeBotPMTypeID); err != nil {
		return fmt.Errorf("unable to decode attachMenuPeerTypeBotPM#c32bfa1a: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *AttachMenuPeerTypeBotPM) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode attachMenuPeerTypeBotPM#c32bfa1a to nil")
	}
	return nil
}

// AttachMenuPeerTypePM represents TL type `attachMenuPeerTypePM#f146d31f`.
type AttachMenuPeerTypePM struct {
}

// AttachMenuPeerTypePMTypeID is TL type id of AttachMenuPeerTypePM.
const AttachMenuPeerTypePMTypeID = 0xf146d31f

// construct implements constructor of AttachMenuPeerTypeClass.
func (a AttachMenuPeerTypePM) construct() AttachMenuPeerTypeClass { return &a }

// Ensuring interfaces in compile-time for AttachMenuPeerTypePM.
var (
	_ bin.Encoder     = &AttachMenuPeerTypePM{}
	_ bin.Decoder     = &AttachMenuPeerTypePM{}
	_ bin.BareEncoder = &AttachMenuPeerTypePM{}
	_ bin.BareDecoder = &AttachMenuPeerTypePM{}

	_ AttachMenuPeerTypeClass = &AttachMenuPeerTypePM{}
)

func (a *AttachMenuPeerTypePM) Zero() bool {
	if a == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (a *AttachMenuPeerTypePM) String() string {
	if a == nil {
		return "AttachMenuPeerTypePM(nil)"
	}
	type Alias AttachMenuPeerTypePM
	return fmt.Sprintf("AttachMenuPeerTypePM%+v", Alias(*a))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AttachMenuPeerTypePM) TypeID() uint32 {
	return AttachMenuPeerTypePMTypeID
}

// TypeName returns name of type in TL schema.
func (*AttachMenuPeerTypePM) TypeName() string {
	return "attachMenuPeerTypePM"
}

// TypeInfo returns info about TL type.
func (a *AttachMenuPeerTypePM) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "attachMenuPeerTypePM",
		ID:   AttachMenuPeerTypePMTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (a *AttachMenuPeerTypePM) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode attachMenuPeerTypePM#f146d31f as nil")
	}
	b.PutID(AttachMenuPeerTypePMTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *AttachMenuPeerTypePM) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode attachMenuPeerTypePM#f146d31f as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (a *AttachMenuPeerTypePM) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode attachMenuPeerTypePM#f146d31f to nil")
	}
	if err := b.ConsumeID(AttachMenuPeerTypePMTypeID); err != nil {
		return fmt.Errorf("unable to decode attachMenuPeerTypePM#f146d31f: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *AttachMenuPeerTypePM) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode attachMenuPeerTypePM#f146d31f to nil")
	}
	return nil
}

// AttachMenuPeerTypeChat represents TL type `attachMenuPeerTypeChat#509113f`.
type AttachMenuPeerTypeChat struct {
}

// AttachMenuPeerTypeChatTypeID is TL type id of AttachMenuPeerTypeChat.
const AttachMenuPeerTypeChatTypeID = 0x509113f

// construct implements constructor of AttachMenuPeerTypeClass.
func (a AttachMenuPeerTypeChat) construct() AttachMenuPeerTypeClass { return &a }

// Ensuring interfaces in compile-time for AttachMenuPeerTypeChat.
var (
	_ bin.Encoder     = &AttachMenuPeerTypeChat{}
	_ bin.Decoder     = &AttachMenuPeerTypeChat{}
	_ bin.BareEncoder = &AttachMenuPeerTypeChat{}
	_ bin.BareDecoder = &AttachMenuPeerTypeChat{}

	_ AttachMenuPeerTypeClass = &AttachMenuPeerTypeChat{}
)

func (a *AttachMenuPeerTypeChat) Zero() bool {
	if a == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (a *AttachMenuPeerTypeChat) String() string {
	if a == nil {
		return "AttachMenuPeerTypeChat(nil)"
	}
	type Alias AttachMenuPeerTypeChat
	return fmt.Sprintf("AttachMenuPeerTypeChat%+v", Alias(*a))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AttachMenuPeerTypeChat) TypeID() uint32 {
	return AttachMenuPeerTypeChatTypeID
}

// TypeName returns name of type in TL schema.
func (*AttachMenuPeerTypeChat) TypeName() string {
	return "attachMenuPeerTypeChat"
}

// TypeInfo returns info about TL type.
func (a *AttachMenuPeerTypeChat) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "attachMenuPeerTypeChat",
		ID:   AttachMenuPeerTypeChatTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (a *AttachMenuPeerTypeChat) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode attachMenuPeerTypeChat#509113f as nil")
	}
	b.PutID(AttachMenuPeerTypeChatTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *AttachMenuPeerTypeChat) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode attachMenuPeerTypeChat#509113f as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (a *AttachMenuPeerTypeChat) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode attachMenuPeerTypeChat#509113f to nil")
	}
	if err := b.ConsumeID(AttachMenuPeerTypeChatTypeID); err != nil {
		return fmt.Errorf("unable to decode attachMenuPeerTypeChat#509113f: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *AttachMenuPeerTypeChat) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode attachMenuPeerTypeChat#509113f to nil")
	}
	return nil
}

// AttachMenuPeerTypeBroadcast represents TL type `attachMenuPeerTypeBroadcast#7bfbdefc`.
type AttachMenuPeerTypeBroadcast struct {
}

// AttachMenuPeerTypeBroadcastTypeID is TL type id of AttachMenuPeerTypeBroadcast.
const AttachMenuPeerTypeBroadcastTypeID = 0x7bfbdefc

// construct implements constructor of AttachMenuPeerTypeClass.
func (a AttachMenuPeerTypeBroadcast) construct() AttachMenuPeerTypeClass { return &a }

// Ensuring interfaces in compile-time for AttachMenuPeerTypeBroadcast.
var (
	_ bin.Encoder     = &AttachMenuPeerTypeBroadcast{}
	_ bin.Decoder     = &AttachMenuPeerTypeBroadcast{}
	_ bin.BareEncoder = &AttachMenuPeerTypeBroadcast{}
	_ bin.BareDecoder = &AttachMenuPeerTypeBroadcast{}

	_ AttachMenuPeerTypeClass = &AttachMenuPeerTypeBroadcast{}
)

func (a *AttachMenuPeerTypeBroadcast) Zero() bool {
	if a == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (a *AttachMenuPeerTypeBroadcast) String() string {
	if a == nil {
		return "AttachMenuPeerTypeBroadcast(nil)"
	}
	type Alias AttachMenuPeerTypeBroadcast
	return fmt.Sprintf("AttachMenuPeerTypeBroadcast%+v", Alias(*a))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AttachMenuPeerTypeBroadcast) TypeID() uint32 {
	return AttachMenuPeerTypeBroadcastTypeID
}

// TypeName returns name of type in TL schema.
func (*AttachMenuPeerTypeBroadcast) TypeName() string {
	return "attachMenuPeerTypeBroadcast"
}

// TypeInfo returns info about TL type.
func (a *AttachMenuPeerTypeBroadcast) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "attachMenuPeerTypeBroadcast",
		ID:   AttachMenuPeerTypeBroadcastTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (a *AttachMenuPeerTypeBroadcast) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode attachMenuPeerTypeBroadcast#7bfbdefc as nil")
	}
	b.PutID(AttachMenuPeerTypeBroadcastTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *AttachMenuPeerTypeBroadcast) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode attachMenuPeerTypeBroadcast#7bfbdefc as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (a *AttachMenuPeerTypeBroadcast) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode attachMenuPeerTypeBroadcast#7bfbdefc to nil")
	}
	if err := b.ConsumeID(AttachMenuPeerTypeBroadcastTypeID); err != nil {
		return fmt.Errorf("unable to decode attachMenuPeerTypeBroadcast#7bfbdefc: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *AttachMenuPeerTypeBroadcast) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode attachMenuPeerTypeBroadcast#7bfbdefc to nil")
	}
	return nil
}

// AttachMenuPeerTypeClassName is schema name of AttachMenuPeerTypeClass.
const AttachMenuPeerTypeClassName = "AttachMenuPeerType"

// AttachMenuPeerTypeClass represents AttachMenuPeerType generic type.
//
// Example:
//
//	g, err := td.DecodeAttachMenuPeerType(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *td.AttachMenuPeerTypeSameBotPM: // attachMenuPeerTypeSameBotPM#7d6be90e
//	case *td.AttachMenuPeerTypeBotPM: // attachMenuPeerTypeBotPM#c32bfa1a
//	case *td.AttachMenuPeerTypePM: // attachMenuPeerTypePM#f146d31f
//	case *td.AttachMenuPeerTypeChat: // attachMenuPeerTypeChat#509113f
//	case *td.AttachMenuPeerTypeBroadcast: // attachMenuPeerTypeBroadcast#7bfbdefc
//	default: panic(v)
//	}
type AttachMenuPeerTypeClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() AttachMenuPeerTypeClass

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

// DecodeAttachMenuPeerType implements binary de-serialization for AttachMenuPeerTypeClass.
func DecodeAttachMenuPeerType(buf *bin.Buffer) (AttachMenuPeerTypeClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case AttachMenuPeerTypeSameBotPMTypeID:
		// Decoding attachMenuPeerTypeSameBotPM#7d6be90e.
		v := AttachMenuPeerTypeSameBotPM{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AttachMenuPeerTypeClass: %w", err)
		}
		return &v, nil
	case AttachMenuPeerTypeBotPMTypeID:
		// Decoding attachMenuPeerTypeBotPM#c32bfa1a.
		v := AttachMenuPeerTypeBotPM{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AttachMenuPeerTypeClass: %w", err)
		}
		return &v, nil
	case AttachMenuPeerTypePMTypeID:
		// Decoding attachMenuPeerTypePM#f146d31f.
		v := AttachMenuPeerTypePM{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AttachMenuPeerTypeClass: %w", err)
		}
		return &v, nil
	case AttachMenuPeerTypeChatTypeID:
		// Decoding attachMenuPeerTypeChat#509113f.
		v := AttachMenuPeerTypeChat{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AttachMenuPeerTypeClass: %w", err)
		}
		return &v, nil
	case AttachMenuPeerTypeBroadcastTypeID:
		// Decoding attachMenuPeerTypeBroadcast#7bfbdefc.
		v := AttachMenuPeerTypeBroadcast{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AttachMenuPeerTypeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode AttachMenuPeerTypeClass: %w", bin.NewUnexpectedID(id))
	}
}

// AttachMenuPeerType boxes the AttachMenuPeerTypeClass providing a helper.
type AttachMenuPeerTypeBox struct {
	AttachMenuPeerType AttachMenuPeerTypeClass
}

// Decode implements bin.Decoder for AttachMenuPeerTypeBox.
func (b *AttachMenuPeerTypeBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode AttachMenuPeerTypeBox to nil")
	}
	v, err := DecodeAttachMenuPeerType(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.AttachMenuPeerType = v
	return nil
}

// Encode implements bin.Encode for AttachMenuPeerTypeBox.
func (b *AttachMenuPeerTypeBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.AttachMenuPeerType == nil {
		return fmt.Errorf("unable to encode AttachMenuPeerTypeClass as nil")
	}
	return b.AttachMenuPeerType.Encode(buf)
}
