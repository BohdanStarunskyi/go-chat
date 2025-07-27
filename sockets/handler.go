package sockets

import (
	"chat/utils"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || len(authHeader) < 8 || authHeader[:7] != "Bearer " {
		http.Error(w, "missing or invalid Authorization header", http.StatusUnauthorized)
		return
	}
	tokenStr := authHeader[7:]

	userID, err := utils.ValidateJwt(tokenStr)
	if err != nil {
		http.Error(w, "invalid token", http.StatusUnauthorized)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	client := NewClient(conn, userID)
	client.UserID = userID
	go client.Read()
	go client.Write()

	HubInstance.Register <- client
}
