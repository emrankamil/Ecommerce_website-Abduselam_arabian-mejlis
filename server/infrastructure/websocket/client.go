package websocket

import (
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct{
	ID		string
	Conn	*websocket.Conn
	Pool	*Pool
	mu 		sync.Mutex
	UserID   string
	IsAdmin  bool
}

type Message struct{
	ChatID       string 	`json:"chat_id" bson:"chat_id"`
	Type		 int 		`json:"type" bson:"type"`
	Body		 string		`json:"body" bson:"body"`
	SenderID     string 	`json:"sender_id" bson:"sender_id"`
    RecipientID  string 	`json:"recipient_id" bson:"redipient_id"`
    Timestamp    time.Time  `json:"timestamp" bson:"timestamp"`
}

func (c *Client) Read(){
	defer func(){
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil{
			log.Println(err)
			return
		}
		message := Message{Type: messageType, Body: string(p), SenderID: c.UserID, }
		c.Pool.Broadcast <- message
		log.Printf("Message Received: %+v\n", message)
	}

}