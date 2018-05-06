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
	"io/ioutil"
	"encoding/json"
)



func main() {
	config = GetConfig()
	fmt.Println(config.String())
	var err error
	var router *mux.Router
	router = GetRouter()
	db, err = sql.Open(config.SqlDriver, config.DatabaseAddress)
	if check(err) {
		return
	}
	defer db.Close()
	var wg sync.WaitGroup
	fmt.Println("Serving on port " + strconv.Itoa(config.Port))
	path, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	fmt.Println("Current Directory: " + path)
	wg.Add(1)
	go prompt(wg)
	if config.HttpsEnabled {
		//HTTPS Server
		log.Fatal(http.ListenAndServeTLS(":"+strconv.Itoa(config.Port),  config.CertFile, config.KeyFile,  router))
	} else {
		//HTTP Server
		log.Fatal(http.ListenAndServe(":"+strconv.Itoa(config.Port),  router))
	}


	wg.Wait()
}

func GetConfig() Configuration {
	raw, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println("config.json not found setting defaults.")
		config := Configuration{8080, false, "cert/server.crt", "cert/server.key", "/site/", "mysql", "root:password@tcp(localhost:3306)/fitnessdb"}
		fmt.Println(config.String())
		return config
	}
	var config Configuration
	json.Unmarshal(raw, &config)
	if (config.RootDir[len(config.RootDir)-1] != '/') {
		config.RootDir += "/";
	}
	return config
}