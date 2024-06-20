package server

import (
	"github.com/Amiraxoba2/weitiu/internal/server/route"
	"github.com/gin-gonic/gin"
	"net/http"
)

func StartServer() {
	r := gin.Default()
	r.LoadHTMLGlob("~/.waitiu/resource/*.go.tmpl")
	r.GET("/", route.Index)
	r.GET("/entry", route.Entry)
	r.GET("/add-entry", func(ctx *gin.Context) { ctx.HTML(http.StatusOK, "add-entry.go.tmpl", nil) })
	r.POST("/add-entry", route.AddEntry)
	r.GET("/tailwind", func(ctx *gin.Context) { ctx.File("./resource/output.css") })
	r.Run(":8080")
}
