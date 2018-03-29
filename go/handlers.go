package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"github.com/gorilla/mux"
	"encoding/json"
	"github.com/satori/go.uuid"
	"github.com/renstrom/shortuuid"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"time"
)


func JSONRequestHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("JSON REQUESTED")
	noCache(&w)
	time.Sleep(time.Second)
	http.ServeFile(w, r, "./json/data.json")
}

func JSONPostHandler(w http.ResponseWriter, r *http.Request) {
	//-------------------------------------------------------------------
	////fmt.Println("JSON RECIEVED")
	//defer r.Body.Close()
	//jsonData, err := ioutil.ReadAll(r.Body)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	err2 := ioutil.WriteFile(".\\json\\data2.json", jsonData, 0644)
	//	if err2 != nil {
	//		fmt.Println(err2)
	//	}
	//}
	//-------------------------------------------------------------------
	var requestType RequestType
	jsonData, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(jsonData, &requestType)
	check(err)
	defer r.Body.Close()
	raw, err := ioutil.ReadFile("./json/workout.json")
	check(err)
	var item Item
	json.Unmarshal(raw, &item)
	item.Value = shortuuid.New()
	for _, si := range item.List {
		si.Value = shortuuid.New()
	}
	finalJson, err := json.Marshal(item)
	check(err)
	time.Sleep(time.Second)
	w.Write(finalJson)
	/*
	if requestType.Request == "workout" {
		http.ServeFile(w, r, "./json/workout.json")
	} else {
		http.ServeFile(w, r, "./json/workout.json")
	}
	*/
}

func getLogout(w http.ResponseWriter, r *http.Request) {
	noCache(&w)
	c, err := r.Cookie("session")
	if err != nil {
		http.Redirect(w, r, "/login", 302)

	} else {
		delete(cookieJar, "moo")
		c = &http.Cookie{
			Name:   "session",
			MaxAge: -1}
		http.SetCookie(w, c)
		http.ServeFile(w, r, "site/logout.html")
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
			cookieJar[sID.String()] = WrappedCookie{*c, id, time.Now()}
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
	//TODO Implement signup POST
	fmt.Println(" Posted to /signup")
}

func getHome(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		if _, ok := cookieJar[c.Value]; ok {
			getIndex(w, r)
		} else {
			c := http.Cookie{
				Name:   "session",
				MaxAge: -1}
			http.SetCookie(w, &c)
		}

	}
}

func getIndex(w http.ResponseWriter, r *http.Request) {
	noCache(&w)
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

func getJSPlugin(w http.ResponseWriter, r *http.Request) {
	value := mux.Vars(r)
	fileName := "site/js/plugins/" + value["plugin"]
	w.Header().Add("Content-Type", "application/javascript")
	http.ServeFile(w, r, fileName)
}

func getImage(w http.ResponseWriter, r *http.Request) {
	value := mux.Vars(r)
	fileName := "site/images/" + value["image"]
	http.ServeFile(w, r, fileName)
	return
}

func noCache(writer *http.ResponseWriter) {
		w := *writer
		w.Header().Set("Cache-Control", "no-cache, private, max-age=0")
		w.Header().Set("Expires", time.Unix(0, 0).Format(http.TimeFormat))
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("X-Accel-Expires", "0")
}