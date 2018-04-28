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
	Name string
}

func (s SignUpInfo) String() string {
	return "{\nUsername: " + s.Username + "\nEmail: " + s.Email +
		"\nName: " + s.Name + "\nPassword: " + s.Password + "\n}"
}

type WrappedCookie struct {
	Cookie http.Cookie
	UserId string
	Created time.Time
}

type RequestType struct {
	Title string
	Id string
}

type SubmitType struct {
	Name string
	Id string
}

type SubmitSet struct {
	Id string
	Update bool
	SetId string
	Reps int
	Weight float64
	Done bool
}

type DeleteRequest struct {
	Id string
}

type Item struct {
	Title string
	Value string
	List []SubItem
}

type WorkoutList struct {
	List []SubItem
}

type UserData struct {
	UserName, FullName string
}

func (i* Item) add(workout SubItem) {
	i.List = append(i.List, workout)
}

type SubItem struct {
	Title string
	Value string
}

type SetList struct {
	Title string
	Value string
	List []Set
}

func (s* SetList) add(set Set) {
	s.List = append(s.List, set)
}

type Set struct {
	Id string
	Reps int
	Weight float64
	Number int
	Done bool
}

type Configuration struct {
	Port int
	HttpsEnabled bool
	CertFile string
	KeyFile string
	RootDir string
	SqlDriver string
	DatabaseAddress string
}

func (c Configuration) String() string {
	return "Port: " + strconv.Itoa(c.Port) + "\tHttpsEnabled: " + strconv.FormatBool(c.HttpsEnabled) +
		"\tRootDir: " + c.RootDir + "\tSqlDriver: " + c.SqlDriver + "\tDatabaseAddress: " + c.DatabaseAddress
}