package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"github.com/go-faster/errors"
	"golang.org/x/sync/errgroup"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/telegram/dcs"
	"github.com/gotd/td/transport"
)

func main()  {
	ctx := context.Background()

	var handler http.Handler
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
	}))
	defer srv.Close()

	fmt.Println(srv.Listener.Addr().String())
	listener, h := transport.WebsocketListener(srv.Listener.Addr())
	handler = h
	list := dcs.List{
		Domains: map[int]string{
			2: srv.URL,
		},
	}

	server := transport.Listen(listener)
	defer server.Close()
	done := make(chan struct{})

	grp, ctx := errgroup.WithContext(ctx)
	grp.Go(func() error {
		defer close(done)

		conn, err := server.Accept()
		if err != nil {
			return errors.Wrap(err, "accept")
		}

		var b bin.Buffer
		fmt.Println("接收消息")
		if err := conn.Recv(ctx, &b); err != nil {
			return errors.Wrap(err, "recv")
		}

		fmt.Println("发送消息")
		if err := conn.Send(ctx, &b); err != nil {
			return errors.Wrap(err, "send")
		}

		return nil
	})

	rs := dcs.Websocket(dcs.WebsocketOptions{})
	conn, err := rs.Primary(ctx, 2, list)
	if err != nil {
		fmt.Println(err)
	}

	data, err := io.ReadAll(io.LimitReader(rand.Reader, 1024))
	fmt.Println("发送的消息",string(data))
	conn.Send(ctx, &bin.Buffer{Buf: data})


	fmt.Println("消息进行接收")
	var b bin.Buffer
	conn.Recv(ctx, &b)

	grp.Wait()
}
