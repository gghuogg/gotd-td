// Code generated by gotdgen, DO NOT EDIT.

package mt

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

// TypesMap returns mapping from type ids to TL type names.
func TypesMap() map[uint32]string {
	return map[uint32]string{
		ResPQTypeID:                      "resPQ#5162463",
		PQInnerDataTypeID:                "p_q_inner_data#83c95aec",
		PQInnerDataDCTypeID:              "p_q_inner_data_dc#a9f55f95",
		PQInnerDataTempDCTypeID:          "p_q_inner_data_temp_dc#56fddf88",
		ServerDHParamsFailTypeID:         "server_DH_params_fail#79cb045d",
		ServerDHParamsOkTypeID:           "server_DH_params_ok#d0e8075c",
		ServerDHInnerDataTypeID:          "server_DH_inner_data#b5890dba",
		ClientDHInnerDataTypeID:          "client_DH_inner_data#6643b654",
		DhGenOkTypeID:                    "dh_gen_ok#3bcbf734",
		DhGenRetryTypeID:                 "dh_gen_retry#46dc1fb9",
		DhGenFailTypeID:                  "dh_gen_fail#a69dae02",
		RPCResultTypeID:                  "rpc_result#f35c6d01",
		RPCErrorTypeID:                   "rpc_error#2144ca19",
		RPCAnswerUnknownTypeID:           "rpc_answer_unknown#5e2ad36e",
		RPCAnswerDroppedRunningTypeID:    "rpc_answer_dropped_running#cd78e586",
		RPCAnswerDroppedTypeID:           "rpc_answer_dropped#a43ad8b7",
		FutureSaltTypeID:                 "future_salt#949d9dc",
		FutureSaltsTypeID:                "future_salts#ae500895",
		PongTypeID:                       "pong#347773c5",
		DestroySessionOkTypeID:           "destroy_session_ok#e22045fc",
		DestroySessionNoneTypeID:         "destroy_session_none#62d350c9",
		NewSessionCreatedTypeID:          "new_session_created#9ec20908",
		MsgContainerTypeID:               "msg_container#73f1f8dc",
		MessageTypeID:                    "message#5bb8e511",
		MsgCopyTypeID:                    "msg_copy#e06046b2",
		GzipPackedTypeID:                 "gzip_packed#3072cfa1",
		MsgsAckTypeID:                    "msgs_ack#62d6b459",
		BadMsgNotificationTypeID:         "bad_msg_notification#a7eff811",
		BadServerSaltTypeID:              "bad_server_salt#edab447b",
		MsgResendReqTypeID:               "msg_resend_req#7d861a08",
		MsgsStateReqTypeID:               "msgs_state_req#da69fb52",
		MsgsStateInfoTypeID:              "msgs_state_info#4deb57d",
		MsgsAllInfoTypeID:                "msgs_all_info#8cc0d131",
		MsgDetailedInfoTypeID:            "msg_detailed_info#276d3ec6",
		MsgNewDetailedInfoTypeID:         "msg_new_detailed_info#809db6df",
		ReqPqRequestTypeID:               "req_pq#60469778",
		ReqPqMultiRequestTypeID:          "req_pq_multi#be7e8ef1",
		ReqDHParamsRequestTypeID:         "req_DH_params#d712e4be",
		SetClientDHParamsRequestTypeID:   "set_client_DH_params#f5045f1f",
		RPCDropAnswerRequestTypeID:       "rpc_drop_answer#58e4a740",
		GetFutureSaltsRequestTypeID:      "get_future_salts#b921bd04",
		PingRequestTypeID:                "ping#7abe77ec",
		PingDelayDisconnectRequestTypeID: "ping_delay_disconnect#f3427b8c",
		DestroySessionRequestTypeID:      "destroy_session#e7512126",
		HTTPWaitRequestTypeID:            "http_wait#9299359f",
	}
}

// NamesMap returns mapping from type names to TL type ids.
func NamesMap() map[string]uint32 {
	return map[string]uint32{
		"resPQ":                      ResPQTypeID,
		"p_q_inner_data":             PQInnerDataTypeID,
		"p_q_inner_data_dc":          PQInnerDataDCTypeID,
		"p_q_inner_data_temp_dc":     PQInnerDataTempDCTypeID,
		"server_DH_params_fail":      ServerDHParamsFailTypeID,
		"server_DH_params_ok":        ServerDHParamsOkTypeID,
		"server_DH_inner_data":       ServerDHInnerDataTypeID,
		"client_DH_inner_data":       ClientDHInnerDataTypeID,
		"dh_gen_ok":                  DhGenOkTypeID,
		"dh_gen_retry":               DhGenRetryTypeID,
		"dh_gen_fail":                DhGenFailTypeID,
		"rpc_result":                 RPCResultTypeID,
		"rpc_error":                  RPCErrorTypeID,
		"rpc_answer_unknown":         RPCAnswerUnknownTypeID,
		"rpc_answer_dropped_running": RPCAnswerDroppedRunningTypeID,
		"rpc_answer_dropped":         RPCAnswerDroppedTypeID,
		"future_salt":                FutureSaltTypeID,
		"future_salts":               FutureSaltsTypeID,
		"pong":                       PongTypeID,
		"destroy_session_ok":         DestroySessionOkTypeID,
		"destroy_session_none":       DestroySessionNoneTypeID,
		"new_session_created":        NewSessionCreatedTypeID,
		"msg_container":              MsgContainerTypeID,
		"message":                    MessageTypeID,
		"msg_copy":                   MsgCopyTypeID,
		"gzip_packed":                GzipPackedTypeID,
		"msgs_ack":                   MsgsAckTypeID,
		"bad_msg_notification":       BadMsgNotificationTypeID,
		"bad_server_salt":            BadServerSaltTypeID,
		"msg_resend_req":             MsgResendReqTypeID,
		"msgs_state_req":             MsgsStateReqTypeID,
		"msgs_state_info":            MsgsStateInfoTypeID,
		"msgs_all_info":              MsgsAllInfoTypeID,
		"msg_detailed_info":          MsgDetailedInfoTypeID,
		"msg_new_detailed_info":      MsgNewDetailedInfoTypeID,
		"req_pq":                     ReqPqRequestTypeID,
		"req_pq_multi":               ReqPqMultiRequestTypeID,
		"req_DH_params":              ReqDHParamsRequestTypeID,
		"set_client_DH_params":       SetClientDHParamsRequestTypeID,
		"rpc_drop_answer":            RPCDropAnswerRequestTypeID,
		"get_future_salts":           GetFutureSaltsRequestTypeID,
		"ping":                       PingRequestTypeID,
		"ping_delay_disconnect":      PingDelayDisconnectRequestTypeID,
		"destroy_session":            DestroySessionRequestTypeID,
		"http_wait":                  HTTPWaitRequestTypeID,
	}
}

// TypesConstructorMap maps type ids to constructors.
func TypesConstructorMap() map[uint32]func() bin.Object {
	return map[uint32]func() bin.Object{
		ResPQTypeID:                      func() bin.Object { return &ResPQ{} },
		PQInnerDataTypeID:                func() bin.Object { return &PQInnerData{} },
		PQInnerDataDCTypeID:              func() bin.Object { return &PQInnerDataDC{} },
		PQInnerDataTempDCTypeID:          func() bin.Object { return &PQInnerDataTempDC{} },
		ServerDHParamsFailTypeID:         func() bin.Object { return &ServerDHParamsFail{} },
		ServerDHParamsOkTypeID:           func() bin.Object { return &ServerDHParamsOk{} },
		ServerDHInnerDataTypeID:          func() bin.Object { return &ServerDHInnerData{} },
		ClientDHInnerDataTypeID:          func() bin.Object { return &ClientDHInnerData{} },
		DhGenOkTypeID:                    func() bin.Object { return &DhGenOk{} },
		DhGenRetryTypeID:                 func() bin.Object { return &DhGenRetry{} },
		DhGenFailTypeID:                  func() bin.Object { return &DhGenFail{} },
		RPCResultTypeID:                  func() bin.Object { return &RPCResult{} },
		RPCErrorTypeID:                   func() bin.Object { return &RPCError{} },
		RPCAnswerUnknownTypeID:           func() bin.Object { return &RPCAnswerUnknown{} },
		RPCAnswerDroppedRunningTypeID:    func() bin.Object { return &RPCAnswerDroppedRunning{} },
		RPCAnswerDroppedTypeID:           func() bin.Object { return &RPCAnswerDropped{} },
		FutureSaltTypeID:                 func() bin.Object { return &FutureSalt{} },
		FutureSaltsTypeID:                func() bin.Object { return &FutureSalts{} },
		PongTypeID:                       func() bin.Object { return &Pong{} },
		DestroySessionOkTypeID:           func() bin.Object { return &DestroySessionOk{} },
		DestroySessionNoneTypeID:         func() bin.Object { return &DestroySessionNone{} },
		NewSessionCreatedTypeID:          func() bin.Object { return &NewSessionCreated{} },
		MsgContainerTypeID:               func() bin.Object { return &MsgContainer{} },
		MessageTypeID:                    func() bin.Object { return &Message{} },
		MsgCopyTypeID:                    func() bin.Object { return &MsgCopy{} },
		GzipPackedTypeID:                 func() bin.Object { return &GzipPacked{} },
		MsgsAckTypeID:                    func() bin.Object { return &MsgsAck{} },
		BadMsgNotificationTypeID:         func() bin.Object { return &BadMsgNotification{} },
		BadServerSaltTypeID:              func() bin.Object { return &BadServerSalt{} },
		MsgResendReqTypeID:               func() bin.Object { return &MsgResendReq{} },
		MsgsStateReqTypeID:               func() bin.Object { return &MsgsStateReq{} },
		MsgsStateInfoTypeID:              func() bin.Object { return &MsgsStateInfo{} },
		MsgsAllInfoTypeID:                func() bin.Object { return &MsgsAllInfo{} },
		MsgDetailedInfoTypeID:            func() bin.Object { return &MsgDetailedInfo{} },
		MsgNewDetailedInfoTypeID:         func() bin.Object { return &MsgNewDetailedInfo{} },
		ReqPqRequestTypeID:               func() bin.Object { return &ReqPqRequest{} },
		ReqPqMultiRequestTypeID:          func() bin.Object { return &ReqPqMultiRequest{} },
		ReqDHParamsRequestTypeID:         func() bin.Object { return &ReqDHParamsRequest{} },
		SetClientDHParamsRequestTypeID:   func() bin.Object { return &SetClientDHParamsRequest{} },
		RPCDropAnswerRequestTypeID:       func() bin.Object { return &RPCDropAnswerRequest{} },
		GetFutureSaltsRequestTypeID:      func() bin.Object { return &GetFutureSaltsRequest{} },
		PingRequestTypeID:                func() bin.Object { return &PingRequest{} },
		PingDelayDisconnectRequestTypeID: func() bin.Object { return &PingDelayDisconnectRequest{} },
		DestroySessionRequestTypeID:      func() bin.Object { return &DestroySessionRequest{} },
		HTTPWaitRequestTypeID:            func() bin.Object { return &HTTPWaitRequest{} },
	}
}

// ClassConstructorsMap maps class schema name to constructors type ids.
func ClassConstructorsMap() map[string][]uint32 {
	return map[string][]uint32{
		BadMsgNotificationClassName: {
			BadMsgNotificationTypeID,
			BadServerSaltTypeID,
		},
		DestroySessionResClassName: {
			DestroySessionOkTypeID,
			DestroySessionNoneTypeID,
		},
		MsgDetailedInfoClassName: {
			MsgDetailedInfoTypeID,
			MsgNewDetailedInfoTypeID,
		},
		PQInnerDataClassName: {
			PQInnerDataTypeID,
			PQInnerDataDCTypeID,
			PQInnerDataTempDCTypeID,
		},
		RPCDropAnswerClassName: {
			RPCAnswerUnknownTypeID,
			RPCAnswerDroppedRunningTypeID,
			RPCAnswerDroppedTypeID,
		},
		ServerDHParamsClassName: {
			ServerDHParamsFailTypeID,
			ServerDHParamsOkTypeID,
		},
		SetClientDHParamsAnswerClassName: {
			DhGenOkTypeID,
			DhGenRetryTypeID,
			DhGenFailTypeID,
		},
	}
}
