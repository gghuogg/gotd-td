package transport

import (
	"context"
	"log"
	"net"
	"sync"
	"time"

	"github.com/go-faster/errors"

	"github.com/gotd/td/bin"
)

// Conn is transport connection.
type Conn interface {
	Send(ctx context.Context, b *bin.Buffer) error
	Recv(ctx context.Context, b *bin.Buffer) error
	Close() error
}

var _ Conn = (*connection)(nil)

// connection is MTProto connection.
type connection struct {
	conn  net.Conn
	codec Codec

	readMux  sync.Mutex
	writeMux sync.Mutex
}

// Send sends message from buffer using MTProto connection.
func (c *connection) Send(ctx context.Context, b *bin.Buffer) error {
	// Serializing access to deadlines.
	c.writeMux.Lock()
	defer c.writeMux.Unlock()
	log.Println("设置写入消息日期")
	if err := c.conn.SetWriteDeadline(time.Time{}); err != nil {
		return errors.Wrap(err, "reset write deadline")
	}
	log.Println("设置写入消息的最后期限")
	if deadline, ok := ctx.Deadline(); ok {
		if err := c.conn.SetWriteDeadline(deadline); err != nil {
			return errors.Wrap(err, "set write deadline")
		}
	}
	log.Println("消息写入")
	//c.conn.Write([]byte("nihaoya,hahahahhahahhahahhah"))
	//bstring,_ := b.String()
	//log.Println(b.Len(),bstring)
	//c.conn.Write([]byte("helloworld"))
	if err := c.codec.Write(c.conn, b); err != nil {
		log.Println("消息写入失败",err)
		return errors.Wrap(err, "write")
	}

	return nil
}

// Recv reads message to buffer using MTProto connection.
func (c *connection) Recv(ctx context.Context, b *bin.Buffer) error {
	// Serializing access to deadlines.
	c.readMux.Lock()
	defer c.readMux.Unlock()

	log.Println("设置读取消息的日期")
	if err := c.conn.SetReadDeadline(time.Time{}); err != nil {
		return errors.Wrap(err, "reset read deadline")
	}
	log.Println("设置消息的最后期限")
	if deadline, ok := ctx.Deadline(); ok {
		if err := c.conn.SetReadDeadline(deadline); err != nil {
			return errors.Wrap(err, "set read deadline")
		}
	}
	log.Println("消息的读取")
	//paddingint,vstring,errerror := bin.DecodeString(b.Copy())
	//log.Println("paddingint,vstring,errerror",paddingint,vstring,errerror)
	log.Println(c.conn.LocalAddr().String())
	//c.conn.Read(b.Buf)

	if err := c.codec.Read(c.conn, b); err != nil {
		log.Println("消息读取失败",err)
		return errors.Wrap(err, "read")
	}

	return nil
}

// Close closes MTProto connection.
func (c *connection) Close() error {
	return c.conn.Close()
}
