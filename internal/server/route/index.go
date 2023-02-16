package route

import (
	"github.com/Amiraxoba2/weitiu/internal/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(ctx *gin.Context) {
	var entries []database.Entry
	database.Db.Find(&entries)

	ctx.HTML(http.StatusOK, "index.go.tmpl", gin.H{
		"ENTRIES": entries,
	})
}
