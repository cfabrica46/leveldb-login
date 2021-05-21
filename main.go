package main

import (
	"fmt"
	"log"

	"github.com/cfabrica46/leveldb-login/database"
	"github.com/cfabrica46/leveldb-login/utils"
)

func main() {
	var eleccion int

	err := database.Migration(database.DB)
	if err != nil {
		log.Fatal(err)
	}

	//users, err := database.GetUsers(database.DB)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(users)

	fmt.Println("Welcome!:")
	fmt.Println("1.Sign In")
	fmt.Println("2.Sign Up")
	fmt.Print("> ")
	fmt.Scan(&eleccion)

	fmt.Println()

	switch eleccion {
	case 1:

		user, err := database.GetUser(utils.AskData())
		if err != nil {
			log.Fatal(err)
		}

		if user == nil {
			fmt.Println("Username Or Password Incorrect")
			return
		}

		utils.PrintUser(*user)

	case 2:

		username, password := utils.AskData()

		check, err := database.CheckIfUserAlreadyExist(username)
		if err != nil {
			log.Fatal(err)
		}
		if check {
			fmt.Println("Username already taken")
			return
		}

		err = database.RegisterUser(username, password, "member")
		if err != nil {
			log.Fatal(err)
		}

		user, err := database.GetUser(username, password)
		if err != nil {
			log.Fatal(err)
		}

		utils.PrintUser(*user)

	default:
		fmt.Println("Invalid Option")
	}

}
