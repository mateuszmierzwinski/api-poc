package notifications

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"newPocApi/endpoints"
	"time"
)

type Update struct {
	Icon 	string	`json:"icon"`
	Description 	string	`json:"description"`
	Title 	string	`json:"title"`
	Link 	string 	`json:"link"`
}

type ctrl struct {

}

func (m *ctrl) BindStream(ctx *gin.Context) {
	ctx.Stream(func(w io.Writer) bool {
		log.Println("Stream opened. Notification will be sent in 10 seconds")
		time.Sleep(10 * time.Second)
		json.NewEncoder(w).Encode(m.getNotification(fmt.Sprintf("https://%s", ctx.Request.Host)))
		log.Println("Notification sent. To get new one reconnect to new stream (restart client app)")
		return true
	})

	ctx.Done()
}

func (m *ctrl) Bind(ctx *gin.Context) {
	m.BindSocket(ctx)
}

func (m *ctrl) BindSocket(ctx *gin.Context) {
	conn, err := wsupgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}

	log.Println("WebSocket opened. Notification will be sent in 10 seconds")
	time.Sleep(10 * time.Second)
	data,_ := json.Marshal(m.getNotification(fmt.Sprintf("https://%s", ctx.Request.Host)))
	conn.WriteMessage(websocket.TextMessage ,data)
	log.Println("Notification sent. To get new one reconnect to new websocket (restart client app)")
}

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (m *ctrl) getNotification(currentUrl string) *Update {
	return &Update{
		Icon: fmt.Sprintf("%s%s", currentUrl, "/static/images/unsplash1.jpg"),
		Description: "Based on your experience we found some interesting topics for you",
		Title: "Notifications",
		Link: "https://math-stg.aida.pearson.com/api/v1/assets/videos/467",
	}
}

func NewBinder() endpoints.Binder {
	m := &ctrl{}
	return m
}