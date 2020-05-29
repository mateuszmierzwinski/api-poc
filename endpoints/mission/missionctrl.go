package mission

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"newPocApi/endpoints"
)

type UpdatesMDL struct {
	Updates []Update `json:"updates"`
}

type Update struct {
	Image 	string	`json:"image"`
	Status 	string	`json:"status"`
	Title 	string	`json:"title"`
	Link 	string 	`json:"link"`
}

type ctrl struct {

}

func (m *ctrl) Bind(ctx *gin.Context) {
	ctx.JSON(200, m.getUptates(fmt.Sprintf("https://%s", ctx.Request.Host)))
	if pusher := ctx.Writer.Pusher(); pusher != nil {
		log.Println("Pushing content")
		pusher.Push("/static/images/unsplash1.jpg", &http.PushOptions{
			Method: "GET",
		})
		pusher.Push("/static/images/unsplash2.jpg", &http.PushOptions{
			Method: "GET",
		})
	}

	ctx.Done()
}

func (m *ctrl) getUptates(currentUrl string) *UpdatesMDL {
	return &UpdatesMDL{
		Updates: []Update{
			{
				Image: fmt.Sprintf("%s%s", currentUrl, "/static/images/unsplash1.jpg"),
				Status: "green",
				Title: "Talk about grammar #1",
				Link: "https://math-stg.aida.pearson.com/api/v1/assets/videos/467",
			},{
				Image: fmt.Sprintf("%s%s", currentUrl, "/static/images/unsplash2.jpg"),
				Status: "yellow",
				Title: "Effective presentation #1",
				Link: "https://math-stg.aida.pearson.com/api/v1/assets/videos/467",
			}, {
				Image: fmt.Sprintf("%s%s", currentUrl, "/static/images/unsplash1.jpg"),
				Status: "red",
				Title: "Talk about grammar #2",
				Link: "https://math-stg.aida.pearson.com/api/v1/assets/videos/467",
			},{
				Image: fmt.Sprintf("%s%s", currentUrl, "/static/images/unsplash2.jpg"),
				Status: "yellow",
				Title: "Effective presentation #2",
				Link: "https://math-stg.aida.pearson.com/api/v1/assets/videos/467",
			},
		},
	}
}

func NewBinder() endpoints.Binder {
	m := &ctrl{}
	return m
}