package tgtest

import (
	"context"
	"log"

	"github.com/go-faster/errors"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/internal/exchange"
	"github.com/gotd/td/internal/proto/codec"
	"github.com/gotd/td/transport"
)

func (s *Server) read(ctx context.Context, conn transport.Conn, b *bin.Buffer) error {
	b.Reset()
	log.Println("开始读取server的连接信息")
	ctx, cancel := context.WithTimeout(ctx, s.readTimeout)
	defer cancel()
	if err := conn.Recv(ctx, b); err != nil {
		log.Println("连接接收数据错误，",err)
		return err
	}

	return nil
}

func (s *Server) sendProtoError(ctx context.Context, conn transport.Conn, e int32) error {
	var buf bin.Buffer
	buf.PutInt32(-e)

	ctx, cancel := context.WithTimeout(ctx, s.writeTimeout)
	defer cancel()

	if err := conn.Send(ctx, &buf); err != nil {
		return errors.Wrap(err, "send")
	}
	return nil
}

func (s *Server) serveConn(ctx context.Context, conn transport.Conn) error {
	log.Println("用户开始连接")
	// s.log.Debug("User connected")
	defer func() {
		_ = conn.Close()
		log.Println("用户断开连接")
		// s.log.Debug("User disconnected")
	}()

	b := new(bin.Buffer)
	for {
		log.Println("开始读取连接信息")
		if err := s.read(ctx, conn, b); err != nil {
			return errors.Wrap(err, "read")
		}

		log.Println("读取authKeyID")
		var authKeyID [8]byte
		if err := b.PeekN(authKeyID[:], len(authKeyID)); err != nil {
			return errors.Wrap(err, "peek id")
		}

		log.Println("通过authKeyID来获取用户的session")
		// TODO(tdakkota): dispatch by type ID instead?
		if _, ok := s.users.getSession(authKeyID); ok {
			log.Println("处理rpc连接")
			if err := s.rpcHandle(ctx, conn, b); err != nil {
				return errors.Wrap(err, "handle")
			}
			continue
		}

		log.Println("如果authKeyID为空或者是0，发送CodeAuthKeyNotFound")
		// If authKeyID not found and is not zero, so send protocol error.
		if authKeyID != [8]byte{} {
			if err := s.sendProtoError(ctx, conn, codec.CodeAuthKeyNotFound); err != nil {
				return errors.Wrap(err, "send AuthKeyNotFound")
			}
			continue
		}

		log.Println("开始key交换")
		s.log.Debug("Starting key exchange")
		c := newBufferedConn(conn)
		c.Push(b)

		log.Println("开始进行key交换")
		key, err := s.exchange(ctx, exchangeConn{Conn: c})
		if err != nil {
			log.Println("server交换出现错误")
			var exchangeErr *exchange.ServerExchangeError
			if errors.As(err, &exchangeErr) {
				code := exchangeErr.Code
				log.Println("发送交换错误码")
				if err := s.sendProtoError(ctx, c, code); err != nil {
					return errors.Wrapf(err, "send proto error %v", code)
				}
				return nil
			}
			log.Println("key交换失败")
			return errors.Wrap(err, "key exchange failed")
		}

		log.Println("添加用户session")
		s.users.addSession(key)
	}
}
