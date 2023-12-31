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

// AutoDownloadSettings represents TL type `autoDownloadSettings#baa57628`.
type AutoDownloadSettings struct {
	// Flags field of AutoDownloadSettings.
	Flags bin.Fields
	// Disabled field of AutoDownloadSettings.
	Disabled bool
	// VideoPreloadLarge field of AutoDownloadSettings.
	VideoPreloadLarge bool
	// AudioPreloadNext field of AutoDownloadSettings.
	AudioPreloadNext bool
	// PhonecallsLessData field of AutoDownloadSettings.
	PhonecallsLessData bool
	// StoriesPreload field of AutoDownloadSettings.
	StoriesPreload bool
	// PhotoSizeMax field of AutoDownloadSettings.
	PhotoSizeMax int
	// VideoSizeMax field of AutoDownloadSettings.
	VideoSizeMax int64
	// FileSizeMax field of AutoDownloadSettings.
	FileSizeMax int64
	// VideoUploadMaxbitrate field of AutoDownloadSettings.
	VideoUploadMaxbitrate int
	// SmallQueueActiveOperationsMax field of AutoDownloadSettings.
	SmallQueueActiveOperationsMax int
	// LargeQueueActiveOperationsMax field of AutoDownloadSettings.
	LargeQueueActiveOperationsMax int
}

// AutoDownloadSettingsTypeID is TL type id of AutoDownloadSettings.
const AutoDownloadSettingsTypeID = 0xbaa57628

// Ensuring interfaces in compile-time for AutoDownloadSettings.
var (
	_ bin.Encoder     = &AutoDownloadSettings{}
	_ bin.Decoder     = &AutoDownloadSettings{}
	_ bin.BareEncoder = &AutoDownloadSettings{}
	_ bin.BareDecoder = &AutoDownloadSettings{}
)

func (a *AutoDownloadSettings) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Flags.Zero()) {
		return false
	}
	if !(a.Disabled == false) {
		return false
	}
	if !(a.VideoPreloadLarge == false) {
		return false
	}
	if !(a.AudioPreloadNext == false) {
		return false
	}
	if !(a.PhonecallsLessData == false) {
		return false
	}
	if !(a.StoriesPreload == false) {
		return false
	}
	if !(a.PhotoSizeMax == 0) {
		return false
	}
	if !(a.VideoSizeMax == 0) {
		return false
	}
	if !(a.FileSizeMax == 0) {
		return false
	}
	if !(a.VideoUploadMaxbitrate == 0) {
		return false
	}
	if !(a.SmallQueueActiveOperationsMax == 0) {
		return false
	}
	if !(a.LargeQueueActiveOperationsMax == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *AutoDownloadSettings) String() string {
	if a == nil {
		return "AutoDownloadSettings(nil)"
	}
	type Alias AutoDownloadSettings
	return fmt.Sprintf("AutoDownloadSettings%+v", Alias(*a))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AutoDownloadSettings) TypeID() uint32 {
	return AutoDownloadSettingsTypeID
}

// TypeName returns name of type in TL schema.
func (*AutoDownloadSettings) TypeName() string {
	return "autoDownloadSettings"
}

// TypeInfo returns info about TL type.
func (a *AutoDownloadSettings) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "autoDownloadSettings",
		ID:   AutoDownloadSettingsTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Disabled",
			SchemaName: "disabled",
			Null:       !a.Flags.Has(0),
		},
		{
			Name:       "VideoPreloadLarge",
			SchemaName: "video_preload_large",
			Null:       !a.Flags.Has(1),
		},
		{
			Name:       "AudioPreloadNext",
			SchemaName: "audio_preload_next",
			Null:       !a.Flags.Has(2),
		},
		{
			Name:       "PhonecallsLessData",
			SchemaName: "phonecalls_less_data",
			Null:       !a.Flags.Has(3),
		},
		{
			Name:       "StoriesPreload",
			SchemaName: "stories_preload",
			Null:       !a.Flags.Has(4),
		},
		{
			Name:       "PhotoSizeMax",
			SchemaName: "photo_size_max",
		},
		{
			Name:       "VideoSizeMax",
			SchemaName: "video_size_max",
		},
		{
			Name:       "FileSizeMax",
			SchemaName: "file_size_max",
		},
		{
			Name:       "VideoUploadMaxbitrate",
			SchemaName: "video_upload_maxbitrate",
		},
		{
			Name:       "SmallQueueActiveOperationsMax",
			SchemaName: "small_queue_active_operations_max",
		},
		{
			Name:       "LargeQueueActiveOperationsMax",
			SchemaName: "large_queue_active_operations_max",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (a *AutoDownloadSettings) SetFlags() {
	if !(a.Disabled == false) {
		a.Flags.Set(0)
	}
	if !(a.VideoPreloadLarge == false) {
		a.Flags.Set(1)
	}
	if !(a.AudioPreloadNext == false) {
		a.Flags.Set(2)
	}
	if !(a.PhonecallsLessData == false) {
		a.Flags.Set(3)
	}
	if !(a.StoriesPreload == false) {
		a.Flags.Set(4)
	}
}

// Encode implements bin.Encoder.
func (a *AutoDownloadSettings) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode autoDownloadSettings#baa57628 as nil")
	}
	b.PutID(AutoDownloadSettingsTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *AutoDownloadSettings) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode autoDownloadSettings#baa57628 as nil")
	}
	a.SetFlags()
	if err := a.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode autoDownloadSettings#baa57628: field flags: %w", err)
	}
	b.PutInt(a.PhotoSizeMax)
	b.PutLong(a.VideoSizeMax)
	b.PutLong(a.FileSizeMax)
	b.PutInt(a.VideoUploadMaxbitrate)
	b.PutInt(a.SmallQueueActiveOperationsMax)
	b.PutInt(a.LargeQueueActiveOperationsMax)
	return nil
}

// Decode implements bin.Decoder.
func (a *AutoDownloadSettings) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode autoDownloadSettings#baa57628 to nil")
	}
	if err := b.ConsumeID(AutoDownloadSettingsTypeID); err != nil {
		return fmt.Errorf("unable to decode autoDownloadSettings#baa57628: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *AutoDownloadSettings) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode autoDownloadSettings#baa57628 to nil")
	}
	{
		if err := a.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode autoDownloadSettings#baa57628: field flags: %w", err)
		}
	}
	a.Disabled = a.Flags.Has(0)
	a.VideoPreloadLarge = a.Flags.Has(1)
	a.AudioPreloadNext = a.Flags.Has(2)
	a.PhonecallsLessData = a.Flags.Has(3)
	a.StoriesPreload = a.Flags.Has(4)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode autoDownloadSettings#baa57628: field photo_size_max: %w", err)
		}
		a.PhotoSizeMax = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode autoDownloadSettings#baa57628: field video_size_max: %w", err)
		}
		a.VideoSizeMax = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode autoDownloadSettings#baa57628: field file_size_max: %w", err)
		}
		a.FileSizeMax = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode autoDownloadSettings#baa57628: field video_upload_maxbitrate: %w", err)
		}
		a.VideoUploadMaxbitrate = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode autoDownloadSettings#baa57628: field small_queue_active_operations_max: %w", err)
		}
		a.SmallQueueActiveOperationsMax = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode autoDownloadSettings#baa57628: field large_queue_active_operations_max: %w", err)
		}
		a.LargeQueueActiveOperationsMax = value
	}
	return nil
}

// SetDisabled sets value of Disabled conditional field.
func (a *AutoDownloadSettings) SetDisabled(value bool) {
	if value {
		a.Flags.Set(0)
		a.Disabled = true
	} else {
		a.Flags.Unset(0)
		a.Disabled = false
	}
}

// GetDisabled returns value of Disabled conditional field.
func (a *AutoDownloadSettings) GetDisabled() (value bool) {
	if a == nil {
		return
	}
	return a.Flags.Has(0)
}

// SetVideoPreloadLarge sets value of VideoPreloadLarge conditional field.
func (a *AutoDownloadSettings) SetVideoPreloadLarge(value bool) {
	if value {
		a.Flags.Set(1)
		a.VideoPreloadLarge = true
	} else {
		a.Flags.Unset(1)
		a.VideoPreloadLarge = false
	}
}

// GetVideoPreloadLarge returns value of VideoPreloadLarge conditional field.
func (a *AutoDownloadSettings) GetVideoPreloadLarge() (value bool) {
	if a == nil {
		return
	}
	return a.Flags.Has(1)
}

// SetAudioPreloadNext sets value of AudioPreloadNext conditional field.
func (a *AutoDownloadSettings) SetAudioPreloadNext(value bool) {
	if value {
		a.Flags.Set(2)
		a.AudioPreloadNext = true
	} else {
		a.Flags.Unset(2)
		a.AudioPreloadNext = false
	}
}

// GetAudioPreloadNext returns value of AudioPreloadNext conditional field.
func (a *AutoDownloadSettings) GetAudioPreloadNext() (value bool) {
	if a == nil {
		return
	}
	return a.Flags.Has(2)
}

// SetPhonecallsLessData sets value of PhonecallsLessData conditional field.
func (a *AutoDownloadSettings) SetPhonecallsLessData(value bool) {
	if value {
		a.Flags.Set(3)
		a.PhonecallsLessData = true
	} else {
		a.Flags.Unset(3)
		a.PhonecallsLessData = false
	}
}

// GetPhonecallsLessData returns value of PhonecallsLessData conditional field.
func (a *AutoDownloadSettings) GetPhonecallsLessData() (value bool) {
	if a == nil {
		return
	}
	return a.Flags.Has(3)
}

// SetStoriesPreload sets value of StoriesPreload conditional field.
func (a *AutoDownloadSettings) SetStoriesPreload(value bool) {
	if value {
		a.Flags.Set(4)
		a.StoriesPreload = true
	} else {
		a.Flags.Unset(4)
		a.StoriesPreload = false
	}
}

// GetStoriesPreload returns value of StoriesPreload conditional field.
func (a *AutoDownloadSettings) GetStoriesPreload() (value bool) {
	if a == nil {
		return
	}
	return a.Flags.Has(4)
}

// GetPhotoSizeMax returns value of PhotoSizeMax field.
func (a *AutoDownloadSettings) GetPhotoSizeMax() (value int) {
	if a == nil {
		return
	}
	return a.PhotoSizeMax
}

// GetVideoSizeMax returns value of VideoSizeMax field.
func (a *AutoDownloadSettings) GetVideoSizeMax() (value int64) {
	if a == nil {
		return
	}
	return a.VideoSizeMax
}

// GetFileSizeMax returns value of FileSizeMax field.
func (a *AutoDownloadSettings) GetFileSizeMax() (value int64) {
	if a == nil {
		return
	}
	return a.FileSizeMax
}

// GetVideoUploadMaxbitrate returns value of VideoUploadMaxbitrate field.
func (a *AutoDownloadSettings) GetVideoUploadMaxbitrate() (value int) {
	if a == nil {
		return
	}
	return a.VideoUploadMaxbitrate
}

// GetSmallQueueActiveOperationsMax returns value of SmallQueueActiveOperationsMax field.
func (a *AutoDownloadSettings) GetSmallQueueActiveOperationsMax() (value int) {
	if a == nil {
		return
	}
	return a.SmallQueueActiveOperationsMax
}

// GetLargeQueueActiveOperationsMax returns value of LargeQueueActiveOperationsMax field.
func (a *AutoDownloadSettings) GetLargeQueueActiveOperationsMax() (value int) {
	if a == nil {
		return
	}
	return a.LargeQueueActiveOperationsMax
}
