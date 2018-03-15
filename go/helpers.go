package main

import (
	"fmt"
)

func check(err error) bool {
	if err != nil {
		fmt.Println(err)
		return true
	}
	return false
}

func testQuery(username string) string {
	queryStatement := "SELECT user_id FROM users WHERE username=? OR email=?"
	rows, err := db.Query(queryStatement, username, username)
	if err != nil {
		return err.Error()
	}
	var full, value string
	for rows.Next() {
		rows.Scan(&value)
		full += value + "\n"
	}
	return full
}