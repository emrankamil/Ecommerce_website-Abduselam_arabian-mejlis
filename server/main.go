package main

import (
	"abduselam-arabianmejlis/bootstrap"
	"abduselam-arabianmejlis/delivery/route"
	"abduselam-arabianmejlis/infrastructure/websocket"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	app := bootstrap.App()
	env := app.Env

	db := app.Mongo.Database(env.DBName)

	if app.Redis == nil {
		panic("Redis is not connected")
	}
	redisClient := app.Redis
	defer app.Close()

	// start the client manager routine
	clientManager := bootstrap.NewClientManager()
	go clientManager.Start()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	// Initialize Gin
	router := gin.Default()

	// Setup HTTP routes
	route.Setup(env, timeout, db, router, redisClient, clientManager)

	// Setup WebSocket routes using Gin
	// setupRoutes(router)

	router.Run(env.ServerAddress) 
}

func serveWS(pool *websocket.Pool, c *gin.Context) {
	fmt.Println("WebSocket endpoint reached")

	conn, err := websocket.Upgrade(c.Writer, c.Request)
	if err != nil {
		fmt.Fprintf(c.Writer, "%+v\n", err)
		return
	}

	client := &websocket.Client{
		Conn:    conn,
		Pool:    pool,
		UserID:  "12345", 
		IsAdmin: false, 
	}
	pool.Register <- client
	client.Read()
}

func setupRoutes(router *gin.Engine) {
	pool := websocket.NewPool()
	go pool.Start()

	// WebSocket route
	router.GET("/ws", func(c *gin.Context) {
		serveWS(pool, c)
	})

	// Other HTTP routes can be added here as well
}
