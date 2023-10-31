package main

import (
	"context"
	"crypto/x509"
	"encoding/pem"
	"github.com/gotd/td/tg"
	"github.com/gotd/td/tgtest"
	"github.com/gotd/td/transport"
	"io/ioutil"
	"log"
	"net"
)


func main() {
	log.SetFlags(log.Ldate|log.Ltime|log.Llongfile)

	// 1、读取私钥文件，获取私钥字节
	privateKeyBytes, err := ioutil.ReadFile(".\\publickeyread\\rsa_private_pkcs1.pem")
	if err != nil {
	}
	// 2、对私钥文件进行编码，生成加密块对象
	block, _ := pem.Decode(privateKeyBytes)
	if block == nil {
	}
	// 3、解析DER编码的私钥，生成私钥对象
	k, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
	}

	log.Println("设置serverOptions")
	opt := tgtest.ServerOptions{
		DC:     2,
	}

	log.Println("定义NewServer")
	a := &application{}
	a.setDispatcher(tg.NewServerDispatcher(a.Fallback))

	srv := tgtest.NewServer(tgtest.NewPrivateKey(k), tgtest.UnpackInvoke(a), opt)
	ctx := context.Background()

	log.Println("定义监听地址")
	var lc net.ListenConfig
	ln, err := lc.Listen(ctx, "tcp","127.0.0.1:10443")

	log.Println("打印KEY",srv.Key().Fingerprint())


	log.Println("运行server")
	//srv.Serve(ctx,  transport.Listen(transport.ObfuscatedListener(ln)))
	srv.Serve(ctx,transport.Listen(ln))


}


