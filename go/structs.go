package main

import (
	"strconv"
	"net/http"
	"time"
)

type LoginInfo struct {
	Username string
	Password string
	Remember bool
}

func (l LoginInfo) String() string {
	return "Username: " + l.Username + "\tpassword: " + l.Password + "\tremember: " + strconv.FormatBool(l.Remember)
}

type SignUpInfo struct {
	Username string
	Email string
	Password string
}

func (s SignUpInfo) String() string {
	return "Username: " + s.Username + "\tpassword: " + s.Password + "\tremember: " + s.Email
}

type WrappedCookie struct {
	Cookie http.Cookie
	UserId string
	Created time.Time
}

type RequestType struct {
	Request string
}

type Item struct {
	Title string
	Value string
	List []SubItem
}

type SubItem struct {
	Title string
	Value string
}