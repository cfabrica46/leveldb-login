package database

import (
	"encoding/json"
	"log"

	"github.com/syndtr/goleveldb/leveldb"
)

type User struct {
	Username, Password, Role string
}

func GetUser(username, password string) (user *User, err error) {

	iter := DB.NewIterator(nil, nil)

	for iter.Next() {

		if string(iter.Key()) == username {
			var userBeta User

			err = json.Unmarshal(iter.Value(), &userBeta)
			if err != nil {
				return
			}

			if userBeta.Password != password {
				return
			}

			userBeta.Username = string(iter.Key())
			user = &userBeta
			return
		}
	}

	return
}

func RegisterUser(username, password, role string) (err error) {

	user := User{Password: password, Role: role}

	userJSON, err := json.Marshal(user)
	if err != nil {
		return
	}

	err = DB.Put([]byte(username), userJSON, nil)
	if err != nil {
		return
	}
	return
}

func CheckIfUserAlreadyExist(username string) (check bool, err error) {

	_, err = DB.Get([]byte(username), nil)

	if err != nil {
		if err == leveldb.ErrNotFound {
			err = nil
			return
		}
		log.Fatal(err)
	}

	check = true
	return
}

func GetUsers(db *leveldb.DB) (users []User, err error) {

	iter := db.NewIterator(nil, nil)

	for iter.Next() {
		var user User
		key := iter.Key()
		value := iter.Value()

		err = json.Unmarshal(value, &user)

		if err != nil {
			return
		}

		user.Username = string(key)

		users = append(users, user)
	}

	return
}
