package main

import (
	"fmt"
	"net/http"
	"time"

	"fitus-chat-service/config"
	"fitus-chat-service/internal/model"
	"fitus-chat-service/internal/repository"
	"fitus-chat-service/internal/service"

	"github.com/gorilla/websocket"
	"gorm.io/gorm"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}


func setupDatabase() *gorm.DB {
	db := config.SetupDatabase();
	db.AutoMigrate(&model.Message{})
	return db
}

func handleConnections(ms *service.MessageService, w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Upgrade error:", err)
		return
	}
	service.RegisterClient(ws)
	defer service.UnregisterClient(ws)
	defer ws.Close()

	for {
		var msg model.Message
		if err := ws.ReadJSON(&msg); err != nil {
			fmt.Println("Read error:", err)
			break
		}
		if err := ms.CreateMessage(msg.UserId, msg.Content); err != nil {
			fmt.Println("Create message error:", err)
		}
	}
}

func main() {
	config.LoadEnv()
	db := setupDatabase()

	messageRepo := repository.NewMessageRepository(db)
	messageService := service.NewMessageService(messageRepo)

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handleConnections(messageService, w, r)
	})

	srv := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Println("WebSocket server started on :8080")
	if err := srv.ListenAndServe(); err != nil {
		fmt.Println("Server error:", err)
	}
}
