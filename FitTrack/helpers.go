package main

import (
	"fmt"
	"net/http"
	"github.com/satori/go.uuid"
	"time"
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

func UsernameExist(user string) bool {
	queryStatement := "SELECT user_id FROM users WHERE username=?"
	row, err := db.Query(queryStatement, user)
	if err != nil {
		return false
	}
	if !row.Next() {
		return false
	} else {
		return true
	}
}

func EmailExist(email string) bool {
	queryStatement := "SELECT user_id FROM users WHERE email=?"
	row, err := db.Query(queryStatement, email)
	if err != nil {
		return false
	}
	if !row.Next() {
		return false
	} else {
		return true
	}
}

func createCookie(userId string) *http.Cookie {
	sID, err := uuid.NewV4()
	if err != nil {
		fmt.Println(err)
	}
	c := &http.Cookie{
		Name: "session",
		Value: sID.String(),
	}
	cookieJar[sID.String()] = WrappedCookie{*c, userId, time.Now()}
	return c
}

func IsCookieInJar(r *http.Request) bool {
	current, err := r.Cookie("session")
	check(err)
	if _, ok := cookieJar[current.Value]; !ok {
		return false
	}
	return true
}

func GetUserIdFromCookieJar(r *http.Request) (string, error) {
	current, err := r.Cookie("session")
	if err != nil {
		return "", err
	}
	return cookieJar[current.Value].UserId, nil
}