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

// PhotosGetUserPhotosRequest represents TL type `photos.getUserPhotos#91cd32a8`.
type PhotosGetUserPhotosRequest struct {
	// UserID field of PhotosGetUserPhotosRequest.
	UserID InputUserClass
	// Offset field of PhotosGetUserPhotosRequest.
	Offset int
	// MaxID field of PhotosGetUserPhotosRequest.
	MaxID int64
	// Limit field of PhotosGetUserPhotosRequest.
	Limit int
}

// PhotosGetUserPhotosRequestTypeID is TL type id of PhotosGetUserPhotosRequest.
const PhotosGetUserPhotosRequestTypeID = 0x91cd32a8

// Ensuring interfaces in compile-time for PhotosGetUserPhotosRequest.
var (
	_ bin.Encoder     = &PhotosGetUserPhotosRequest{}
	_ bin.Decoder     = &PhotosGetUserPhotosRequest{}
	_ bin.BareEncoder = &PhotosGetUserPhotosRequest{}
	_ bin.BareDecoder = &PhotosGetUserPhotosRequest{}
)

func (g *PhotosGetUserPhotosRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.UserID == nil) {
		return false
	}
	if !(g.Offset == 0) {
		return false
	}
	if !(g.MaxID == 0) {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *PhotosGetUserPhotosRequest) String() string {
	if g == nil {
		return "PhotosGetUserPhotosRequest(nil)"
	}
	type Alias PhotosGetUserPhotosRequest
	return fmt.Sprintf("PhotosGetUserPhotosRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PhotosGetUserPhotosRequest) TypeID() uint32 {
	return PhotosGetUserPhotosRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PhotosGetUserPhotosRequest) TypeName() string {
	return "photos.getUserPhotos"
}

// TypeInfo returns info about TL type.
func (g *PhotosGetUserPhotosRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "photos.getUserPhotos",
		ID:   PhotosGetUserPhotosRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "Offset",
			SchemaName: "offset",
		},
		{
			Name:       "MaxID",
			SchemaName: "max_id",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *PhotosGetUserPhotosRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode photos.getUserPhotos#91cd32a8 as nil")
	}
	b.PutID(PhotosGetUserPhotosRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *PhotosGetUserPhotosRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode photos.getUserPhotos#91cd32a8 as nil")
	}
	if g.UserID == nil {
		return fmt.Errorf("unable to encode photos.getUserPhotos#91cd32a8: field user_id is nil")
	}
	if err := g.UserID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode photos.getUserPhotos#91cd32a8: field user_id: %w", err)
	}
	b.PutInt(g.Offset)
	b.PutLong(g.MaxID)
	b.PutInt(g.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (g *PhotosGetUserPhotosRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode photos.getUserPhotos#91cd32a8 to nil")
	}
	if err := b.ConsumeID(PhotosGetUserPhotosRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode photos.getUserPhotos#91cd32a8: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *PhotosGetUserPhotosRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode photos.getUserPhotos#91cd32a8 to nil")
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode photos.getUserPhotos#91cd32a8: field user_id: %w", err)
		}
		g.UserID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode photos.getUserPhotos#91cd32a8: field offset: %w", err)
		}
		g.Offset = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode photos.getUserPhotos#91cd32a8: field max_id: %w", err)
		}
		g.MaxID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode photos.getUserPhotos#91cd32a8: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// GetUserID returns value of UserID field.
func (g *PhotosGetUserPhotosRequest) GetUserID() (value InputUserClass) {
	if g == nil {
		return
	}
	return g.UserID
}

// GetOffset returns value of Offset field.
func (g *PhotosGetUserPhotosRequest) GetOffset() (value int) {
	if g == nil {
		return
	}
	return g.Offset
}

// GetMaxID returns value of MaxID field.
func (g *PhotosGetUserPhotosRequest) GetMaxID() (value int64) {
	if g == nil {
		return
	}
	return g.MaxID
}

// GetLimit returns value of Limit field.
func (g *PhotosGetUserPhotosRequest) GetLimit() (value int) {
	if g == nil {
		return
	}
	return g.Limit
}

// PhotosGetUserPhotos invokes method photos.getUserPhotos#91cd32a8 returning error if any.
func (c *Client) PhotosGetUserPhotos(ctx context.Context, request *PhotosGetUserPhotosRequest) (PhotosPhotosClass, error) {
	var result PhotosPhotosBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Photos, nil
}
