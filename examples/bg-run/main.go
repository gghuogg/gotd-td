// Binary bg-run implements alternative to Run pattern.
// 二进制bg-run实现了run模式的替代方案。
package main

import (
	"context"
	"fmt"
	"os"

	"go.uber.org/zap"

	"github.com/gotd/contrib/bg"

	"github.com/gotd/td/examples"
	"github.com/gotd/td/telegram"
)

func main() {
	// Some users find explicit client.Run(ctx, f) pattern not very convenient.
	//
	// However, it is possible to implement wrapper and use classic "Connect"
	// pattern instead.
	//
	// The `contrib/bg` package is example implementation of such pattern.

	//一些用户找到显式客户端。运行（ctx，f）模式不是很方便。
	//但是，可以实现包装器并使用经典的“Connect” 模式。
	//“contrib/bg”包就是这种模式的示例实现。

	APP_ID := 29453132
	APP_HASH := "60e35bbbefb9d5cb6b395f1d04b1d588"
	fmt.Println(APP_ID,APP_HASH)
	os.Setenv("APP_ID", "29453132")
	os.Setenv("APP_HASH", "60e35bbbefb9d5cb6b395f1d04b1d588")

	examples.Run(func(ctx context.Context, log *zap.Logger) error {
		client, err := telegram.ClientFromEnvironment(telegram.Options{
			Logger: log,
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
