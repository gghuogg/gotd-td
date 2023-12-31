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

// MessagesSendEncryptedFileRequest represents TL type `messages.sendEncryptedFile#5559481d`.
type MessagesSendEncryptedFileRequest struct {
	// Flags field of MessagesSendEncryptedFileRequest.
	Flags bin.Fields
	// Silent field of MessagesSendEncryptedFileRequest.
	Silent bool
	// Peer field of MessagesSendEncryptedFileRequest.
	Peer InputEncryptedChat
	// RandomID field of MessagesSendEncryptedFileRequest.
	RandomID int64
	// Data field of MessagesSendEncryptedFileRequest.
	Data []byte
	// File field of MessagesSendEncryptedFileRequest.
	File InputEncryptedFileClass
}

// MessagesSendEncryptedFileRequestTypeID is TL type id of MessagesSendEncryptedFileRequest.
const MessagesSendEncryptedFileRequestTypeID = 0x5559481d

// Ensuring interfaces in compile-time for MessagesSendEncryptedFileRequest.
var (
	_ bin.Encoder     = &MessagesSendEncryptedFileRequest{}
	_ bin.Decoder     = &MessagesSendEncryptedFileRequest{}
	_ bin.BareEncoder = &MessagesSendEncryptedFileRequest{}
	_ bin.BareDecoder = &MessagesSendEncryptedFileRequest{}
)

func (s *MessagesSendEncryptedFileRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.Silent == false) {
		return false
	}
	if !(s.Peer.Zero()) {
		return false
	}
	if !(s.RandomID == 0) {
		return false
	}
	if !(s.Data == nil) {
		return false
	}
	if !(s.File == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesSendEncryptedFileRequest) String() string {
	if s == nil {
		return "MessagesSendEncryptedFileRequest(nil)"
	}
	type Alias MessagesSendEncryptedFileRequest
	return fmt.Sprintf("MessagesSendEncryptedFileRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesSendEncryptedFileRequest) TypeID() uint32 {
	return MessagesSendEncryptedFileRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesSendEncryptedFileRequest) TypeName() string {
	return "messages.sendEncryptedFile"
}

// TypeInfo returns info about TL type.
func (s *MessagesSendEncryptedFileRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.sendEncryptedFile",
		ID:   MessagesSendEncryptedFileRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Silent",
			SchemaName: "silent",
			Null:       !s.Flags.Has(0),
		},
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "RandomID",
			SchemaName: "random_id",
		},
		{
			Name:       "Data",
			SchemaName: "data",
		},
		{
			Name:       "File",
			SchemaName: "file",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (s *MessagesSendEncryptedFileRequest) SetFlags() {
	if !(s.Silent == false) {
		s.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (s *MessagesSendEncryptedFileRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.sendEncryptedFile#5559481d as nil")
	}
	b.PutID(MessagesSendEncryptedFileRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *MessagesSendEncryptedFileRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.sendEncryptedFile#5559481d as nil")
	}
	s.SetFlags()
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.sendEncryptedFile#5559481d: field flags: %w", err)
	}
	if err := s.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.sendEncryptedFile#5559481d: field peer: %w", err)
	}
	b.PutLong(s.RandomID)
	b.PutBytes(s.Data)
	if s.File == nil {
		return fmt.Errorf("unable to encode messages.sendEncryptedFile#5559481d: field file is nil")
	}
	if err := s.File.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.sendEncryptedFile#5559481d: field file: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *MessagesSendEncryptedFileRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.sendEncryptedFile#5559481d to nil")
	}
	if err := b.ConsumeID(MessagesSendEncryptedFileRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.sendEncryptedFile#5559481d: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *MessagesSendEncryptedFileRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.sendEncryptedFile#5559481d to nil")
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.sendEncryptedFile#5559481d: field flags: %w", err)
		}
	}
	s.Silent = s.Flags.Has(0)
	{
		if err := s.Peer.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.sendEncryptedFile#5559481d: field peer: %w", err)
		}
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendEncryptedFile#5559481d: field random_id: %w", err)
		}
		s.RandomID = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendEncryptedFile#5559481d: field data: %w", err)
		}
		s.Data = value
	}
	{
		value, err := DecodeInputEncryptedFile(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendEncryptedFile#5559481d: field file: %w", err)
		}
		s.File = value
	}
	return nil
}

// SetSilent sets value of Silent conditional field.
func (s *MessagesSendEncryptedFileRequest) SetSilent(value bool) {
	if value {
		s.Flags.Set(0)
		s.Silent = true
	} else {
		s.Flags.Unset(0)
		s.Silent = false
	}
}

// GetSilent returns value of Silent conditional field.
func (s *MessagesSendEncryptedFileRequest) GetSilent() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(0)
}

// GetPeer returns value of Peer field.
func (s *MessagesSendEncryptedFileRequest) GetPeer() (value InputEncryptedChat) {
	if s == nil {
		return
	}
	return s.Peer
}

// GetRandomID returns value of RandomID field.
func (s *MessagesSendEncryptedFileRequest) GetRandomID() (value int64) {
	if s == nil {
		return
	}
	return s.RandomID
}

// GetData returns value of Data field.
func (s *MessagesSendEncryptedFileRequest) GetData() (value []byte) {
	if s == nil {
		return
	}
	return s.Data
}

// GetFile returns value of File field.
func (s *MessagesSendEncryptedFileRequest) GetFile() (value InputEncryptedFileClass) {
	if s == nil {
		return
	}
	return s.File
}

// MessagesSendEncryptedFile invokes method messages.sendEncryptedFile#5559481d returning error if any.
func (c *Client) MessagesSendEncryptedFile(ctx context.Context, request *MessagesSendEncryptedFileRequest) (MessagesSentEncryptedMessageClass, error) {
	var result MessagesSentEncryptedMessageBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.SentEncryptedMessage, nil
}
