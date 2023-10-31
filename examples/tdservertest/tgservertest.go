// Binary bg-run implements alternative to Run pattern.
// 二进制bg-run实现了run模式的替代方案。
package main

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"

	"go.uber.org/zap"

	"github.com/gotd/contrib/bg"

	"github.com/go-faster/errors"
	"github.com/gotd/td/examples"
	"github.com/gotd/td/telegram"
)

func tgservertest() {

	log.SetFlags(log.Ldate|log.Ltime|log.Llongfile)
	log.Println("client starting ")


	// Some users find explicit client.Run(ctx, f) pattern not very convenient.
	//
	// However, it is possible to implement wrapper and use classic "Connect"
	// pattern instead.
	//
	// The `contrib/bg` package is example implementation of such pattern.

	//一些用户找到显式客户端。运行（ctx，f）模式不是很方便。
	//但是，可以实现包装器并使用经典的“Connect” 模式。
	//“contrib/bg”包就是这种模式的示例实现。

	APP_ID := 17349
	APP_HASH := "344583e45741c457fe1862106095a5eb"
	fmt.Println(APP_ID,APP_HASH)
	os.Setenv("APP_ID", "17349")
	os.Setenv("APP_HASH", "344583e45741c457fe1862106095a5eb")

	/*
	const char *kPublicRSAKeys[] = { "\
	-----BEGIN RSA PUBLIC KEY-----\n\
	MIIBCgKCAQEAwX3W9oxq8FdwkMOA/HNeeDJOx8pF6fIw4MrtbfsYPhc5yTs0lg6H\n\
	3ia4SmUTnpbJ+NhGxQX2nhDR6911SETBIq1Iz0OUuOPyDVhVhyZzF+vzBukUE75s\n\
	lPy3UzAzAyruOT5iCoX7mCDhRHWh9Rhm2HHVC02LAKM7RYx6u3WyjI4LV1a1Fpla\n\
	D0wTh1ElJQcHkRsCVJzfD2fCSrkO/F2QNuOdEe9uoyIFGZjTX82GJNd/H6DXKB4X\n\
	Q158yJifZBwZohH9q/7IYrTEGUreQyMjeL1CiROwYKSEmHOaDnLWqq4u8KSmjhqN\n\
	eQnhbTFuwG958vSX1zJOqkDkXlue/1XeMQIDAQAB\n\
	-----END RSA PUBLIC KEY-----" };
	*/

//	var publicKeyStrs = `-----BEGIN RSA PUBLIC KEY-----
//MIGJAoGBANqUsY2pv+Ljd27T++7TIVKNydih1wemHujb5TydbUD1mttS10S+oCh8
//Fc/WNK1m6fVLtJVbTLcb1l8Xw+DHQT8uVbYAbDDGhbLfTxyQ0OOaAI44GleAy0Wi
//llDCZcXrB+ZQEqvfXmanKBN++yOkfN+lOSL1RSXfEJlz3rjmWz9rAgMBAAE=
//-----END RSA PUBLIC KEY-----
//`
	//// 解析pem格式的公钥数据
	//blockPub, _ := pem.Decode([]byte(publicKeyStrs))
	//publicKey,_ := x509.ParsePKCS1PublicKey(blockPub.Bytes)
	//fmt.Println(publicKey)

	//privateKeyEncoded, err := os.ReadFile(".\\publickeyread\\rsa_public.pem")
	//privateKeyEncoded, err := os.ReadFile(".\\publickeyread\\rsa_public_pkcs1.pem")
	privateKeyEncoded, err := os.ReadFile(".\\publickeyread\\rsapublic.pem")
	if err != nil {
		fmt.Println(err)
	}
	log.Println("进行私钥编码")
	publicKey, err := ParsePrivateKey(privateKeyEncoded)
	log.Println("打印publickey的大小",publicKey.Size())



	var tgpublickeys []telegram.PublicKey
	var tgpublickey telegram.PublicKey

	tgpublickey.RSA = publicKey
	tgpublickeys = append(tgpublickeys,tgpublickey)
	fmt.Println("打印tgpublickeys",tgpublickeys[0].Fingerprint())

	examples.Run(func(ctx context.Context, log *zap.Logger) error {
		client, err := telegram.ClientFromEnvironment(telegram.Options{
			Logger: log,
			PublicKeys: tgpublickeys,
		})
		if err != nil {
			return err
		}

		// bg.Connect will call Run in background.
		// Call stop() to disconnect and release resources.
		//bg.Connect将在后台调用Run。
		//调用stop（）断开连接并释放资源。
		fmt.Println("bg.Connect")
		stop, err := bg.Connect(client)

		if err != nil {
			return err
		}
		defer func() { _ = stop() }()

		// Now you can use client.
		//现在您可以使用客户端了。
		if _, err := client.Auth().Status(ctx); err != nil {
			return err
		}

		fmt.Println("ping")
		if pingerr := client.Ping(ctx); pingerr != nil{
			fmt.Println("pingerr",pingerr)
		}

		//fmt.Println("phoneGetCallConfig")
		//datejson,dateerr := client.API().PhoneGetCallConfig(ctx)
		//if dateerr != nil {
		//	return dateerr
		//}
		//fmt.Println(datejson.String())

		fmt.Println("AccountCheckUsername")
		bool,boolerr := client.API().AccountCheckUsername(ctx,"gghhxx")
		if boolerr != nil {
			fmt.Println( "boolerr",boolerr)
		}
		fmt.Println(bool)


		return nil
	})
}


// ParsePrivateKey parses PEM encoded private key.
func ParsePrivateKey(data []byte) (*rsa.PublicKey, error) {
	block,_ := pem.Decode(data)
	if block.Type != "RSA PUBLIC KEY" {
		return nil, errors.Errorf("unsupported key type: %q", block.Type)
	}

	k, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, errors.Wrap(err, "PKCS1")
	}

	return k, nil
}
