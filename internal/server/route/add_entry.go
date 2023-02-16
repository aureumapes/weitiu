package route

import (
	"github.com/Amiraxoba2/weitiu/internal/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddEntry(ctx *gin.Context) {
	name := ctx.PostForm("name")
	content := ctx.PostForm("content")
	enteredPwd := ctx.PostForm("admin")

	var settings database.AdminSettings
	database.Db.Last(&settings)

	if enteredPwd == settings.Password {
		database.Db.Create(&database.Entry{Name: name, Content: content})
		ctx.JSON(http.StatusCreated, gin.H{
			"message": "Your new entry has been successfully added to the wikis database!",
			"link":    "/entry?name=" + name,
		})
	} else {
		ctx.JSON(http.StatusForbidden, gin.H{
			"message": "You entered the wrong password!",
		})
	}
}
