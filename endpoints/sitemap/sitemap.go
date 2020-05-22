package sitemap

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"newPocApi/endpoints"
)

type sitemap map[string] siteDetails

type siteDetails struct {
	Title 	string	`json:"title"`
	Link 	string 	`json:"link"`
}

type ctrl struct {

}

func (m *ctrl) Bind(ctx *gin.Context) {
	ctx.JSON(200, m.getSitemap(fmt.Sprintf("https://%s", ctx.Request.Host)))
	ctx.Done()
}

func (m *ctrl) getSitemap(currentUrl string) *sitemap {
	return &sitemap{
		"about_page": {
			Title: "About information",
			Link: fmt.Sprintf("%s%s", currentUrl, "/static/about.html"),
		},
	}
}

func NewBinder() endpoints.Binder {
	m := &ctrl{}
	return m
}