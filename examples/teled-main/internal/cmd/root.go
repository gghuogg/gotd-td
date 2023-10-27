// Package cmd implements commands of teled binary.
package cmd

import (
	"log"
	"net"
	"os"

	"github.com/go-faster/errors"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/gotd/td/tg"

	"github.com/gotd/td/tgtest"
	"github.com/gotd/td/transport"

	"github.com/gotd/td/examples/teled-main/internal/key"
)

func newRoot(a *application) *cobra.Command {
	log.Println("定义rootCmd 命令")
	var rootCmd = &cobra.Command{
		Use:   "teled",
		Short: "Telegram Server in Go",
		Long: `Telegram Server in Go, including auxiliary commands.
				Not affiliated with official Telegram.
				Apache License 2.0, The GoTD Authors.
				Based on https://gotd.dev Telegram protocol implementation.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			log.Println("读取私钥文件路径中的私钥")
			privateKeyEncoded, err := os.ReadFile(a.PrivateKeyPath)
			if err != nil {
				return errors.Wrap(err, "failed to read private key")
			}
			log.Println("进行私钥编码")
			k, err := key.ParsePrivateKey(privateKeyEncoded)
			if err != nil {
				return errors.Wrap(err, "failed to parse private key")
			}

			log.Println("监听地址和端口设置")
			var lc net.ListenConfig
			ln, err := lc.Listen(ctx, "tcp", a.Addr())
			if err != nil {
				return errors.Wrap(err, "failed to listen")
			}
			opt := tgtest.ServerOptions{
				DC:     1,
				Logger: a.lg,
			}
			a.lg.Info("Listening",
				zap.String("addr", a.Addr()),
				zap.Int("dc", opt.DC),
			)
			log.Println("开始运行NewServer")
			srv := tgtest.NewServer(tgtest.NewPrivateKey(k), tgtest.UnpackInvoke(a), opt)
			return srv.Serve(ctx, transport.Listen(transport.ObfuscatedListener(ln)))
		},
	}

	{
		log.Println("命令行解析，参数设置")
		f := rootCmd.Flags()
		f.StringVar(&a.Host, "host", "localhost", "Hostname of the server")
		f.IntVar(&a.Port, "port", 10443, "Port of the server")
		f.StringVar(&a.PrivateKeyPath, "key", ".\\_testdata\\test.key.pem", "Path to PEM-encoded private key")

		markFlagsRequired(f, "key")
	}

	log.Println("整合整体的命令行")
	rootCmd.AddCommand(
		newKeys(a),
	)
	log.Println("整体命令行整理完毕，后面进行整体命令执行")
	return rootCmd
}

// Execute executes root command.
func Execute() {
	log.Println("server 开始执行。。。。。。。")
	lg, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	a := &application{
		lg: lg,
	}
	log.Println("日志设置")
	a.setDispatcher(tg.NewServerDispatcher(a.Fallback))
	log.Println("newRoot 开始执行")
	if err := newRoot(a).Execute(); err != nil {
		panic(err)
	}
}
