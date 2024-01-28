package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/villekuosmanen/physiology-sim/src/body"
	"github.com/villekuosmanen/physiology-sim/src/systems/metabolism"
	"github.com/villekuosmanen/physiology-sim/src/ws"
)

func main() {
	connManager := ws.NewConnectionManager()

	// Start WebSocket server
	http.HandleFunc("/ws", connManager.HandleConnections)
	go connManager.HandleMessages()

	go http.ListenAndServe(":8080", nil)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	body := body.ConstructBody(connManager)
	body.SetMetabolicRate(metabolism.METHeavyCardio)

	// TODO
	// Basically make the messages below broadcast over the websockets.
	// In addition make the simulation controllable via mebsocket messages.
	body.Run(10, false, false, sigs)
}
