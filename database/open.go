package database

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/syndtr/goleveldb/leveldb"
)

var DB *leveldb.DB

func init() {
	var err error
	DB, err = open()
	if err != nil {
		log.Fatal(err)
	}
}

func open() (DB *leveldb.DB, err error) {
	DB, err = leveldb.OpenFile("database-dir", nil)
	if err != nil {
		return
	}
	return
}

func Migration(db *leveldb.DB) (err error) {

	var users []User

	usersJSON, err := ioutil.ReadFile("migration.json")
	if err != nil {
		return
	}

	err = json.Unmarshal(usersJSON, &users)
	if err != nil {
		return
	}

	for i := range users {
		var userJSON []byte

		key := []byte(users[i].Username)
		users[i].Username = ""

		userJSON, err = json.Marshal(users[i])
		if err != nil {
			return
		}

		err = db.Put(key, userJSON, nil)
		if err != nil {
			return
		}

	}

	return
}
