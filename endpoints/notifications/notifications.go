package notifications

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
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

func (m *ctrl) Bind(ctx *gin.Context) {
	ctx.Stream(func(w io.Writer) bool {
		log.Println("Stream opened. Notification will be sent in 10 seconds")
		time.Sleep(10 * time.Second)
		json.NewEncoder(w).Encode(m.getNotification(fmt.Sprintf("https://%s", ctx.Request.Host)))
		log.Println("Notification sent. To get new one reconnect to new stream (restart client app)")
		return true
	})

	ctx.Done()
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