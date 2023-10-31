package tgtest

import (
	"context"
	"encoding/binary"
	"github.com/go-faster/errors"
	"go.uber.org/multierr"
	"go.uber.org/zap"
	"log"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/internal/crypto"
	"github.com/gotd/td/internal/mt"
	"github.com/gotd/td/internal/proto"
	"github.com/gotd/td/tgerr"
	"github.com/gotd/td/transport"
)

func (s *Server) rpcHandle(ctx context.Context, c transport.Conn, b *bin.Buffer) error {
	log.Println("Server rpcHandle")
	m := &crypto.EncryptedMessage{}
	if err := m.DecodeWithoutCopy(b); err != nil {
		return errors.Wrap(err, "decode encrypted message")
	}

	log.Println("根据AuthKeyID获取用户session")
	key, ok := s.users.getSession(m.AuthKeyID)
	if !ok {
		return errors.New("invalid session")
	}

	log.Println("Server 密码解密")
	msg, err := s.cipher.Decrypt(key, m)
	if err != nil {
		return errors.Wrap(err, "decrypt message")
	}

	log.Println("定义session的值")
	session := Session{
		ID:      msg.SessionID,
		AuthKey: key,
	}
	log.Println("server根据session创建连接")
	if conn := s.users.createConnection(msg.SessionID, c); !conn.sentCreated() {
		s.log.Debug("Send handleSessionCreated event", zap.Inline(session))
		salt := int64(binary.LittleEndian.Uint64(key.ID[:]))
		if err := s.sendSessionCreated(ctx, session, salt); err != nil {
			return err
		}
	}

	log.Println("缓冲区现在包含明文消息负载")
	// Buffer now contains plaintext message payload.
	b.ResetTo(msg.Data())

	log.Println("Server处理请求")
	if err := s.handle(&Request{
		DC:         s.dcID,
		Session:    session,
		MsgID:      msg.MessageID,
		Buf:        b,
		RequestCtx: ctx,
	}); err != nil {
		return errors.Wrap(err, "handle")
	}

	return nil
}

func (s *Server) handle(req *Request) error {
	log.Println("Server处理具体的请求")
	in := req.Buf
	id, err := in.PeekID()
	if err != nil {
		return errors.Wrap(err, "peek id")
	}
	log.Println("使用日志进行标记")
	s.log.Debug("Got request",
		zap.Inline(req.Session),
		zap.Int64("msg_id", req.MsgID),
		zap.String("type", s.types.Get(id)),
	)
	log.Println("session",req.Session)
	log.Println("msg_id", req.MsgID)
	log.Println("type", s.types.Get(id))
	log.Printf("id, %p,%v,%v",&id,mt.MsgsAckTypeID,mt.MsgsAckTypeID)


	// TODO(tdakkota): unpack all containers
	switch id {
	//
	case mt.PingDelayDisconnectRequestTypeID:
		pingReq := mt.PingDelayDisconnectRequest{}
		if err := pingReq.Decode(in); err != nil {
			return err
		}

		return s.SendPong(req, pingReq.PingID)
	case mt.PingRequestTypeID:
		pingReq := mt.PingRequest{}
		if err := pingReq.Decode(in); err != nil {
			return err
		}

		return s.SendPong(req, pingReq.PingID)

	case mt.GetFutureSaltsRequestTypeID:
		saltsRequest := mt.GetFutureSaltsRequest{}
		if err := saltsRequest.Decode(in); err != nil {
			return err
		}

		return s.SendEternalSalt(req)

	case mt.RPCDropAnswerRequestTypeID:
		drop := mt.RPCDropAnswerRequest{}
		if err := drop.Decode(in); err != nil {
			return err
		}

		return s.SendResult(req, &mt.RPCAnswerDroppedRunning{})

	case proto.GZIPTypeID:
		var content proto.GZIP
		if err := content.Decode(in); err != nil {
			return errors.Wrap(err, "gzip")
		}
		req.Buf = &bin.Buffer{Buf: content.Data}

	case proto.MessageContainerTypeID:
		var container proto.MessageContainer
		if err := container.Decode(in); err != nil {
			return errors.Wrap(err, "container")
		}

		var err error
		for _, msg := range container.Messages {
			err = multierr.Append(err, s.handle(&Request{
				DC:         req.DC,
				Session:    req.Session,
				MsgID:      msg.ID,
				Buf:        &bin.Buffer{Buf: msg.Body},
				RequestCtx: req.RequestCtx,
			}))
		}
		return err


	//
	//case td.InvokeWithoutUpdatesRequestTypeID:
	//	log.Println("td.InvokeWithoutUpdatesRequestTypeID")

	}

	log.Println("若是没有对应的消息ID，则进行handler处理")
	if err := s.handler.OnMessage(s, req); err != nil {
		log.Println("hander OnMessage is err")
		var rpcErr *tgerr.Error
		if errors.As(err, &rpcErr) {
			return s.SendErr(req, rpcErr)
		}
		return err
	}

	return nil
}
