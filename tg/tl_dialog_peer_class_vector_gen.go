// Code generated by gotdgen, DO NOT EDIT.

package tg

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

// DialogPeerClassVector is a box for Vector<DialogPeer>
type DialogPeerClassVector struct {
	// Elements of Vector<DialogPeer>
	Elems []DialogPeerClass
}

// DialogPeerClassVectorTypeID is TL type id of DialogPeerClassVector.
const DialogPeerClassVectorTypeID = bin.TypeVector

// Ensuring interfaces in compile-time for DialogPeerClassVector.
var (
	_ bin.Encoder     = &DialogPeerClassVector{}
	_ bin.Decoder     = &DialogPeerClassVector{}
	_ bin.BareEncoder = &DialogPeerClassVector{}
	_ bin.BareDecoder = &DialogPeerClassVector{}
)

func (vec *DialogPeerClassVector) Zero() bool {
	if vec == nil {
		return true
	}
	if !(vec.Elems == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (vec *DialogPeerClassVector) String() string {
	if vec == nil {
		return "DialogPeerClassVector(nil)"
	}
	type Alias DialogPeerClassVector
	return fmt.Sprintf("DialogPeerClassVector%+v", Alias(*vec))
}

// FillFrom fills DialogPeerClassVector from given interface.
func (vec *DialogPeerClassVector) FillFrom(from interface {
	GetElems() (value []DialogPeerClass)
}) {
	vec.Elems = from.GetElems()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*DialogPeerClassVector) TypeID() uint32 {
	return DialogPeerClassVectorTypeID
}

// TypeName returns name of type in TL schema.
func (*DialogPeerClassVector) TypeName() string {
	return ""
}

// TypeInfo returns info about TL type.
func (vec *DialogPeerClassVector) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "",
		ID:   DialogPeerClassVectorTypeID,
	}
	if vec == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Elems",
			SchemaName: "Elems",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (vec *DialogPeerClassVector) Encode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't encode Vector<DialogPeer> as nil")
	}

	return vec.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (vec *DialogPeerClassVector) EncodeBare(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't encode Vector<DialogPeer> as nil")
	}
	b.PutVectorHeader(len(vec.Elems))
	for idx, v := range vec.Elems {
		if v == nil {
			return fmt.Errorf("unable to encode Vector<DialogPeer>: field Elems element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode Vector<DialogPeer>: field Elems element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (vec *DialogPeerClassVector) Decode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't decode Vector<DialogPeer> to nil")
	}

	return vec.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (vec *DialogPeerClassVector) DecodeBare(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't decode Vector<DialogPeer> to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode Vector<DialogPeer>: field Elems: %w", err)
		}

		if headerLen > 0 {
			vec.Elems = make([]DialogPeerClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeDialogPeer(b)
			if err != nil {
				return fmt.Errorf("unable to decode Vector<DialogPeer>: field Elems: %w", err)
			}
			vec.Elems = append(vec.Elems, value)
		}
	}
	return nil
}

// GetElems returns value of Elems field.
func (vec *DialogPeerClassVector) GetElems() (value []DialogPeerClass) {
	if vec == nil {
		return
	}
	return vec.Elems
}

// MapElems returns field Elems wrapped in DialogPeerClassArray helper.
func (vec *DialogPeerClassVector) MapElems() (value DialogPeerClassArray) {
	return DialogPeerClassArray(vec.Elems)
}
