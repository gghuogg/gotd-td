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

// AutosaveSettings represents TL type `autosaveSettings#c0d11a28`.
type AutosaveSettings struct {
	// Default autosave settings for private chats
	PrivateChatSettings ScopeAutosaveSettings
	// Default autosave settings for basic group and supergroup chats
	GroupSettings ScopeAutosaveSettings
	// Default autosave settings for channel chats
	ChannelSettings ScopeAutosaveSettings
	// Autosave settings for specific chats
	Exceptions []AutosaveSettingsException
}

// AutosaveSettingsTypeID is TL type id of AutosaveSettings.
const AutosaveSettingsTypeID = 0xc0d11a28

// Ensuring interfaces in compile-time for AutosaveSettings.
var (
	_ bin.Encoder     = &AutosaveSettings{}
	_ bin.Decoder     = &AutosaveSettings{}
	_ bin.BareEncoder = &AutosaveSettings{}
	_ bin.BareDecoder = &AutosaveSettings{}
)

func (a *AutosaveSettings) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.PrivateChatSettings.Zero()) {
		return false
	}
	if !(a.GroupSettings.Zero()) {
		return false
	}
	if !(a.ChannelSettings.Zero()) {
		return false
	}
	if !(a.Exceptions == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *AutosaveSettings) String() string {
	if a == nil {
		return "AutosaveSettings(nil)"
	}
	type Alias AutosaveSettings
	return fmt.Sprintf("AutosaveSettings%+v", Alias(*a))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AutosaveSettings) TypeID() uint32 {
	return AutosaveSettingsTypeID
}

// TypeName returns name of type in TL schema.
func (*AutosaveSettings) TypeName() string {
	return "autosaveSettings"
}

// TypeInfo returns info about TL type.
func (a *AutosaveSettings) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "autosaveSettings",
		ID:   AutosaveSettingsTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "PrivateChatSettings",
			SchemaName: "private_chat_settings",
		},
		{
			Name:       "GroupSettings",
			SchemaName: "group_settings",
		},
		{
			Name:       "ChannelSettings",
			SchemaName: "channel_settings",
		},
		{
			Name:       "Exceptions",
			SchemaName: "exceptions",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *AutosaveSettings) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode autosaveSettings#c0d11a28 as nil")
	}
	b.PutID(AutosaveSettingsTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *AutosaveSettings) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode autosaveSettings#c0d11a28 as nil")
	}
	if err := a.PrivateChatSettings.Encode(b); err != nil {
		return fmt.Errorf("unable to encode autosaveSettings#c0d11a28: field private_chat_settings: %w", err)
	}
	if err := a.GroupSettings.Encode(b); err != nil {
		return fmt.Errorf("unable to encode autosaveSettings#c0d11a28: field group_settings: %w", err)
	}
	if err := a.ChannelSettings.Encode(b); err != nil {
		return fmt.Errorf("unable to encode autosaveSettings#c0d11a28: field channel_settings: %w", err)
	}
	b.PutInt(len(a.Exceptions))
	for idx, v := range a.Exceptions {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare autosaveSettings#c0d11a28: field exceptions element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (a *AutosaveSettings) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode autosaveSettings#c0d11a28 to nil")
	}
	if err := b.ConsumeID(AutosaveSettingsTypeID); err != nil {
		return fmt.Errorf("unable to decode autosaveSettings#c0d11a28: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *AutosaveSettings) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode autosaveSettings#c0d11a28 to nil")
	}
	{
		if err := a.PrivateChatSettings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode autosaveSettings#c0d11a28: field private_chat_settings: %w", err)
		}
	}
	{
		if err := a.GroupSettings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode autosaveSettings#c0d11a28: field group_settings: %w", err)
		}
	}
	{
		if err := a.ChannelSettings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode autosaveSettings#c0d11a28: field channel_settings: %w", err)
		}
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode autosaveSettings#c0d11a28: field exceptions: %w", err)
		}

		if headerLen > 0 {
			a.Exceptions = make([]AutosaveSettingsException, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value AutosaveSettingsException
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare autosaveSettings#c0d11a28: field exceptions: %w", err)
			}
			a.Exceptions = append(a.Exceptions, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (a *AutosaveSettings) EncodeTDLibJSON(b tdjson.Encoder) error {
	if a == nil {
		return fmt.Errorf("can't encode autosaveSettings#c0d11a28 as nil")
	}
	b.ObjStart()
	b.PutID("autosaveSettings")
	b.Comma()
	b.FieldStart("private_chat_settings")
	if err := a.PrivateChatSettings.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode autosaveSettings#c0d11a28: field private_chat_settings: %w", err)
	}
	b.Comma()
	b.FieldStart("group_settings")
	if err := a.GroupSettings.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode autosaveSettings#c0d11a28: field group_settings: %w", err)
	}
	b.Comma()
	b.FieldStart("channel_settings")
	if err := a.ChannelSettings.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode autosaveSettings#c0d11a28: field channel_settings: %w", err)
	}
	b.Comma()
	b.FieldStart("exceptions")
	b.ArrStart()
	for idx, v := range a.Exceptions {
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode autosaveSettings#c0d11a28: field exceptions element with index %d: %w", idx, err)
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
func (a *AutosaveSettings) DecodeTDLibJSON(b tdjson.Decoder) error {
	if a == nil {
		return fmt.Errorf("can't decode autosaveSettings#c0d11a28 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("autosaveSettings"); err != nil {
				return fmt.Errorf("unable to decode autosaveSettings#c0d11a28: %w", err)
			}
		case "private_chat_settings":
			if err := a.PrivateChatSettings.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode autosaveSettings#c0d11a28: field private_chat_settings: %w", err)
			}
		case "group_settings":
			if err := a.GroupSettings.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode autosaveSettings#c0d11a28: field group_settings: %w", err)
			}
		case "channel_settings":
			if err := a.ChannelSettings.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode autosaveSettings#c0d11a28: field channel_settings: %w", err)
			}
		case "exceptions":
			if err := b.Arr(func(b tdjson.Decoder) error {
				var value AutosaveSettingsException
				if err := value.DecodeTDLibJSON(b); err != nil {
					return fmt.Errorf("unable to decode autosaveSettings#c0d11a28: field exceptions: %w", err)
				}
				a.Exceptions = append(a.Exceptions, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode autosaveSettings#c0d11a28: field exceptions: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetPrivateChatSettings returns value of PrivateChatSettings field.
func (a *AutosaveSettings) GetPrivateChatSettings() (value ScopeAutosaveSettings) {
	if a == nil {
		return
	}
	return a.PrivateChatSettings
}

// GetGroupSettings returns value of GroupSettings field.
func (a *AutosaveSettings) GetGroupSettings() (value ScopeAutosaveSettings) {
	if a == nil {
		return
	}
	return a.GroupSettings
}

// GetChannelSettings returns value of ChannelSettings field.
func (a *AutosaveSettings) GetChannelSettings() (value ScopeAutosaveSettings) {
	if a == nil {
		return
	}
	return a.ChannelSettings
}

// GetExceptions returns value of Exceptions field.
func (a *AutosaveSettings) GetExceptions() (value []AutosaveSettingsException) {
	if a == nil {
		return
	}
	return a.Exceptions
}
