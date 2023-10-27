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

// WallPaperClassVector is a box for Vector<WallPaper>
type WallPaperClassVector struct {
	// Elements of Vector<WallPaper>
	Elems []WallPaperClass
}

// WallPaperClassVectorTypeID is TL type id of WallPaperClassVector.
const WallPaperClassVectorTypeID = bin.TypeVector

// Ensuring interfaces in compile-time for WallPaperClassVector.
var (
	_ bin.Encoder     = &WallPaperClassVector{}
	_ bin.Decoder     = &WallPaperClassVector{}
	_ bin.BareEncoder = &WallPaperClassVector{}
	_ bin.BareDecoder = &WallPaperClassVector{}
)

func (vec *WallPaperClassVector) Zero() bool {
	if vec == nil {
		return true
	}
	if !(vec.Elems == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (vec *WallPaperClassVector) String() string {
	if vec == nil {
		return "WallPaperClassVector(nil)"
	}
	type Alias WallPaperClassVector
	return fmt.Sprintf("WallPaperClassVector%+v", Alias(*vec))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*WallPaperClassVector) TypeID() uint32 {
	return WallPaperClassVectorTypeID
}

// TypeName returns name of type in TL schema.
func (*WallPaperClassVector) TypeName() string {
	return ""
}

// TypeInfo returns info about TL type.
func (vec *WallPaperClassVector) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "",
		ID:   WallPaperClassVectorTypeID,
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
func (vec *WallPaperClassVector) Encode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't encode Vector<WallPaper> as nil")
	}

	return vec.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (vec *WallPaperClassVector) EncodeBare(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't encode Vector<WallPaper> as nil")
	}
	b.PutVectorHeader(len(vec.Elems))
	for idx, v := range vec.Elems {
		if v == nil {
			return fmt.Errorf("unable to encode Vector<WallPaper>: field Elems element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode Vector<WallPaper>: field Elems element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (vec *WallPaperClassVector) Decode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't decode Vector<WallPaper> to nil")
	}

	return vec.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (vec *WallPaperClassVector) DecodeBare(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't decode Vector<WallPaper> to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode Vector<WallPaper>: field Elems: %w", err)
		}

		if headerLen > 0 {
			vec.Elems = make([]WallPaperClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeWallPaper(b)
			if err != nil {
				return fmt.Errorf("unable to decode Vector<WallPaper>: field Elems: %w", err)
			}
			vec.Elems = append(vec.Elems, value)
		}
	}
	return nil
}

// GetElems returns value of Elems field.
func (vec *WallPaperClassVector) GetElems() (value []WallPaperClass) {
	if vec == nil {
		return
	}
	return vec.Elems
}
