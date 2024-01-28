package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/villekuosmanen/physiology-sim/src/body"
	"github.com/villekuosmanen/physiology-sim/src/systems/metabolism"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	body := body.ConstructBody()
	body.SetMetabolicRate(metabolism.METMediumCardio)

	body.Run(10, true, false, sigs)
}
