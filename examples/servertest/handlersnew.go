



package main

import (
	"context"
	"github.com/gotd/td/bin"
	"github.com/gotd/td/tg"
)

// OnAuthSendCode(f func(ctx context.Context, request *AuthSendCodeRequest) (AuthSentCodeClass, error)) {
func (a *application) authSendCodebak(ctx context.Context, request *tg.AuthSendCodeRequest) (tg.AuthSentCodeClass, error) {
        return &tg.AuthSentCode{
    	PhoneCodeHash: "theLostLamb3344",
    	Timeout: 444333,
	}, nil
}


// OnAuthSignUp(f func(ctx context.Context, request *AuthSignUpRequest) (AuthAuthorizationClass, error)) {
func (a *application) authSignUp(ctx context.Context, request *tg.AuthSignUpRequest) (tg.AuthAuthorizationClass, error) {
    return &tg.AuthAuthorization{
    	SetupPasswordRequired: true,
    	OtherwiseReloginDays: 333444,
    	TmpSessions: 444333,
    	FutureAuthToken: []byte("theLostLamb"),
	}, nil
}


// OnAuthSignIn(f func(ctx context.Context, request *AuthSignInRequest) (AuthAuthorizationClass, error)) {
func (a *application) authSignI(ctx context.Context, request *tg.AuthSignInRequest) (tg.AuthAuthorizationClass, error) {
        return &tg.AuthAuthorization{
    	SetupPasswordRequired: true,
    	OtherwiseReloginDays: 333444,
    	TmpSessions: 444333,
    	FutureAuthToken: []byte("theLostLamb"),
	}, nil
}


// OnAuthLogOut(f func(ctx context.Context) (*AuthLoggedOut, error)) {
func (a *application) authLogOut(ctx context.Context) (*tg.AuthLoggedOut, error) {
    return &tg.AuthLoggedOut{
    	FutureAuthToken:  []byte("theLostLamb"),
	}, nil
}


// OnAuthResetAuthorizations(f func(ctx context.Context) (bool, error)) {
func (a *application) authResetAuthorizations(ctx context.Context) (bool, error) {
    return true, nil
}


// OnAuthExportAuthorization(f func(ctx context.Context, dcid int) (*AuthExportedAuthorization, error)) {
func (a *application) authExportAuthorizatio(ctx context.Context, dcid int) (*tg.AuthExportedAuthorization, error) {
    return &tg.AuthExportedAuthorization{
    	ID: 333444,
    	Bytes: []byte("theLostLamb"),
	}, nil
}


// OnAuthImportAuthorization(f func(ctx context.Context, request *AuthImportAuthorizationRequest) (AuthAuthorizationClass, error)) {
func (a *application) authImportAuthorizatio(ctx context.Context, request *tg.AuthImportAuthorizationRequest) (tg.AuthAuthorizationClass, error) {
        return &tg.AuthAuthorization{
    	SetupPasswordRequired: true,
    	OtherwiseReloginDays: 333444,
    	TmpSessions: 444333,
    	FutureAuthToken: []byte("theLostLamb"),
	}, nil
}


// OnAuthBindTempAuthKey(f func(ctx context.Context, request *AuthBindTempAuthKeyRequest) (bool, error)) {
func (a *application) authBindTempAuthKeybak(ctx context.Context, request *tg.AuthBindTempAuthKeyRequest) (bool , error) {
    return true, nil
}


// OnAuthImportBotAuthorization(f func(ctx context.Context, request *AuthImportBotAuthorizationRequest) (AuthAuthorizationClass, error)) {
func (a *application) authImportBotAuthorizatio(ctx context.Context, request *tg.AuthImportBotAuthorizationRequest) (tg.AuthAuthorizationClass, error) {
        return &tg.AuthAuthorization{
    	SetupPasswordRequired: true,
    	OtherwiseReloginDays: 333444,
    	TmpSessions: 444333,
    	FutureAuthToken: []byte("theLostLamb"),
	}, nil
}


// OnAuthCheckPassword(f func(ctx context.Context, password InputCheckPasswordSRPClass) (AuthAuthorizationClass, error)) {
func (a *application) authCheckPassword(ctx context.Context, password tg.InputCheckPasswordSRPClass) (tg.AuthAuthorizationClass, error) {
        return &tg.AuthAuthorization{
    	SetupPasswordRequired: true,
    	OtherwiseReloginDays: 333444,
    	TmpSessions: 444333,
    	FutureAuthToken: []byte("theLostLamb"),
	}, nil
}


// OnAuthRequestPasswordRecovery(f func(ctx context.Context) (*AuthPasswordRecovery, error)) {
func (a *application) authRequestPasswordRecovery(ctx context.Context) (*tg.AuthPasswordRecovery, error) {
    return &tg.AuthPasswordRecovery{
    	EmailPattern: "333444",
	}, nil
}


// OnAuthRecoverPassword(f func(ctx context.Context, request *AuthRecoverPasswordRequest) (AuthAuthorizationClass, error)) {
func (a *application) authRecoverPassword(ctx context.Context, request *tg.AuthRecoverPasswordRequest) (tg.AuthAuthorizationClass, error) {
        return &tg.AuthAuthorization{
    	SetupPasswordRequired: true,
    	OtherwiseReloginDays: 333444,
    	TmpSessions: 444333,
    	FutureAuthToken: []byte("theLostLamb"),
	}, nil
}


// OnAuthResendCode(f func(ctx context.Context, request *AuthResendCodeRequest) (AuthSentCodeClass, error)) {
func (a *application) authResendCode(ctx context.Context, request *tg.AuthResendCodeRequest) (tg.AuthSentCodeClass, error) {
    return &tg.AuthSentCode{
    	PhoneCodeHash: "theLostLamb3344",
    	Timeout: 444333,
	}, nil
}


// OnAuthCancelCode(f func(ctx context.Context, request *AuthCancelCodeRequest) (bool, error)) {
func (a *application) authCancelCode(ctx context.Context, request *tg.AuthCancelCodeRequest) (bool , error) {
    return true, nil
}


// OnAuthDropTempAuthKeys(f func(ctx context.Context, exceptauthkeys []int64) (bool, error)) {
func (a *application) authDropTempAuthKeys(ctx context.Context, exceptauthkeys []int64) (bool , error) {
    return true, nil
}


// OnAuthExportLoginToken(f func(ctx context.Context, request *AuthExportLoginTokenRequest) (AuthLoginTokenClass, error)) {
func (a *application) authExportLoginToke(ctx context.Context, request *tg.AuthExportLoginTokenRequest) (tg.AuthLoginTokenClass, error) {
    return &tg.AuthLoginToken{
    	Expires: 333444,
	}, nil
}


// OnAuthImportLoginToken(f func(ctx context.Context, token []byte) (AuthLoginTokenClass, error)) {
func (a *application) authImportLoginToke(ctx context.Context, token []byte) (tg.AuthLoginTokenClass, error) {
    return &tg.AuthLoginToken{
    	Expires: 333444,
	}, nil
}


// OnAuthAcceptLoginToken(f func(ctx context.Context, token []byte) (*Authorization, error)) {
func (a *application) authAcceptLoginToke(ctx context.Context, token []byte) (*tg.Authorization, error) {
    return &tg.Authorization{
    	Hash: 333444,
    	APIID: 444333,
    	AppName: "333444",
    	AppVersion: "444333",
	}, nil
}


// OnAuthCheckRecoveryPassword(f func(ctx context.Context, code string) (bool, error)) {
func (a *application) authCheckRecoveryPassword(ctx context.Context, code string) (bool , error) {
    return true, nil
}


// OnAuthImportWebTokenAuthorization(f func(ctx context.Context, request *AuthImportWebTokenAuthorizationRequest) (AuthAuthorizationClass, error)) {
func (a *application) authImportWebTokenAuthorizatio(ctx context.Context, request *tg.AuthImportWebTokenAuthorizationRequest) (tg.AuthAuthorizationClass, error) {
        return &tg.AuthAuthorization{
    	SetupPasswordRequired: true,
    	OtherwiseReloginDays: 333444,
    	TmpSessions: 444333,
    	FutureAuthToken: []byte("theLostLamb"),
	}, nil
}


// OnAuthRequestFirebaseSMS(f func(ctx context.Context, request *AuthRequestFirebaseSMSRequest) (bool, error)) {
func (a *application) authRequestFirebaseSMS(ctx context.Context, request *tg.AuthRequestFirebaseSMSRequest) (bool , error) {
    return true, nil
}


// OnAuthResetLoginEmail(f func(ctx context.Context, request *AuthResetLoginEmailRequest) (AuthSentCodeClass, error)) {
func (a *application) authResetLoginEmail(ctx context.Context, request *tg.AuthResetLoginEmailRequest) (tg.AuthSentCodeClass, error) {
        return &tg.AuthSentCode{
    	PhoneCodeHash: "theLostLamb3344",
    	Timeout: 444333,
	}, nil
}


// OnAccountRegisterDevice(f func(ctx context.Context, request *AccountRegisterDeviceRequest) (bool, error)) {
func (a *application) accountRegisterDevice(ctx context.Context, request *tg.AccountRegisterDeviceRequest) (bool , error) {
    return true, nil
}


// OnAccountUnregisterDevice(f func(ctx context.Context, request *AccountUnregisterDeviceRequest) (bool, error)) {
func (a *application) accountUnregisterDevice(ctx context.Context, request *tg.AccountUnregisterDeviceRequest) (bool , error) {
    return true, nil
}


// OnAccountUpdateNotifySettings(f func(ctx context.Context, request *AccountUpdateNotifySettingsRequest) (bool, error)) {
func (a *application) accountUpdateNotifySettings(ctx context.Context, request *tg.AccountUpdateNotifySettingsRequest) (bool , error) {
    return true, nil
}


// OnAccountGetNotifySettings(f func(ctx context.Context, peer InputNotifyPeerClass) (*PeerNotifySettings, error)) {
func (a *application) accountGetNotifySettingsbak(ctx context.Context, peer *tg.InputNotifyPeerClass) (*tg.PeerNotifySettings, error) {
    return &tg.PeerNotifySettings{
    	ShowPreviews: true,
    	Silent: true,
    	MuteUntil: 333444,
    	StoriesHideSender: true,
	}, nil
}


// OnAccountResetNotifySettings(f func(ctx context.Context) (bool, error)) {
func (a *application) accountResetNotifySettings(ctx context.Context) (bool, error) {
    return true, nil
}


// OnAccountUpdateProfile(f func(ctx context.Context, request *AccountUpdateProfileRequest) (UserClass, error)) {
func (a *application) accountUpdateProfile(ctx context.Context, request *tg.AccountUpdateProfileRequest) (tg.UserClass, error) {
    return &tg.User{
    	Self: true,
    	AccessHash: 333444,
    	FirstName: "theLostLamb",
    	Username: "333444theLostLamb",
    	Phone: "333444333444",
	}, nil
}


// OnAccountUpdateStatus(f func(ctx context.Context, offline bool) (bool, error)) {
func (a *application) accountUpdateStatus(ctx context.Context, offline bool) (bool , error) {
    return true, nil
}


// OnAccountGetWallPapers(f func(ctx context.Context, hash int64) (AccountWallPapersClass, error)) {
func (a *application) accountGetWallPapers(ctx context.Context, hash int64) (tg.AccountWallPapersClass, error) {
    return &tg.AccountWallPapers{
    	Hash: 333444,
	}, nil
}


// OnAccountReportPeer(f func(ctx context.Context, request *AccountReportPeerRequest) (bool, error)) {
func (a *application) accountReportPeer(ctx context.Context, request *tg.AccountReportPeerRequest) (bool , error) {
    return true, nil
}


// OnAccountCheckUsername(f func(ctx context.Context, username string) (bool, error)) {
func (a *application) accountCheckUsername(ctx context.Context, username string) (bool , error) {
    return true, nil
}


// OnAccountUpdateUsername(f func(ctx context.Context, username string) (UserClass, error)) {
func (a *application) accountUpdateUsername(ctx context.Context, username string) (tg.UserClass, error) {
        return &tg.User{
    	Self: true,
    	AccessHash: 333444,
    	FirstName: "theLostLamb",
    	Username: "333444theLostLamb",
    	Phone: "333444333444",
	}, nil
}


// OnAccountGetPrivacy(f func(ctx context.Context, key InputPrivacyKeyClass) (*AccountPrivacyRules, error)) {
func (a *application) accountGetPrivacy(ctx context.Context, key tg.InputPrivacyKeyClass) (*tg.AccountPrivacyRules, error) {
    return &tg.AccountPrivacyRules{
		Users: []tg.UserClass{
			&tg.User{
				Username: "333444",
				Phone: "444333",
			},
		},
	}, nil
}


// OnAccountSetPrivacy(f func(ctx context.Context, request *AccountSetPrivacyRequest) (*AccountPrivacyRules, error)) {
func (a *application) accountSetPrivacy(ctx context.Context, request *tg.AccountSetPrivacyRequest) (*tg.AccountPrivacyRules, error) {
    return &tg.AccountPrivacyRules{
		Users: []tg.UserClass{
			&tg.User{
				Username: "333444",
				Phone: "444333",
			},
		},
	}, nil
}


// OnAccountDeleteAccount(f func(ctx context.Context, request *AccountDeleteAccountRequest) (bool, error)) {
func (a *application) accountDeleteAccount(ctx context.Context, request *tg.AccountDeleteAccountRequest) (bool , error) {
    return true, nil
}


// OnAccountGetAccountTTL(f func(ctx context.Context) (*AccountDaysTTL, error)) {
func (a *application) accountGetAccountTTL(ctx context.Context, ) (*tg.AccountDaysTTL, error) {
    return &tg.AccountDaysTTL{
    	Days: 333444,
	}, nil
}


// OnAccountSetAccountTTL(f func(ctx context.Context, ttl AccountDaysTTL) (bool, error)) {
func (a *application) accountSetAccountTTL(ctx context.Context, ttl tg.AccountDaysTTL) (bool , error) {
    return true, nil
}


// OnAccountSendChangePhoneCode(f func(ctx context.Context, request *AccountSendChangePhoneCodeRequest) (AuthSentCodeClass, error)) {
func (a *application) accountSendChangePhoneCode(ctx context.Context, request *tg.AccountSendChangePhoneCodeRequest) (tg.AuthSentCodeClass, error) {
        return &tg.AuthSentCode{
    	PhoneCodeHash: "theLostLamb3344",
    	Timeout: 444333,
	}, nil
}


// OnAccountChangePhone(f func(ctx context.Context, request *AccountChangePhoneRequest) (UserClass, error)) {
func (a *application) accountChangePhone(ctx context.Context, request *tg.AccountChangePhoneRequest) (tg.UserClass, error) {
        return &tg.User{
    	Self: true,
    	AccessHash: 333444,
    	FirstName: "theLostLamb",
    	Username: "333444theLostLamb",
    	Phone: "333444333444",
	}, nil
}


// OnAccountUpdateDeviceLocked(f func(ctx context.Context, period int) (bool, error)) {
func (a *application) accountUpdateDeviceLocked(ctx context.Context, period int) (bool , error) {
    return true, nil
}


// OnAccountGetAuthorizations(f func(ctx context.Context) (*AccountAuthorizations, error)) {
func (a *application) accountGetAuthorizations(ctx context.Context,) (*tg.AccountAuthorizations, error) {
    return &tg.AccountAuthorizations{
    	AuthorizationTTLDays: 333444,
	}, nil
}


// OnAccountResetAuthorization(f func(ctx context.Context, hash int64) (bool, error)) {
func (a *application) accountResetAuthorizatio(ctx context.Context, hash int64) (bool , error) {
    return true, nil
}


// OnAccountGetPassword(f func(ctx context.Context) (*AccountPassword, error)) {
func (a *application) accountGetPassword(ctx context.Context) (*tg.AccountPassword, error) {
    return &tg.AccountPassword{
    	HasPassword: true,
    	Hint: "333444",
    	SecureRandom: []byte("theLostLamb333444"),
    	PendingResetDate: 333444,
    	LoginEmailPattern: "theLostLamb333444",
	}, nil
}


// OnAccountGetPasswordSettings(f func(ctx context.Context, password InputCheckPasswordSRPClass) (*AccountPasswordSettings, error)) {
func (a *application) accountGetPasswordSettings(ctx context.Context, password tg.InputCheckPasswordSRPClass) (*tg.AccountPasswordSettings, error) {
    return &tg.AccountPasswordSettings{
    	Email: "theLostLamb3344",
	}, nil
}


// OnAccountUpdatePasswordSettings(f func(ctx context.Context, request *AccountUpdatePasswordSettingsRequest) (bool, error)) {
func (a *application) accountUpdatePasswordSettings(ctx context.Context, request *tg.AccountUpdatePasswordSettingsRequest) (bool , error) {
    return true, nil
}


// OnAccountSendConfirmPhoneCode(f func(ctx context.Context, request *AccountSendConfirmPhoneCodeRequest) (AuthSentCodeClass, error)) {
func (a *application) accountSendConfirmPhoneCode(ctx context.Context, request *tg.AccountSendConfirmPhoneCodeRequest) (tg.AuthSentCodeClass, error) {
        return &tg.AuthSentCode{
    	PhoneCodeHash: "theLostLamb3344",
    	Timeout: 444333,
	}, nil
}


// OnAccountConfirmPhone(f func(ctx context.Context, request *AccountConfirmPhoneRequest) (bool, error)) {
func (a *application) accountConfirmPhone(ctx context.Context, request *tg.AccountConfirmPhoneRequest) (bool , error) {
    return true, nil
}


// OnAccountGetTmpPassword(f func(ctx context.Context, request *AccountGetTmpPasswordRequest) (*AccountTmpPassword, error)) {
func (a *application) accountGetTmpPassword(ctx context.Context, request *tg.AccountGetTmpPasswordRequest) (*tg.AccountTmpPassword, error) {
    return &tg.AccountTmpPassword{
    	TmpPassword: []byte("theLostLamb3344"),
    	ValidUntil: 333444,
	}, nil
}


// OnAccountGetWebAuthorizations(f func(ctx context.Context) (*AccountWebAuthorizations, error)) {
func (a *application) accountGetWebAuthorizations(ctx context.Context) (*tg.AccountWebAuthorizations, error) {
    return &tg.AccountWebAuthorizations{
    	Authorizations: []tg.WebAuthorization{
    		{
    			Hash: 333444,
    			Domain: "theLostLamb",
    			DateActive: 444333,
			},
		},
	}, nil
}


// OnAccountResetWebAuthorization(f func(ctx context.Context, hash int64) (bool, error)) {
func (a *application) accountResetWebAuthorizatio(ctx context.Context, hash int64) (bool , error) {
    return true, nil
}


// OnAccountResetWebAuthorizations(f func(ctx context.Context) (bool, error)) {
func (a *application) accountResetWebAuthorizations(ctx context.Context) (bool, error) {
    return true, nil
}


// OnAccountGetAllSecureValues(f func(ctx context.Context) ([]SecureValue, error)) {
func (a *application) accountGetAllSecureValues(ctx context.Context) ([]tg.SecureValue, error) {
    return []tg.SecureValue{
    	{
    		Hash: []byte("theLostLamb333444"),
		},
	}, nil
}


// OnAccountGetSecureValue(f func(ctx context.Context, types []SecureValueTypeClass) ([]SecureValue, error)) {
func (a *application) accountGetSecureValue(ctx context.Context, types []tg.SecureValueTypeClass) ([]tg.SecureValue, error) {
	return []tg.SecureValue{
		{
			Hash: []byte("theLostLamb333444"),
		},
	}, nil
}


// OnAccountSaveSecureValue(f func(ctx context.Context, request *AccountSaveSecureValueRequest) (*SecureValue, error)) {
func (a *application) accountSaveSecureValue(ctx context.Context, request *tg.AccountSaveSecureValueRequest) (*tg.SecureValue, error) {
    return &tg.SecureValue{
		Hash: []byte("theLostLamb333444"),
	}, nil
}


// OnAccountDeleteSecureValue(f func(ctx context.Context, types []SecureValueTypeClass) (bool, error)) {
func (a *application) accountDeleteSecureValue(ctx context.Context, types []tg.SecureValueTypeClass) (bool , error) {
    return true, nil
}


// OnAccountGetAuthorizationForm(f func(ctx context.Context, request *AccountGetAuthorizationFormRequest) (*AccountAuthorizationForm, error)) {
func (a *application) accountGetAuthorizationForm(ctx context.Context, request *tg.AccountGetAuthorizationFormRequest) (*tg.AccountAuthorizationForm, error) {
    return &tg.AccountAuthorizationForm{
    	PrivacyPolicyURL: "theLostLamb333444",
	}, nil
}


// OnAccountAcceptAuthorization(f func(ctx context.Context, request *AccountAcceptAuthorizationRequest) (bool, error)) {
func (a *application) accountAcceptAuthorizatio(ctx context.Context, request *tg.AccountAcceptAuthorizationRequest) (bool , error) {
    return true, nil
}


// OnAccountSendVerifyPhoneCode(f func(ctx context.Context, request *AccountSendVerifyPhoneCodeRequest) (AuthSentCodeClass, error)) {
func (a *application) accountSendVerifyPhoneCode(ctx context.Context, request *tg.AccountSendVerifyPhoneCodeRequest) (tg.AuthSentCodeClass, error) {
        return &tg.AuthSentCode{
    	PhoneCodeHash: "theLostLamb3344",
    	Timeout: 444333,
	}, nil
}


// OnAccountVerifyPhone(f func(ctx context.Context, request *AccountVerifyPhoneRequest) (bool, error)) {
func (a *application) accountVerifyPhone(ctx context.Context, request *tg.AccountVerifyPhoneRequest) (bool , error) {
    return true, nil
}


// OnAccountSendVerifyEmailCode(f func(ctx context.Context, request *AccountSendVerifyEmailCodeRequest) (*AccountSentEmailCode, error)) {
func (a *application) accountSendVerifyEmailCode(ctx context.Context, request *tg.AccountSendVerifyEmailCodeRequest) (*tg.AccountSentEmailCode, error) {
    return &tg.AccountSentEmailCode{
    	EmailPattern: "333444",
    	Length: 444333,
	}, nil
}


// OnAccountVerifyEmail(f func(ctx context.Context, request *AccountVerifyEmailRequest) (AccountEmailVerifiedClass, error)) {
func (a *application) accountVerifyEmail(ctx context.Context, request *tg.AccountVerifyEmailRequest) (tg.AccountEmailVerifiedClass, error) {
    return &tg.AccountEmailVerified{Email: "4443333"}, nil
}


// OnAccountInitTakeoutSession(f func(ctx context.Context, request *AccountInitTakeoutSessionRequest) (*AccountTakeout, error)) {
func (a *application) accountInitTakeoutSessio(ctx context.Context, request *tg.AccountInitTakeoutSessionRequest) (*tg.AccountTakeout, error) {
    return &tg.AccountTakeout{ID: 333444}, nil
}


// OnAccountFinishTakeoutSession(f func(ctx context.Context, request *AccountFinishTakeoutSessionRequest) (bool, error)) {
func (a *application) accountFinishTakeoutSessio(ctx context.Context, request *tg.AccountFinishTakeoutSessionRequest) (bool , error) {
    return true, nil
}


// OnAccountConfirmPasswordEmail(f func(ctx context.Context, code string) (bool, error)) {
func (a *application) accountConfirmPasswordEmail(ctx context.Context, code string) (bool , error) {
    return true, nil
}


// OnAccountResendPasswordEmail(f func(ctx context.Context) (bool, error)) {
func (a *application) accountResendPasswordEmail(ctx context.Context) (bool, error) {
    return true, nil
}


// OnAccountCancelPasswordEmail(f func(ctx context.Context) (bool, error)) {
func (a *application) accountCancelPasswordEmail(ctx context.Context) (bool, error) {
    return true, nil
}


// OnAccountGetContactSignUpNotification(f func(ctx context.Context) (bool, error)) {
func (a *application) accountGetContactSignUpNotificatio(ctx context.Context) (bool, error) {
    return true, nil
}


// OnAccountSetContactSignUpNotification(f func(ctx context.Context, silent bool) (bool, error)) {
func (a *application) accountSetContactSignUpNotificatio(ctx context.Context, silent bool) (bool , error) {
    return true, nil
}


// OnAccountGetNotifyExceptions(f func(ctx context.Context, request *AccountGetNotifyExceptionsRequest) (UpdatesClass, error)) {
func (a *application) accountGetNotifyExceptions(ctx context.Context, request *tg.AccountGetNotifyExceptionsRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnAccountGetWallPaper(f func(ctx context.Context, wallpaper InputWallPaperClass) (WallPaperClass, error)) {
func (a *application) accountGetWallPaper(ctx context.Context, wallpaper tg.InputWallPaperClass) (tg.WallPaperClass, error) {
    return &tg.WallPaper{
		Dark: true,
		AccessHash: 33444,
		Slug: "444333",
	}, nil
}


// OnAccountUploadWallPaper(f func(ctx context.Context, request *AccountUploadWallPaperRequest) (WallPaperClass, error)) {
func (a *application) accountUploadWallPaper(ctx context.Context, request *tg.AccountUploadWallPaperRequest) (tg.WallPaperClass, error) {
    return &tg.WallPaper{
    	Dark: true,
    	AccessHash: 33444,
    	Slug: "444333",
	}, nil
}


// OnAccountSaveWallPaper(f func(ctx context.Context, request *AccountSaveWallPaperRequest) (bool, error)) {
func (a *application) accountSaveWallPaper(ctx context.Context, request *tg.AccountSaveWallPaperRequest) (bool , error) {
    return true, nil
}


// OnAccountInstallWallPaper(f func(ctx context.Context, request *AccountInstallWallPaperRequest) (bool, error)) {
func (a *application) accountInstallWallPaper(ctx context.Context, request *tg.AccountInstallWallPaperRequest) (bool , error) {
    return true, nil
}


// OnAccountResetWallPapers(f func(ctx context.Context) (bool, error)) {
func (a *application) accountResetWallPapers(ctx context.Context) (bool, error) {
    return true, nil
}


// OnAccountGetAutoDownloadSettings(f func(ctx context.Context) (*AccountAutoDownloadSettings, error)) {
func (a *application) accountGetAutoDownloadSettings(ctx context.Context) (*tg.AccountAutoDownloadSettings, error) {
    return &tg.AccountAutoDownloadSettings{
    	Low: struct {
			Flags                         bin.Fields
			Disabled                      bool
			VideoPreloadLarge             bool
			AudioPreloadNext              bool
			PhonecallsLessData            bool
			StoriesPreload                bool
			PhotoSizeMax                  int
			VideoSizeMax                  int64
			FileSizeMax                   int64
			VideoUploadMaxbitrate         int
			SmallQueueActiveOperationsMax int
			LargeQueueActiveOperationsMax int
		}{ Disabled: true, VideoPreloadLarge: true, PhotoSizeMax: 3344, VideoSizeMax: 4433, FileSizeMax: 3434, },
	}, nil
}


// OnAccountSaveAutoDownloadSettings(f func(ctx context.Context, request *AccountSaveAutoDownloadSettingsRequest) (bool, error)) {
func (a *application) accountSaveAutoDownloadSettings(ctx context.Context, request *tg.AccountSaveAutoDownloadSettingsRequest) (bool , error) {
    return true, nil
}


// OnAccountUploadTheme(f func(ctx context.Context, request *AccountUploadThemeRequest) (DocumentClass, error)) {
func (a *application) accountUploadTheme(ctx context.Context, request *tg.AccountUploadThemeRequest) (tg.DocumentClass, error) {
    return &tg.Document{
		ID: 333444,
		MimeType: "444333",
	}, nil
}


// OnAccountCreateTheme(f func(ctx context.Context, request *AccountCreateThemeRequest) (*Theme, error)) {
func (a *application) accountCreateTheme(ctx context.Context, request *tg.AccountCreateThemeRequest) (*tg.Theme, error) {
	return &tg.Theme{
		Default: true,
		ID: 33344,
		Title: "3334444",
	}, nil
}


// OnAccountUpdateTheme(f func(ctx context.Context, request *AccountUpdateThemeRequest) (*Theme, error)) {
func (a *application) accountUpdateTheme(ctx context.Context, request *tg.AccountUpdateThemeRequest) (*tg.Theme, error) {
	return &tg.Theme{
		Default: true,
		ID: 33344,
		Title: "3334444",
	}, nil
}


// OnAccountSaveTheme(f func(ctx context.Context, request *AccountSaveThemeRequest) (bool, error)) {
func (a *application) accountSaveTheme(ctx context.Context, request *tg.AccountSaveThemeRequest) (bool , error) {
    return true, nil
}


// OnAccountInstallTheme(f func(ctx context.Context, request *AccountInstallThemeRequest) (bool, error)) {
func (a *application) accountInstallTheme(ctx context.Context, request *tg.AccountInstallThemeRequest) (bool , error) {
    return true, nil
}


// OnAccountGetTheme(f func(ctx context.Context, request *AccountGetThemeRequest) (*Theme, error)) {
func (a *application) accountGetTheme(ctx context.Context, request *tg.AccountGetThemeRequest) (*tg.Theme, error) {
    return &tg.Theme{
    	Default: true,
    	ID: 33344,
    	Title: "3334444",
	}, nil
}


// OnAccountGetThemes(f func(ctx context.Context, request *AccountGetThemesRequest) (AccountThemesClass, error)) {
func (a *application) accountGetThemes(ctx context.Context, request *tg.AccountGetThemesRequest) (tg.AccountThemesClass, error) {
    return &tg.AccountThemes{
    	Hash: 333444,
	}, nil
}


// OnAccountSetContentSettings(f func(ctx context.Context, request *AccountSetContentSettingsRequest) (bool, error)) {
func (a *application) accountSetContentSettings(ctx context.Context, request *tg.AccountSetContentSettingsRequest) (bool , error) {
    return true, nil
}


// OnAccountGetContentSettings(f func(ctx context.Context) (*AccountContentSettings, error)) {
func (a *application) accountGetContentSettings(ctx context.Context) (*tg.AccountContentSettings, error) {
    return &tg.AccountContentSettings{
    	SensitiveCanChange: true,
    	SensitiveEnabled: true,
	}, nil
}


// OnAccountGetMultiWallPapers(f func(ctx context.Context, wallpapers []InputWallPaperClass) ([]WallPaperClass, error)) {
func (a *application) accountGetMultiWallPapers(ctx context.Context, wallpapers []tg.InputWallPaperClass) ([]tg.WallPaperClass, error) {
	return []tg.WallPaperClass{
		&tg.WallPaper{
			ID: 333444,
			Slug: "444333",
		},
	},nil
}


// OnAccountGetGlobalPrivacySettings(f func(ctx context.Context) (*GlobalPrivacySettings, error)) {
func (a *application) accountGetGlobalPrivacySettingsbak(ctx context.Context) (*tg.GlobalPrivacySettings, error) {
    return &tg.GlobalPrivacySettings{
		KeepArchivedUnmuted: true,
		KeepArchivedFolders: true,
	}, nil
}


// OnAccountSetGlobalPrivacySettings(f func(ctx context.Context, settings GlobalPrivacySettings) (*GlobalPrivacySettings, error)) {
func (a *application) accountSetGlobalPrivacySettings(ctx context.Context, settings tg.GlobalPrivacySettings) (*tg.GlobalPrivacySettings, error) {
    return &tg.GlobalPrivacySettings{
    	KeepArchivedUnmuted: true,
    	KeepArchivedFolders: true,
	}, nil
}


// OnAccountReportProfilePhoto(f func(ctx context.Context, request *AccountReportProfilePhotoRequest) (bool, error)) {
func (a *application) accountReportProfilePhoto(ctx context.Context, request *tg.AccountReportProfilePhotoRequest) (bool , error) {
    return true, nil
}


// OnAccountResetPassword(f func(ctx context.Context) (AccountResetPasswordResultClass, error)) {
func (a *application) accountResetPassword(ctx context.Context) (tg.AccountResetPasswordResultClass, error) {
    return &tg.AccountResetPasswordOk{}, nil
}


// OnAccountDeclinePasswordReset(f func(ctx context.Context) (bool, error)) {
func (a *application) accountDeclinePasswordReset(ctx context.Context) (bool, error) {
    return true, nil
}


// OnAccountGetChatThemes(f func(ctx context.Context, hash int64) (AccountThemesClass, error)) {
func (a *application) accountGetChatThemes(ctx context.Context, hash int64) (tg.AccountThemesClass, error) {
    return &tg.AccountThemes{
		Hash: 333444,
	}, nil
}


// OnAccountSetAuthorizationTTL(f func(ctx context.Context, authorizationttldays int) (bool, error)) {
func (a *application) accountSetAuthorizationTTL(ctx context.Context, authorizationttldays int) (bool , error) {
    return true, nil
}


// OnAccountChangeAuthorizationSettings(f func(ctx context.Context, request *AccountChangeAuthorizationSettingsRequest) (bool, error)) {
func (a *application) accountChangeAuthorizationSettings(ctx context.Context, request *tg.AccountChangeAuthorizationSettingsRequest) (bool , error) {
    return true, nil
}


// OnAccountGetSavedRingtones(f func(ctx context.Context, hash int64) (AccountSavedRingtonesClass, error)) {
func (a *application) accountGetSavedRingtones(ctx context.Context, hash int64) (tg.AccountSavedRingtonesClass, error) {
    return &tg.AccountSavedRingtones{
		Hash: 333444,
	}, nil
}


// OnAccountSaveRingtone(f func(ctx context.Context, request *AccountSaveRingtoneRequest) (AccountSavedRingtoneClass, error)) {
func (a *application) accountSaveRingtone(ctx context.Context, request *tg.AccountSaveRingtoneRequest) (tg.AccountSavedRingtoneClass, error) {
    return &tg.AccountSavedRingtone{}, nil
}


// OnAccountUploadRingtone(f func(ctx context.Context, request *AccountUploadRingtoneRequest) (DocumentClass, error)) {
func (a *application) accountUploadRingtone(ctx context.Context, request *tg.AccountUploadRingtoneRequest) (tg.DocumentClass, error) {
	return &tg.Document{
		ID: 333444,
		MimeType: "444333",
	}, nil
}


// OnAccountUpdateEmojiStatus(f func(ctx context.Context, emojistatus EmojiStatusClass) (bool, error)) {
func (a *application) accountUpdateEmojiStatus(ctx context.Context, emojistatus tg.EmojiStatusClass) (bool , error) {
    return true, nil
}


// OnAccountGetDefaultEmojiStatuses(f func(ctx context.Context, hash int64) (AccountEmojiStatusesClass, error)) {
func (a *application) accountGetDefaultEmojiStatuses(ctx context.Context, hash int64) (tg.AccountEmojiStatusesClass, error) {
    return &tg.AccountEmojiStatuses{
		Hash: 333444,
	}, nil
}


// OnAccountGetRecentEmojiStatuses(f func(ctx context.Context, hash int64) (AccountEmojiStatusesClass, error)) {
func (a *application) accountGetRecentEmojiStatuses(ctx context.Context, hash int64) (tg.AccountEmojiStatusesClass, error) {
    return &tg.AccountEmojiStatuses{
		Hash: 333444,
	}, nil
}


// OnAccountClearRecentEmojiStatuses(f func(ctx context.Context) (bool, error)) {
func (a *application) accountClearRecentEmojiStatuses(ctx context.Context) (bool, error) {
    return true, nil
}


// OnAccountReorderUsernames(f func(ctx context.Context, order []string) (bool, error)) {
func (a *application) accountReorderUsernames(ctx context.Context, order []string) (bool , error) {
    return true, nil
}


// OnAccountToggleUsername(f func(ctx context.Context, request *AccountToggleUsernameRequest) (bool, error)) {
func (a *application) accountToggleUsername(ctx context.Context, request *tg.AccountToggleUsernameRequest) (bool , error) {
    return true, nil
}


// OnAccountGetDefaultProfilePhotoEmojis(f func(ctx context.Context, hash int64) (EmojiListClass, error)) {
func (a *application) accountGetDefaultProfilePhotoEmojis(ctx context.Context, hash int64) (tg.EmojiListClass, error) {
    return &tg.EmojiList{
    	Hash: 333444,
	}, nil
}


// OnAccountGetDefaultGroupPhotoEmojis(f func(ctx context.Context, hash int64) (EmojiListClass, error)) {
func (a *application) accountGetDefaultGroupPhotoEmojis(ctx context.Context, hash int64) (tg.EmojiListClass, error) {
    return &tg.EmojiList{
    	Hash: 33334444,
	}, nil
}


// OnAccountGetAutoSaveSettings(f func(ctx context.Context) (*AccountAutoSaveSettings, error)) {
func (a *application) accountGetAutoSaveSettings(ctx context.Context) (*tg.AccountAutoSaveSettings, error) {
    return &tg.AccountAutoSaveSettings{
		Users: []tg.UserClass{
			&tg.User{
				Username: "333444",
				Phone: "444333",
			},
		},
	}, nil
}


// OnAccountSaveAutoSaveSettings(f func(ctx context.Context, request *AccountSaveAutoSaveSettingsRequest) (bool, error)) {
func (a *application) accountSaveAutoSaveSettings(ctx context.Context, request *tg.AccountSaveAutoSaveSettingsRequest) (bool , error) {
    return true, nil
}


// OnAccountDeleteAutoSaveExceptions(f func(ctx context.Context) (bool, error)) {
func (a *application) accountDeleteAutoSaveExceptions(ctx context.Context) (bool, error) {
    return true, nil
}


// OnAccountInvalidateSignInCodes(f func(ctx context.Context, codes []string) (bool, error)) {
func (a *application) accountInvalidateSignInCodes(ctx context.Context, codes []string) (bool , error) {
    return true, nil
}


// OnUsersGetUsers(f func(ctx context.Context, id []InputUserClass) ([]UserClass, error)) {
func (a *application) usersGetUsers(ctx context.Context, id []tg.InputUserClass) ([]tg.UserClass, error) {
	return []tg.UserClass{
		&tg.User{
			Username: "333444",
			Phone: "444333",
		},
	},nil
}


// OnUsersGetFullUser(f func(ctx context.Context, id InputUserClass) (*UsersUserFull, error)) {
func (a *application) usersGetFullUserbak(ctx context.Context, id *tg.InputUserClass) (*tg.UsersUserFull, error) {
    return &tg.UsersUserFull{
		Users: []tg.UserClass{
			&tg.User{
				Username: "333444",
				Phone: "444333",
			},
		},
	}, nil
}


// OnUsersSetSecureValueErrors(f func(ctx context.Context, request *UsersSetSecureValueErrorsRequest) (bool, error)) {
func (a *application) usersSetSecureValueErrors(ctx context.Context, request *tg.UsersSetSecureValueErrorsRequest) (bool , error) {
    return true, nil
}


// OnContactsGetContactIDs(f func(ctx context.Context, hash int64) ([]int, error)) {
func (a *application) contactsGetContactIDs(ctx context.Context, hash int64) ([]int, error) {
	return []int{333444}, nil
}


// OnContactsGetStatuses(f func(ctx context.Context) ([]ContactStatus, error)) {
func (a *application) contactsGetStatuses(ctx context.Context) ([]tg.ContactStatus, error) {
    return []tg.ContactStatus{
    	{
    		UserID: 333444,
		},
	}, nil
}


// OnContactsGetContacts(f func(ctx context.Context, hash int64) (ContactsContactsClass, error)) {
func (a *application) contactsGetContactsbak(ctx context.Context, hash int64) (tg.ContactsContactsClass, error) {
    return &tg.ContactsContacts{
		Users: []tg.UserClass{
			&tg.User{
				Username: "333444",
				Phone: "444333",
			},
		},
	}, nil
}


// OnContactsImportContacts(f func(ctx context.Context, contacts []InputPhoneContact) (*ContactsImportedContacts, error)) {
func (a *application) contactsImportContacts(ctx context.Context, contacts []tg.InputPhoneContact) (*tg.ContactsImportedContacts, error) {
    return &tg.ContactsImportedContacts{
		Users: []tg.UserClass{
			&tg.User{
				Username: "333444",
				Phone: "444333",
			},
		},
	}, nil
}


// OnContactsDeleteContacts(f func(ctx context.Context, id []InputUserClass) (UpdatesClass, error)) {
func (a *application) contactsDeleteContacts(ctx context.Context, id []tg.InputUserClass) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnContactsDeleteByPhones(f func(ctx context.Context, phones []string) (bool, error)) {
func (a *application) contactsDeleteByPhones(ctx context.Context, phones []string) (bool , error) {
    return true, nil
}


// OnContactsBlock(f func(ctx context.Context, request *ContactsBlockRequest) (bool, error)) {
func (a *application) contactsBlock(ctx context.Context, request *tg.ContactsBlockRequest) (bool , error) {
    return true, nil
}


// OnContactsUnblock(f func(ctx context.Context, request *ContactsUnblockRequest) (bool, error)) {
func (a *application) contactsUnblock(ctx context.Context, request *tg.ContactsUnblockRequest) (bool , error) {
    return true, nil
}


// OnContactsGetBlocked(f func(ctx context.Context, request *ContactsGetBlockedRequest) (ContactsBlockedClass, error)) {
func (a *application) contactsGetBlocked(ctx context.Context, request *tg.ContactsGetBlockedRequest) (tg.ContactsBlockedClass, error) {
    return &tg.ContactsBlocked{
		Users: []tg.UserClass{
			&tg.User{
				Username: "333444",
				Phone: "444333",
			},
		},
	}, nil
}


// OnContactsSearch(f func(ctx context.Context, request *ContactsSearchRequest) (*ContactsFound, error)) {
func (a *application) contactsSearch(ctx context.Context, request *tg.ContactsSearchRequest) (*tg.ContactsFound, error) {
    return &tg.ContactsFound{
		Users: []tg.UserClass{
			&tg.User{
				Username: "333444",
				Phone: "444333",
			},
		},
	}, nil
}


// OnContactsResolveUsername(f func(ctx context.Context, username string) (*ContactsResolvedPeer, error)) {
func (a *application) contactsResolveUsernamebak(ctx context.Context, username string) (*tg.ContactsResolvedPeer, error) {
    return &tg.ContactsResolvedPeer{
		Users: []tg.UserClass{
			&tg.User{
				Username: "333444",
				Phone: "444333",
			},
		},
	}, nil
}


// OnContactsGetTopPeers(f func(ctx context.Context, request *ContactsGetTopPeersRequest) (ContactsTopPeersClass, error)) {
func (a *application) contactsGetTopPeers(ctx context.Context, request *tg.ContactsGetTopPeersRequest) (tg.ContactsTopPeersClass, error) {
    return &tg.ContactsTopPeers{
		Users: []tg.UserClass{
			&tg.User{
				Username: "333444",
				Phone: "444333",
			},
		},
	}, nil
}


// OnContactsResetTopPeerRating(f func(ctx context.Context, request *ContactsResetTopPeerRatingRequest) (bool, error)) {
func (a *application) contactsResetTopPeerRating(ctx context.Context, request *tg.ContactsResetTopPeerRatingRequest) (bool , error) {
    return true, nil
}


// OnContactsResetSaved(f func(ctx context.Context) (bool, error)) {
func (a *application) contactsResetSaved(ctx context.Context) (bool, error) {
    return true, nil
}


// OnContactsGetSaved(f func(ctx context.Context) ([]SavedPhoneContact, error)) {
func (a *application) contactsGetSaved(ctx context.Context) ([]tg.SavedPhoneContact, error) {
    return []tg.SavedPhoneContact{
    	{
    		Phone: "333444",
    		FirstName: "theLostLamb",
    		LastName: "theLostLamb3344",
    		Date: 444333,
		},
	}, nil
}


// OnContactsToggleTopPeers(f func(ctx context.Context, enabled bool) (bool, error)) {
func (a *application) contactsToggleTopPeers(ctx context.Context, enabled bool) (bool , error) {
    return true, nil
}


// OnContactsAddContact(f func(ctx context.Context, request *ContactsAddContactRequest) (UpdatesClass, error)) {
func (a *application) contactsAddContact(ctx context.Context, request *tg.ContactsAddContactRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnContactsAcceptContact(f func(ctx context.Context, id InputUserClass) (UpdatesClass, error)) {
func (a *application) contactsAcceptContact(ctx context.Context, id tg.InputUserClass) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnContactsGetLocated(f func(ctx context.Context, request *ContactsGetLocatedRequest) (UpdatesClass, error)) {
func (a *application) contactsGetLocated(ctx context.Context, request *tg.ContactsGetLocatedRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnContactsBlockFromReplies(f func(ctx context.Context, request *ContactsBlockFromRepliesRequest) (UpdatesClass, error)) {
func (a *application) contactsBlockFromReplies(ctx context.Context, request *tg.ContactsBlockFromRepliesRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnContactsResolvePhone(f func(ctx context.Context, phone string) (*ContactsResolvedPeer, error)) {
func (a *application) contactsResolvePhone(ctx context.Context, phone string) (*tg.ContactsResolvedPeer, error) {
    return &tg.ContactsResolvedPeer{
		Users: []tg.UserClass{
			&tg.User{
				Username: "333444",
				Phone: "444333",
			},
		},
	}, nil
}


// OnContactsExportContactToken(f func(ctx context.Context) (*ExportedContactToken, error)) {
func (a *application) contactsExportContactToke(ctx context.Context) (*tg.ExportedContactToken, error) {
    return &tg.ExportedContactToken{
    	URL: "333444",
    	Expires: 444333,
	}, nil
}


// OnContactsImportContactToken(f func(ctx context.Context, token string) (UserClass, error)) {
func (a *application) contactsImportContactToke(ctx context.Context, token string) (tg.UserClass, error) {
        return &tg.User{
    	Self: true,
    	AccessHash: 333444,
    	FirstName: "theLostLamb",
    	Username: "333444theLostLamb",
    	Phone: "333444333444",
	}, nil
}


// OnContactsEditCloseFriends(f func(ctx context.Context, id []int64) (bool, error)) {
func (a *application) contactsEditCloseFriends(ctx context.Context, id []int64) (bool , error) {
    return true, nil
}


// OnContactsSetBlocked(f func(ctx context.Context, request *ContactsSetBlockedRequest) (bool, error)) {
func (a *application) contactsSetBlocked(ctx context.Context, request *tg.ContactsSetBlockedRequest) (bool , error) {
    return true, nil
}


// OnMessagesGetMessages(f func(ctx context.Context, id []InputMessageClass) (MessagesMessagesClass, error)) {
func (a *application) messagesGetMessages(ctx context.Context, id []tg.InputMessageClass) (tg.MessagesMessagesClass, error) {
	return &tg.MessagesMessages{
		Users: []tg.UserClass{
			&tg.User{
				Username: "333444",
				Phone: "444333",
			},
		},
	}, nil
}


// OnMessagesGetDialogs(f func(ctx context.Context, request *MessagesGetDialogsRequest) (MessagesDialogsClass, error)) {
func (a *application) messagesGetDialogsbak(ctx context.Context, request *tg.MessagesGetDialogsRequest) (tg.MessagesDialogsClass, error) {
    return &tg.MessagesDialogs{
		Users: []tg.UserClass{
			&tg.User{
				Username: "333444",
				Phone: "444333",
			},
		},

	}, nil
}


// OnMessagesGetHistory(f func(ctx context.Context, request *MessagesGetHistoryRequest) (MessagesMessagesClass, error)) {
func (a *application) messagesGetHistory(ctx context.Context, request *tg.MessagesGetHistoryRequest) (tg.MessagesMessagesClass, error) {
	return &tg.MessagesMessages{
		Users: []tg.UserClass{
			&tg.User{
				Username: "333444",
				Phone: "444333",
			},
		},
	}, nil
}


// OnMessagesSearch(f func(ctx context.Context, request *MessagesSearchRequest) (MessagesMessagesClass, error)) {
func (a *application) messagesSearch(ctx context.Context, request *tg.MessagesSearchRequest) (tg.MessagesMessagesClass, error) {
	return &tg.MessagesMessages{
		Users: []tg.UserClass{
			&tg.User{
				Username: "333444",
				Phone: "444333",
			},
		},
	}, nil
}


// OnMessagesReadHistory(f func(ctx context.Context, request *MessagesReadHistoryRequest) (*MessagesAffectedMessages, error)) {
func (a *application) messagesReadHistory(ctx context.Context, request *tg.MessagesReadHistoryRequest) (*tg.MessagesAffectedMessages, error) {
    return &tg.MessagesAffectedMessages{
		Pts: 333444,
		PtsCount: 444333,
	}, nil
}


// OnMessagesDeleteHistory(f func(ctx context.Context, request *MessagesDeleteHistoryRequest) (*MessagesAffectedHistory, error)) {
func (a *application) messagesDeleteHistory(ctx context.Context, request *tg.MessagesDeleteHistoryRequest) (*tg.MessagesAffectedHistory, error) {
    return &tg.MessagesAffectedHistory{
		Pts: 333444,
		PtsCount: 444333,
		Offset: 343434,
	}, nil
}


// OnMessagesDeleteMessages(f func(ctx context.Context, request *MessagesDeleteMessagesRequest) (*MessagesAffectedMessages, error)) {
func (a *application) messagesDeleteMessages(ctx context.Context, request *tg.MessagesDeleteMessagesRequest) (*tg.MessagesAffectedMessages, error) {
    return &tg.MessagesAffectedMessages{
    	Pts: 333444,
    	PtsCount: 444333,
	}, nil
}


// OnMessagesReceivedMessages(f func(ctx context.Context, maxid int) ([]ReceivedNotifyMessage, error)) {
func (a *application) messagesReceivedMessages(ctx context.Context, maxid int) ([]tg.ReceivedNotifyMessage, error) {
    return []tg.ReceivedNotifyMessage{
    	{
    		ID: 333444,
    		Flags: 444333,
		},
	}, nil
}


// OnMessagesSetTyping(f func(ctx context.Context, request *MessagesSetTypingRequest) (bool, error)) {
func (a *application) messagesSetTyping(ctx context.Context, request *tg.MessagesSetTypingRequest) (bool , error) {
    return true, nil
}


// OnMessagesSendMessage(f func(ctx context.Context, request *MessagesSendMessageRequest) (UpdatesClass, error)) {
func (a *application) messagesSendMessage(ctx context.Context, request *tg.MessagesSendMessageRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnMessagesSendMedia(f func(ctx context.Context, request *MessagesSendMediaRequest) (UpdatesClass, error)) {
func (a *application) messagesSendMedia(ctx context.Context, request *tg.MessagesSendMediaRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnMessagesForwardMessages(f func(ctx context.Context, request *MessagesForwardMessagesRequest) (UpdatesClass, error)) {
func (a *application) messagesForwardMessages(ctx context.Context, request *tg.MessagesForwardMessagesRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnMessagesReportSpam(f func(ctx context.Context, peer InputPeerClass) (bool, error)) {
func (a *application) messagesReportSpam(ctx context.Context, peer tg.InputPeerClass) (bool , error) {
    return true, nil
}


// OnMessagesGetPeerSettings(f func(ctx context.Context, peer InputPeerClass) (*MessagesPeerSettings, error)) {
func (a *application) messagesGetPeerSettings(ctx context.Context, peer tg.InputPeerClass) (*tg.MessagesPeerSettings, error) {
    return &tg.MessagesPeerSettings{
		Chats: []tg.ChatClass{
			&tg.Chat{
				ID: 333444,
				Title: "333444",
			},
		},
		Users: []tg.UserClass{
			&tg.User{
				Username: "333444",
				Phone: "444333",
			},
		},
	}, nil
}


// OnMessagesReport(f func(ctx context.Context, request *MessagesReportRequest) (bool, error)) {
func (a *application) messagesReport(ctx context.Context, request *tg.MessagesReportRequest) (bool , error) {
    return true, nil
}


// OnMessagesGetChats(f func(ctx context.Context, id []int64) (MessagesChatsClass, error)) {
func (a *application) messagesGetChats(ctx context.Context, id []int64) (tg.MessagesChatsClass, error) {
    return &tg.MessagesChats{
		Chats: []tg.ChatClass{
			&tg.Chat{
				ID: 333444,
				Title: "333444",
			},
		},
	}, nil
}


// OnMessagesGetFullChat(f func(ctx context.Context, chatid int64) (*MessagesChatFull, error)) {
func (a *application) messagesGetFullChat(ctx context.Context, chatid int64) (*tg.MessagesChatFull, error) {
    return &tg.MessagesChatFull{
		Users: []tg.UserClass{
			&tg.User{
				Username: "333444",
				Phone: "444333",
			},
		},

	}, nil
}


// OnMessagesEditChatTitle(f func(ctx context.Context, request *MessagesEditChatTitleRequest) (UpdatesClass, error)) {
func (a *application) messagesEditChatTitle(ctx context.Context, request *tg.MessagesEditChatTitleRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnMessagesEditChatPhoto(f func(ctx context.Context, request *MessagesEditChatPhotoRequest) (UpdatesClass, error)) {
func (a *application) messagesEditChatPhoto(ctx context.Context, request *tg.MessagesEditChatPhotoRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnMessagesAddChatUser(f func(ctx context.Context, request *MessagesAddChatUserRequest) (UpdatesClass, error)) {
func (a *application) messagesAddChatUser(ctx context.Context, request *tg.MessagesAddChatUserRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnMessagesDeleteChatUser(f func(ctx context.Context, request *MessagesDeleteChatUserRequest) (UpdatesClass, error)) {
func (a *application) messagesDeleteChatUser(ctx context.Context, request *tg.MessagesDeleteChatUserRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnMessagesCreateChat(f func(ctx context.Context, request *MessagesCreateChatRequest) (UpdatesClass, error)) {
func (a *application) messagesCreateChat(ctx context.Context, request *tg.MessagesCreateChatRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnMessagesGetDhConfig(f func(ctx context.Context, request *MessagesGetDhConfigRequest) (MessagesDhConfigClass, error)) {
func (a *application) messagesGetDhConfig(ctx context.Context, request *tg.MessagesGetDhConfigRequest) (tg.MessagesDhConfigClass, error) {
    return &tg.MessagesDhConfig{
    	G: 333444,
    	Version: 444333,
	}, nil
}


// OnMessagesRequestEncryption(f func(ctx context.Context, request *MessagesRequestEncryptionRequest) (EncryptedChatClass, error)) {
func (a *application) messagesRequestEncryptio(ctx context.Context, request *tg.MessagesRequestEncryptionRequest) (tg.EncryptedChatClass, error) {
    return &tg.EncryptedChat{
		ID: 333444,
		AccessHash: 444333,
		Date: 343434,
	}, nil
}


// OnMessagesAcceptEncryption(f func(ctx context.Context, request *MessagesAcceptEncryptionRequest) (EncryptedChatClass, error)) {
func (a *application) messagesAcceptEncryptio(ctx context.Context, request *tg.MessagesAcceptEncryptionRequest) (tg.EncryptedChatClass, error) {
    return &tg.EncryptedChat{
    	ID: 333444,
    	AccessHash: 444333,
    	Date: 343434,
	}, nil
}


// OnMessagesDiscardEncryption(f func(ctx context.Context, request *MessagesDiscardEncryptionRequest) (bool, error)) {
func (a *application) messagesDiscardEncryptio(ctx context.Context, request *tg.MessagesDiscardEncryptionRequest) (bool , error) {
    return true, nil
}


// OnMessagesSetEncryptedTyping(f func(ctx context.Context, request *MessagesSetEncryptedTypingRequest) (bool, error)) {
func (a *application) messagesSetEncryptedTyping(ctx context.Context, request *tg.MessagesSetEncryptedTypingRequest) (bool , error) {
    return true, nil
}


// OnMessagesReadEncryptedHistory(f func(ctx context.Context, request *MessagesReadEncryptedHistoryRequest) (bool, error)) {
func (a *application) messagesReadEncryptedHistory(ctx context.Context, request *tg.MessagesReadEncryptedHistoryRequest) (bool , error) {
    return true, nil
}


// OnMessagesSendEncrypted(f func(ctx context.Context, request *MessagesSendEncryptedRequest) (MessagesSentEncryptedMessageClass, error)) {
func (a *application) messagesSendEncrypted(ctx context.Context, request *tg.MessagesSendEncryptedRequest) (tg.MessagesSentEncryptedMessageClass, error) {
    return &tg.MessagesSentEncryptedMessage{
		Date: 333444,
	}, nil
}


// OnMessagesSendEncryptedFile(f func(ctx context.Context, request *MessagesSendEncryptedFileRequest) (MessagesSentEncryptedMessageClass, error)) {
func (a *application) messagesSendEncryptedFile(ctx context.Context, request *tg.MessagesSendEncryptedFileRequest) (tg.MessagesSentEncryptedMessageClass, error) {
    return &tg.MessagesSentEncryptedMessage{
		Date: 333444,
	}, nil
}


// OnMessagesSendEncryptedService(f func(ctx context.Context, request *MessagesSendEncryptedServiceRequest) (MessagesSentEncryptedMessageClass, error)) {
func (a *application) messagesSendEncryptedService(ctx context.Context, request *tg.MessagesSendEncryptedServiceRequest) (tg.MessagesSentEncryptedMessageClass, error) {
    return &tg.MessagesSentEncryptedMessage{
    	Date: 333444,
	}, nil
}


// OnMessagesReceivedQueue(f func(ctx context.Context, maxqts int) ([]int64, error)) {
func (a *application) messagesReceivedQueue(ctx context.Context, maxqts int) ([]int64, error) {
    return []int64{333444}, nil
}


// OnMessagesReportEncryptedSpam(f func(ctx context.Context, peer InputEncryptedChat) (bool, error)) {
func (a *application) messagesReportEncryptedSpam(ctx context.Context, peer tg.InputEncryptedChat) (bool , error) {
    return true, nil
}


// OnMessagesReadMessageContents(f func(ctx context.Context, id []int) (*MessagesAffectedMessages, error)) {
func (a *application) messagesReadMessageContents(ctx context.Context, id []int) (*tg.MessagesAffectedMessages, error) {
    return &tg.MessagesAffectedMessages{
    	Pts: 333444,
    	PtsCount: 444333,
	}, nil
}


// OnMessagesGetStickers(f func(ctx context.Context, request *MessagesGetStickersRequest) (MessagesStickersClass, error)) {
func (a *application) messagesGetStickersbak(ctx context.Context, request *tg.MessagesGetStickersRequest) (tg.MessagesStickersClass, error) {
    return &tg.MessagesStickers{
    	Hash: 333444,
    	Stickers: []tg.DocumentClass{
			&tg.Document{
				ID: 333444,
				MimeType: "444333",
			},
		},
	}, nil
}


// OnMessagesGetAllStickers(f func(ctx context.Context, hash int64) (MessagesAllStickersClass, error)) {
func (a *application) messagesGetAllStickers(ctx context.Context, hash int64) (tg.MessagesAllStickersClass, error) {
	return &tg.MessagesAllStickers{
		Hash: 333444,
		Sets: []tg.StickerSet{
			{
				ID: 333444,
				Title: "444333",
			},
		},
	}, nil
}


// OnMessagesGetWebPagePreview(f func(ctx context.Context, request *MessagesGetWebPagePreviewRequest) (MessageMediaClass, error)) {
func (a *application) messagesGetWebPagePreview(ctx context.Context, request *tg.MessagesGetWebPagePreviewRequest) (tg.MessageMediaClass, error) {
    return &tg.MessageMediaPhoto{
    	Spoiler: true,
    	TTLSeconds: 333444,
	}, nil
}


// OnMessagesExportChatInvite(f func(ctx context.Context, request *MessagesExportChatInviteRequest) (ExportedChatInviteClass, error)) {
func (a *application) messagesExportChatInvite(ctx context.Context, request *tg.MessagesExportChatInviteRequest) (tg.ExportedChatInviteClass, error) {
    return &tg.ChatInvitePublicJoinRequests{}, nil
}


// OnMessagesCheckChatInvite(f func(ctx context.Context, hash string) (ChatInviteClass, error)) {
func (a *application) messagesCheckChatInvite(ctx context.Context, hash string) (tg.ChatInviteClass, error) {
    return &tg.ChatInvite{
    	Title: "333444",
    	About: "444333",
	}, nil
}


// OnMessagesImportChatInvite(f func(ctx context.Context, hash string) (UpdatesClass, error)) {
func (a *application) messagesImportChatInvite(ctx context.Context, hash string) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnMessagesGetStickerSet(f func(ctx context.Context, request *MessagesGetStickerSetRequest) (MessagesStickerSetClass, error)) {
func (a *application) messagesGetStickerSetbak(ctx context.Context, request *tg.MessagesGetStickerSetRequest) (tg.MessagesStickerSetClass, error) {
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
		}{ Archived:true , Official: true, Masks: true, Animated: true, Videos: true, Emojis: true, InstalledDate: 555,
			ID: 555, AccessHash: 345345, Title: "ksdkjfskf", ShortName: "klsjdflsj",  ThumbDCID: 2, ThumbVersion: 454, ThumbDocumentID: 34534, Count:234 , Hash: 435345},
	}, nil
}


// OnMessagesInstallStickerSet(f func(ctx context.Context, request *MessagesInstallStickerSetRequest) (MessagesStickerSetInstallResultClass, error)) {
func (a *application) messagesInstallStickerSet(ctx context.Context, request *tg.MessagesInstallStickerSetRequest) (tg.MessagesStickerSetInstallResultClass, error) {
    return &tg.MessagesStickerSetInstallResultSuccess{}, nil
}


// OnMessagesUninstallStickerSet(f func(ctx context.Context, stickerset InputStickerSetClass) (bool, error)) {
func (a *application) messagesUninstallStickerSet(ctx context.Context, stickerset tg.InputStickerSetClass) (bool , error) {
    return true, nil
}


// OnMessagesStartBot(f func(ctx context.Context, request *MessagesStartBotRequest) (UpdatesClass, error)) {
func (a *application) messagesStartBot(ctx context.Context, request *tg.MessagesStartBotRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnMessagesGetMessagesViews(f func(ctx context.Context, request *MessagesGetMessagesViewsRequest) (*MessagesMessageViews, error)) {
func (a *application) messagesGetMessagesViews(ctx context.Context, request *tg.MessagesGetMessagesViewsRequest) (*tg.MessagesMessageViews, error) {
    return &tg.MessagesMessageViews{
		Users: []tg.UserClass{
			&tg.User{
				Username: "333444",
				Phone: "444333",
			},
		},
	}, nil
}


// OnMessagesEditChatAdmin(f func(ctx context.Context, request *MessagesEditChatAdminRequest) (bool, error)) {
func (a *application) messagesEditChatAdmi(ctx context.Context, request *tg.MessagesEditChatAdminRequest) (bool , error) {
    return true, nil
}


// OnMessagesMigrateChat(f func(ctx context.Context, chatid int64) (UpdatesClass, error)) {
func (a *application) messagesMigrateChat(ctx context.Context, chatid int64) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnMessagesSearchGlobal(f func(ctx context.Context, request *MessagesSearchGlobalRequest) (MessagesMessagesClass, error)) {
func (a *application) messagesSearchGlobal(ctx context.Context, request *tg.MessagesSearchGlobalRequest) (tg.MessagesMessagesClass, error) {
	return &tg.MessagesMessages{
		Users: []tg.UserClass{
			&tg.User{
				Username: "333444",
				Phone: "444333",
			},
		},
	}, nil
}


// OnMessagesReorderStickerSets(f func(ctx context.Context, request *MessagesReorderStickerSetsRequest) (bool, error)) {
func (a *application) messagesReorderStickerSets(ctx context.Context, request *tg.MessagesReorderStickerSetsRequest) (bool , error) {
    return true, nil
}


// OnMessagesGetDocumentByHash(f func(ctx context.Context, request *MessagesGetDocumentByHashRequest) (DocumentClass, error)) {
func (a *application) messagesGetDocumentByHash(ctx context.Context, request *tg.MessagesGetDocumentByHashRequest) (tg.DocumentClass, error) {
	return &tg.Document{
		ID: 333444,
		MimeType: "444333",
	}, nil
}


// OnMessagesGetSavedGifs(f func(ctx context.Context, hash int64) (MessagesSavedGifsClass, error)) {
func (a *application) messagesGetSavedGifs(ctx context.Context, hash int64) (tg.MessagesSavedGifsClass, error) {
    return &tg.MessagesSavedGifs{
    	Hash: 333444,
	}, nil
}


// OnMessagesSaveGif(f func(ctx context.Context, request *MessagesSaveGifRequest) (bool, error)) {
func (a *application) messagesSaveGif(ctx context.Context, request *tg.MessagesSaveGifRequest) (bool , error) {
    return true, nil
}


// OnMessagesGetInlineBotResults(f func(ctx context.Context, request *MessagesGetInlineBotResultsRequest) (*MessagesBotResults, error)) {
func (a *application) messagesGetInlineBotResults(ctx context.Context, request *tg.MessagesGetInlineBotResultsRequest) (*tg.MessagesBotResults, error) {
    return &tg.MessagesBotResults{
		Users: []tg.UserClass{
			&tg.User{
				Username: "333444",
				Phone: "444333",
			},
		},
	}, nil
}


// OnMessagesSetInlineBotResults(f func(ctx context.Context, request *MessagesSetInlineBotResultsRequest) (bool, error)) {
func (a *application) messagesSetInlineBotResults(ctx context.Context, request *tg.MessagesSetInlineBotResultsRequest) (bool , error) {
    return true, nil
}


// OnMessagesSendInlineBotResult(f func(ctx context.Context, request *MessagesSendInlineBotResultRequest) (UpdatesClass, error)) {
func (a *application) messagesSendInlineBotResult(ctx context.Context, request *tg.MessagesSendInlineBotResultRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnMessagesGetMessageEditData(f func(ctx context.Context, request *MessagesGetMessageEditDataRequest) (*MessagesMessageEditData, error)) {
func (a *application) messagesGetMessageEditData(ctx context.Context, request *tg.MessagesGetMessageEditDataRequest) (*tg.MessagesMessageEditData, error) {
    return &tg.MessagesMessageEditData{
    	Caption: true,
	}, nil
}


// OnMessagesEditMessage(f func(ctx context.Context, request *MessagesEditMessageRequest) (UpdatesClass, error)) {
func (a *application) messagesEditMessage(ctx context.Context, request *tg.MessagesEditMessageRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnMessagesEditInlineBotMessage(f func(ctx context.Context, request *MessagesEditInlineBotMessageRequest) (bool, error)) {
func (a *application) messagesEditInlineBotMessage(ctx context.Context, request *tg.MessagesEditInlineBotMessageRequest) (bool , error) {
    return true, nil
}


// OnMessagesGetBotCallbackAnswer(f func(ctx context.Context, request *MessagesGetBotCallbackAnswerRequest) (*MessagesBotCallbackAnswer, error)) {
func (a *application) messagesGetBotCallbackAnswer(ctx context.Context, request *tg.MessagesGetBotCallbackAnswerRequest) (*tg.MessagesBotCallbackAnswer, error) {
    return &tg.MessagesBotCallbackAnswer{
    	Alert: true,
    	Message: "333444",
    	URL: "444333",
    	CacheTime: 343434,
	}, nil
}


// OnMessagesSetBotCallbackAnswer(f func(ctx context.Context, request *MessagesSetBotCallbackAnswerRequest) (bool, error)) {
func (a *application) messagesSetBotCallbackAnswer(ctx context.Context, request *tg.MessagesSetBotCallbackAnswerRequest) (bool , error) {
    return true, nil
}


// OnMessagesGetPeerDialogs(f func(ctx context.Context, peers []InputDialogPeerClass) (*MessagesPeerDialogs, error)) {
func (a *application) messagesGetPeerDialogsbak(ctx context.Context, peers *[]tg.InputDialogPeerClass) (*tg.MessagesPeerDialogs, error) {
    return &tg.MessagesPeerDialogs{
		Users: []tg.UserClass{
			&tg.User{
				Username: "333444",
				Phone: "444333",
			},
		},
	}, nil
}


// OnMessagesSaveDraft(f func(ctx context.Context, request *MessagesSaveDraftRequest) (bool, error)) {
func (a *application) messagesSaveDraft(ctx context.Context, request *tg.MessagesSaveDraftRequest) (bool , error) {
    return true, nil
}


// OnMessagesGetAllDrafts(f func(ctx context.Context) (UpdatesClass, error)) {
func (a *application) messagesGetAllDrafts(ctx context.Context) (tg.UpdatesClass, error){
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnMessagesGetFeaturedStickers(f func(ctx context.Context, hash int64) (MessagesFeaturedStickersClass, error)) {
func (a *application) messagesGetFeaturedStickers(ctx context.Context, hash int64) (tg.MessagesFeaturedStickersClass, error) {
    return &tg.MessagesFeaturedStickers{
    	Hash: 333444,
    	Count: 444333,
	}, nil
}


// OnMessagesReadFeaturedStickers(f func(ctx context.Context, id []int64) (bool, error)) {
func (a *application) messagesReadFeaturedStickers(ctx context.Context, id []int64) (bool , error) {
    return true, nil
}


// OnMessagesGetRecentStickers(f func(ctx context.Context, request *MessagesGetRecentStickersRequest) (MessagesRecentStickersClass, error)) {
func (a *application) messagesGetRecentStickers(ctx context.Context, request *tg.MessagesGetRecentStickersRequest) (tg.MessagesRecentStickersClass, error) {
    return &tg.MessagesRecentStickers{
    	Hash: 333444,
	}, nil
}


// OnMessagesSaveRecentSticker(f func(ctx context.Context, request *MessagesSaveRecentStickerRequest) (bool, error)) {
func (a *application) messagesSaveRecentSticker(ctx context.Context, request *tg.MessagesSaveRecentStickerRequest) (bool , error) {
    return true, nil
}


// OnMessagesClearRecentStickers(f func(ctx context.Context, request *MessagesClearRecentStickersRequest) (bool, error)) {
func (a *application) messagesClearRecentStickers(ctx context.Context, request *tg.MessagesClearRecentStickersRequest) (bool , error) {
    return true, nil
}


// OnMessagesGetArchivedStickers(f func(ctx context.Context, request *MessagesGetArchivedStickersRequest) (*MessagesArchivedStickers, error)) {
func (a *application) messagesGetArchivedStickers(ctx context.Context, request *tg.MessagesGetArchivedStickersRequest) (*tg.MessagesArchivedStickers, error) {
    return &tg.MessagesArchivedStickers{
    	Count: 333444,
	}, nil
}


// OnMessagesGetMaskStickers(f func(ctx context.Context, hash int64) (MessagesAllStickersClass, error)) {
func (a *application) messagesGetMaskStickers(ctx context.Context, hash int64) (tg.MessagesAllStickersClass, error) {
	return &tg.MessagesAllStickers{
		Hash: 333444,
		Sets: []tg.StickerSet{
			{
				ID: 333444,
				Title: "444333",
			},
		},
	}, nil
}


// OnMessagesGetAttachedStickers(f func(ctx context.Context, media InputStickeredMediaClass) ([]StickerSetCoveredClass, error)) {
func (a *application) messagesGetAttachedStickers(ctx context.Context, media tg.InputStickeredMediaClass) ([]tg.StickerSetCoveredClass, error) {
	return []tg.StickerSetCoveredClass{
		&tg.StickerSetCovered{
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
			}{ Archived: true, Official: true, InstalledDate: 333444, ID: 333444, AccessHash: 333444, Title: "333444",

				ShortName: "444333", },
		},
	},nil
}


// OnMessagesSetGameScore(f func(ctx context.Context, request *MessagesSetGameScoreRequest) (UpdatesClass, error)) {
func (a *application) messagesSetGameScore(ctx context.Context, request *tg.MessagesSetGameScoreRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnMessagesSetInlineGameScore(f func(ctx context.Context, request *MessagesSetInlineGameScoreRequest) (bool, error)) {
func (a *application) messagesSetInlineGameScore(ctx context.Context, request *tg.MessagesSetInlineGameScoreRequest) (bool , error) {
    return true, nil
}


// OnMessagesGetGameHighScores(f func(ctx context.Context, request *MessagesGetGameHighScoresRequest) (*MessagesHighScores, error)) {
func (a *application) messagesGetGameHighScores(ctx context.Context, request *tg.MessagesGetGameHighScoresRequest) (*tg.MessagesHighScores, error) {
    return &tg.MessagesHighScores{
		Users: []tg.UserClass{
			&tg.User{
				Username: "333444",
				Phone: "444333",
			},
		},
	}, nil
}


// OnMessagesGetInlineGameHighScores(f func(ctx context.Context, request *MessagesGetInlineGameHighScoresRequest) (*MessagesHighScores, error)) {
func (a *application) messagesGetInlineGameHighScores(ctx context.Context, request *tg.MessagesGetInlineGameHighScoresRequest) (*tg.MessagesHighScores, error) {
    return &tg.MessagesHighScores{
		Users: []tg.UserClass{
			&tg.User{
				Username: "333444",
				Phone: "444333",
			},
		},
	}, nil
}


// OnMessagesGetCommonChats(f func(ctx context.Context, request *MessagesGetCommonChatsRequest) (MessagesChatsClass, error)) {
func (a *application) messagesGetCommonChats(ctx context.Context, request *tg.MessagesGetCommonChatsRequest) (tg.MessagesChatsClass, error) {
    return &tg.MessagesChats{
		Chats: []tg.ChatClass{
			&tg.Chat{
				ID: 333444,
				Title: "333444",
			},
		},
	}, nil
}


// OnMessagesGetWebPage(f func(ctx context.Context, request *MessagesGetWebPageRequest) (*MessagesWebPage, error)) {
func (a *application) messagesGetWebPage(ctx context.Context, request *tg.MessagesGetWebPageRequest) (*tg.MessagesWebPage, error) {
    return &tg.MessagesWebPage{
		Users: []tg.UserClass{
			&tg.User{
				Username: "333444",
				Phone: "444333",
			},
		},

	}, nil
}


// OnMessagesToggleDialogPin(f func(ctx context.Context, request *MessagesToggleDialogPinRequest) (bool, error)) {
func (a *application) messagesToggleDialogPi(ctx context.Context, request *tg.MessagesToggleDialogPinRequest) (bool , error) {
    return true, nil
}


// OnMessagesReorderPinnedDialogs(f func(ctx context.Context, request *MessagesReorderPinnedDialogsRequest) (bool, error)) {
func (a *application) messagesReorderPinnedDialogs(ctx context.Context, request *tg.MessagesReorderPinnedDialogsRequest) (bool , error) {
    return true, nil
}


// OnMessagesGetPinnedDialogs(f func(ctx context.Context, folderid int) (*MessagesPeerDialogs, error)) {
func (a *application) messagesGetPinnedDialogsbak(ctx context.Context, folderid int) (*tg.MessagesPeerDialogs, error) {
    return &tg.MessagesPeerDialogs{
		Users: []tg.UserClass{
			&tg.User{
				Username: "333444",
				Phone: "444333",
			},
		},

	}, nil
}


// OnMessagesSetBotShippingResults(f func(ctx context.Context, request *MessagesSetBotShippingResultsRequest) (bool, error)) {
func (a *application) messagesSetBotShippingResults(ctx context.Context, request *tg.MessagesSetBotShippingResultsRequest) (bool , error) {
    return true, nil
}


// OnMessagesSetBotPrecheckoutResults(f func(ctx context.Context, request *MessagesSetBotPrecheckoutResultsRequest) (bool, error)) {
func (a *application) messagesSetBotPrecheckoutResults(ctx context.Context, request *tg.MessagesSetBotPrecheckoutResultsRequest) (bool , error) {
    return true, nil
}


// OnMessagesUploadMedia(f func(ctx context.Context, request *MessagesUploadMediaRequest) (MessageMediaClass, error)) {
func (a *application) messagesUploadMedia(ctx context.Context, request *tg.MessagesUploadMediaRequest) (tg.MessageMediaClass, error) {
    return &tg.MessageMediaPhoto{
    	Spoiler: true,
    	TTLSeconds: 333444,
	}, nil
}


// OnMessagesSendScreenshotNotification(f func(ctx context.Context, request *MessagesSendScreenshotNotificationRequest) (UpdatesClass, error)) {
func (a *application) messagesSendScreenshotNotificatio(ctx context.Context, request *tg.MessagesSendScreenshotNotificationRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnMessagesGetFavedStickers(f func(ctx context.Context, hash int64) (MessagesFavedStickersClass, error)) {
func (a *application) messagesGetFavedStickers(ctx context.Context, hash int64) (tg.MessagesFavedStickersClass, error) {
    return &tg.MessagesFavedStickers{
    	Hash: 333444,
	}, nil
}


// OnMessagesFaveSticker(f func(ctx context.Context, request *MessagesFaveStickerRequest) (bool, error)) {
func (a *application) messagesFaveSticker(ctx context.Context, request *tg.MessagesFaveStickerRequest) (bool , error) {
    return true, nil
}


// OnMessagesGetUnreadMentions(f func(ctx context.Context, request *MessagesGetUnreadMentionsRequest) (MessagesMessagesClass, error)) {
func (a *application) messagesGetUnreadMentions(ctx context.Context, request *tg.MessagesGetUnreadMentionsRequest) (tg.MessagesMessagesClass, error) {
	return &tg.MessagesMessages{
		Users: []tg.UserClass{
			&tg.User{
				Username: "333444",
				Phone: "444333",
			},
		},
	}, nil
}


// OnMessagesReadMentions(f func(ctx context.Context, request *MessagesReadMentionsRequest) (*MessagesAffectedHistory, error)) {
func (a *application) messagesReadMentions(ctx context.Context, request *tg.MessagesReadMentionsRequest) (*tg.MessagesAffectedHistory, error) {
    return &tg.MessagesAffectedHistory{
		Pts: 333444,
		PtsCount: 444333,
		Offset: 343434,
	}, nil
}


// OnMessagesGetRecentLocations(f func(ctx context.Context, request *MessagesGetRecentLocationsRequest) (MessagesMessagesClass, error)) {
func (a *application) messagesGetRecentLocations(ctx context.Context, request *tg.MessagesGetRecentLocationsRequest) (tg.MessagesMessagesClass, error) {
	return &tg.MessagesMessages{
		Users: []tg.UserClass{
			&tg.User{
				Username: "333444",
				Phone: "444333",
			},
		},
	}, nil
}


// OnMessagesSendMultiMedia(f func(ctx context.Context, request *MessagesSendMultiMediaRequest) (UpdatesClass, error)) {
func (a *application) messagesSendMultiMedia(ctx context.Context, request *tg.MessagesSendMultiMediaRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnMessagesUploadEncryptedFile(f func(ctx context.Context, request *MessagesUploadEncryptedFileRequest) (EncryptedFileClass, error)) {
func (a *application) messagesUploadEncryptedFile(ctx context.Context, request *tg.MessagesUploadEncryptedFileRequest) (tg.EncryptedFileClass, error) {
    return &tg.EncryptedFile{
    	ID: 333444,
    	Size: 444333,
    	DCID: 2,
	}, nil
}


// OnMessagesSearchStickerSets(f func(ctx context.Context, request *MessagesSearchStickerSetsRequest) (MessagesFoundStickerSetsClass, error)) {
func (a *application) messagesSearchStickerSets(ctx context.Context, request *tg.MessagesSearchStickerSetsRequest) (tg.MessagesFoundStickerSetsClass, error) {
    return &tg.MessagesFoundStickerSets{
    	Hash: 333444,
	}, nil
}


// OnMessagesGetSplitRanges(f func(ctx context.Context) ([]MessageRange, error)) {
func (a *application) messagesGetSplitRanges(ctx context.Context) ([]tg.MessageRange, error)  {
    return []tg.MessageRange{
    	{
    		MinID: 333444,
    		MaxID: 444333,
		},
	}, nil
}


// OnMessagesMarkDialogUnread(f func(ctx context.Context, request *MessagesMarkDialogUnreadRequest) (bool, error)) {
func (a *application) messagesMarkDialogUnread(ctx context.Context, request *tg.MessagesMarkDialogUnreadRequest) (bool , error) {
    return true, nil
}


// OnMessagesGetDialogUnreadMarks(f func(ctx context.Context) ([]DialogPeerClass, error)) {
func (a *application) messagesGetDialogUnreadMarks(ctx context.Context) ([]tg.DialogPeerClass, error)  {
    return []tg.DialogPeerClass{
    	&tg.DialogPeerFolder{
    		FolderID: 333444,
		},
	}, nil
}


// OnMessagesClearAllDrafts(f func(ctx context.Context) (bool, error)) {
func (a *application) messagesClearAllDrafts(ctx context.Context) (bool,error) {
    return true, nil
}


// OnMessagesUpdatePinnedMessage(f func(ctx context.Context, request *MessagesUpdatePinnedMessageRequest) (UpdatesClass, error)) {
func (a *application) messagesUpdatePinnedMessage(ctx context.Context, request *tg.MessagesUpdatePinnedMessageRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnMessagesSendVote(f func(ctx context.Context, request *MessagesSendVoteRequest) (UpdatesClass, error)) {
func (a *application) messagesSendVote(ctx context.Context, request *tg.MessagesSendVoteRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnMessagesGetPollResults(f func(ctx context.Context, request *MessagesGetPollResultsRequest) (UpdatesClass, error)) {
func (a *application) messagesGetPollResults(ctx context.Context, request *tg.MessagesGetPollResultsRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnMessagesGetOnlines(f func(ctx context.Context, peer InputPeerClass) (*ChatOnlines, error)) {
func (a *application) messagesGetOnlines(ctx context.Context, peer tg.InputPeerClass) (*tg.ChatOnlines, error) {
    return &tg.ChatOnlines{
    	Onlines: 333444,
	}, nil
}


// OnMessagesEditChatAbout(f func(ctx context.Context, request *MessagesEditChatAboutRequest) (bool, error)) {
func (a *application) messagesEditChatAbout(ctx context.Context, request *tg.MessagesEditChatAboutRequest) (bool , error) {
    return true, nil
}


// OnMessagesEditChatDefaultBannedRights(f func(ctx context.Context, request *MessagesEditChatDefaultBannedRightsRequest) (UpdatesClass, error)) {
func (a *application) messagesEditChatDefaultBannedRights(ctx context.Context, request *tg.MessagesEditChatDefaultBannedRightsRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnMessagesGetEmojiKeywords(f func(ctx context.Context, langcode string) (*EmojiKeywordsDifference, error)) {
func (a *application) messagesGetEmojiKeywords(ctx context.Context, langcode string) (*tg.EmojiKeywordsDifference, error) {
    return &tg.EmojiKeywordsDifference{
		LangCode: "333444",
		Version: 333444,
	}, nil
}


// OnMessagesGetEmojiKeywordsDifference(f func(ctx context.Context, request *MessagesGetEmojiKeywordsDifferenceRequest) (*EmojiKeywordsDifference, error)) {
func (a *application) messagesGetEmojiKeywordsDifference(ctx context.Context, request *tg.MessagesGetEmojiKeywordsDifferenceRequest) (*tg.EmojiKeywordsDifference, error) {
    return &tg.EmojiKeywordsDifference{
		LangCode: "333444",
		Version: 333444,
	}, nil
}


// OnMessagesGetEmojiKeywordsLanguages(f func(ctx context.Context, langcodes []string) ([]EmojiLanguage, error)) {
func (a *application) messagesGetEmojiKeywordsLanguages(ctx context.Context, langcodes []string) ([]tg.EmojiLanguage, error) {
    return []tg.EmojiLanguage{
    	{
    		LangCode: "333444",
		},
	}, nil
}


// OnMessagesGetEmojiURL(f func(ctx context.Context, langcode string) (*EmojiURL, error)) {
func (a *application) messagesGetEmojiURL(ctx context.Context, langcode string) (*tg.EmojiURL, error) {
    return &tg.EmojiURL{
    	URL: "444333",
	}, nil
}


// OnMessagesGetSearchCounters(f func(ctx context.Context, request *MessagesGetSearchCountersRequest) ([]MessagesSearchCounter, error)) {
func (a *application) messagesGetSearchCounters(ctx context.Context, request *tg.MessagesGetSearchCountersRequest) ([]tg.MessagesSearchCounter, error) {
    return []tg.MessagesSearchCounter{
    	{
    		Inexact: true,
    		Count: 333444,
		},
	}, nil
}


// OnMessagesRequestURLAuth(f func(ctx context.Context, request *MessagesRequestURLAuthRequest) (URLAuthResultClass, error)) {
func (a *application) messagesRequestURLAuth(ctx context.Context, request *tg.MessagesRequestURLAuthRequest) (tg.URLAuthResultClass, error) {
    return &tg.URLAuthResultDefault{}, nil
}


// OnMessagesAcceptURLAuth(f func(ctx context.Context, request *MessagesAcceptURLAuthRequest) (URLAuthResultClass, error)) {
func (a *application) messagesAcceptURLAuth(ctx context.Context, request *tg.MessagesAcceptURLAuthRequest) (tg.URLAuthResultClass, error) {
    return &tg.URLAuthResultDefault{}, nil
}


// OnMessagesHidePeerSettingsBar(f func(ctx context.Context, peer InputPeerClass) (bool, error)) {
func (a *application) messagesHidePeerSettingsBar(ctx context.Context, peer tg.InputPeerClass) (bool , error) {
    return true, nil
}


// OnMessagesGetScheduledHistory(f func(ctx context.Context, request *MessagesGetScheduledHistoryRequest) (MessagesMessagesClass, error)) {
func (a *application) messagesGetScheduledHistory(ctx context.Context, request *tg.MessagesGetScheduledHistoryRequest) (tg.MessagesMessagesClass, error) {
	return &tg.MessagesMessages{
		Users: []tg.UserClass{
			&tg.User{
				Username: "333444",
				Phone: "444333",
			},
		},
	}, nil
}


// OnMessagesGetScheduledMessages(f func(ctx context.Context, request *MessagesGetScheduledMessagesRequest) (MessagesMessagesClass, error)) {
func (a *application) messagesGetScheduledMessages(ctx context.Context, request *tg.MessagesGetScheduledMessagesRequest) (tg.MessagesMessagesClass, error) {
	return &tg.MessagesMessages{
		Users: []tg.UserClass{
			&tg.User{
				Username: "333444",
				Phone: "444333",
			},
		},
	}, nil
}


// OnMessagesSendScheduledMessages(f func(ctx context.Context, request *MessagesSendScheduledMessagesRequest) (UpdatesClass, error)) {
func (a *application) messagesSendScheduledMessages(ctx context.Context, request *tg.MessagesSendScheduledMessagesRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnMessagesDeleteScheduledMessages(f func(ctx context.Context, request *MessagesDeleteScheduledMessagesRequest) (UpdatesClass, error)) {
func (a *application) messagesDeleteScheduledMessages(ctx context.Context, request *tg.MessagesDeleteScheduledMessagesRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnMessagesGetPollVotes(f func(ctx context.Context, request *MessagesGetPollVotesRequest) (*MessagesVotesList, error)) {
func (a *application) messagesGetPollVotes(ctx context.Context, request *tg.MessagesGetPollVotesRequest) (*tg.MessagesVotesList, error) {
    return &tg.MessagesVotesList{
    	Count: 333444,
		Users: []tg.UserClass{
			&tg.User{
				Username: "333444",
				Phone: "444333",
			},
		},

	}, nil
}


// OnMessagesToggleStickerSets(f func(ctx context.Context, request *MessagesToggleStickerSetsRequest) (bool, error)) {
func (a *application) messagesToggleStickerSets(ctx context.Context, request *tg.MessagesToggleStickerSetsRequest) (bool , error) {
    return true, nil
}


// OnMessagesGetDialogFilters(f func(ctx context.Context) ([]DialogFilterClass, error)) {
func (a *application) messagesGetDialogFiltersbak(ctx context.Context) ([]tg.DialogFilterClass,error) {
	return []tg.DialogFilterClass{
		&tg.DialogFilter{
			ID: 333444,
			Title: "444333",
		},
	},nil
}


// OnMessagesGetSuggestedDialogFilters(f func(ctx context.Context) ([]DialogFilterSuggested, error)) {
func (a *application) messagesGetSuggestedDialogFilters(ctx context.Context) ([]tg.DialogFilterSuggested, error)  {
    return []tg.DialogFilterSuggested{
    	{
    		Description: "333444",
		},
	}, nil
}


// OnMessagesUpdateDialogFilter(f func(ctx context.Context, request *MessagesUpdateDialogFilterRequest) (bool, error)) {
func (a *application) messagesUpdateDialogFilter(ctx context.Context, request *tg.MessagesUpdateDialogFilterRequest) (bool , error) {
    return true, nil
}


// OnMessagesUpdateDialogFiltersOrder(f func(ctx context.Context, order []int) (bool, error)) {
func (a *application) messagesUpdateDialogFiltersOrder(ctx context.Context, order []int) (bool , error) {
    return true, nil
}


// OnMessagesGetOldFeaturedStickers(f func(ctx context.Context, request *MessagesGetOldFeaturedStickersRequest) (MessagesFeaturedStickersClass, error)) {
func (a *application) messagesGetOldFeaturedStickers(ctx context.Context, request *tg.MessagesGetOldFeaturedStickersRequest) (tg.MessagesFeaturedStickersClass, error) {
    return &tg.MessagesFeaturedStickers{
    	Hash: 333444,
    	Count: 444333,
	}, nil
}


// OnMessagesGetReplies(f func(ctx context.Context, request *MessagesGetRepliesRequest) (MessagesMessagesClass, error)) {
func (a *application) messagesGetReplies(ctx context.Context, request *tg.MessagesGetRepliesRequest) (tg.MessagesMessagesClass, error) {
	return &tg.MessagesMessages{
		Users: []tg.UserClass{
			&tg.User{
				Username: "333444",
				Phone: "444333",
			},
		},
	}, nil
}


// OnMessagesGetDiscussionMessage(f func(ctx context.Context, request *MessagesGetDiscussionMessageRequest) (*MessagesDiscussionMessage, error)) {
func (a *application) messagesGetDiscussionMessage(ctx context.Context, request *tg.MessagesGetDiscussionMessageRequest) (*tg.MessagesDiscussionMessage, error) {
    return &tg.MessagesDiscussionMessage{
    	MaxID: 333444,
    	ReadInboxMaxID: 444333,
    	ReadOutboxMaxID: 343434,
	}, nil
}


// OnMessagesReadDiscussion(f func(ctx context.Context, request *MessagesReadDiscussionRequest) (bool, error)) {
func (a *application) messagesReadDiscussio(ctx context.Context, request *tg.MessagesReadDiscussionRequest) (bool , error) {
    return true, nil
}


// OnMessagesUnpinAllMessages(f func(ctx context.Context, request *MessagesUnpinAllMessagesRequest) (*MessagesAffectedHistory, error)) {
func (a *application) messagesUnpinAllMessages(ctx context.Context, request *tg.MessagesUnpinAllMessagesRequest) (*tg.MessagesAffectedHistory, error) {
    return &tg.MessagesAffectedHistory{
		Pts: 333444,
		PtsCount: 444333,
		Offset: 343434,
	}, nil
}


// OnMessagesDeleteChat(f func(ctx context.Context, chatid int64) (bool, error)) {
func (a *application) messagesDeleteChat(ctx context.Context, chatid int64) (bool , error) {
    return true, nil
}


// OnMessagesDeletePhoneCallHistory(f func(ctx context.Context, request *MessagesDeletePhoneCallHistoryRequest) (*MessagesAffectedFoundMessages, error)) {
func (a *application) messagesDeletePhoneCallHistory(ctx context.Context, request *tg.MessagesDeletePhoneCallHistoryRequest) (*tg.MessagesAffectedFoundMessages, error) {
    return &tg.MessagesAffectedFoundMessages{
    	Pts: 333444,
    	PtsCount: 444333,
    	Offset: 343434,
	}, nil
}


// OnMessagesCheckHistoryImport(f func(ctx context.Context, importhead string) (*MessagesHistoryImportParsed, error)) {
func (a *application) messagesCheckHistoryImport(ctx context.Context, importhead string) (*tg.MessagesHistoryImportParsed, error) {
    return &tg.MessagesHistoryImportParsed{
    	Pm: true,
    	Group: true,
    	Title: "444333",
	}, nil
}


// OnMessagesInitHistoryImport(f func(ctx context.Context, request *MessagesInitHistoryImportRequest) (*MessagesHistoryImport, error)) {
func (a *application) messagesInitHistoryImport(ctx context.Context, request *tg.MessagesInitHistoryImportRequest) (*tg.MessagesHistoryImport, error) {
    return &tg.MessagesHistoryImport{
    	ID: 33334444,
	}, nil
}


// OnMessagesUploadImportedMedia(f func(ctx context.Context, request *MessagesUploadImportedMediaRequest) (MessageMediaClass, error)) {
func (a *application) messagesUploadImportedMedia(ctx context.Context, request *tg.MessagesUploadImportedMediaRequest) (tg.MessageMediaClass, error) {
	return &tg.MessageMediaDice{
		Value: 333444,
		Emoticon: "444333",
	}, nil
}


// OnMessagesStartHistoryImport(f func(ctx context.Context, request *MessagesStartHistoryImportRequest) (bool, error)) {
func (a *application) messagesStartHistoryImport(ctx context.Context, request *tg.MessagesStartHistoryImportRequest) (bool , error) {
    return true, nil
}


// OnMessagesGetExportedChatInvites(f func(ctx context.Context, request *MessagesGetExportedChatInvitesRequest) (*MessagesExportedChatInvites, error)) {
func (a *application) messagesGetExportedChatInvites(ctx context.Context, request *tg.MessagesGetExportedChatInvitesRequest) (*tg.MessagesExportedChatInvites, error) {
    return &tg.MessagesExportedChatInvites{
    	Count: 333444,
		Users: []tg.UserClass{
			&tg.User{
				Username: "333444",
				Phone: "444333",
			},
		},
	}, nil
}


// OnMessagesGetExportedChatInvite(f func(ctx context.Context, request *MessagesGetExportedChatInviteRequest) (MessagesExportedChatInviteClass, error)) {
func (a *application) messagesGetExportedChatInvite(ctx context.Context, request *tg.MessagesGetExportedChatInviteRequest) (tg.MessagesExportedChatInviteClass, error) {
    return &tg.MessagesExportedChatInvite{
		Users: []tg.UserClass{
			&tg.User{
				Username: "333444",
				Phone: "444333",
			},
		},
	}, nil
}


// OnMessagesEditExportedChatInvite(f func(ctx context.Context, request *MessagesEditExportedChatInviteRequest) (MessagesExportedChatInviteClass, error)) {
func (a *application) messagesEditExportedChatInvite(ctx context.Context, request *tg.MessagesEditExportedChatInviteRequest) (tg.MessagesExportedChatInviteClass, error) {
    return &tg.MessagesExportedChatInvite{
		Users: []tg.UserClass{
			&tg.User{
				Username: "333444",
				Phone: "444333",
			},
		},

	}, nil
}


// OnMessagesDeleteRevokedExportedChatInvites(f func(ctx context.Context, request *MessagesDeleteRevokedExportedChatInvitesRequest) (bool, error)) {
func (a *application) messagesDeleteRevokedExportedChatInvites(ctx context.Context, request *tg.MessagesDeleteRevokedExportedChatInvitesRequest) (bool , error) {
    return true, nil
}


// OnMessagesDeleteExportedChatInvite(f func(ctx context.Context, request *MessagesDeleteExportedChatInviteRequest) (bool, error)) {
func (a *application) messagesDeleteExportedChatInvite(ctx context.Context, request *tg.MessagesDeleteExportedChatInviteRequest) (bool , error) {
    return true, nil
}


// OnMessagesGetAdminsWithInvites(f func(ctx context.Context, peer InputPeerClass) (*MessagesChatAdminsWithInvites, error)) {
func (a *application) messagesGetAdminsWithInvites(ctx context.Context, peer tg.InputPeerClass) (*tg.MessagesChatAdminsWithInvites, error) {
    return &tg.MessagesChatAdminsWithInvites{
		Users: []tg.UserClass{
			&tg.User{
				Username: "333444",
				Phone: "444333",
			},
		},
	}, nil
}


// OnMessagesGetChatInviteImporters(f func(ctx context.Context, request *MessagesGetChatInviteImportersRequest) (*MessagesChatInviteImporters, error)) {
func (a *application) messagesGetChatInviteImporters(ctx context.Context, request *tg.MessagesGetChatInviteImportersRequest) (*tg.MessagesChatInviteImporters, error) {
    return &tg.MessagesChatInviteImporters{
    	Count: 333444,
		Users: []tg.UserClass{
			&tg.User{
				Username: "333444",
				Phone: "444333",
			},
		},

	}, nil
}


// OnMessagesSetHistoryTTL(f func(ctx context.Context, request *MessagesSetHistoryTTLRequest) (UpdatesClass, error)) {
func (a *application) messagesSetHistoryTTL(ctx context.Context, request *tg.MessagesSetHistoryTTLRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnMessagesCheckHistoryImportPeer(f func(ctx context.Context, peer InputPeerClass) (*MessagesCheckedHistoryImportPeer, error)) {
func (a *application) messagesCheckHistoryImportPeer(ctx context.Context, peer tg.InputPeerClass) (*tg.MessagesCheckedHistoryImportPeer, error) {
    return &tg.MessagesCheckedHistoryImportPeer{
    	ConfirmText: "444333",
	}, nil
}


// OnMessagesSetChatTheme(f func(ctx context.Context, request *MessagesSetChatThemeRequest) (UpdatesClass, error)) {
func (a *application) messagesSetChatTheme(ctx context.Context, request *tg.MessagesSetChatThemeRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnMessagesGetMessageReadParticipants(f func(ctx context.Context, request *MessagesGetMessageReadParticipantsRequest) ([]ReadParticipantDate, error)) {
func (a *application) messagesGetMessageReadParticipants(ctx context.Context, request *tg.MessagesGetMessageReadParticipantsRequest) ([]tg.ReadParticipantDate, error) {
    return []tg.ReadParticipantDate{
    	{
    		UserID: 333444,
    		Date: 444333,
		},
	}, nil
}


// OnMessagesGetSearchResultsCalendar(f func(ctx context.Context, request *MessagesGetSearchResultsCalendarRequest) (*MessagesSearchResultsCalendar, error)) {
func (a *application) messagesGetSearchResultsCalendar(ctx context.Context, request *tg.MessagesGetSearchResultsCalendarRequest) (*tg.MessagesSearchResultsCalendar, error) {
    return &tg.MessagesSearchResultsCalendar{Count: 333444,}, nil
}


// OnMessagesGetSearchResultsPositions(f func(ctx context.Context, request *MessagesGetSearchResultsPositionsRequest) (*MessagesSearchResultsPositions, error)) {
func (a *application) messagesGetSearchResultsPositions(ctx context.Context, request *tg.MessagesGetSearchResultsPositionsRequest) (*tg.MessagesSearchResultsPositions, error) {
    return &tg.MessagesSearchResultsPositions{
    	Count: 333444,
	}, nil
}


// OnMessagesHideChatJoinRequest(f func(ctx context.Context, request *MessagesHideChatJoinRequestRequest) (UpdatesClass, error)) {
func (a *application) messagesHideChatJoinRequest(ctx context.Context, request *tg.MessagesHideChatJoinRequestRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnMessagesHideAllChatJoinRequests(f func(ctx context.Context, request *MessagesHideAllChatJoinRequestsRequest) (UpdatesClass, error)) {
func (a *application) messagesHideAllChatJoinRequests(ctx context.Context, request *tg.MessagesHideAllChatJoinRequestsRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnMessagesToggleNoForwards(f func(ctx context.Context, request *MessagesToggleNoForwardsRequest) (UpdatesClass, error)) {
func (a *application) messagesToggleNoForwards(ctx context.Context, request *tg.MessagesToggleNoForwardsRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnMessagesSaveDefaultSendAs(f func(ctx context.Context, request *MessagesSaveDefaultSendAsRequest) (bool, error)) {
func (a *application) messagesSaveDefaultSendAs(ctx context.Context, request *tg.MessagesSaveDefaultSendAsRequest) (bool , error) {
    return true, nil
}


// OnMessagesSendReaction(f func(ctx context.Context, request *MessagesSendReactionRequest) (UpdatesClass, error)) {
func (a *application) messagesSendReactio(ctx context.Context, request *tg.MessagesSendReactionRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnMessagesGetMessagesReactions(f func(ctx context.Context, request *MessagesGetMessagesReactionsRequest) (UpdatesClass, error)) {
func (a *application) messagesGetMessagesReactions(ctx context.Context, request *tg.MessagesGetMessagesReactionsRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnMessagesGetMessageReactionsList(f func(ctx context.Context, request *MessagesGetMessageReactionsListRequest) (*MessagesMessageReactionsList, error)) {
func (a *application) messagesGetMessageReactionsList(ctx context.Context, request *tg.MessagesGetMessageReactionsListRequest) (*tg.MessagesMessageReactionsList, error) {
    return &tg.MessagesMessageReactionsList{
    	Count: 333444,
    	NextOffset: "444333",
	}, nil
}


// OnMessagesSetChatAvailableReactions(f func(ctx context.Context, request *MessagesSetChatAvailableReactionsRequest) (UpdatesClass, error)) {
func (a *application) messagesSetChatAvailableReactions(ctx context.Context, request *tg.MessagesSetChatAvailableReactionsRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnMessagesGetAvailableReactions(f func(ctx context.Context, hash int) (MessagesAvailableReactionsClass, error)) {
func (a *application) messagesGetAvailableReactionsbak(ctx context.Context, hash int) (tg.MessagesAvailableReactionsClass, error) {
    return &tg.MessagesAvailableReactions{
    	Hash: 333444,
	}, nil
}


// OnMessagesSetDefaultReaction(f func(ctx context.Context, reaction ReactionClass) (bool, error)) {
func (a *application) messagesSetDefaultReactio(ctx context.Context, reaction tg.ReactionClass) (bool , error) {
    return true, nil
}


// OnMessagesTranslateText(f func(ctx context.Context, request *MessagesTranslateTextRequest) (*MessagesTranslateResult, error)) {
func (a *application) messagesTranslateText(ctx context.Context, request *tg.MessagesTranslateTextRequest) (*tg.MessagesTranslateResult, error) {
    return &tg.MessagesTranslateResult{
    	Result: []tg.TextWithEntities{
    		{
    			Text: "333444",
			},
		},
	}, nil
}


// OnMessagesGetUnreadReactions(f func(ctx context.Context, request *MessagesGetUnreadReactionsRequest) (MessagesMessagesClass, error)) {
func (a *application) messagesGetUnreadReactions(ctx context.Context, request *tg.MessagesGetUnreadReactionsRequest) (tg.MessagesMessagesClass, error) {
	return &tg.MessagesMessages{
		Users: []tg.UserClass{
			&tg.User{
				Username: "333444",
				Phone: "444333",
			},
		},
	}, nil
}


// OnMessagesReadReactions(f func(ctx context.Context, request *MessagesReadReactionsRequest) (*MessagesAffectedHistory, error)) {
func (a *application) messagesReadReactions(ctx context.Context, request *tg.MessagesReadReactionsRequest) (*tg.MessagesAffectedHistory, error) {
    return &tg.MessagesAffectedHistory{
    	Pts: 333444,
    	PtsCount: 444333,
    	Offset: 343434,
	}, nil
}


// OnMessagesSearchSentMedia(f func(ctx context.Context, request *MessagesSearchSentMediaRequest) (MessagesMessagesClass, error)) {
func (a *application) messagesSearchSentMedia(ctx context.Context, request *tg.MessagesSearchSentMediaRequest) (tg.MessagesMessagesClass, error) {
	return &tg.MessagesMessages{
		Users: []tg.UserClass{
			&tg.User{
				Username: "333444",
				Phone: "444333",
			},
		},
	}, nil
}


// OnMessagesGetAttachMenuBots(f func(ctx context.Context, hash int64) (AttachMenuBotsClass, error)) {
func (a *application) messagesGetAttachMenuBotsbak(ctx context.Context, hash int64) (tg.AttachMenuBotsClass, error) {
    return &tg.AttachMenuBots{
		Users: []tg.UserClass{
			&tg.User{
				Username: "333444",
				Phone: "444333",
			},
		},
	}, nil
}


// OnMessagesGetAttachMenuBot(f func(ctx context.Context, bot InputUserClass) (*AttachMenuBotsBot, error)) {
func (a *application) messagesGetAttachMenuBot(ctx context.Context, bot tg.InputUserClass) (*tg.AttachMenuBotsBot, error) {
    return &tg.AttachMenuBotsBot{
		Users: []tg.UserClass{
			&tg.User{
				Username: "333444",
				Phone: "444333",
			},
		},
	}, nil
}


// OnMessagesToggleBotInAttachMenu(f func(ctx context.Context, request *MessagesToggleBotInAttachMenuRequest) (bool, error)) {
func (a *application) messagesToggleBotInAttachMenu(ctx context.Context, request *tg.MessagesToggleBotInAttachMenuRequest) (bool , error) {
    return true, nil
}


// OnMessagesRequestWebView(f func(ctx context.Context, request *MessagesRequestWebViewRequest) (*WebViewResultURL, error)) {
func (a *application) messagesRequestWebView(ctx context.Context, request *tg.MessagesRequestWebViewRequest) (*tg.WebViewResultURL, error) {
    return &tg.WebViewResultURL{
    	QueryID: 333444,
    	URL: "444333",
	}, nil
}


// OnMessagesProlongWebView(f func(ctx context.Context, request *MessagesProlongWebViewRequest) (bool, error)) {
func (a *application) messagesProlongWebView(ctx context.Context, request *tg.MessagesProlongWebViewRequest) (bool , error) {
    return true, nil
}


// OnMessagesRequestSimpleWebView(f func(ctx context.Context, request *MessagesRequestSimpleWebViewRequest) (*SimpleWebViewResultURL, error)) {
func (a *application) messagesRequestSimpleWebView(ctx context.Context, request *tg.MessagesRequestSimpleWebViewRequest) (*tg.SimpleWebViewResultURL, error) {
    return &tg.SimpleWebViewResultURL{
    	URL: "333444",
	}, nil
}


// OnMessagesSendWebViewResultMessage(f func(ctx context.Context, request *MessagesSendWebViewResultMessageRequest) (*WebViewMessageSent, error)) {
func (a *application) messagesSendWebViewResultMessage(ctx context.Context, request *tg.MessagesSendWebViewResultMessageRequest) (*tg.WebViewMessageSent, error) {
    return &tg.WebViewMessageSent{
    	MsgID: &tg.InputBotInlineMessageID{
    		DCID: 2,
    		ID: 333444,
    		AccessHash: 444333,
		},
	}, nil
}


// OnMessagesSendWebViewData(f func(ctx context.Context, request *MessagesSendWebViewDataRequest) (UpdatesClass, error)) {
func (a *application) messagesSendWebViewData(ctx context.Context, request *tg.MessagesSendWebViewDataRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnMessagesTranscribeAudio(f func(ctx context.Context, request *MessagesTranscribeAudioRequest) (*MessagesTranscribedAudio, error)) {
func (a *application) messagesTranscribeAudio(ctx context.Context, request *tg.MessagesTranscribeAudioRequest) (*tg.MessagesTranscribedAudio, error) {
    return &tg.MessagesTranscribedAudio{
    	TranscriptionID: 333444,
    	Text: "444333",
	}, nil
}


// OnMessagesRateTranscribedAudio(f func(ctx context.Context, request *MessagesRateTranscribedAudioRequest) (bool, error)) {
func (a *application) messagesRateTranscribedAudio(ctx context.Context, request *tg.MessagesRateTranscribedAudioRequest) (bool , error) {
    return true, nil
}


// OnMessagesGetCustomEmojiDocuments(f func(ctx context.Context, documentid []int64) ([]DocumentClass, error)) {
func (a *application) messagesGetCustomEmojiDocuments(ctx context.Context, documentid []int64) ([]tg.DocumentClass, error) {
	return []tg.DocumentClass{
		&tg.Document{
			ID: 333444,
			MimeType: "444333",
		},
	},nil
}


// OnMessagesGetEmojiStickers(f func(ctx context.Context, hash int64) (MessagesAllStickersClass, error)) {
func (a *application) messagesGetEmojiStickers(ctx context.Context, hash int64) (tg.MessagesAllStickersClass, error) {
    return &tg.MessagesAllStickers{
    	Hash: 333444,
    	Sets: []tg.StickerSet{
    		{
    			ID: 333444,
    			Title: "444333",
			},
		},
	}, nil
}


// OnMessagesGetFeaturedEmojiStickers(f func(ctx context.Context, hash int64) (MessagesFeaturedStickersClass, error)) {
func (a *application) messagesGetFeaturedEmojiStickers(ctx context.Context, hash int64) (tg.MessagesFeaturedStickersClass, error) {
    return &tg.MessagesFeaturedStickers{
    	Hash: 333444,
    	Count: 444333,
	}, nil
}


// OnMessagesReportReaction(f func(ctx context.Context, request *MessagesReportReactionRequest) (bool, error)) {
func (a *application) messagesReportReactio(ctx context.Context, request *tg.MessagesReportReactionRequest) (bool , error) {
    return true, nil
}


// OnMessagesGetTopReactions(f func(ctx context.Context, request *MessagesGetTopReactionsRequest) (MessagesReactionsClass, error)) {
func (a *application) messagesGetTopReactions(ctx context.Context, request *tg.MessagesGetTopReactionsRequest) (tg.MessagesReactionsClass, error) {
    return &tg.MessagesReactions{
    	Hash: 333444,
	}, nil
}


// OnMessagesGetRecentReactions(f func(ctx context.Context, request *MessagesGetRecentReactionsRequest) (MessagesReactionsClass, error)) {
func (a *application) messagesGetRecentReactions(ctx context.Context, request *tg.MessagesGetRecentReactionsRequest) (tg.MessagesReactionsClass, error) {
    return &tg.MessagesReactions{
    	Hash: 444333,
	}, nil
}


// OnMessagesClearRecentReactions(f func(ctx context.Context) (bool, error)) {
func (a *application) messagesClearRecentReactions(ctx context.Context) (bool, error){
    return true, nil
}


// OnMessagesGetExtendedMedia(f func(ctx context.Context, request *MessagesGetExtendedMediaRequest) (UpdatesClass, error)) {
func (a *application) messagesGetExtendedMedia(ctx context.Context, request *tg.MessagesGetExtendedMediaRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnMessagesSetDefaultHistoryTTL(f func(ctx context.Context, period int) (bool, error)) {
func (a *application) messagesSetDefaultHistoryTTL(ctx context.Context, period int) (bool , error) {
    return true, nil
}


// OnMessagesGetDefaultHistoryTTL(f func(ctx context.Context) (*DefaultHistoryTTL, error)) {
func (a *application) messagesGetDefaultHistoryTTL(ctx context.Context) (*tg.DefaultHistoryTTL, error)  {
    return &tg.DefaultHistoryTTL{
    	Period: 333444,
	}, nil
}


// OnMessagesSendBotRequestedPeer(f func(ctx context.Context, request *MessagesSendBotRequestedPeerRequest) (UpdatesClass, error)) {
func (a *application) messagesSendBotRequestedPeer(ctx context.Context, request *tg.MessagesSendBotRequestedPeerRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnMessagesGetEmojiGroups(f func(ctx context.Context, hash int) (MessagesEmojiGroupsClass, error)) {
func (a *application) messagesGetEmojiGroups(ctx context.Context, hash int) (tg.MessagesEmojiGroupsClass, error) {
    return &tg.MessagesEmojiGroups{
    	Hash: 333444,
	}, nil
}


// OnMessagesGetEmojiStatusGroups(f func(ctx context.Context, hash int) (MessagesEmojiGroupsClass, error)) {
func (a *application) messagesGetEmojiStatusGroups(ctx context.Context, hash int) (tg.MessagesEmojiGroupsClass, error) {
    return &tg.MessagesEmojiGroups{
		Hash: 333444,
	}, nil
}


// OnMessagesGetEmojiProfilePhotoGroups(f func(ctx context.Context, hash int) (MessagesEmojiGroupsClass, error)) {
func (a *application) messagesGetEmojiProfilePhotoGroups(ctx context.Context, hash int) (tg.MessagesEmojiGroupsClass, error) {
    return &tg.MessagesEmojiGroups{
		Hash: 333444,
	}, nil
}


// OnMessagesSearchCustomEmoji(f func(ctx context.Context, request *MessagesSearchCustomEmojiRequest) (EmojiListClass, error)) {
func (a *application) messagesSearchCustomEmoji(ctx context.Context, request *tg.MessagesSearchCustomEmojiRequest) (tg.EmojiListClass, error) {
    return &tg.EmojiList{
		Hash: 333444,
	}, nil
}


// OnMessagesTogglePeerTranslations(f func(ctx context.Context, request *MessagesTogglePeerTranslationsRequest) (bool, error)) {
func (a *application) messagesTogglePeerTranslations(ctx context.Context, request *tg.MessagesTogglePeerTranslationsRequest) (bool , error) {
    return true, nil
}


// OnMessagesGetBotApp(f func(ctx context.Context, request *MessagesGetBotAppRequest) (*MessagesBotApp, error)) {
func (a *application) messagesGetBotApp(ctx context.Context, request *tg.MessagesGetBotAppRequest) (*tg.MessagesBotApp, error) {
    return &tg.MessagesBotApp{
    	App: &tg.BotApp{
    		AccessHash: 333444,
    		Title: "444333",
		},
	}, nil
}


// OnMessagesRequestAppWebView(f func(ctx context.Context, request *MessagesRequestAppWebViewRequest) (*AppWebViewResultURL, error)) {
func (a *application) messagesRequestAppWebView(ctx context.Context, request *tg.MessagesRequestAppWebViewRequest) (*tg.AppWebViewResultURL, error) {
    return &tg.AppWebViewResultURL{
    	URL: "444333",
	}, nil
}


// OnMessagesSetChatWallPaper(f func(ctx context.Context, request *MessagesSetChatWallPaperRequest) (UpdatesClass, error)) {
func (a *application) messagesSetChatWallPaper(ctx context.Context, request *tg.MessagesSetChatWallPaperRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnUpdatesGetState(f func(ctx context.Context) (*UpdatesState, error)) {
func (a *application) updatesGetStatebak(ctx context.Context) (*tg.UpdatesState, error){
    return &tg.UpdatesState{
    	Pts: 333444,
    	Qts: 444333,
    	Date: 34334,
    	Seq: 43334,
	}, nil
}


// OnUpdatesGetDifference(f func(ctx context.Context, request *UpdatesGetDifferenceRequest) (UpdatesDifferenceClass, error)) {
func (a *application) updatesGetDifferencebak(ctx context.Context, request *tg.UpdatesGetDifferenceRequest) (tg.UpdatesDifferenceClass, error) {
    return &tg.UpdatesDifference{
		Users: []tg.UserClass{
			&tg.User{
				Username: "333444",
				Phone: "444333",
			},
		},
	}, nil
}


// OnUpdatesGetChannelDifference(f func(ctx context.Context, request *UpdatesGetChannelDifferenceRequest) (UpdatesChannelDifferenceClass, error)) {
func (a *application) updatesGetChannelDifference(ctx context.Context, request *tg.UpdatesGetChannelDifferenceRequest) (tg.UpdatesChannelDifferenceClass, error) {
    return &tg.UpdatesChannelDifference{
		Users: []tg.UserClass{
			&tg.User{
				Username: "333444",
				Phone: "444333",
			},
		},
	}, nil
}


// OnPhotosUpdateProfilePhoto(f func(ctx context.Context, request *PhotosUpdateProfilePhotoRequest) (*PhotosPhoto, error)) {
func (a *application) photosUpdateProfilePhoto(ctx context.Context, request *tg.PhotosUpdateProfilePhotoRequest) (*tg.PhotosPhoto, error) {
    return &tg.PhotosPhoto{

	}, nil
}


// OnPhotosUploadProfilePhoto(f func(ctx context.Context, request *PhotosUploadProfilePhotoRequest) (*PhotosPhoto, error)) {
func (a *application) photosUploadProfilePhoto(ctx context.Context, request *tg.PhotosUploadProfilePhotoRequest) (*tg.PhotosPhoto, error) {
    return &tg.PhotosPhoto{
		Users: []tg.UserClass{
			&tg.User{
				Username: "333444",
				Phone: "444333",
			},
		},
	}, nil
}


// OnPhotosDeletePhotos(f func(ctx context.Context, id []InputPhotoClass) ([]int64, error)) {
func (a *application) photosDeletePhotos(ctx context.Context, id []tg.InputPhotoClass) ([]int64, error) {
    return []int64{333444}, nil
}


// OnPhotosGetUserPhotos(f func(ctx context.Context, request *PhotosGetUserPhotosRequest) (PhotosPhotosClass, error)) {
func (a *application) photosGetUserPhotos(ctx context.Context, request *tg.PhotosGetUserPhotosRequest) (tg.PhotosPhotosClass, error) {
    return &tg.PhotosPhotos{
		Users: []tg.UserClass{
			&tg.User{
				Username: "333444",
				Phone: "444333",
			},
		},
	}, nil
}


// OnPhotosUploadContactProfilePhoto(f func(ctx context.Context, request *PhotosUploadContactProfilePhotoRequest) (*PhotosPhoto, error)) {
func (a *application) photosUploadContactProfilePhoto(ctx context.Context, request *tg.PhotosUploadContactProfilePhotoRequest) (*tg.PhotosPhoto, error) {
    return &tg.PhotosPhoto{
		Users: []tg.UserClass{
			&tg.User{
				Username: "333444",
				Phone: "444333",
			},
		},
	}, nil
}


// OnUploadSaveFilePart(f func(ctx context.Context, request *UploadSaveFilePartRequest) (bool, error)) {
func (a *application) uploadSaveFilePart(ctx context.Context, request *tg.UploadSaveFilePartRequest) (bool , error) {
    return true, nil
}


// OnUploadGetFile(f func(ctx context.Context, request *UploadGetFileRequest) (UploadFileClass, error)) {
func (a *application) uploadGetFile(ctx context.Context, request *tg.UploadGetFileRequest) (tg.UploadFileClass, error) {
    return &tg.UploadFile{
    	Mtime: 333444,
	}, nil
}


// OnUploadSaveBigFilePart(f func(ctx context.Context, request *UploadSaveBigFilePartRequest) (bool, error)) {
func (a *application) uploadSaveBigFilePart(ctx context.Context, request *tg.UploadSaveBigFilePartRequest) (bool , error) {
    return true, nil
}


// OnUploadGetWebFile(f func(ctx context.Context, request *UploadGetWebFileRequest) (*UploadWebFile, error)) {
func (a *application) uploadGetWebFile(ctx context.Context, request *tg.UploadGetWebFileRequest) (*tg.UploadWebFile, error) {
    return &tg.UploadWebFile{
		Mtime: 333444,
	}, nil
}


// OnUploadGetCDNFile(f func(ctx context.Context, request *UploadGetCDNFileRequest) (UploadCDNFileClass, error)) {
func (a *application) uploadGetCDNFile(ctx context.Context, request *tg.UploadGetCDNFileRequest) (tg.UploadCDNFileClass, error) {
    return &tg.UploadCDNFile{
	}, nil
}


// OnUploadReuploadCDNFile(f func(ctx context.Context, request *UploadReuploadCDNFileRequest) ([]FileHash, error)) {
func (a *application) uploadReuploadCDNFile(ctx context.Context, request *tg.UploadReuploadCDNFileRequest) ([]tg.FileHash, error) {
    return []tg.FileHash{
    	{
    		Offset: 333444,
    		Limit: 444333,
		},
	}, nil
}


// OnUploadGetCDNFileHashes(f func(ctx context.Context, request *UploadGetCDNFileHashesRequest) ([]FileHash, error)) {
func (a *application) uploadGetCDNFileHashes(ctx context.Context, request *tg.UploadGetCDNFileHashesRequest) ([]tg.FileHash, error) {
    return []tg.FileHash{
		{
			Offset: 333444,
			Limit: 444333,
		},
	}, nil
}


// OnUploadGetFileHashes(f func(ctx context.Context, request *UploadGetFileHashesRequest) ([]FileHash, error)) {
func (a *application) uploadGetFileHashes(ctx context.Context, request *tg.UploadGetFileHashesRequest) ([]tg.FileHash, error) {
    return []tg.FileHash{
		{
			Offset: 333444,
			Limit: 444333,
		},
	}, nil
}


// OnHelpGetConfig(f func(ctx context.Context) (*Config, error)) {
func (a *application) helpGetConfigbak(ctx context.Context) (*tg.Config, error) {
    return &tg.Config{
    	Date: 343434,
    	DCTxtDomainName: "444333",
    	OnlineCloudTimeoutMs: 333444,
	}, nil
}


// OnHelpGetNearestDC(f func(ctx context.Context) (*NearestDC, error)) {
func (a *application) helpGetNearestDCbak(ctx context.Context) (*tg.NearestDC, error)  {
    return &tg.NearestDC{
    	Country: "4444333",
    	ThisDC: 2,
    	NearestDC: 1,
	}, nil
}


// OnHelpGetAppUpdate(f func(ctx context.Context, source string) (HelpAppUpdateClass, error)) {
func (a *application) helpGetAppUpdate(ctx context.Context, source string) (tg.HelpAppUpdateClass, error) {
    return &tg.HelpAppUpdate{
    	ID: 333444,
    	Version: "444333",
    	Text: "34334",
	}, nil
}


// OnHelpGetInviteText(f func(ctx context.Context) (*HelpInviteText, error)) {
func (a *application) helpGetInviteText(ctx context.Context) (*tg.HelpInviteText, error)  {
    return &tg.HelpInviteText{
    	Message: "33334444",
	}, nil
}


// OnHelpGetSupport(f func(ctx context.Context) (*HelpSupport, error)) {
func (a *application) helpGetSupport(ctx context.Context) (*tg.HelpSupport,error) {
    return &tg.HelpSupport{
    	PhoneNumber: "343332222",
	}, nil
}


// OnHelpGetAppChangelog(f func(ctx context.Context, prevappversion string) (UpdatesClass, error)) {
func (a *application) helpGetAppChangelog(ctx context.Context, prevappversion string) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnHelpSetBotUpdatesStatus(f func(ctx context.Context, request *HelpSetBotUpdatesStatusRequest) (bool, error)) {
func (a *application) helpSetBotUpdatesStatus(ctx context.Context, request *tg.HelpSetBotUpdatesStatusRequest) (bool , error) {
    return true, nil
}


// OnHelpGetCDNConfig(f func(ctx context.Context) (*CDNConfig, error)) {
func (a *application) helpGetCDNConfig(ctx context.Context) (*tg.CDNConfig,error) {
    return &tg.CDNConfig{
    	PublicKeys: []tg.CDNPublicKey{
    		{
    			DCID: 33344,
    			PublicKey: "444333",
			},
		},
	}, nil
}


// OnHelpGetRecentMeURLs(f func(ctx context.Context, referer string) (*HelpRecentMeURLs, error)) {
func (a *application) helpGetRecentMeURLs(ctx context.Context, referer string) (*tg.HelpRecentMeURLs, error) {
    return &tg.HelpRecentMeURLs{
		Users: []tg.UserClass{
			&tg.User{
				Username: "333444",
				Phone: "444333",
			},
		},
	}, nil
}


// OnHelpGetTermsOfServiceUpdate(f func(ctx context.Context) (HelpTermsOfServiceUpdateClass, error)) {
func (a *application) helpGetTermsOfServiceUpdatebak(ctx context.Context) (tg.HelpTermsOfServiceUpdateClass, error) {
    return &tg.HelpTermsOfServiceUpdate{
    	Expires: 333444,
	},nil
}


// OnHelpAcceptTermsOfService(f func(ctx context.Context, id DataJSON) (bool, error)) {
func (a *application) helpAcceptTermsOfServicebak(ctx context.Context, id *tg.DataJSON) (bool , error) {
    return true, nil
}


// OnHelpGetDeepLinkInfo(f func(ctx context.Context, path string) (HelpDeepLinkInfoClass, error)) {
func (a *application) helpGetDeepLinkInfo(ctx context.Context, path string) (tg.HelpDeepLinkInfoClass, error) {
    return &tg.HelpDeepLinkInfo{
    	UpdateApp: true,
    	Message: "444333",
	}, nil
}


// OnHelpGetAppConfig(f func(ctx context.Context, hash int) (HelpAppConfigClass, error)) {
func (a *application) helpGetAppConfigbak(ctx context.Context, hash int) (tg.HelpAppConfigClass, error) {
    return &tg.HelpAppConfig{
    	Hash: 333444,
	}, nil
}


// OnHelpSaveAppLog(f func(ctx context.Context, events []InputAppEvent) (bool, error)) {
func (a *application) helpSaveAppLog(ctx context.Context, events []tg.InputAppEvent) (bool , error) {
    return true, nil
}


// OnHelpGetPassportConfig(f func(ctx context.Context, hash int) (HelpPassportConfigClass, error)) {
func (a *application) helpGetPassportConfig(ctx context.Context, hash int) (tg.HelpPassportConfigClass, error) {
    return &tg.HelpPassportConfig{
		Hash: 333444,
	}, nil
}


// OnHelpGetSupportName(f func(ctx context.Context) (*HelpSupportName, error)) {
func (a *application) helpGetSupportName(ctx context.Context) (*tg.HelpSupportName,error) {
    return &tg.HelpSupportName{
    	Name: "theLostLamb",
	}, nil
}


// OnHelpGetUserInfo(f func(ctx context.Context, userid InputUserClass) (HelpUserInfoClass, error)) {
func (a *application) helpGetUserInfo(ctx context.Context, userid tg.InputUserClass) (tg.HelpUserInfoClass, error) {
    return &tg.HelpUserInfo{
    	Message: "333444",
    	Author: "444333",
    	Date: 34334,
	}, nil
}


// OnHelpEditUserInfo(f func(ctx context.Context, request *HelpEditUserInfoRequest) (HelpUserInfoClass, error)) {
func (a *application) helpEditUserInfo(ctx context.Context, request *tg.HelpEditUserInfoRequest) (tg.HelpUserInfoClass, error) {
    return &tg.HelpUserInfo{
		Message: "333444",
		Author: "444333",
		Date: 34334,
	}, nil
}


// OnHelpGetPromoData(f func(ctx context.Context) (HelpPromoDataClass, error)) {
func (a *application) helpGetPromoDatabak(ctx context.Context) (tg.HelpPromoDataClass,error)  {
    return &tg.HelpPromoData{
		Users: []tg.UserClass{
			&tg.User{
				Username: "333444",
				Phone: "444333",
			},
		},
	}, nil
}


// OnHelpHidePromoData(f func(ctx context.Context, peer InputPeerClass) (bool, error)) {
func (a *application) helpHidePromoData(ctx context.Context, peer tg.InputPeerClass) (bool , error) {
    return true, nil
}


// OnHelpDismissSuggestion(f func(ctx context.Context, request *HelpDismissSuggestionRequest) (bool, error)) {
func (a *application) helpDismissSuggestio(ctx context.Context, request *tg.HelpDismissSuggestionRequest) (bool , error) {
    return true, nil
}


// OnHelpGetCountriesList(f func(ctx context.Context, request *HelpGetCountriesListRequest) (HelpCountriesListClass, error)) {
func (a *application) helpGetCountriesListbak(ctx context.Context, request *tg.HelpGetCountriesListRequest) (tg.HelpCountriesListClass, error) {
    return &tg.HelpCountriesList{
    	Hash: 333444,
	}, nil
}


// OnHelpGetPremiumPromo(f func(ctx context.Context) (*HelpPremiumPromo, error)) {
func (a *application) helpGetPremiumPromobak(ctx context.Context) (*tg.HelpPremiumPromo, error)  {
    return &tg.HelpPremiumPromo{
		Users: []tg.UserClass{
			&tg.User{
				Username: "333444",
				Phone: "444333",
			},
		},
	}, nil
}


// OnChannelsReadHistory(f func(ctx context.Context, request *ChannelsReadHistoryRequest) (bool, error)) {
func (a *application) channelsReadHistory(ctx context.Context, request *tg.ChannelsReadHistoryRequest) (bool , error) {
    return true, nil
}


// OnChannelsDeleteMessages(f func(ctx context.Context, request *ChannelsDeleteMessagesRequest) (*MessagesAffectedMessages, error)) {
func (a *application) channelsDeleteMessages(ctx context.Context, request *tg.ChannelsDeleteMessagesRequest) (*tg.MessagesAffectedMessages, error) {
    return &tg.MessagesAffectedMessages{
    	Pts: 333444,
    	PtsCount: 444333,
	}, nil
}


// OnChannelsReportSpam(f func(ctx context.Context, request *ChannelsReportSpamRequest) (bool, error)) {
func (a *application) channelsReportSpam(ctx context.Context, request *tg.ChannelsReportSpamRequest) (bool , error) {
    return true, nil
}


// OnChannelsGetMessages(f func(ctx context.Context, request *ChannelsGetMessagesRequest) (MessagesMessagesClass, error)) {
func (a *application) channelsGetMessagesbak(ctx context.Context, request *tg.ChannelsGetMessagesRequest) (tg.MessagesMessagesClass, error) {
	return &tg.MessagesMessages{
		Users: []tg.UserClass{
			&tg.User{
				Username: "333444",
				Phone: "444333",
			},
		},
	}, nil
}


// OnChannelsGetParticipants(f func(ctx context.Context, request *ChannelsGetParticipantsRequest) (ChannelsChannelParticipantsClass, error)) {
func (a *application) channelsGetParticipants(ctx context.Context, request *tg.ChannelsGetParticipantsRequest) (tg.ChannelsChannelParticipantsClass, error) {
    return &tg.ChannelsChannelParticipants{
		Users: []tg.UserClass{
			&tg.User{
				Username: "333444",
				Phone: "444333",
			},
		},
	}, nil
}


// OnChannelsGetParticipant(f func(ctx context.Context, request *ChannelsGetParticipantRequest) (*ChannelsChannelParticipant, error)) {
func (a *application) channelsGetParticipant(ctx context.Context, request *tg.ChannelsGetParticipantRequest) (*tg.ChannelsChannelParticipant, error) {
    return &tg.ChannelsChannelParticipant{
		Users: []tg.UserClass{
			&tg.User{
				Username: "333444",
				Phone: "444333",
			},
		},
	}, nil
}


// OnChannelsGetChannels(f func(ctx context.Context, id []InputChannelClass) (MessagesChatsClass, error)) {
func (a *application) channelsGetChannels(ctx context.Context, id []tg.InputChannelClass) (tg.MessagesChatsClass, error) {
    return &tg.MessagesChats{
		Chats: []tg.ChatClass{
			&tg.Chat{
				ID: 333444,
				Title: "333444",
			},
		},
	}, nil
}


// OnChannelsGetFullChannel(f func(ctx context.Context, channel InputChannelClass) (*MessagesChatFull, error)) {
func (a *application) channelsGetFullChannel(ctx context.Context, channel tg.InputChannelClass) (*tg.MessagesChatFull, error) {
    return &tg.MessagesChatFull{
		Chats: []tg.ChatClass{
			&tg.Chat{
				ID: 333444,
				Title: "333444",
			},
		},
	}, nil
}


// OnChannelsCreateChannel(f func(ctx context.Context, request *ChannelsCreateChannelRequest) (UpdatesClass, error)) {
func (a *application) channelsCreateChannel(ctx context.Context, request *tg.ChannelsCreateChannelRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnChannelsEditAdmin(f func(ctx context.Context, request *ChannelsEditAdminRequest) (UpdatesClass, error)) {
func (a *application) channelsEditAdmi(ctx context.Context, request *tg.ChannelsEditAdminRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnChannelsEditTitle(f func(ctx context.Context, request *ChannelsEditTitleRequest) (UpdatesClass, error)) {
func (a *application) channelsEditTitle(ctx context.Context, request *tg.ChannelsEditTitleRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnChannelsEditPhoto(f func(ctx context.Context, request *ChannelsEditPhotoRequest) (UpdatesClass, error)) {
func (a *application) channelsEditPhoto(ctx context.Context, request *tg.ChannelsEditPhotoRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnChannelsCheckUsername(f func(ctx context.Context, request *ChannelsCheckUsernameRequest) (bool, error)) {
func (a *application) channelsCheckUsername(ctx context.Context, request *tg.ChannelsCheckUsernameRequest) (bool , error) {
    return true, nil
}


// OnChannelsUpdateUsername(f func(ctx context.Context, request *ChannelsUpdateUsernameRequest) (bool, error)) {
func (a *application) channelsUpdateUsername(ctx context.Context, request *tg.ChannelsUpdateUsernameRequest) (bool , error) {
    return true, nil
}


// OnChannelsJoinChannel(f func(ctx context.Context, channel InputChannelClass) (UpdatesClass, error)) {
func (a *application) channelsJoinChannel(ctx context.Context, channel tg.InputChannelClass) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnChannelsLeaveChannel(f func(ctx context.Context, channel InputChannelClass) (UpdatesClass, error)) {
func (a *application) channelsLeaveChannel(ctx context.Context, channel tg.InputChannelClass) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnChannelsInviteToChannel(f func(ctx context.Context, request *ChannelsInviteToChannelRequest) (UpdatesClass, error)) {
func (a *application) channelsInviteToChannel(ctx context.Context, request *tg.ChannelsInviteToChannelRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnChannelsDeleteChannel(f func(ctx context.Context, channel InputChannelClass) (UpdatesClass, error)) {
func (a *application) channelsDeleteChannel(ctx context.Context, channel tg.InputChannelClass) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnChannelsExportMessageLink(f func(ctx context.Context, request *ChannelsExportMessageLinkRequest) (*ExportedMessageLink, error)) {
func (a *application) channelsExportMessageLink(ctx context.Context, request *tg.ChannelsExportMessageLinkRequest) (*tg.ExportedMessageLink, error) {
    return &tg.ExportedMessageLink{
    	Link: "333444",
    	HTML: "444333",
	}, nil
}


// OnChannelsToggleSignatures(f func(ctx context.Context, request *ChannelsToggleSignaturesRequest) (UpdatesClass, error)) {
func (a *application) channelsToggleSignatures(ctx context.Context, request *tg.ChannelsToggleSignaturesRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnChannelsGetAdminedPublicChannels(f func(ctx context.Context, request *ChannelsGetAdminedPublicChannelsRequest) (MessagesChatsClass, error)) {
func (a *application) channelsGetAdminedPublicChannels(ctx context.Context, request *tg.ChannelsGetAdminedPublicChannelsRequest) (tg.MessagesChatsClass, error) {
    return &tg.MessagesChats{
		Chats: []tg.ChatClass{
			&tg.Chat{
				ID: 333444,
				Title: "333444",
			},
		},
	}, nil
}


// OnChannelsEditBanned(f func(ctx context.Context, request *ChannelsEditBannedRequest) (UpdatesClass, error)) {
func (a *application) channelsEditBanned(ctx context.Context, request *tg.ChannelsEditBannedRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnChannelsGetAdminLog(f func(ctx context.Context, request *ChannelsGetAdminLogRequest) (*ChannelsAdminLogResults, error)) {
func (a *application) channelsGetAdminLog(ctx context.Context, request *tg.ChannelsGetAdminLogRequest) (*tg.ChannelsAdminLogResults, error) {
    return &tg.ChannelsAdminLogResults{
		Chats: []tg.ChatClass{
			&tg.Chat{
				ID: 333444,
				Title: "333444",
			},
		},
	}, nil
}


// OnChannelsSetStickers(f func(ctx context.Context, request *ChannelsSetStickersRequest) (bool, error)) {
func (a *application) channelsSetStickers(ctx context.Context, request *tg.ChannelsSetStickersRequest) (bool , error) {
    return true, nil
}


// OnChannelsReadMessageContents(f func(ctx context.Context, request *ChannelsReadMessageContentsRequest) (bool, error)) {
func (a *application) channelsReadMessageContents(ctx context.Context, request *tg.ChannelsReadMessageContentsRequest) (bool , error) {
    return true, nil
}


// OnChannelsDeleteHistory(f func(ctx context.Context, request *ChannelsDeleteHistoryRequest) (UpdatesClass, error)) {
func (a *application) channelsDeleteHistory(ctx context.Context, request *tg.ChannelsDeleteHistoryRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnChannelsTogglePreHistoryHidden(f func(ctx context.Context, request *ChannelsTogglePreHistoryHiddenRequest) (UpdatesClass, error)) {
func (a *application) channelsTogglePreHistoryHidde(ctx context.Context, request *tg.ChannelsTogglePreHistoryHiddenRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnChannelsGetLeftChannels(f func(ctx context.Context, offset int) (MessagesChatsClass, error)) {
func (a *application) channelsGetLeftChannels(ctx context.Context, offset int) (tg.MessagesChatsClass, error) {
    return &tg.MessagesChats{
		Chats: []tg.ChatClass{
			&tg.Chat{
				ID: 333444,
				Title: "333444",
			},
		},
	}, nil
}


// OnChannelsGetGroupsForDiscussion(f func(ctx context.Context) (MessagesChatsClass, error)) {
func (a *application) channelsGetGroupsForDiscussio(ctx context.Context) (tg.MessagesChatsClass,error){
    return &tg.MessagesChats{
		Chats: []tg.ChatClass{
			&tg.Chat{
				ID: 333444,
				Title: "333444",
			},
		},
	}, nil
}


// OnChannelsSetDiscussionGroup(f func(ctx context.Context, request *ChannelsSetDiscussionGroupRequest) (bool, error)) {
func (a *application) channelsSetDiscussionGroup(ctx context.Context, request *tg.ChannelsSetDiscussionGroupRequest) (bool , error) {
    return true, nil
}


// OnChannelsEditCreator(f func(ctx context.Context, request *ChannelsEditCreatorRequest) (UpdatesClass, error)) {
func (a *application) channelsEditCreator(ctx context.Context, request *tg.ChannelsEditCreatorRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnChannelsEditLocation(f func(ctx context.Context, request *ChannelsEditLocationRequest) (bool, error)) {
func (a *application) channelsEditLocatio(ctx context.Context, request *tg.ChannelsEditLocationRequest) (bool , error) {
    return true, nil
}


// OnChannelsToggleSlowMode(f func(ctx context.Context, request *ChannelsToggleSlowModeRequest) (UpdatesClass, error)) {
func (a *application) channelsToggleSlowMode(ctx context.Context, request *tg.ChannelsToggleSlowModeRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnChannelsGetInactiveChannels(f func(ctx context.Context) (*MessagesInactiveChats, error)) {
func (a *application) channelsGetInactiveChannels(ctx context.Context) (*tg.MessagesInactiveChats, error)  {
    return &tg.MessagesInactiveChats{
		Chats: []tg.ChatClass{
			&tg.Chat{
				ID: 333444,
				Title: "333444",
			},
		},
	}, nil
}


// OnChannelsConvertToGigagroup(f func(ctx context.Context, channel InputChannelClass) (UpdatesClass, error)) {
func (a *application) channelsConvertToGigagroup(ctx context.Context, channel tg.InputChannelClass) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnChannelsViewSponsoredMessage(f func(ctx context.Context, request *ChannelsViewSponsoredMessageRequest) (bool, error)) {
func (a *application) channelsViewSponsoredMessage(ctx context.Context, request *tg.ChannelsViewSponsoredMessageRequest) (bool , error) {
    return true, nil
}


// OnChannelsGetSponsoredMessages(f func(ctx context.Context, channel InputChannelClass) (MessagesSponsoredMessagesClass, error)) {
func (a *application) channelsGetSponsoredMessages(ctx context.Context, channel tg.InputChannelClass) (tg.MessagesSponsoredMessagesClass, error) {
    return &tg.MessagesSponsoredMessages{
		Chats: []tg.ChatClass{
			&tg.Chat{
				ID: 333444,
				Title: "333444",
			},
		},
	}, nil
}


// OnChannelsGetSendAs(f func(ctx context.Context, peer InputPeerClass) (*ChannelsSendAsPeers, error)) {
func (a *application) channelsGetSendAs(ctx context.Context, peer tg.InputPeerClass) (*tg.ChannelsSendAsPeers, error) {
    return &tg.ChannelsSendAsPeers{
		Chats: []tg.ChatClass{
			&tg.Chat{
				ID: 333444,
				Title: "333444",
			},
		},
	}, nil
}


// OnChannelsDeleteParticipantHistory(f func(ctx context.Context, request *ChannelsDeleteParticipantHistoryRequest) (*MessagesAffectedHistory, error)) {
func (a *application) channelsDeleteParticipantHistory(ctx context.Context, request *tg.ChannelsDeleteParticipantHistoryRequest) (*tg.MessagesAffectedHistory, error) {
    return &tg.MessagesAffectedHistory{
    	PtsCount: 3334444,
    	Pts: 444333,
    	Offset: 343434,
	}, nil
}


// OnChannelsToggleJoinToSend(f func(ctx context.Context, request *ChannelsToggleJoinToSendRequest) (UpdatesClass, error)) {
func (a *application) channelsToggleJoinToSend(ctx context.Context, request *tg.ChannelsToggleJoinToSendRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnChannelsToggleJoinRequest(f func(ctx context.Context, request *ChannelsToggleJoinRequestRequest) (UpdatesClass, error)) {
func (a *application) channelsToggleJoinRequest(ctx context.Context, request *tg.ChannelsToggleJoinRequestRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnChannelsReorderUsernames(f func(ctx context.Context, request *ChannelsReorderUsernamesRequest) (bool, error)) {
func (a *application) channelsReorderUsernames(ctx context.Context, request *tg.ChannelsReorderUsernamesRequest) (bool , error) {
    return true, nil
}


// OnChannelsToggleUsername(f func(ctx context.Context, request *ChannelsToggleUsernameRequest) (bool, error)) {
func (a *application) channelsToggleUsername(ctx context.Context, request *tg.ChannelsToggleUsernameRequest) (bool , error) {
    return true, nil
}


// OnChannelsDeactivateAllUsernames(f func(ctx context.Context, channel InputChannelClass) (bool, error)) {
func (a *application) channelsDeactivateAllUsernames(ctx context.Context, channel tg.InputChannelClass) (bool , error) {
    return true, nil
}


// OnChannelsToggleForum(f func(ctx context.Context, request *ChannelsToggleForumRequest) (UpdatesClass, error)) {
func (a *application) channelsToggleForum(ctx context.Context, request *tg.ChannelsToggleForumRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnChannelsCreateForumTopic(f func(ctx context.Context, request *ChannelsCreateForumTopicRequest) (UpdatesClass, error)) {
func (a *application) channelsCreateForumTopic(ctx context.Context, request *tg.ChannelsCreateForumTopicRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnChannelsGetForumTopics(f func(ctx context.Context, request *ChannelsGetForumTopicsRequest) (*MessagesForumTopics, error)) {
func (a *application) channelsGetForumTopics(ctx context.Context, request *tg.ChannelsGetForumTopicsRequest) (*tg.MessagesForumTopics, error) {
    return &tg.MessagesForumTopics{
		Chats: []tg.ChatClass{
			&tg.Chat{
				ID: 333444,
				Title: "333444",
			},
		},
	}, nil
}


// OnChannelsGetForumTopicsByID(f func(ctx context.Context, request *ChannelsGetForumTopicsByIDRequest) (*MessagesForumTopics, error)) {
func (a *application) channelsGetForumTopicsByID(ctx context.Context, request *tg.ChannelsGetForumTopicsByIDRequest) (*tg.MessagesForumTopics, error) {
    return &tg.MessagesForumTopics{
		Chats: []tg.ChatClass{
			&tg.Chat{
				ID: 333444,
				Title: "333444",
			},
		},
	}, nil
}


// OnChannelsEditForumTopic(f func(ctx context.Context, request *ChannelsEditForumTopicRequest) (UpdatesClass, error)) {
func (a *application) channelsEditForumTopic(ctx context.Context, request *tg.ChannelsEditForumTopicRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnChannelsUpdatePinnedForumTopic(f func(ctx context.Context, request *ChannelsUpdatePinnedForumTopicRequest) (UpdatesClass, error)) {
func (a *application) channelsUpdatePinnedForumTopic(ctx context.Context, request *tg.ChannelsUpdatePinnedForumTopicRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnChannelsDeleteTopicHistory(f func(ctx context.Context, request *ChannelsDeleteTopicHistoryRequest) (*MessagesAffectedHistory, error)) {
func (a *application) channelsDeleteTopicHistory(ctx context.Context, request *tg.ChannelsDeleteTopicHistoryRequest) (*tg.MessagesAffectedHistory, error) {
    return &tg.MessagesAffectedHistory{
    	Pts: 333444,
    	PtsCount: 444333,
    	Offset: 34334,
	}, nil
}


// OnChannelsReorderPinnedForumTopics(f func(ctx context.Context, request *ChannelsReorderPinnedForumTopicsRequest) (UpdatesClass, error)) {
func (a *application) channelsReorderPinnedForumTopics(ctx context.Context, request *tg.ChannelsReorderPinnedForumTopicsRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnChannelsToggleAntiSpam(f func(ctx context.Context, request *ChannelsToggleAntiSpamRequest) (UpdatesClass, error)) {
func (a *application) channelsToggleAntiSpam(ctx context.Context, request *tg.ChannelsToggleAntiSpamRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnChannelsReportAntiSpamFalsePositive(f func(ctx context.Context, request *ChannelsReportAntiSpamFalsePositiveRequest) (bool, error)) {
func (a *application) channelsReportAntiSpamFalsePositive(ctx context.Context, request *tg.ChannelsReportAntiSpamFalsePositiveRequest) (bool , error) {
    return true, nil
}


// OnChannelsToggleParticipantsHidden(f func(ctx context.Context, request *ChannelsToggleParticipantsHiddenRequest) (UpdatesClass, error)) {
func (a *application) channelsToggleParticipantsHidde(ctx context.Context, request *tg.ChannelsToggleParticipantsHiddenRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnChannelsClickSponsoredMessage(f func(ctx context.Context, request *ChannelsClickSponsoredMessageRequest) (bool, error)) {
func (a *application) channelsClickSponsoredMessage(ctx context.Context, request *tg.ChannelsClickSponsoredMessageRequest) (bool , error) {
    return true, nil
}


// OnBotsSendCustomRequest(f func(ctx context.Context, request *BotsSendCustomRequestRequest) (*DataJSON, error)) {
func (a *application) botsSendCustomRequest(ctx context.Context, request *tg.BotsSendCustomRequestRequest) (*tg.DataJSON, error) {
	return &tg.DataJSON{
		Data: "theLostLamb333",
	}, nil
}


// OnBotsAnswerWebhookJSONQuery(f func(ctx context.Context, request *BotsAnswerWebhookJSONQueryRequest) (bool, error)) {
func (a *application) botsAnswerWebhookJSONQuery(ctx context.Context, request *tg.BotsAnswerWebhookJSONQueryRequest) (bool , error) {
    return true, nil
}


// OnBotsSetBotCommands(f func(ctx context.Context, request *BotsSetBotCommandsRequest) (bool, error)) {
func (a *application) botsSetBotCommands(ctx context.Context, request *tg.BotsSetBotCommandsRequest) (bool , error) {
    return true, nil
}


// OnBotsResetBotCommands(f func(ctx context.Context, request *BotsResetBotCommandsRequest) (bool, error)) {
func (a *application) botsResetBotCommands(ctx context.Context, request *tg.BotsResetBotCommandsRequest) (bool , error) {
    return true, nil
}


// OnBotsGetBotCommands(f func(ctx context.Context, request *BotsGetBotCommandsRequest) ([]BotCommand, error)) {
func (a *application) botsGetBotCommands(ctx context.Context, request *tg.BotsGetBotCommandsRequest) ([]tg.BotCommand, error) {
    return []tg.BotCommand{
	}, nil
}


// OnBotsSetBotMenuButton(f func(ctx context.Context, request *BotsSetBotMenuButtonRequest) (bool, error)) {
func (a *application) botsSetBotMenuButto(ctx context.Context, request *tg.BotsSetBotMenuButtonRequest) (bool , error) {
    return true, nil
}


// OnBotsGetBotMenuButton(f func(ctx context.Context, userid InputUserClass) (BotMenuButtonClass, error)) {
func (a *application) botsGetBotMenuButto(ctx context.Context, userid tg.InputUserClass) (tg.BotMenuButtonClass, error) {
    return &tg.BotMenuButton{
    	Text: "333444",
    	URL: "444333",
	}, nil
}


// OnBotsSetBotBroadcastDefaultAdminRights(f func(ctx context.Context, adminrights ChatAdminRights) (bool, error)) {
func (a *application) botsSetBotBroadcastDefaultAdminRights(ctx context.Context, adminrights tg.ChatAdminRights) (bool , error) {
    return true, nil
}


// OnBotsSetBotGroupDefaultAdminRights(f func(ctx context.Context, adminrights ChatAdminRights) (bool, error)) {
func (a *application) botsSetBotGroupDefaultAdminRights(ctx context.Context, adminrights tg.ChatAdminRights) (bool , error) {
    return true, nil
}


// OnBotsSetBotInfo(f func(ctx context.Context, request *BotsSetBotInfoRequest) (bool, error)) {
func (a *application) botsSetBotInfo(ctx context.Context, request *tg.BotsSetBotInfoRequest) (bool , error) {
    return true, nil
}


// OnBotsGetBotInfo(f func(ctx context.Context, request *BotsGetBotInfoRequest) (*BotsBotInfo, error)) {
func (a *application) botsGetBotInfo(ctx context.Context, request *tg.BotsGetBotInfoRequest) (*tg.BotsBotInfo, error) {
    return &tg.BotsBotInfo{
    	Name: "theLostLamb",
    	About: "333444",
    	Description: "444333",
	}, nil
}


// OnBotsReorderUsernames(f func(ctx context.Context, request *BotsReorderUsernamesRequest) (bool, error)) {
func (a *application) botsReorderUsernames(ctx context.Context, request *tg.BotsReorderUsernamesRequest) (bool , error) {
    return true, nil
}


// OnBotsToggleUsername(f func(ctx context.Context, request *BotsToggleUsernameRequest) (bool, error)) {
func (a *application) botsToggleUsername(ctx context.Context, request *tg.BotsToggleUsernameRequest) (bool , error) {
    return true, nil
}


// OnBotsCanSendMessage(f func(ctx context.Context, bot InputUserClass) (bool, error)) {
func (a *application) botsCanSendMessage(ctx context.Context, bot tg.InputUserClass) (bool , error) {
    return true, nil
}


// OnBotsAllowSendMessage(f func(ctx context.Context, bot InputUserClass) (UpdatesClass, error)) {
func (a *application) botsAllowSendMessage(ctx context.Context, bot tg.InputUserClass) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnBotsInvokeWebViewCustomMethod(f func(ctx context.Context, request *BotsInvokeWebViewCustomMethodRequest) (*DataJSON, error)) {
func (a *application) botsInvokeWebViewCustomMethod(ctx context.Context, request *tg.BotsInvokeWebViewCustomMethodRequest) (*tg.DataJSON, error) {
	return &tg.DataJSON{
		Data: "theLostLamb333",
	}, nil
}


// OnPaymentsGetPaymentForm(f func(ctx context.Context, request *PaymentsGetPaymentFormRequest) (*PaymentsPaymentForm, error)) {
func (a *application) paymentsGetPaymentForm(ctx context.Context, request *tg.PaymentsGetPaymentFormRequest) (*tg.PaymentsPaymentForm, error) {
    return &tg.PaymentsPaymentForm{
		Users: []tg.UserClass{
			&tg.User{
				Username: "333444",
				Phone: "444333",
			},
		},
	}, nil
}


// OnPaymentsGetPaymentReceipt(f func(ctx context.Context, request *PaymentsGetPaymentReceiptRequest) (*PaymentsPaymentReceipt, error)) {
func (a *application) paymentsGetPaymentReceipt(ctx context.Context, request *tg.PaymentsGetPaymentReceiptRequest) (*tg.PaymentsPaymentReceipt, error) {
    return &tg.PaymentsPaymentReceipt{
		Users: []tg.UserClass{
			&tg.User{
				Username: "333444",
				Phone: "444333",
			},
		},
	}, nil
}


// OnPaymentsValidateRequestedInfo(f func(ctx context.Context, request *PaymentsValidateRequestedInfoRequest) (*PaymentsValidatedRequestedInfo, error)) {
func (a *application) paymentsValidateRequestedInfo(ctx context.Context, request *tg.PaymentsValidateRequestedInfoRequest) (*tg.PaymentsValidatedRequestedInfo, error) {
    return &tg.PaymentsValidatedRequestedInfo{
    	ID: "3333444",
	}, nil
}


// OnPaymentsSendPaymentForm(f func(ctx context.Context, request *PaymentsSendPaymentFormRequest) (PaymentsPaymentResultClass, error)) {
func (a *application) paymentsSendPaymentForm(ctx context.Context, request *tg.PaymentsSendPaymentFormRequest) (tg.PaymentsPaymentResultClass, error) {
    return &tg.PaymentsPaymentResult{
    	Updates: &tg.Updates{
    		Date: 333444,
    		Seq: 444333,
		},
	}, nil
}


// OnPaymentsGetSavedInfo(f func(ctx context.Context) (*PaymentsSavedInfo, error)) {
func (a *application) paymentsGetSavedInfo(ctx context.Context) (*tg.PaymentsSavedInfo,error) {
    return &tg.PaymentsSavedInfo{
    	SavedInfo: struct {
			Flags           bin.Fields
			Name            string
			Phone           string
			Email           string
			ShippingAddress tg.PostAddress
		}{ Name: "theLostLamb", Phone: "333444", Email: "444333", },
	}, nil
}


// OnPaymentsClearSavedInfo(f func(ctx context.Context, request *PaymentsClearSavedInfoRequest) (bool, error)) {
func (a *application) paymentsClearSavedInfo(ctx context.Context, request *tg.PaymentsClearSavedInfoRequest) (bool , error) {
    return true, nil
}


// OnPaymentsGetBankCardData(f func(ctx context.Context, number string) (*PaymentsBankCardData, error)) {
func (a *application) paymentsGetBankCardData(ctx context.Context, number string) (*tg.PaymentsBankCardData, error) {
    return &tg.PaymentsBankCardData{
    	Title: "333444",
	}, nil
}


// OnPaymentsExportInvoice(f func(ctx context.Context, invoicemedia InputMediaClass) (*PaymentsExportedInvoice, error)) {
func (a *application) paymentsExportInvoice(ctx context.Context, invoicemedia tg.InputMediaClass) (*tg.PaymentsExportedInvoice, error) {
    return &tg.PaymentsExportedInvoice{
    	URL: "333444",
	}, nil
}


// OnPaymentsAssignAppStoreTransaction(f func(ctx context.Context, request *PaymentsAssignAppStoreTransactionRequest) (UpdatesClass, error)) {
func (a *application) paymentsAssignAppStoreTransactio(ctx context.Context, request *tg.PaymentsAssignAppStoreTransactionRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnPaymentsAssignPlayMarketTransaction(f func(ctx context.Context, request *PaymentsAssignPlayMarketTransactionRequest) (UpdatesClass, error)) {
func (a *application) paymentsAssignPlayMarketTransactio(ctx context.Context, request *tg.PaymentsAssignPlayMarketTransactionRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnPaymentsCanPurchasePremium(f func(ctx context.Context, purpose InputStorePaymentPurposeClass) (bool, error)) {
func (a *application) paymentsCanPurchasePremium(ctx context.Context, purpose tg.InputStorePaymentPurposeClass) (bool , error) {
    return true, nil
}


// OnStickersCreateStickerSet(f func(ctx context.Context, request *StickersCreateStickerSetRequest) (MessagesStickerSetClass, error)) {
func (a *application) stickersCreateStickerSet(ctx context.Context, request *tg.StickersCreateStickerSetRequest) (tg.MessagesStickerSetClass, error) {
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
		}{ Archived:true , Official: true, Masks: true, Animated: true, Videos: true, Emojis: true, InstalledDate: 555,
			ID: 555, AccessHash: 345345, Title: "ksdkjfskf", ShortName: "klsjdflsj",  ThumbDCID: 2, ThumbVersion: 454, ThumbDocumentID: 34534, Count:234 , Hash: 435345},
	}, nil
}


// OnStickersRemoveStickerFromSet(f func(ctx context.Context, sticker InputDocumentClass) (MessagesStickerSetClass, error)) {
func (a *application) stickersRemoveStickerFromSet(ctx context.Context, sticker tg.InputDocumentClass) (tg.MessagesStickerSetClass, error) {
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
		}{ Archived:true , Official: true, Masks: true, Animated: true, Videos: true, Emojis: true, InstalledDate: 555,
			ID: 555, AccessHash: 345345, Title: "ksdkjfskf", ShortName: "klsjdflsj",  ThumbDCID: 2, ThumbVersion: 454, ThumbDocumentID: 34534, Count:234 , Hash: 435345},
	}, nil
}


// OnStickersChangeStickerPosition(f func(ctx context.Context, request *StickersChangeStickerPositionRequest) (MessagesStickerSetClass, error)) {
func (a *application) stickersChangeStickerPositio(ctx context.Context, request *tg.StickersChangeStickerPositionRequest) (tg.MessagesStickerSetClass, error) {
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
		}{ Archived:true , Official: true, Masks: true, Animated: true, Videos: true, Emojis: true, InstalledDate: 555,
			ID: 555, AccessHash: 345345, Title: "ksdkjfskf", ShortName: "klsjdflsj",  ThumbDCID: 2, ThumbVersion: 454, ThumbDocumentID: 34534, Count:234 , Hash: 435345},
	}, nil
}


// OnStickersAddStickerToSet(f func(ctx context.Context, request *StickersAddStickerToSetRequest) (MessagesStickerSetClass, error)) {
func (a *application) stickersAddStickerToSet(ctx context.Context, request *tg.StickersAddStickerToSetRequest) (tg.MessagesStickerSetClass, error) {
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
		}{ Archived:true , Official: true, Masks: true, Animated: true, Videos: true, Emojis: true, InstalledDate: 555,
			ID: 555, AccessHash: 345345, Title: "ksdkjfskf", ShortName: "klsjdflsj",  ThumbDCID: 2, ThumbVersion: 454, ThumbDocumentID: 34534, Count:234 , Hash: 435345},
	}, nil
}


// OnStickersSetStickerSetThumb(f func(ctx context.Context, request *StickersSetStickerSetThumbRequest) (MessagesStickerSetClass, error)) {
func (a *application) stickersSetStickerSetThumb(ctx context.Context, request *tg.StickersSetStickerSetThumbRequest) (tg.MessagesStickerSetClass, error) {
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
		}{ Archived:true , Official: true, Masks: true, Animated: true, Videos: true, Emojis: true, InstalledDate: 555,
			ID: 555, AccessHash: 345345, Title: "ksdkjfskf", ShortName: "klsjdflsj",  ThumbDCID: 2, ThumbVersion: 454, ThumbDocumentID: 34534, Count:234 , Hash: 435345},
	}, nil
}


// OnStickersCheckShortName(f func(ctx context.Context, shortname string) (bool, error)) {
func (a *application) stickersCheckShortName(ctx context.Context, shortname string) (bool , error) {
    return true, nil
}


// OnStickersSuggestShortName(f func(ctx context.Context, title string) (*StickersSuggestedShortName, error)) {
func (a *application) stickersSuggestShortName(ctx context.Context, title string) (*tg.StickersSuggestedShortName, error) {
    return &tg.StickersSuggestedShortName{
    	ShortName: "theLostLamb333",
	}, nil
}


// OnStickersChangeSticker(f func(ctx context.Context, request *StickersChangeStickerRequest) (MessagesStickerSetClass, error)) {
func (a *application) stickersChangeSticker(ctx context.Context, request *tg.StickersChangeStickerRequest) (tg.MessagesStickerSetClass, error) {
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
		}{ Archived:true , Official: true, Masks: true, Animated: true, Videos: true, Emojis: true, InstalledDate: 555,
			ID: 555, AccessHash: 345345, Title: "ksdkjfskf", ShortName: "klsjdflsj",  ThumbDCID: 2, ThumbVersion: 454, ThumbDocumentID: 34534, Count:234 , Hash: 435345},
	}, nil
}


// OnStickersRenameStickerSet(f func(ctx context.Context, request *StickersRenameStickerSetRequest) (MessagesStickerSetClass, error)) {
func (a *application) stickersRenameStickerSet(ctx context.Context, request *tg.StickersRenameStickerSetRequest) (tg.MessagesStickerSetClass, error) {
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
		}{ Archived:true , Official: true, Masks: true, Animated: true, Videos: true, Emojis: true, InstalledDate: 555,
			ID: 555, AccessHash: 345345, Title: "ksdkjfskf", ShortName: "klsjdflsj",  ThumbDCID: 2, ThumbVersion: 454, ThumbDocumentID: 34534, Count:234 , Hash: 435345},
	}, nil
}


// OnStickersDeleteStickerSet(f func(ctx context.Context, stickerset InputStickerSetClass) (bool, error)) {
func (a *application) stickersDeleteStickerSet(ctx context.Context, stickerset tg.InputStickerSetClass) (bool , error) {
    return true, nil
}


// OnPhoneGetCallConfig(f func(ctx context.Context) (*DataJSON, error)) {
func (a *application) phoneGetCallConfig(ctx context.Context) (*tg.DataJSON,error) {
    return &tg.DataJSON{
    	Data: "theLostLamb333",
	}, nil
}


// OnPhoneRequestCall(f func(ctx context.Context, request *PhoneRequestCallRequest) (*PhonePhoneCall, error)) {
func (a *application) phoneRequestCall(ctx context.Context, request *tg.PhoneRequestCallRequest) (*tg.PhonePhoneCall, error) {
    return &tg.PhonePhoneCall{
    	PhoneCall: &tg.PhoneCallAccepted{
    		Video: true,
    		ID: 333444,
    		AccessHash: 444333,
    		Date: 34334,
    		AdminID: 333,
    		ParticipantID: 444,
    		GB: []byte{'4'},
    		Protocol: struct {
				Flags           bin.Fields
				UDPP2P          bool
				UDPReflector    bool
				MinLayer        int
				MaxLayer        int
				LibraryVersions []string
			}{UDPP2P:true , UDPReflector: true, MinLayer: 33, MaxLayer: 44, LibraryVersions: []string{"333444"}},
		},
		Users: []tg.UserClass{
			&tg.User{
				Username: "333444",
				Phone: "444333",
			},
		},

	}, nil
}


// OnPhoneAcceptCall(f func(ctx context.Context, request *PhoneAcceptCallRequest) (*PhonePhoneCall, error)) {
func (a *application) phoneAcceptCall(ctx context.Context, request *tg.PhoneAcceptCallRequest) (*tg.PhonePhoneCall, error) {
	return &tg.PhonePhoneCall{
		PhoneCall: &tg.PhoneCallAccepted{
			Video: true,
			ID: 333444,
			AccessHash: 444333,
			Date: 34334,
			AdminID: 333,
			ParticipantID: 444,
			GB: []byte{'4'},
			Protocol: struct {
				Flags           bin.Fields
				UDPP2P          bool
				UDPReflector    bool
				MinLayer        int
				MaxLayer        int
				LibraryVersions []string
			}{UDPP2P:true , UDPReflector: true, MinLayer: 33, MaxLayer: 44, LibraryVersions: []string{"333444"}},
		},
		Users: []tg.UserClass{
			&tg.User{
				Username: "333444",
				Phone: "444333",
			},
		},

	}, nil
}


// OnPhoneConfirmCall(f func(ctx context.Context, request *PhoneConfirmCallRequest) (*PhonePhoneCall, error)) {
func (a *application) phoneConfirmCall(ctx context.Context, request *tg.PhoneConfirmCallRequest) (*tg.PhonePhoneCall, error) {
	return &tg.PhonePhoneCall{
		PhoneCall: &tg.PhoneCallAccepted{
			Video: true,
			ID: 333444,
			AccessHash: 444333,
			Date: 34334,
			AdminID: 333,
			ParticipantID: 444,
			GB: []byte{'4'},
			Protocol: struct {
				Flags           bin.Fields
				UDPP2P          bool
				UDPReflector    bool
				MinLayer        int
				MaxLayer        int
				LibraryVersions []string
			}{UDPP2P:true , UDPReflector: true, MinLayer: 33, MaxLayer: 44, LibraryVersions: []string{"333444"}},
		},
		Users: []tg.UserClass{
			&tg.User{
				Username: "333444",
				Phone: "444333",
			},
		},

	}, nil
}


// OnPhoneReceivedCall(f func(ctx context.Context, peer InputPhoneCall) (bool, error)) {
func (a *application) phoneReceivedCall(ctx context.Context, peer tg.InputPhoneCall) (bool , error) {
    return true, nil
}


// OnPhoneDiscardCall(f func(ctx context.Context, request *PhoneDiscardCallRequest) (UpdatesClass, error)) {
func (a *application) phoneDiscardCall(ctx context.Context, request *tg.PhoneDiscardCallRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnPhoneSetCallRating(f func(ctx context.Context, request *PhoneSetCallRatingRequest) (UpdatesClass, error)) {
func (a *application) phoneSetCallRating(ctx context.Context, request *tg.PhoneSetCallRatingRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnPhoneSaveCallDebug(f func(ctx context.Context, request *PhoneSaveCallDebugRequest) (bool, error)) {
func (a *application) phoneSaveCallDebug(ctx context.Context, request *tg.PhoneSaveCallDebugRequest) (bool , error) {
    return true, nil
}


// OnPhoneSendSignalingData(f func(ctx context.Context, request *PhoneSendSignalingDataRequest) (bool, error)) {
func (a *application) phoneSendSignalingData(ctx context.Context, request *tg.PhoneSendSignalingDataRequest) (bool , error) {
    return true, nil
}


// OnPhoneCreateGroupCall(f func(ctx context.Context, request *PhoneCreateGroupCallRequest) (UpdatesClass, error)) {
func (a *application) phoneCreateGroupCall(ctx context.Context, request *tg.PhoneCreateGroupCallRequest) (tg.UpdatesClass, error) {
    return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnPhoneJoinGroupCall(f func(ctx context.Context, request *PhoneJoinGroupCallRequest) (UpdatesClass, error)) {
func (a *application) phoneJoinGroupCall(ctx context.Context, request *tg.PhoneJoinGroupCallRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnPhoneLeaveGroupCall(f func(ctx context.Context, request *PhoneLeaveGroupCallRequest) (UpdatesClass, error)) {
func (a *application) phoneLeaveGroupCall(ctx context.Context, request *tg.PhoneLeaveGroupCallRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnPhoneInviteToGroupCall(f func(ctx context.Context, request *PhoneInviteToGroupCallRequest) (UpdatesClass, error)) {
func (a *application) phoneInviteToGroupCall(ctx context.Context, request *tg.PhoneInviteToGroupCallRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnPhoneDiscardGroupCall(f func(ctx context.Context, call InputGroupCall) (UpdatesClass, error)) {
func (a *application) phoneDiscardGroupCall(ctx context.Context, call tg.InputGroupCall) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnPhoneToggleGroupCallSettings(f func(ctx context.Context, request *PhoneToggleGroupCallSettingsRequest) (UpdatesClass, error)) {
func (a *application) phoneToggleGroupCallSettings(ctx context.Context, request *tg.PhoneToggleGroupCallSettingsRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnPhoneGetGroupCall(f func(ctx context.Context, request *PhoneGetGroupCallRequest) (*PhoneGroupCall, error)) {
func (a *application) phoneGetGroupCall(ctx context.Context, request *tg.PhoneGetGroupCallRequest) (*tg.PhoneGroupCall, error) {
    return &tg.PhoneGroupCall{
    	ParticipantsNextOffset: "333444",
	}, nil
}


// OnPhoneGetGroupParticipants(f func(ctx context.Context, request *PhoneGetGroupParticipantsRequest) (*PhoneGroupParticipants, error)) {
func (a *application) phoneGetGroupParticipants(ctx context.Context, request *tg.PhoneGetGroupParticipantsRequest) (*tg.PhoneGroupParticipants, error) {
    return &tg.PhoneGroupParticipants{
    	Count: 3344,
    	NextOffset: "444333",
	}, nil
}


// OnPhoneCheckGroupCall(f func(ctx context.Context, request *PhoneCheckGroupCallRequest) ([]int, error)) {
func (a *application) phoneCheckGroupCall(ctx context.Context, request *tg.PhoneCheckGroupCallRequest) ([]int, error) {
	return []int{333444}, nil
}


// OnPhoneToggleGroupCallRecord(f func(ctx context.Context, request *PhoneToggleGroupCallRecordRequest) (UpdatesClass, error)) {
func (a *application) phoneToggleGroupCallRecord(ctx context.Context, request *tg.PhoneToggleGroupCallRecordRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnPhoneEditGroupCallParticipant(f func(ctx context.Context, request *PhoneEditGroupCallParticipantRequest) (UpdatesClass, error)) {
func (a *application) phoneEditGroupCallParticipant(ctx context.Context, request *tg.PhoneEditGroupCallParticipantRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnPhoneEditGroupCallTitle(f func(ctx context.Context, request *PhoneEditGroupCallTitleRequest) (UpdatesClass, error)) {
func (a *application) phoneEditGroupCallTitle(ctx context.Context, request *tg.PhoneEditGroupCallTitleRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnPhoneGetGroupCallJoinAs(f func(ctx context.Context, peer InputPeerClass) (*PhoneJoinAsPeers, error)) {
func (a *application) phoneGetGroupCallJoinAs(ctx context.Context, peer tg.InputPeerClass) (*tg.PhoneJoinAsPeers, error) {
    return &tg.PhoneJoinAsPeers{
    	Users: []tg.UserClass{
    		&tg.User{
    			Username: "333444",
    			Phone: "444333",
			},
		},
	},nil
}


// OnPhoneExportGroupCallInvite(f func(ctx context.Context, request *PhoneExportGroupCallInviteRequest) (*PhoneExportedGroupCallInvite, error)) {
func (a *application) phoneExportGroupCallInvite(ctx context.Context, request *tg.PhoneExportGroupCallInviteRequest) (*tg.PhoneExportedGroupCallInvite, error) {
    return &tg.PhoneExportedGroupCallInvite{
    	Link: "333444",
	}, nil
}


// OnPhoneToggleGroupCallStartSubscription(f func(ctx context.Context, request *PhoneToggleGroupCallStartSubscriptionRequest) (UpdatesClass, error)) {
func (a *application) phoneToggleGroupCallStartSubscriptio(ctx context.Context, request *tg.PhoneToggleGroupCallStartSubscriptionRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnPhoneStartScheduledGroupCall(f func(ctx context.Context, call InputGroupCall) (UpdatesClass, error)) {
func (a *application) phoneStartScheduledGroupCall(ctx context.Context, call tg.InputGroupCall) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnPhoneSaveDefaultGroupCallJoinAs(f func(ctx context.Context, request *PhoneSaveDefaultGroupCallJoinAsRequest) (bool, error)) {
func (a *application) phoneSaveDefaultGroupCallJoinAs(ctx context.Context, request *tg.PhoneSaveDefaultGroupCallJoinAsRequest) (bool , error) {
    return true, nil
}


// OnPhoneJoinGroupCallPresentation(f func(ctx context.Context, request *PhoneJoinGroupCallPresentationRequest) (UpdatesClass, error)) {
func (a *application) phoneJoinGroupCallPresentatio(ctx context.Context, request *tg.PhoneJoinGroupCallPresentationRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnPhoneLeaveGroupCallPresentation(f func(ctx context.Context, call InputGroupCall) (UpdatesClass, error)) {
func (a *application) phoneLeaveGroupCallPresentatio(ctx context.Context, call tg.InputGroupCall) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnPhoneGetGroupCallStreamChannels(f func(ctx context.Context, call InputGroupCall) (*PhoneGroupCallStreamChannels, error)) {
func (a *application) phoneGetGroupCallStreamChannels(ctx context.Context, call tg.InputGroupCall) (*tg.PhoneGroupCallStreamChannels, error) {
    return &tg.PhoneGroupCallStreamChannels{
    	Channels: []tg.GroupCallStreamChannel{
    		{
    			Channel: 333444,
    			Scale: 444333,
    			LastTimestampMs: 343434,
			},
		},
	}, nil
}


// OnPhoneGetGroupCallStreamRtmpURL(f func(ctx context.Context, request *PhoneGetGroupCallStreamRtmpURLRequest) (*PhoneGroupCallStreamRtmpURL, error)) {
func (a *application) phoneGetGroupCallStreamRtmpURL(ctx context.Context, request *tg.PhoneGetGroupCallStreamRtmpURLRequest) (*tg.PhoneGroupCallStreamRtmpURL, error) {
    return &tg.PhoneGroupCallStreamRtmpURL{
    	URL: "333444",
    	Key: "444333",
	}, nil
}


// OnPhoneSaveCallLog(f func(ctx context.Context, request *PhoneSaveCallLogRequest) (bool, error)) {
func (a *application) phoneSaveCallLog(ctx context.Context, request *tg.PhoneSaveCallLogRequest) (bool , error) {
    return true, nil
}


// OnLangpackGetLangPack(f func(ctx context.Context, request *LangpackGetLangPackRequest) (*LangPackDifference, error)) {
func (a *application) langpackGetLangPack(ctx context.Context, request *tg.LangpackGetLangPackRequest) (*tg.LangPackDifference, error) {
    return &tg.LangPackDifference{
    	LangCode: "theLostLamb",
    	FromVersion: 444333,
    	Version: 333444,
	}, nil
}


// OnLangpackGetStrings(f func(ctx context.Context, request *LangpackGetStringsRequest) ([]LangPackStringClass, error)) {
func (a *application) langpackGetStrings(ctx context.Context, request *tg.LangpackGetStringsRequest) ([]tg.LangPackStringClass, error) {
	return []tg.LangPackStringClass{
		&tg.LangPackString{
			Key: "333444",
			Value: "444333",
		},
	},nil
}


// OnLangpackGetDifference(f func(ctx context.Context, request *LangpackGetDifferenceRequest) (*LangPackDifference, error)) {
func (a *application) langpackGetDifference(ctx context.Context, request *tg.LangpackGetDifferenceRequest) (*tg.LangPackDifference, error) {
	return &tg.LangPackDifference{
		LangCode: "theLostLamb",
		FromVersion: 444333,
		Version: 333444,
	}, nil
}


// OnLangpackGetLanguages(f func(ctx context.Context, langpack string) ([]LangPackLanguage, error)) {
func (a *application) langpackGetLanguages(ctx context.Context, langpack string) ([]tg.LangPackLanguage, error) {
    return []tg.LangPackLanguage{
    	tg.LangPackLanguage{
			Official: true,
			Name: "theLostLamb",
			NativeName: "lostLamb",
			LangCode: "333444",
			BaseLangCode: "444333",
		},
	}, nil
}


// OnLangpackGetLanguage(f func(ctx context.Context, request *LangpackGetLanguageRequest) (*LangPackLanguage, error)) {
func (a *application) langpackGetLanguage(ctx context.Context, request *tg.LangpackGetLanguageRequest) (*tg.LangPackLanguage, error) {
    return &tg.LangPackLanguage{
    	Official: true,
    	Name: "theLostLamb",
    	NativeName: "lostLamb",
    	LangCode: "333444",
    	BaseLangCode: "444333",
	}, nil
}


// OnFoldersEditPeerFolders(f func(ctx context.Context, folderpeers []InputFolderPeer) (UpdatesClass, error)) {
func (a *application) foldersEditPeerFolders(ctx context.Context, folderpeers []tg.InputFolderPeer) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnStatsGetBroadcastStats(f func(ctx context.Context, request *StatsGetBroadcastStatsRequest) (*StatsBroadcastStats, error)) {
func (a *application) statsGetBroadcastStats(ctx context.Context, request *tg.StatsGetBroadcastStatsRequest) (*tg.StatsBroadcastStats, error) {
    return &tg.StatsBroadcastStats{
    	EnabledNotifications: struct {
			Part  float64
			Total float64
		}{Part: 333.444, Total: 444.333},
	}, nil
}


// OnStatsLoadAsyncGraph(f func(ctx context.Context, request *StatsLoadAsyncGraphRequest) (StatsGraphClass, error)) {
func (a *application) statsLoadAsyncGraph(ctx context.Context, request *tg.StatsLoadAsyncGraphRequest) (tg.StatsGraphClass, error) {
    return &tg.StatsGraph{
    	ZoomToken: "333444",
	}, nil
}


// OnStatsGetMegagroupStats(f func(ctx context.Context, request *StatsGetMegagroupStatsRequest) (*StatsMegagroupStats, error)) {
func (a *application) statsGetMegagroupStats(ctx context.Context, request *tg.StatsGetMegagroupStatsRequest) (*tg.StatsMegagroupStats, error) {
    return &tg.StatsMegagroupStats{
    	Messages: struct {
			Current  float64
			Previous float64
		}{Current: 333.444, Previous: 444.333},
	}, nil
}


// OnStatsGetMessagePublicForwards(f func(ctx context.Context, request *StatsGetMessagePublicForwardsRequest) (MessagesMessagesClass, error)) {
func (a *application) statsGetMessagePublicForwards(ctx context.Context, request *tg.StatsGetMessagePublicForwardsRequest) (tg.MessagesMessagesClass, error) {
    return &tg.MessagesMessages{
    	Users: []tg.UserClass{
    		&tg.User{
    			Username: "333444",
    			Phone: "444333",
			},
		},
	}, nil
}


// OnStatsGetMessageStats(f func(ctx context.Context, request *StatsGetMessageStatsRequest) (*StatsMessageStats, error)) {
func (a *application) statsGetMessageStats(ctx context.Context, request *tg.StatsGetMessageStatsRequest) (*tg.StatsMessageStats, error) {
    return &tg.StatsMessageStats{
    	ViewsGraph: &tg.StatsGraph{
    		ZoomToken: "333444",
		},
	}, nil
}


// OnChatlistsExportChatlistInvite(f func(ctx context.Context, request *ChatlistsExportChatlistInviteRequest) (*ChatlistsExportedChatlistInvite, error)) {
func (a *application) chatlistsExportChatlistInvite(ctx context.Context, request *tg.ChatlistsExportChatlistInviteRequest) (*tg.ChatlistsExportedChatlistInvite, error) {
    return &tg.ChatlistsExportedChatlistInvite{
    	Invite: tg.ExportedChatlistInvite{
    		Title: "333444",
    		URL: "444333",
		},
	}, nil
}


// OnChatlistsDeleteExportedInvite(f func(ctx context.Context, request *ChatlistsDeleteExportedInviteRequest) (bool, error)) {
func (a *application) chatlistsDeleteExportedInvite(ctx context.Context, request *tg.ChatlistsDeleteExportedInviteRequest) (bool , error) {
    return true, nil
}


// OnChatlistsEditExportedInvite(f func(ctx context.Context, request *ChatlistsEditExportedInviteRequest) (*ExportedChatlistInvite, error)) {
func (a *application) chatlistsEditExportedInvite(ctx context.Context, request *tg.ChatlistsEditExportedInviteRequest) (*tg.ExportedChatlistInvite, error) {
    return &tg.ExportedChatlistInvite{
		Title: "333444",
		URL: "444333",
	}, nil
}


// OnChatlistsGetExportedInvites(f func(ctx context.Context, chatlist InputChatlistDialogFilter) (*ChatlistsExportedInvites, error)) {
func (a *application) chatlistsGetExportedInvites(ctx context.Context, chatlist tg.InputChatlistDialogFilter) (*tg.ChatlistsExportedInvites, error) {
    return &tg.ChatlistsExportedInvites{
    	Invites: []tg.ExportedChatlistInvite{
    		{
    			Title: "333444",
    			URL: "444333",
			},
		},
	}, nil
}


// OnChatlistsCheckChatlistInvite(f func(ctx context.Context, slug string) (ChatlistsChatlistInviteClass, error)) {
func (a *application) chatlistsCheckChatlistInvite(ctx context.Context, slug string) (tg.ChatlistsChatlistInviteClass, error) {
    return &tg.ChatlistsChatlistInvite{
    	Title: "333444",
    	Emoticon: "444333",
	}, nil
}


// OnChatlistsJoinChatlistInvite(f func(ctx context.Context, request *ChatlistsJoinChatlistInviteRequest) (UpdatesClass, error)) {
func (a *application) chatlistsJoinChatlistInvite(ctx context.Context, request *tg.ChatlistsJoinChatlistInviteRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnChatlistsGetChatlistUpdates(f func(ctx context.Context, chatlist InputChatlistDialogFilter) (*ChatlistsChatlistUpdates, error)) {
func (a *application) chatlistsGetChatlistUpdates(ctx context.Context, chatlist tg.InputChatlistDialogFilter) (*tg.ChatlistsChatlistUpdates, error) {
    return &tg.ChatlistsChatlistUpdates{
    	Users: []tg.UserClass{
			&tg.User{
				Username: "333444",
				Phone: "444333",
			},
		},
	}, nil
}


// OnChatlistsJoinChatlistUpdates(f func(ctx context.Context, request *ChatlistsJoinChatlistUpdatesRequest) (UpdatesClass, error)) {
func (a *application) chatlistsJoinChatlistUpdates(ctx context.Context, request *tg.ChatlistsJoinChatlistUpdatesRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnChatlistsHideChatlistUpdates(f func(ctx context.Context, chatlist InputChatlistDialogFilter) (bool, error)) {
func (a *application) chatlistsHideChatlistUpdates(ctx context.Context, chatlist tg.InputChatlistDialogFilter) (bool , error) {
    return true, nil
}


// OnChatlistsGetLeaveChatlistSuggestions(f func(ctx context.Context, chatlist InputChatlistDialogFilter) ([]PeerClass, error)) {
func (a *application) chatlistsGetLeaveChatlistSuggestions(ctx context.Context, chatlist tg.InputChatlistDialogFilter) ([]tg.PeerClass, error) {
	return nil, nil
}


// OnChatlistsLeaveChatlist(f func(ctx context.Context, request *ChatlistsLeaveChatlistRequest) (UpdatesClass, error)) {
func (a *application) chatlistsLeaveChatlist(ctx context.Context, request *tg.ChatlistsLeaveChatlistRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnStoriesCanSendStory(f func(ctx context.Context, peer InputPeerClass) (bool, error)) {
func (a *application) storiesCanSendStory(ctx context.Context, peer tg.InputPeerClass) (bool , error) {
    return true, nil
}


// OnStoriesSendStory(f func(ctx context.Context, request *StoriesSendStoryRequest) (UpdatesClass, error)) {
func (a *application) storiesSendStory(ctx context.Context, request *tg.StoriesSendStoryRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnStoriesEditStory(f func(ctx context.Context, request *StoriesEditStoryRequest) (UpdatesClass, error)) {
func (a *application) storiesEditStory(ctx context.Context, request *tg.StoriesEditStoryRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnStoriesDeleteStories(f func(ctx context.Context, request *StoriesDeleteStoriesRequest) ([]int, error)) {
func (a *application) storiesDeleteStories(ctx context.Context, request *tg.StoriesDeleteStoriesRequest) ([]int, error) {
	return []int{333444}, nil
}


// OnStoriesTogglePinned(f func(ctx context.Context, request *StoriesTogglePinnedRequest) ([]int, error)) {
func (a *application) storiesTogglePinned(ctx context.Context, request *tg.StoriesTogglePinnedRequest) ([]int, error) {
	return []int{333444}, nil
}


// OnStoriesGetAllStories(f func(ctx context.Context, request *StoriesGetAllStoriesRequest) (StoriesAllStoriesClass, error)) {
func (a *application) storiesGetAllStories(ctx context.Context, request *tg.StoriesGetAllStoriesRequest) (tg.StoriesAllStoriesClass, error) {
    return &tg.StoriesAllStories{
		Users: []tg.UserClass{
			&tg.User{
				Username: "333444",
				Phone: "444333",
			},
		},
	}, nil
}


// OnStoriesGetPinnedStories(f func(ctx context.Context, request *StoriesGetPinnedStoriesRequest) (*StoriesStories, error)) {
func (a *application) storiesGetPinnedStories(ctx context.Context, request *tg.StoriesGetPinnedStoriesRequest) (*tg.StoriesStories, error) {
    return &tg.StoriesStories{
    	Count: 333444,
	}, nil
}


// OnStoriesGetStoriesArchive(f func(ctx context.Context, request *StoriesGetStoriesArchiveRequest) (*StoriesStories, error)) {
func (a *application) storiesGetStoriesArchive(ctx context.Context, request *tg.StoriesGetStoriesArchiveRequest) (*tg.StoriesStories, error) {
	return &tg.StoriesStories{
		Count: 333444,
	}, nil
}


// OnStoriesGetStoriesByID(f func(ctx context.Context, request *StoriesGetStoriesByIDRequest) (*StoriesStories, error)) {
func (a *application) storiesGetStoriesByID(ctx context.Context, request *tg.StoriesGetStoriesByIDRequest) (*tg.StoriesStories, error) {
	return &tg.StoriesStories{
		Count: 333444,
	}, nil
}


// OnStoriesToggleAllStoriesHidden(f func(ctx context.Context, hidden bool) (bool, error)) {
func (a *application) storiesToggleAllStoriesHidde(ctx context.Context, hidden bool) (bool , error) {
    return true, nil
}


// OnStoriesReadStories(f func(ctx context.Context, request *StoriesReadStoriesRequest) ([]int, error)) {
func (a *application) storiesReadStories(ctx context.Context, request *tg.StoriesReadStoriesRequest) ([]int, error) {
	return []int{333444}, nil
}


// OnStoriesIncrementStoryViews(f func(ctx context.Context, request *StoriesIncrementStoryViewsRequest) (bool, error)) {
func (a *application) storiesIncrementStoryViews(ctx context.Context, request *tg.StoriesIncrementStoryViewsRequest) (bool , error) {
    return true, nil
}


// OnStoriesGetStoryViewsList(f func(ctx context.Context, request *StoriesGetStoryViewsListRequest) (*StoriesStoryViewsList, error)) {
func (a *application) storiesGetStoryViewsList(ctx context.Context, request *tg.StoriesGetStoryViewsListRequest) (*tg.StoriesStoryViewsList, error) {
    return &tg.StoriesStoryViewsList{
    	Count: 333444,
    	NextOffset: "333444",
	}, nil
}


// OnStoriesGetStoriesViews(f func(ctx context.Context, request *StoriesGetStoriesViewsRequest) (*StoriesStoryViews, error)) {
func (a *application) storiesGetStoriesViews(ctx context.Context, request *tg.StoriesGetStoriesViewsRequest) (*tg.StoriesStoryViews, error) {
    return &tg.StoriesStoryViews{
		Users: []tg.UserClass{
			&tg.User{
				Username: "333444",
				Phone: "444333",
			},
		},
	}, nil
}


// OnStoriesExportStoryLink(f func(ctx context.Context, request *StoriesExportStoryLinkRequest) (*ExportedStoryLink, error)) {
func (a *application) storiesExportStoryLink(ctx context.Context, request *tg.StoriesExportStoryLinkRequest) (*tg.ExportedStoryLink, error) {
    return &tg.ExportedStoryLink{
    	Link: "333444",
	}, nil
}


// OnStoriesReport(f func(ctx context.Context, request *StoriesReportRequest) (bool, error)) {
func (a *application) storiesReport(ctx context.Context, request *tg.StoriesReportRequest) (bool , error) {
    return true, nil
}


// OnStoriesActivateStealthMode(f func(ctx context.Context, request *StoriesActivateStealthModeRequest) (UpdatesClass, error)) {
func (a *application) storiesActivateStealthMode(ctx context.Context, request *tg.StoriesActivateStealthModeRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnStoriesSendReaction(f func(ctx context.Context, request *StoriesSendReactionRequest) (UpdatesClass, error)) {
func (a *application) storiesSendReactio(ctx context.Context, request *tg.StoriesSendReactionRequest) (tg.UpdatesClass, error) {
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnStoriesGetPeerStories(f func(ctx context.Context, peer InputPeerClass) (*StoriesPeerStories, error)) {
func (a *application) storiesGetPeerStories(ctx context.Context, peer tg.InputPeerClass) (*tg.StoriesPeerStories, error) {
    return &tg.StoriesPeerStories{
		Users: []tg.UserClass{
			&tg.User{
				Username: "333444",
				Phone: "444333",
			},
		},
	}, nil
}


// OnStoriesGetAllReadPeerStories(f func(ctx context.Context) (UpdatesClass, error)) {
func (a *application) storiesGetAllReadPeerStories(ctx context.Context) (tg.UpdatesClass,error){
        return &tg.Updates{
    	Date: 33334444,
    	Seq: 444333,
	}, nil
}


// OnStoriesGetPeerMaxIDs(f func(ctx context.Context, id []InputPeerClass) ([]int, error)) {
func (a *application) storiesGetPeerMaxIDs(ctx context.Context, id []tg.InputPeerClass) ([]int, error) {
	return []int{333444}, nil
}


// OnStoriesGetChatsToSend(f func(ctx context.Context) (MessagesChatsClass, error)) {
func (a *application) storiesGetChatsToSend(ctx context.Context) (tg.MessagesChatsClass, error) {
	return &tg.MessagesChats{
		Chats: []tg.ChatClass{
			&tg.Chat{
				ID: 333444,
				Title: "333444",
			},
		},
	}, nil
}


// OnStoriesTogglePeerStoriesHidden(f func(ctx context.Context, request *StoriesTogglePeerStoriesHiddenRequest) (bool, error)) {
func (a *application) storiesTogglePeerStoriesHidde(ctx context.Context, request *tg.StoriesTogglePeerStoriesHiddenRequest) (bool , error) {
    return true, nil
}


// OnStoriesGetBoostsStatus(f func(ctx context.Context, peer InputPeerClass) (*StoriesBoostsStatus, error)) {
func (a *application) storiesGetBoostsStatus(ctx context.Context, peer tg.InputPeerClass) (*tg.StoriesBoostsStatus, error) {
    return &tg.StoriesBoostsStatus{
    	MyBoost: true,
    	Level: 333444,
    	BoostURL: "333444",
	}, nil
}


// OnStoriesGetBoostersList(f func(ctx context.Context, request *StoriesGetBoostersListRequest) (*StoriesBoostersList, error)) {
func (a *application) storiesGetBoostersList(ctx context.Context, request *tg.StoriesGetBoostersListRequest) (*tg.StoriesBoostersList, error) {
    return &tg.StoriesBoostersList{
		Users: []tg.UserClass{
			&tg.User{
				Username: "333444",
				Phone: "444333",
			},
		},
	}, nil
}


// OnStoriesCanApplyBoost(f func(ctx context.Context, peer InputPeerClass) (StoriesCanApplyBoostResultClass, error)) {
func (a *application) storiesCanApplyBoost(ctx context.Context, peer tg.InputPeerClass) (tg.StoriesCanApplyBoostResultClass, error) {
	return &tg.StoriesCanApplyBoostReplace{
		Chats: []tg.ChatClass{
			&tg.Chat{
				ID: 333444,
				Title: "333444",
			},
		},

	},nil
}


// OnStoriesApplyBoost(f func(ctx context.Context, peer InputPeerClass) (bool, error)) {
func (a *application) storiesApplyBoost(ctx context.Context, peer tg.InputPeerClass) (bool , error) {
    return true, nil
}


// OnTestUseError(f func(ctx context.Context) (*Error, error)) {
func (a *application) testUseError(ctx context.Context) (*tg.Error,error) {
    return &tg.Error{
    	Text: "333444",
	}, nil
}


// OnTestUseConfigSimple(f func(ctx context.Context) (*HelpConfigSimple, error)) {
func (a *application) testUseConfigSimple(ctx context.Context) (*tg.HelpConfigSimple, error){
    return &tg.HelpConfigSimple{
    	Date: 333444,
    	Expires: 444333,
	}, nil
}


