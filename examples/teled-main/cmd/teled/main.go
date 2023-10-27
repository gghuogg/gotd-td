// Binary teled implements Telegram Server in Go.
package main

import (
	"github.com/gotd/td/examples/teled-main/internal/cmd"
	"log"
)

func main() {
	log.SetFlags(log.Ldate|log.Ltime|log.Llongfile)
	log.Println("server starting ")
	cmd.Execute()
}
