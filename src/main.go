package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/villekuosmanen/physiology-sim/src/body"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	body := body.ConstructBody()
	body.Run(10, true, sigs)
}
