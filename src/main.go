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
	results, _ := db.Query("SELECT * FROM user")
	for results.Next() {
		var userData = UserData{}
		err = results.Scan(
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
		fmt.Println(userData)
	}
}
