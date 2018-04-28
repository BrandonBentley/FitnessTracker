package main

import (
	"database/sql"
)

var cookieJar = map[string]WrappedCookie{}
var db *sql.DB
var config Configuration