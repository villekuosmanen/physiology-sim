package ws

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/villekuosmanen/physiology-sim/src/simulation"
)

type ConnectionManager struct {
	clients      map[*websocket.Conn]bool
	bloodStats   chan simulation.BloodStatistics
	generalStats chan simulation.GeneralStatistics
	mutex        sync.RWMutex
	upgrader     websocket.Upgrader
}

func NewConnectionManager() *ConnectionManager {
	return &ConnectionManager{
		clients:      make(map[*websocket.Conn]bool),
		bloodStats:   make(chan simulation.BloodStatistics),
		generalStats: make(chan simulation.GeneralStatistics),
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true // For development purposes only
			},
		},
	}
}

func (m *ConnectionManager) HandleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := m.upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		ws.Close()

		m.mutex.Lock()
		delete(m.clients, ws)
		m.mutex.Unlock()
	}()

	m.mutex.Lock()
	m.clients[ws] = true
	m.mutex.Unlock()

	for {
		// Read (and ignore) messages from the WebSocket
		// TODO process these messages
		_, _, err := ws.ReadMessage()
		if err != nil {
			// error, return
			break
		}
	}
}

func (m *ConnectionManager) HandleMessages() {
	for {
		select {
		case notification := <-m.bloodStats:
			msg, err := json.Marshal(notification)
			if err != nil {
				fmt.Printf("error marshalling: %v\n", err)
				return
			}
			m.broadcast(msg)

		case notification := <-m.generalStats:
			msg, err := json.Marshal(notification)
			if err != nil {
				fmt.Printf("error marshalling: %v\n", err)
				return
			}
			m.broadcast(msg)
		}

	}

}

func (m *ConnectionManager) BroadcastBloodStats(s simulation.BloodStatistics) {
	m.bloodStats <- s
}

func (m *ConnectionManager) BroadcastGeneralStats(s simulation.GeneralStatistics) {
	m.generalStats <- s
}

func (m *ConnectionManager) broadcast(msg []byte) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	for client := range m.clients {
		err := client.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			// close connection, it will be automatically cleaned up
			client.Close()
		}
	}
}
