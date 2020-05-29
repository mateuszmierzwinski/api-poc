package practice

import (
	"github.com/gin-gonic/gin"
	"newPocApi/endpoints"
)

type practiceDetails struct {
	Title 	string	`json:"title"`
	Description  string	`json:"description"`
	Icon 	string	`json:"icon"`
	Link 	string 	`json:"link"`
}

type ctrl struct {

}

func (m *ctrl) Bind(ctx *gin.Context) {
	ctx.JSON(200, m.getFocused())
	ctx.Done()
}

func (m *ctrl) getFocused() *practiceDetails {
	return &practiceDetails{
		Title:       "Talk about grammar #1",
		Description: "Talk about grammar improves your basic grammar skills. You will study basic phonology, morphology and syntax, often complemented by phonetics, semantics and pragmatics.",
		Icon:        "icon_globe",
		Link:        "https://math-stg.aida.pearson.com/api/v1/assets/videos/467",
	}
}

func NewBinder() endpoints.Binder {
	m := &ctrl{}
	return m
}