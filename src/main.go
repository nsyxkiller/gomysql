package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type UserData struct {
	Id              int
	CitizentId      string
	Firstname       string
	Lastname        string
	BirthYear       int
	FirstnameFather string
	LastnameFather  string
	FirstnameMother string
	LastnameMother  string
	SoldierId       int
	AddressId       int
}

func main() {
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/testsck")
	if err != nil {
		fmt.Println("connect fail")
	}
	fmt.Println("connect success")
	defer db.Close() //ทำงานจบจะปิดการทำงานของfunctionทันที
	//fmt.Println(add(db))
	// fmt.Println(read(db))
	//fmt.Println(remove(db, "1"))
	//fmt.Println(readByCitizentId(db, "3380074028245"))
	fmt.Println(readByCitizentId2(db, "3380074028245"))
	//fmt.Println(edit(db, "3", "กก"))
	//fmt.Println(read(db))
}

func read(db *sql.DB) []UserData {
	var userDataList []UserData
	results, _ := db.Query("SELECT * FROM user")
	for results.Next() {
		var userData = UserData{}
		err := results.Scan(
			&userData.Id,
			&userData.CitizentId,
			&userData.Firstname,
			&userData.Lastname,
			&userData.BirthYear,
			&userData.FirstnameFather,
			&userData.LastnameFather,
			&userData.FirstnameMother,
			&userData.LastnameMother,
			&userData.SoldierId,
			&userData.AddressId,
		)
		if err != nil {
			panic(err.Error())
		}
		userDataList = append(userDataList, userData)
	}
	return userDataList
}

func add(db *sql.DB) bool {
	statement, _ := db.Prepare(`INSERT INTO user
		(citizen_id, firstname,
		lastname, birthyear, firstname_father,
		lastname_father, firstname_mother,
		lastname_mother, soldier_id,address_id)
		VALUES(?,?,?,?,?,?,?,?,?,?)`) //คำสั่งเก็บในชุดคำสั่งไว้ excute
	defer statement.Close()
	_, err := statement.Exec("5573183074474", "สวัสดี", "ประเทศไทย", "1999", "เดินหน้า", "ประเทศไทย", "สบายดี", "ประเทศไทย", "50", "1")
	if err != nil {
		panic(err.Error())
		return false
	}
	return true
}

func remove(db *sql.DB, id string) bool {
	statement, _ := db.Prepare("DELETE FROM user WHERE user_id=?")
	defer statement.Close()
	_, err := statement.Exec(id)
	if err != nil {
		panic(err.Error())
		return false
	}
	return true
}

func edit(db *sql.DB, id, firstnameFather string) bool {
	statement, _ := db.Prepare("UPDATE user SET firstname_father=? WHERE user_id=?")
	defer statement.Close()
	_, err := statement.Exec(firstnameFather, id)
	if err != nil {
		panic(err.Error())
		return false
	}
	return true
}

func readByCitizentId(db *sql.DB, citizentId string) UserData {
	var userData UserData
	results, _ := db.Query("SELECT * FROM user WHERE citizen_id=?", citizentId)
	for results.Next() {
		err := results.Scan(
			&userData.Id,
			&userData.CitizentId,
			&userData.Firstname,
			&userData.Lastname,
			&userData.BirthYear,
			&userData.FirstnameFather,
			&userData.LastnameFather,
			&userData.FirstnameMother,
			&userData.LastnameMother,
			&userData.SoldierId,
			&userData.AddressId,
		)
		if err != nil {
			panic(err.Error())
		}
	}
	return userData
}

func readByCitizentId2(db *sql.DB, citizentId string) UserData {
	var userData UserData
	err := db.QueryRow("SELECT * FROM user WHERE citizen_id=?", citizentId).Scan(&userData.Id,
		&userData.CitizentId,
		&userData.Firstname,
		&userData.Lastname,
		&userData.BirthYear,
		&userData.FirstnameFather,
		&userData.LastnameFather,
		&userData.FirstnameMother,
		&userData.LastnameMother,
		&userData.SoldierId,
		&userData.AddressId,
	)
	if err != nil {
		panic(err.Error())
	}
	return userData
}
