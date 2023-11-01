package main

import (
	"github.com/go-faster/jx"
	"github.com/gotd/td/bin"
	"github.com/gotd/td/telegram/tljson"
	"github.com/gotd/td/tg"
	"go.uber.org/zap"
	"time"
	"context"
	"github.com/go-faster/errors"
)

func (a application) messagesGetDialogFilters(ctx context.Context) ([]tg.DialogFilterClass, error) {
	return []tg.DialogFilterClass{
		&tg.DialogFilterDefault{},
	}, nil
}

func (a *application) helpGetNearestDC(ctx context.Context) (*tg.NearestDC, error) {
	return &tg.NearestDC{
		ThisDC:    1,
		NearestDC: 1,
		Country:   "Kekistan",
	}, nil
}

func (a *application) helpGetCountriesList(ctx context.Context, req *tg.HelpGetCountriesListRequest) (tg.HelpCountriesListClass, error) {
	return &tg.HelpCountriesList{
		Hash: 1337,
		Countries: []tg.HelpCountry{
			{
				ISO2:        "AF",
				Name:        "Kekistan",
				DefaultName: "Kekistan",
				CountryCodes: []tg.HelpCountryCode{
					{
						CountryCode: "1337",
						Patterns: []string{
							"XXX XXX",
						},
						Prefixes: []string{
							"+1337",
							"1337",
						},
					},
				},
			},
		},
	}, nil
}

func (a *application) helpGetConfig(ctx context.Context) (*tg.Config, error) {
	return &tg.Config{
		Date:    int(time.Now().Unix()),
		Expires: int(time.Now().AddDate(0, 0, 1).Unix()),

		DCOptions: []tg.DCOption{
			{
				ID:           1,
				Port:         a.Port,
				IPAddress:    "127.0.0.1",
				ThisPortOnly: true,
			},
		},
	}, nil
}

func (a *application) helpGetAppConfig(ctx context.Context, h int) (tg.HelpAppConfigClass, error) {
	d := jx.DecodeStr(defaultAppConfig)
	v, err := tljson.Decode(d)
	if err != nil {
		return nil, errors.Wrap(err, "decode")
	}

	return &tg.HelpAppConfig{
		Config: v,
	}, nil
}

func (a *application) authExportLoginToken(ctx context.Context, req *tg.AuthExportLoginTokenRequest) (tg.AuthLoginTokenClass, error) {
	return &tg.AuthLoginToken{
		Expires: int(time.Now().AddDate(0, 0, 1).Unix()),
		Token:   []byte("token"),
	}, nil
}

func (a *application) authBindTempAuthKey(ctx context.Context, req *tg.AuthBindTempAuthKeyRequest) (bool, error) {
	a.lg.Debug("authBindTempAuthKey")
	return true, nil
}

func (a *application) authSendCode(ctx context.Context, req *tg.AuthSendCodeRequest) (tg.AuthSentCodeClass, error) {
	return &tg.AuthSentCode{
		Type:          &tg.AuthSentCodeTypeSMS{Length: 2},
		PhoneCodeHash: "do-u-knw-da-way",
		Timeout:       1337,
	}, nil
}

func (a *application) authSignIn(ctx context.Context, req *tg.AuthSignInRequest) (tg.AuthAuthorizationClass, error) {
	return &tg.AuthAuthorization{
		User: &tg.User{
			Self:      true,
			Verified:  true,
			Phone:     req.PhoneNumber,
			ID:        69,
			Username:  "nice",
			FirstName: "Nice",
		},
	}, nil
}

func (a *application) updatesGetState(ctx context.Context) (*tg.UpdatesState, error) {
	return &tg.UpdatesState{
		Date: int(time.Now().Unix()),
	}, nil
}

func (a *application) messagesGetAvailableReactions(ctx context.Context, hash int) (tg.MessagesAvailableReactionsClass, error) {
	return &tg.MessagesAvailableReactions{
		Hash:      1337,
		Reactions: []tg.AvailableReaction{},
	}, nil
}

func (a *application) messagesGetStickerSet(ctx context.Context, req *tg.MessagesGetStickerSetRequest) (tg.MessagesStickerSetClass, error) {
	return &tg.MessagesStickerSet{
		Set: struct {
			Flags           bin.Fields
			Archived        bool
			Official        bool
			Masks           bool
			Animated        bool
			Videos          bool
			Emojis          bool
			InstalledDate   int
			ID              int64
			AccessHash      int64
			Title           string
			ShortName       string
			Thumbs          []tg.PhotoSizeClass
			ThumbDCID       int
			ThumbVersion    int
			ThumbDocumentID int64
			Count           int
			Hash            int
		}{ Archived: true, Official: true, InstalledDate: 333444, ID: 333444, AccessHash: 333444, Title: "333444", ShortName: "333444",
			ThumbDCID: 3344, ThumbVersion: 3344, ThumbDocumentID: 3344, Count: 3344, Hash: 3344},
	}, nil
}

func (a *application) helpGetTermsOfServiceUpdate(ctx context.Context) (tg.HelpTermsOfServiceUpdateClass, error) {
	return &tg.HelpTermsOfServiceUpdate{
		Expires: int((time.Hour * 24 * 365).Seconds()),
		TermsOfService: tg.HelpTermsOfService{
			ID: tg.DataJSON{
				Data: "1",
			},
			Text: "Go Domination Ensured",
		},
	}, nil
}

func (a *application) usersGetFullUser(ctx context.Context, id tg.InputUserClass) (*tg.UsersUserFull, error) {
	return &tg.UsersUserFull{
		FullUser: tg.UserFull{
			ID:    69,
			About: "About",
		},
	}, nil
}

func (a *application) accountGetNotifySettings(ctx context.Context, peer tg.InputNotifyPeerClass) (*tg.PeerNotifySettings, error) {
	return &tg.PeerNotifySettings{
		ShowPreviews: true,
		Silent: true,
		MuteUntil: 333444,
	}, nil
}

func (a *application) accountGetGlobalPrivacySettings(ctx context.Context) (*tg.GlobalPrivacySettings, error) {
	return &tg.GlobalPrivacySettings{
		ArchiveAndMuteNewNoncontactPeers: true,
		KeepArchivedFolders: true,
		KeepArchivedUnmuted: true,
	}, nil
}

func (a *application) messagesGetAttachMenuBots(ctx context.Context, hash int64) (tg.AttachMenuBotsClass, error) {
	return &tg.AttachMenuBots{
		Hash: 333444,
	}, nil
}

func (a *application) messagesGetDialogs(ctx context.Context, request *tg.MessagesGetDialogsRequest) (tg.MessagesDialogsClass, error) {
	return &tg.MessagesDialogs{
		Users: []tg.UserClass{
			&tg.User{
				Username: "theLostLamb",
				Phone: "theLostLamb3344",
				Bot: true,
			},
		},
	}, nil
}

func (a *application) messagesGetPinnedDialogs(ctx context.Context, folderid int) (*tg.MessagesPeerDialogs, error) {
	return &tg.MessagesPeerDialogs{
		State: struct {
			Pts         int
			Qts         int
			Date        int
			Seq         int
			UnreadCount int
		}{Pts: 3344, Qts: 3344, Date: 3344, Seq: 3344, UnreadCount:3344 },
	}, nil
}

func (a *application) helpGetPremiumPromo(ctx context.Context) (*tg.HelpPremiumPromo, error) {
	return &tg.HelpPremiumPromo{
		StatusText: "MORE GOLD",
	}, nil
}

func (a *application) messagesGetStickers(ctx context.Context, request *tg.MessagesGetStickersRequest) (tg.MessagesStickersClass, error) {
	return &tg.MessagesStickers{Hash: 1337}, nil
}

func (a *application) helpGetPromoData(ctx context.Context) (tg.HelpPromoDataClass, error) {
	return &tg.HelpPromoDataEmpty{
		Expires: 3344,
	}, nil
}

func (a *application) updatesGetDifference(ctx context.Context, req *tg.UpdatesGetDifferenceRequest) (tg.UpdatesDifferenceClass, error) {
	return &tg.UpdatesDifferenceEmpty{
		Date: 3344,
		Seq: 3344,
	}, nil
}

func (a *application) contactsGetContacts(ctx context.Context, hash int64) (tg.ContactsContactsClass, error) {
	return &tg.ContactsContacts{
		SavedCount: 333444,
	}, nil
}

func (a *application) helpDismissSuggestion(ctx context.Context, request *tg.HelpDismissSuggestionRequest) (bool, error) {
	return true, nil
}

func (a *application) helpAcceptTermsOfService(ctx context.Context, id tg.DataJSON) (bool, error) {
	return true, nil
}

func (a *application) messagesGetPeerDialogs(ctx context.Context, peers []tg.InputDialogPeerClass) (*tg.MessagesPeerDialogs, error) {
	return &tg.MessagesPeerDialogs{
		State: struct {
			Pts         int
			Qts         int
			Date        int
			Seq         int
			UnreadCount int
		}{Pts: 3344, Qts: 3344, Date: 3344, Seq: 3344, UnreadCount: 3344},
	}, nil
}

func (a *application) contactsResolveUsername(ctx context.Context, username string) (*tg.ContactsResolvedPeer, error) {
	// Only tdhbcfiles is supported.
	a.lg.Debug("ContactsResolveUsername", zap.String("username", username))
	const peerID = 1300
	return &tg.ContactsResolvedPeer{
		Peer: &tg.PeerChannel{ChannelID: peerID},
		Chats: []tg.ChatClass{
			&tg.Channel{
				Left:       true,
				AccessHash: 1337,
				Title:      username,
				Username:   username,
				Photo:      &tg.ChatPhotoEmpty{},
				Date:       int(time.Now().Unix()),
				ID:         peerID,
			},
		},
	}, nil
}

func (a *application) channelsGetMessages(ctx context.Context, request *tg.ChannelsGetMessagesRequest) (tg.MessagesMessagesClass, error) {
	return &tg.MessagesChannelMessages{
		Inexact: true,
		Pts: 333444,
		Count: 333444,
	}, nil
}
