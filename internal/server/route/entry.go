package route

import (
	"github.com/Amiraxoba2/weitiu/internal/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Entry(ctx *gin.Context) {
	name := ctx.Query("name")
	var entry database.Entry
	database.Db.Last(&entry, "name = ?", name)
	ctx.HTML(http.StatusOK, "entry.go.tmpl", gin.H{
		"ENTRY_NAME": entry.Name,
		"CONTENT":    entry.Content,
	})
}
