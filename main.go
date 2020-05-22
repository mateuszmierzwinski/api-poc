package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"newPocApi/endpoints/mission"
	"newPocApi/endpoints/notifications"
	"newPocApi/endpoints/practice"
	"newPocApi/endpoints/sitemap"
	"newPocApi/middleware"
	"os"
	"path/filepath"
)

var version = "0"

func main() {
	log.Println("Version", version)

	// new gin
	newGin := gin.New()

	newGin.Use(middleware.New(version).Middleware)

	newGin.GET("/info", func(ctx *gin.Context){
		ctx.JSON(200, struct {
			Version 	string	`json:"version"`
		}{
			Version: version,
		})
	})

	mCtrl := mission.NewBinder()
	newGin.GET("/mission/:learnerId/latest", mCtrl.Bind)

	pCtrl := practice.NewBinder()
	newGin.GET("/practice/:learnerId/focused", pCtrl.Bind)

	sCtrl := sitemap.NewBinder()
	newGin.GET("/sitemap", sCtrl.Bind)

	nCtrl := notifications.NewBinder()
	newGin.GET("/notification", nCtrl.Bind)


	newGin.Static("/static/", getLocalPath("static/"))

	newGin.RunTLS(":8443", getLocalPath("domain.crt"), getLocalPath("domain.key"))
}

func getLocalPath(filename string) string {
	abs,_ := filepath.Abs(os.Args[0])
	appPath := filepath.Dir(abs)
	return filepath.Join(appPath, filename)
}
