package utils

import (
	"fmt"

	"github.com/cfabrica46/leveldb/login/database"
)

func AskData() (username, password string) {
	fmt.Print("Username: ")
	fmt.Scan(&username)
	fmt.Print("Password: ")
	fmt.Scan(&password)
	return
}

func PrintUser(user database.User) {

	fmt.Printf("Bienvenido %s %s\n", user.Role, user.Username)

}
