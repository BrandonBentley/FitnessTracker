package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"github.com/gorilla/mux"
	"encoding/json"
	"github.com/satori/go.uuid"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)


func JSONRequestHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("JSON REQUESTED")
	http.ServeFile(w, r, "./json/data.json")
}

func JSONPostHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("JSON RECIEVED")
	defer r.Body.Close()
	jsonData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	} else {
		err2 := ioutil.WriteFile(".\\json\\data2.json", jsonData, 0644)
		if err2 != nil {
			fmt.Println(err2)
		}
	}
}

func getLogin(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("session")
	if err != nil {
		http.ServeFile(w, r, "site/login.html")
	} else {
		getHome(w, r)
	}

}

func postLogin(w http.ResponseWriter, r *http.Request) {
	jsonData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		check(err)
		return
	}
	loginInfo := LoginInfo{}
	err = json.Unmarshal(jsonData, &loginInfo)
	defer r.Body.Close()
	var c *http.Cookie
	queryStatement := "SELECT user_id, pass_hash FROM users WHERE (username=? OR email=?)"
	rows, err := db.Query(queryStatement, loginInfo.Username, loginInfo.Username)
	if err != nil {
		check(err)
		return
	}
	var id, passhash string
	if rows.Next() {
		rows.Scan(&id, &passhash)
		misMatch := bcrypt.CompareHashAndPassword([]byte(passhash), []byte(loginInfo.Password)) != nil
		if misMatch {
			w.WriteHeader(401)
			return
		}
		sID, err := uuid.NewV4()
		if err != nil {
			fmt.Println(err)
		} else {
			c = &http.Cookie{
				Name: "session",
				Value: sID.String(),
			}
			cookieJar[sID.String()] = loginInfo
			http.SetCookie(w, c)
			fmt.Println("Session Cookie Set")
		}
		http.Redirect(w, r, "/", 302)
	}
}

func getSignUp(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "site/signup.html")
}

func postSignUp(w http.ResponseWriter, r *http.Request) {
	//TODO Inplement signup POST
	fmt.Println(" Posted to /signup")
}

func getHome(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("session")
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		getIndex(w, r)
	}
}

func getIndex(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.Dir("site/")).ServeHTTP(w, r)
}

func getCSS(w http.ResponseWriter, r *http.Request) {
	value := mux.Vars(r)
	w.Header().Add("Content-Type", "text/css")
	fileName := "site/css/" + value["stylesheet"]
	http.ServeFile(w, r, fileName)
	return
}

func getJS(w http.ResponseWriter, r *http.Request) {
	value := mux.Vars(r)
	fileName := "site/js/" + value["script"]
	w.Header().Add("Content-Type", "application/javascript")
	http.ServeFile(w, r, fileName)
	return
}

func GetImage(w http.ResponseWriter, r *http.Request) {
	value := mux.Vars(r)
	fileName := "site/images/" + value["image"]
	http.ServeFile(w, r, fileName)
	return
}