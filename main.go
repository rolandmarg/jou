package main

import (
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// db, err := sql.Open("sqlite3", "./foo.db")
	// e(err)

	// _, err = db.Exec(`CREATE TABLE IF NOT EXISTS 'userinfo' (
	// 	'uid' INTEGER PRIMARY KEY AUTOINCREMENT,
	// 	'username' VARCHAR(64) NULL,
	// 	'departname' VARCHAR(64) NULL,
	// 	'created' DATE NULL
	// 	)
	// `)
	// e(err)

	// stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?, ?, ?)")
	// e(err)

	// res, err := stmt.Exec("roma", "IT", time.Now())
	// e(err)

	// id, err := res.LastInsertId()
	// e(err)

	// fmt.Println(id)

	// stmt, err = db.Prepare("update userinfo set username=? where uid=?")

	// res, err = stmt.Exec("Roland", id)
	// e(err)

	// affect, err := res.RowsAffected()
	// e(err)

	// fmt.Println("rows affected", affect)

	// rows, err := db.Query("SELECT * FROM userinfo")
	// e(err)

	// var uid int
	// var username string
	// var departname string
	// var created time.Time

	// for rows.Next() {
	// 	err = rows.Scan(&uid, &username, &departname, &created)
	// 	e(err)
	// 	fmt.Println(uid, username, departname, created)
	// }

	// rows.Close()

	// stmt, err = db.Prepare("delete from userinfo where uid=?")
	// e(err)

	// res, err = stmt.Exec(id)
	// e(err)

	// affect, err = res.RowsAffected()
	// e(err)

	// fmt.Println("rows affected", affect)

	// db.Close()
}

func e(err error) {
	if err != nil {
		panic(err)
	}
}

// func main() {
// 	journal := NewJournal("journal")

// 	journal.AddEntry(&JournalEntryInput{
// 		title: "my first entry",
// 		body:  "I'm feeling good that I write Go",
// 	})
// }

/*
TODO list

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
