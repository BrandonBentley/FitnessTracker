package main

import (
	"github.com/gorilla/mux"
	"database/sql"
)

var port int
var router *mux.Router
var cookieJar = map[string]WrappedCookie{}
var db *sql.DB