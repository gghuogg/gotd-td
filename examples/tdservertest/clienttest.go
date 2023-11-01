package main

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/gotd/td/bin"
	"io/ioutil"
	"log"

	"github.com/gotd/td/telegram"
	"github.com/gotd/td/tg"
)

func main()  {
	log.SetFlags(log.Ldate|log.Ltime|log.Llongfile)
	log.Println("client starting ")

	var client *telegram.Client
	ctx := context.Background()

	d := tg.NewUpdateDispatcher()
	// 接收消息，私聊，群聊
	log.Println("OnNewMessage")
	d.OnNewMessage(func(ctx context.Context, e tg.Entities, update *tg.UpdateNewMessage) error {
		msg, ok1 := update.Message.(*tg.Message)
		if !ok1 {
			return nil
		}
		self, err := client.Self(ctx)
		if err != nil {
			return err
		}
		fmt.Println("msg:", msg,self)
		return nil
	})

	 log.Println("OnNewChannelMessage")
	d.OnNewChannelMessage(func(ctx context.Context, e tg.Entities, update *tg.UpdateNewChannelMessage) error {
		msg, ok1 := update.Message.(*tg.Message)
		if !ok1 {
			return nil
		}
		self, err := client.Self(ctx)
		if err != nil {
			return err
		}

		fmt.Println("msg:", msg,self)
		return nil
	})


	log.Println("publicKeyBytes get")
	// 1、读取公钥文件，获取公钥字节
	publicKeyBytes, err := ioutil.ReadFile(".\\publickeyread\\rsa_public.pem")
	if err != nil {
	}
	// 2、解码公钥字节，生成加密块对象
	block, _ := pem.Decode(publicKeyBytes)
	if block == nil {
	}
	// 3、解析DER编码的公钥，生成公钥接口
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
	}
	// 4、公钥接口转型成公钥对象
	publicKey := publicKeyInterface.(*rsa.PublicKey)


	log.Println("tgpublickeys get")
	var tgpublickeys []telegram.PublicKey
	var tgpublickey telegram.PublicKey
	tgpublickey.RSA = publicKey
	tgpublickeys = append(tgpublickeys,tgpublickey)
	fmt.Println("打印tgpublickeys",tgpublickeys[0].Fingerprint())

	options := telegram.Options{
		PublicKeys: tgpublickeys,
	}
	client = telegram.NewClient(int(17349), "344583e45741c457fe1862106095a5eb",options)

	log.Println("client RUN")
	if err := client.Run(ctx, func(ctx context.Context) error {
		log.Println("Client Ping")
		if err = client.Ping(ctx); err != nil {
			return err
		}

		log.Println("HelpGetNearestDC")
		getNearestDc,getNearestDcerr := client.API().HelpGetNearestDC(ctx)
		log.Println(getNearestDc,getNearestDcerr)

		log.Println("HelpGetConfig")
		a,b := client.API().HelpGetConfig(ctx)
		log.Println(a,b)
		log.Println("HelpGetAppConfig")
		a1,b1 := client.API().HelpGetAppConfig(ctx,44555)
		log.Println(a1,b1)
		log.Println("AuthExportLoginToken")
		a2,b2 := client.API().AuthExportLoginToken(ctx,&tg.AuthExportLoginTokenRequest{})
		log.Println(a2,b2)

		//log.Println("AuthBindTempAuthKey")
		//client.API().AuthBindTempAuthKey(ctx,&tg.AuthBindTempAuthKeyRequest{})
		log.Println("AuthSendCode")
		a3,b3 := client.API().AuthSendCode(ctx,&tg.AuthSendCodeRequest{})
		log.Println(a3,b3)
		log.Println("AuthSignIn")
		a4,b4 := client.API().AuthSignIn(ctx,&tg.AuthSignInRequest{})
		log.Println(a4,b4)

		log.Println("MessagesGetDialogFilters")
		a5,b5 := client.API().MessagesGetDialogFilters(ctx)
		log.Println(a5,b5)
		log.Println("UpdatesGetState")
		a6,b6 := client.API().UpdatesGetState(ctx)
		log.Println(a6,b6)
		log.Println("MessagesGetAvailableReactions")
		a7,b7 := client.API().MessagesGetAvailableReactions(ctx,55633)
		log.Println(a7,b7)

		log.Println("MessagesGetStickerSet")
		a8,b8 := client.API().MessagesGetStickerSet(ctx,&tg.MessagesGetStickerSetRequest{
			Stickerset: &tg.InputStickerSetID{
				ID: 333444,
				AccessHash: 444333,
			},
		})
		log.Println(a8,b8)
		log.Println("HelpGetTermsOfServiceUpdate")
		a9,b9 := client.API().HelpGetTermsOfServiceUpdate(ctx)
		log.Println(a9,b9)
		// client.API().UsersGetFullUser(ctx,tg.InputUserClass())
		log.Println("AccountGetNotifySettings")
		a10,b10 := client.API().AccountGetNotifySettings(ctx,&tg.InputNotifyPeer{
			Peer: &tg.InputPeerUser{
				UserID: 333444,
				AccessHash: 444333,
			},
		})
		log.Println(a10,b10)

		log.Println("AccountGetGlobalPrivacySettings")
		a11,b11 := client.API().AccountGetGlobalPrivacySettings(ctx)
		log.Println(a11,b11)
		log.Println("MessagesGetAttachMenuBots")
		a12,b12 := client.API().MessagesGetAttachMenuBots(ctx,555555)
		log.Println(a12,b12)
		log.Println("MessagesGetDialogs")
		a13,b13 := client.API().MessagesGetDialogs(ctx,&tg.MessagesGetDialogsRequest{
			OffsetPeer: &tg.InputPeerUser{
				UserID: 333444,
				AccessHash: 444333,
			},
		})
		log.Println(a13,b13)

		log.Println("MessagesGetPinnedDialogs")
		a14,b14 := client.API().MessagesGetPinnedDialogs(ctx,555343)
		log.Println(a14,b14)
		log.Println("HelpGetPremiumPromo")
		a15,b15 := client.API().HelpGetPremiumPromo(ctx)
		log.Println(a15,b15)
		log.Println("MessagesGetStickers")
		a16,b16 := client.API().MessagesGetStickers(ctx,&tg.MessagesGetStickersRequest{})
		log.Println(a16,b16)


		log.Println("HelpGetPromoData")
		a17,b17 := client.API().HelpGetPromoData(ctx)
		log.Println(a17,b17)
		log.Println("UpdatesGetDifference")
		a18,b18 := client.API().UpdatesGetDifference(ctx,&tg.UpdatesGetDifferenceRequest{})
		log.Println(a18,b18)
		log.Println("ContactsGetContacts")
		a19,b19 := client.API().ContactsGetContacts(ctx,5555)
		log.Println(a19,b19)

		log.Println("HelpDismissSuggestion")
		a20,b20 := client.API().HelpDismissSuggestion(ctx,&tg.HelpDismissSuggestionRequest{
			Peer: &tg.InputPeerUser{
				UserID: 333444,
				AccessHash: 444333,
			},
		})
		log.Println(a20,b20)
		log.Println("HelpAcceptTermsOfService")
		a21,b21 := client.API().HelpAcceptTermsOfService(ctx,tg.DataJSON{})
		log.Println(a21,b21)
		log.Println("MessagesGetPeerDialogs")
		a22,b22 := client.API().MessagesGetPeerDialogs(ctx,[]tg.InputDialogPeerClass{})
		log.Println(a22,b22)

		//log.Println("ContactsResolveUsername")
		//a23,b23 := client.API().ContactsResolveUsername(ctx,"gghhxx")
		//log.Println(a23,b23)
		log.Println("ChannelsGetMessages")
		a24,b24 := client.API().ChannelsGetMessages(ctx,&tg.ChannelsGetMessagesRequest{
			Channel: &tg.InputChannel{
				ChannelID: 333444,
				AccessHash: 444333,
			},
			ID: []tg.InputMessageClass{
				&tg.InputMessageID{
					ID: 333444,
				},
			},
		})
		log.Println(a24,b24)






		// OnStickersChangeStickerPosition
		aa,bb := client.API().StickersChangeStickerPosition(ctx,&tg.StickersChangeStickerPositionRequest{
			Position: 333444,
			Sticker: &tg.InputDocument{
				ID: 333444,
				AccessHash: 444333,
				FileReference: []byte("333444"),
			},
			})
		log.Println(aa,bb)
		// OnStickersAddStickerToSet
		aa2,bb2 := client.API().StickersAddStickerToSet(ctx,&tg.StickersAddStickerToSetRequest{Sticker: struct {
			Flags      bin.Fields
			Document   tg.InputDocumentClass
			Emoji      string
			MaskCoords tg.MaskCoords
			Keywords   string
		}{Emoji: "333444", Keywords:"444333" }})
		log.Println(aa2,bb2)
		// OnStickersSetStickerSetThumb
		aa3,bb3 := client.API().StickersSetStickerSetThumb(ctx,&tg.StickersSetStickerSetThumbRequest{
			Stickerset: &tg.InputStickerSetID{
				ID: 333444,
				AccessHash: 444333,
			},
			Thumb: &tg.InputDocument{
				ID: 333444,
				AccessHash: 444333,
				FileReference: []byte("333444"),
			},
			ThumbDocumentID: 333444,
		})
		log.Println(aa3,bb3)
		// OnStickersCheckShortName
		aa4,bb4 := client.API().StickersCheckShortName(ctx,"klskldflsdklfjkl")
		log.Println(aa4,bb4)
		// OnStickersSuggestShortName
		aa5,bb5 := client.API().StickersSuggestShortName(ctx,"lsjdkkfasl")
		log.Println(aa5,bb5)
		// OnStickersChangeSticker
		aa6,bb6 := client.API().StickersChangeSticker(ctx,&tg.StickersChangeStickerRequest{
			Sticker: &tg.InputDocument{
				ID: 333444,
				AccessHash: 444333,
				FileReference: []byte("333444"),
			},
			MaskCoords: struct {
				N    int
				X    float64
				Y    float64
				Zoom float64
			}{N: 33, X: 33.33, Y: 33.44, Zoom: 44.44},
			Keywords: "444333",
			Emoji: "333444",
		})
		log.Println(aa6,bb6)
		// OnStickersRenameStickerSet
		aa7,bb7 := client.API().StickersRenameStickerSet(ctx,&tg.StickersRenameStickerSetRequest{
			Stickerset: &tg.InputStickerSetID{
				ID: 333444,
				AccessHash: 444333,
			},
			Title: "444333",
		})
		log.Println(aa7,bb7)
		// OnPhoneGetCallConfig
		aa8,bb8 := client.API().PhoneGetCallConfig(ctx)
		log.Println(aa8,bb8)
		// OnPhoneRequestCall
		aa9,bb9 := client.API().PhoneRequestCall(ctx,&tg.PhoneRequestCallRequest{
			Video: true,
			RandomID: 333444,
			UserID: &tg.InputUser{
				UserID: 333444,
				AccessHash: 444333,
			},
			GAHash: []byte{'4'},
			Protocol: struct {
				Flags           bin.Fields
				UDPP2P          bool
				UDPReflector    bool
				MinLayer        int
				MaxLayer        int
				LibraryVersions []string
			}{ UDPP2P: true, UDPReflector: true, MinLayer: 33, MaxLayer: 44, LibraryVersions: []string{"dddsss"}},
		},)
		log.Println(aa9,bb9)
		// OnPhoneAcceptCall
		aa10,bb10 := client.API().PhoneAcceptCall(ctx,&tg.PhoneAcceptCallRequest{
			Peer: struct {
				ID         int64
				AccessHash int64
			}{ID: 333444, AccessHash: 555},
			GB: []byte("333444"),
			Protocol: struct {
				Flags           bin.Fields
				UDPP2P          bool
				UDPReflector    bool
				MinLayer        int
				MaxLayer        int
				LibraryVersions []string
			}{ UDPP2P: true, UDPReflector: true, MinLayer: 33, MaxLayer: 44, LibraryVersions: []string{"dddsss"}},
		})
		log.Println(aa10,bb10)
		// OnPhoneConfirmCall
		aa11,bb11 := client.API().PhoneConfirmCall(ctx,&tg.PhoneConfirmCallRequest{
			Peer: struct {
				ID         int64
				AccessHash int64
			}{ID: 333444, AccessHash: 555},
			GA: []byte("333444"),
			Protocol: struct {
				Flags           bin.Fields
				UDPP2P          bool
				UDPReflector    bool
				MinLayer        int
				MaxLayer        int
				LibraryVersions []string
			}{ UDPP2P: true, UDPReflector: true, MinLayer: 33, MaxLayer: 44, LibraryVersions: []string{"dddsss"}},
			KeyFingerprint: 333444,
			})
		log.Println(aa11,bb11)
		// OnPhoneReceivedCall
		aa12,bb12 := client.API().PhoneReceivedCall(ctx,tg.InputPhoneCall{
			ID: 333444,
			AccessHash: 444333,
		})
		log.Println(aa12,bb12)
		// OnPhoneDiscardCall
		aa13,bb13 := client.API().PhoneDiscardCall(ctx,&tg.PhoneDiscardCallRequest{
			Video: true,
			Peer: struct {
				ID         int64
				AccessHash int64
			}{ID: 333444, AccessHash: 555},
			Duration: 333444,
			Reason: &tg.PhoneCallDiscardReasonBusy{},
			ConnectionID: 333444,
		})
		log.Println(aa13,bb13)
		// OnPhoneSetCallRating
		aa14,bb14 := client.API().PhoneSetCallRating(ctx,&tg.PhoneSetCallRatingRequest{
			UserInitiative: true,
			Peer: struct {
				ID         int64
				AccessHash int64
			}{ID: 333444, AccessHash: 555},
			Rating: 333444,
			Comment: "444333",
		})
		log.Println(aa14,bb14)
		// OnPhoneSaveCallDebug
		aa15,bb15 := client.API().PhoneSaveCallDebug(ctx,&tg.PhoneSaveCallDebugRequest{
			Peer: struct {
				ID         int64
				AccessHash int64
			}{ID: 333444, AccessHash: 555},
			Debug: struct{
				Data string }{Data: "333444"},
		})
		log.Println(aa15,bb15)
		// OnPhoneSendSignalingData
		aa16,bb16 := client.API().PhoneSendSignalingData(ctx,&tg.PhoneSendSignalingDataRequest{
			Peer: struct {
				ID         int64
				AccessHash int64
			}{ID: 333444, AccessHash: 555},
			Data: []byte("333444"),
		})
		log.Println(aa16,bb16)
		// OnPhoneCreateGroupCall
		aa17,bb17 := client.API().PhoneCreateGroupCall(ctx,&tg.PhoneCreateGroupCallRequest{
			RtmpStream: true,
			Peer: &tg.InputPeerUser{
				UserID: 333444,
				AccessHash: 4444333,
			},
			RandomID: 33344,
			Title: "3334444",
			ScheduleDate: 444,
		})
		log.Println(aa17,bb17)
		// OnPhoneJoinGroupCall
		aa18,bb18 := client.API().PhoneJoinGroupCall(ctx,&tg.PhoneJoinGroupCallRequest{
			Muted: true,
			VideoStopped: true,
			Call: struct {
				ID         int64
				AccessHash int64
			}{ID: 333, AccessHash: 444},
			JoinAs: &tg.InputPeerUser{
				UserID: 333444,
				AccessHash: 4444333,
			},
			InviteHash: "444333",
			Params: struct{ Data string }{Data: "333444"},
		})
		log.Println(aa18,bb18)
		// OnPhoneLeaveGroupCall
		aa19,bb19 := client.API().PhoneLeaveGroupCall(ctx,&tg.PhoneLeaveGroupCallRequest{
			Call: struct {
				ID         int64
				AccessHash int64
			}{ID: 333, AccessHash: 444},
			Source: 333444,
		})
		log.Println(aa19,bb19)





		log.Println("client Auth Status")
		status, err := client.Auth().Status(ctx)
		if err != nil {
			return err
		}
		if status.Authorized {

		} else {
			log.Println("认证结束")
		}

		log.Println("AccountCheckUsername")
		checkNameBools,checkNameErrs := client.API().AccountCheckUsername(ctx,"gghhxx")
		log.Println(checkNameBools,checkNameErrs)

		log.Println("client handler err")
		return err
	}); err != nil {
		log.Println("client yunxing err")
		log.Println(err)
	}




}
