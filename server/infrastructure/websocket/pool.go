package websocket

import (
	"fmt"
)

type Pool struct {
    Register   chan *Client
    Unregister chan *Client
    Clients    map[string]*Client
    Broadcast  chan Message
}

func NewPool() *Pool {
	return &Pool{
		Register: make(chan *Client),
		Unregister: make(chan *Client),
		Clients: make(map[string]*Client),
		Broadcast: make(chan Message),
	}
}

func (pool *Pool) Start(){
	for {
		select{
			case client := <- pool.Register:
				pool.Clients[client.UserID] = client
				fmt.Println("Size of Connection Pool: ", len(pool.Clients))
				break
			case client := <- pool.Unregister:
				delete(pool.Clients, client.UserID)
				fmt.Println("Size of Connection Pool: ", len(pool.Clients))
				break
			case message := <- pool.Broadcast:
				fmt.Println("Sending message to recipient client in Pool")
				recipient, ok := pool.Clients[message.RecipientID]
				if ok {
					if err := recipient.Conn.WriteJSON(message); err != nil {
						fmt.Println(err)
						return
					}
				}
				break

		}
	}
}