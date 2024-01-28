package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/villekuosmanen/physiology-sim/src/manager"
	"github.com/villekuosmanen/physiology-sim/src/ws"
)

func main() {
	simManager := manager.BodySimManager{}
	connManager := ws.NewConnectionManager(simManager)

	// Start WebSocket server
	http.HandleFunc("/ws", connManager.HandleConnections)
	go connManager.HandleMessages()

	go http.ListenAndServe(":7766", nil)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// run until completition
	<-sigs
}
