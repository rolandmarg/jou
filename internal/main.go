package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	DB, err := sql.Open("sqlite3", "./journal.db")
	if err != nil {
		panic(err)
	}

	_, err = DB.Exec(Schema)
	if err != nil {
		panic(err)
	}
}

/*
TODO list

add import export data
add redis
add sessions in redis
add benchmark tests
add logging
add optional data encryption in db
add pincode on journal
add backups periodically
add geolocation
add audio to autoplay during specific entry
add beautiful elements/ornaments for each entry customization

Actions
	journal
		create
		edit
		delete
		get with entries
	entry
		create
		edit
		delete
		get with tags
*/
