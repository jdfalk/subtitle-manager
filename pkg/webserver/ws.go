// file: pkg/webserver/ws.go
package webserver

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/jdfalk/subtitle-manager/pkg/tasks"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

// tasksWebSocketHandler streams task updates over WebSocket.
func tasksWebSocketHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			http.Error(w, "Failed to upgrade to WebSocket: "+err.Error(), http.StatusInternalServerError)
			log.Printf("WebSocket upgrade failed: %v", err)
			return
		}
		ch := tasks.Subscribe()
		defer tasks.Unsubscribe(ch)
		for {
			t, ok := <-ch
			if !ok {
				break
			}
			if err := c.WriteJSON(t); err != nil {
				break
			}
		}
		_ = c.Close()
	})
}
