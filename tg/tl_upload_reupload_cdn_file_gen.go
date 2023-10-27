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

// UploadReuploadCDNFileRequest represents TL type `upload.reuploadCdnFile#9b2754a8`.
// Request a reupload of a certain file to a CDN DC¹.
//
// Links:
//  1. https://core.telegram.org/cdn
//
// See https://core.telegram.org/method/upload.reuploadCdnFile for reference.
type UploadReuploadCDNFileRequest struct {
	// File token
	FileToken []byte
	// Request token
	RequestToken []byte
}

// UploadReuploadCDNFileRequestTypeID is TL type id of UploadReuploadCDNFileRequest.
const UploadReuploadCDNFileRequestTypeID = 0x9b2754a8

// Ensuring interfaces in compile-time for UploadReuploadCDNFileRequest.
var (
	_ bin.Encoder     = &UploadReuploadCDNFileRequest{}
	_ bin.Decoder     = &UploadReuploadCDNFileRequest{}
	_ bin.BareEncoder = &UploadReuploadCDNFileRequest{}
	_ bin.BareDecoder = &UploadReuploadCDNFileRequest{}
)

func (r *UploadReuploadCDNFileRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.FileToken == nil) {
		return false
	}
	if !(r.RequestToken == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *UploadReuploadCDNFileRequest) String() string {
	if r == nil {
		return "UploadReuploadCDNFileRequest(nil)"
	}
	type Alias UploadReuploadCDNFileRequest
	return fmt.Sprintf("UploadReuploadCDNFileRequest%+v", Alias(*r))
}

// FillFrom fills UploadReuploadCDNFileRequest from given interface.
func (r *UploadReuploadCDNFileRequest) FillFrom(from interface {
	GetFileToken() (value []byte)
	GetRequestToken() (value []byte)
}) {
	r.FileToken = from.GetFileToken()
	r.RequestToken = from.GetRequestToken()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*UploadReuploadCDNFileRequest) TypeID() uint32 {
	return UploadReuploadCDNFileRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*UploadReuploadCDNFileRequest) TypeName() string {
	return "upload.reuploadCdnFile"
}

// TypeInfo returns info about TL type.
func (r *UploadReuploadCDNFileRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "upload.reuploadCdnFile",
		ID:   UploadReuploadCDNFileRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "FileToken",
			SchemaName: "file_token",
		},
		{
			Name:       "RequestToken",
			SchemaName: "request_token",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *UploadReuploadCDNFileRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode upload.reuploadCdnFile#9b2754a8 as nil")
	}
	b.PutID(UploadReuploadCDNFileRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *UploadReuploadCDNFileRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode upload.reuploadCdnFile#9b2754a8 as nil")
	}
	b.PutBytes(r.FileToken)
	b.PutBytes(r.RequestToken)
	return nil
}

// Decode implements bin.Decoder.
func (r *UploadReuploadCDNFileRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode upload.reuploadCdnFile#9b2754a8 to nil")
	}
	if err := b.ConsumeID(UploadReuploadCDNFileRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode upload.reuploadCdnFile#9b2754a8: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *UploadReuploadCDNFileRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode upload.reuploadCdnFile#9b2754a8 to nil")
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode upload.reuploadCdnFile#9b2754a8: field file_token: %w", err)
		}
		r.FileToken = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode upload.reuploadCdnFile#9b2754a8: field request_token: %w", err)
		}
		r.RequestToken = value
	}
	return nil
}

// GetFileToken returns value of FileToken field.
func (r *UploadReuploadCDNFileRequest) GetFileToken() (value []byte) {
	if r == nil {
		return
	}
	return r.FileToken
}

// GetRequestToken returns value of RequestToken field.
func (r *UploadReuploadCDNFileRequest) GetRequestToken() (value []byte) {
	if r == nil {
		return
	}
	return r.RequestToken
}

// UploadReuploadCDNFile invokes method upload.reuploadCdnFile#9b2754a8 returning error if any.
// Request a reupload of a certain file to a CDN DC¹.
//
// Links:
//  1. https://core.telegram.org/cdn
//
// Possible errors:
//
//	500 CDN_UPLOAD_TIMEOUT: A server-side timeout occurred while reuploading the file to the CDN DC.
//	400 FILE_TOKEN_INVALID: The specified file token is invalid.
//	400 RSA_DECRYPT_FAILED: Internal RSA decryption failed.
//
// See https://core.telegram.org/method/upload.reuploadCdnFile for reference.
// Can be used by bots.
func (c *Client) UploadReuploadCDNFile(ctx context.Context, request *UploadReuploadCDNFileRequest) ([]FileHash, error) {
	var result FileHashVector

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return []FileHash(result.Elems), nil
}
