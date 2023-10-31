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

		//log.Println("HelpGetConfig")
		//a,b := client.API().HelpGetConfig(ctx)
		//log.Println(a,b)
		//log.Println("HelpGetAppConfig")
		//a1,b1 := client.API().HelpGetAppConfig(ctx,44555)
		//log.Println(a1,b1)
		//log.Println("AuthExportLoginToken")
		//a2,b2 := client.API().AuthExportLoginToken(ctx,&tg.AuthExportLoginTokenRequest{})
		//log.Println(a2,b2)
		//
		////log.Println("AuthBindTempAuthKey")
		////client.API().AuthBindTempAuthKey(ctx,&tg.AuthBindTempAuthKeyRequest{})
		//log.Println("AuthSendCode")
		//a3,b3 := client.API().AuthSendCode(ctx,&tg.AuthSendCodeRequest{})
		//log.Println(a3,b3)
		//log.Println("AuthSignIn")
		//a4,b4 := client.API().AuthSignIn(ctx,&tg.AuthSignInRequest{})
		//log.Println(a4,b4)
		//
		//log.Println("MessagesGetDialogFilters")
		//a5,b5 := client.API().MessagesGetDialogFilters(ctx)
		//log.Println(a5,b5)
		//log.Println("UpdatesGetState")
		//a6,b6 := client.API().UpdatesGetState(ctx)
		//log.Println(a6,b6)
		//log.Println("MessagesGetAvailableReactions")
		//a7,b7 := client.API().MessagesGetAvailableReactions(ctx,55633)
		//log.Println(a7,b7)
		//
		//log.Println("MessagesGetStickerSet")
		//a8,b8 := client.API().MessagesGetStickerSet(ctx,&tg.MessagesGetStickerSetRequest{})
		//log.Println(a8,b8)
		//log.Println("HelpGetTermsOfServiceUpdate")
		//a9,b9 := client.API().HelpGetTermsOfServiceUpdate(ctx)
		//log.Println(a9,b9)
		//// client.API().UsersGetFullUser(ctx,tg.InputUserClass())
		//log.Println("AccountGetNotifySettings")
		//a10,b10 := client.API().AccountGetNotifySettings(ctx,&tg.InputNotifyPeer{})
		//log.Println(a10,b10)
		//
		//log.Println("AccountGetGlobalPrivacySettings")
		//a11,b11 := client.API().AccountGetGlobalPrivacySettings(ctx)
		//log.Println(a11,b11)
		//log.Println("MessagesGetAttachMenuBots")
		//a12,b12 := client.API().MessagesGetAttachMenuBots(ctx,555555)
		//log.Println(a12,b12)
		//log.Println("MessagesGetDialogs")
		//a13,b13 := client.API().MessagesGetDialogs(ctx,&tg.MessagesGetDialogsRequest{})
		//log.Println(a13,b13)
		//
		//log.Println("MessagesGetPinnedDialogs")
		//a14,b14 := client.API().MessagesGetPinnedDialogs(ctx,555343)
		//log.Println(a14,b14)
		//log.Println("HelpGetPremiumPromo")
		//a15,b15 := client.API().HelpGetPremiumPromo(ctx)
		//log.Println(a15,b15)
		//log.Println("MessagesGetStickers")
		//a16,b16 := client.API().MessagesGetStickers(ctx,&tg.MessagesGetStickersRequest{})
		//log.Println(a16,b16)
		//
		//
		//log.Println("HelpGetPromoData")
		//a17,b17 := client.API().HelpGetPromoData(ctx)
		//log.Println(a17,b17)
		//log.Println("UpdatesGetDifference")
		//a18,b18 := client.API().UpdatesGetDifference(ctx,&tg.UpdatesGetDifferenceRequest{})
		//log.Println(a18,b18)
		//log.Println("ContactsGetContacts")
		//a19,b19 := client.API().ContactsGetContacts(ctx,5555)
		//log.Println(a19,b19)
		//
		//log.Println("HelpDismissSuggestion")
		//a20,b20 := client.API().HelpDismissSuggestion(ctx,&tg.HelpDismissSuggestionRequest{})
		//log.Println(a20,b20)
		//log.Println("HelpAcceptTermsOfService")
		//a21,b21 := client.API().HelpAcceptTermsOfService(ctx,tg.DataJSON{})
		//log.Println(a21,b21)
		//log.Println("MessagesGetPeerDialogs")
		//a22,b22 := client.API().MessagesGetPeerDialogs(ctx,[]tg.InputDialogPeerClass{})
		//log.Println(a22,b22)
		//
		////log.Println("ContactsResolveUsername")
		////a23,b23 := client.API().ContactsResolveUsername(ctx,"gghhxx")
		////log.Println(a23,b23)
		//log.Println("ChannelsGetMessages")
		//a24,b24 := client.API().ChannelsGetMessages(ctx,&tg.ChannelsGetMessagesRequest{})
		//log.Println(a24,b24)






		// OnStickersChangeStickerPosition
		a,b := client.API().StickersChangeStickerPosition(ctx,&tg.StickersChangeStickerPositionRequest{Position: 333444})
		log.Println(a,b)
		// OnStickersAddStickerToSet
		a2,b2 := client.API().StickersAddStickerToSet(ctx,&tg.StickersAddStickerToSetRequest{Sticker: struct {
			Flags      bin.Fields
			Document   tg.InputDocumentClass
			Emoji      string
			MaskCoords tg.MaskCoords
			Keywords   string
		}{Emoji: "333444", Keywords:"444333" }})
		log.Println(a2,b2)
		// OnStickersSetStickerSetThumb
		a3,b3 := client.API().StickersSetStickerSetThumb(ctx,&tg.StickersSetStickerSetThumbRequest{ThumbDocumentID: 333444})
		log.Println(a3,b3)
		// OnStickersCheckShortName
		a4,b4 := client.API().StickersCheckShortName(ctx,"klskldflsdklfjkl")
		log.Println(a4,b4)
		// OnStickersSuggestShortName
		a5,b5 := client.API().StickersSuggestShortName(ctx,"lsjdkkfasl")
		log.Println(a5,b5)
		// OnStickersChangeSticker
		a6,b6 := client.API().StickersChangeSticker(ctx,&tg.StickersChangeStickerRequest{Keywords: "444333",Emoji: "333444"})
		log.Println(a6,b6)
		// OnStickersRenameStickerSet
		a7,b7 := client.API().StickersRenameStickerSet(ctx,&tg.StickersRenameStickerSetRequest{Title: "444333"})
		log.Println(a7,b7)
		// OnPhoneGetCallConfig
		a8,b8 := client.API().PhoneGetCallConfig(ctx)
		log.Println(a8,b8)
		// OnPhoneRequestCall
		a9,b9 := client.API().PhoneRequestCall(ctx,&tg.PhoneRequestCallRequest{RandomID: 333444})
		log.Println(a9,b9)
		// OnPhoneAcceptCall
		a10,b10 := client.API().PhoneAcceptCall(ctx,&tg.PhoneAcceptCallRequest{GB: []byte("333444")})
		log.Println(a10,b10)
		// OnPhoneConfirmCall
		a11,b11 := client.API().PhoneConfirmCall(ctx,&tg.PhoneConfirmCallRequest{KeyFingerprint: 333444})
		log.Println(a11,b11)
		// OnPhoneReceivedCall
		a12,b12 := client.API().PhoneReceivedCall(ctx,tg.InputPhoneCall{ID: 333444,AccessHash: 444333})
		log.Println(a12,b12)
		// OnPhoneDiscardCall
		a13,b13 := client.API().PhoneDiscardCall(ctx,&tg.PhoneDiscardCallRequest{Duration: 333444})
		log.Println(a13,b13)
		// OnPhoneSetCallRating
		a14,b14 := client.API().PhoneSetCallRating(ctx,&tg.PhoneSetCallRatingRequest{Rating: 333444})
		log.Println(a14,b14)
		// OnPhoneSaveCallDebug
		a15,b15 := client.API().PhoneSaveCallDebug(ctx,&tg.PhoneSaveCallDebugRequest{Debug: struct{ Data string }{Data: "333444"}})
		log.Println(a15,b15)
		// OnPhoneSendSignalingData
		a16,b16 := client.API().PhoneSendSignalingData(ctx,&tg.PhoneSendSignalingDataRequest{Data: []byte("333444")})
		log.Println(a16,b16)
		// OnPhoneCreateGroupCall
		a17,b17 := client.API().PhoneCreateGroupCall(ctx,&tg.PhoneCreateGroupCallRequest{Title: "3334444"})
		log.Println(a17,b17)
		// OnPhoneJoinGroupCall
		a18,b18 := client.API().PhoneJoinGroupCall(ctx,&tg.PhoneJoinGroupCallRequest{InviteHash: "444333"})
		log.Println(a18,b18)
		// OnPhoneLeaveGroupCall
		a19,b19 := client.API().PhoneLeaveGroupCall(ctx,&tg.PhoneLeaveGroupCallRequest{Source: 333444})
		log.Println(a19,b19)


















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
