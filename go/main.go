package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"github.com/gorilla/mux"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var err error
	db, err = sql.Open("mysql", "root:password@tcp(localhost:3306)/fitnessdb")
	if check(err) {
		return
	}
	defer db.Close()
	port = 8080
	var wg sync.WaitGroup
	fmt.Println("Serving on port " + strconv.Itoa(port))
	path, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	fmt.Println("Current Directory: " + path)
	wg.Add(1)
	go prompt(wg)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), router))
	wg.Wait()
}

func init() {
	router = mux.NewRouter()
	router.HandleFunc("/", getHome).Methods("GET")
	router.HandleFunc("/js/{script}", getJS).Methods("GET")
	router.HandleFunc("/css/{stylesheet}", getCSS).Methods("GET")
	router.HandleFunc("/images/{image}", GetImage).Methods("GET")
	router.HandleFunc("/json", JSONRequestHandler)
	router.HandleFunc("/post", JSONPostHandler)
	router.HandleFunc("/login", getLogin).Methods("GET")
	router.HandleFunc("/login", postLogin).Methods("POST")
	router.HandleFunc("/signup", getSignUp).Methods("GET")
	router.HandleFunc("/signup", postSignUp).Methods("POST")
}