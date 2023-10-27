package main

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"time"

	"go.uber.org/zap"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/examples"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/telegram"
	"github.com/gotd/td/tg"
)

// prettyMiddleware pretty-prints request and response.
func prettyMiddleware() telegram.MiddlewareFunc {
	return func(next tg.Invoker) telegram.InvokeFunc {
		return func(ctx context.Context, input bin.Encoder, output bin.Decoder) error {
			fmt.Println("→", formatObject(input))
			start := time.Now()
			if err := next.Invoke(ctx, input, output); err != nil {
				fmt.Println("←", err)
				return err
			}

			fmt.Printf("← (%s) %s\n", time.Since(start).Round(time.Millisecond), formatObject(output))

			return nil
		}
	}
}

func formatObject(input interface{}) string {
	o, ok := input.(tdp.Object)
	if !ok {
		// Handle tg.*Box values.
		rv := reflect.Indirect(reflect.ValueOf(input))
		for i := 0; i < rv.NumField(); i++ {
			if v, ok := rv.Field(i).Interface().(tdp.Object); ok {
				return formatObject(v)
			}
		}

		return fmt.Sprintf("%T (not object)", input)
	}
	return tdp.Format(o)
}

func main() {
	os.Setenv("APP_ID", "29453132")
	os.Setenv("APP_HASH", "60e35bbbefb9d5cb6b395f1d04b1d588")
	examples.Run(func(ctx context.Context, log *zap.Logger) error {
		return telegram.BotFromEnvironment(ctx, telegram.Options{
			UpdateHandler: telegram.UpdateHandlerFunc(func(ctx context.Context, u tg.UpdatesClass) error {
				// Print all incoming updates.
				fmt.Println("u:", formatObject(u))
				return nil
			}),
			Middlewares: []telegram.Middleware{
				prettyMiddleware(),
			},
		}, nil, telegram.RunUntilCanceled)
	})
}
