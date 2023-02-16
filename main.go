package main

import (
	"bufio"
	"github.com/Amiraxoba2/weitiu/internal/database"
	"github.com/Amiraxoba2/weitiu/internal/server"
	"os"
	"strings"
)

func main() {
	database.Db.AutoMigrate(database.Entry{})
	database.Db.AutoMigrate(database.AdminSettings{})
	println("What do you want to do?")
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	switch strings.Split(input, "\n")[0] {
	case "cp":
		var settings database.AdminSettings
		database.Db.Last(&settings)
		println("Current Password: " + settings.Password)
		println("New Password:")
		newPwdRead, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		newPwd := strings.Split(newPwdRead, "\n")[0]
		database.Db.Create(&database.AdminSettings{Password: newPwd})
		break
	case "s":
		server.StartServer()
		break
	}
}
