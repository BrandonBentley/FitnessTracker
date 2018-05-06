package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"github.com/gorilla/mux"
	"encoding/json"

	"github.com/renstrom/shortuuid"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"time"
)


func JSONRequestHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("JSON REQUESTED")
	noCache(&w)
	//time.Sleep(time.Second)
	http.ServeFile(w, r, "./json/data.json")

}

func getUserData(w http.ResponseWriter, r *http.Request) {
	if IsCookieInJar(r) {
		userId, err := GetUserIdFromCookieJar(r)
		check(err)
		finalJson, err := json.Marshal(getUserInfo(userId))
		w.Write(finalJson)
	} else {
		finalJson, err := json.Marshal(UserData{"INVALID", "INVALID"})
		check(err)
		w.Write(finalJson)
	}
}
func JSONPostHandler(w http.ResponseWriter, r *http.Request) {
	requestcategory := r.Header.Get("category")
	jsonData, _ := ioutil.ReadAll(r.Body)
	current, err := r.Cookie("session")
	check(err)
	if IsCookieInJar(r){
		fmt.Println("cookie not found")
		return;
	}
	userId := cookieJar[current.Value].UserId
	defer r.Body.Close()
	var reqType RequestType
	switch requestcategory {
	case "workoutList":
		GetUserWorkoutList(userId)
		break
	case "workout":
		err := json.Unmarshal(jsonData, &reqType)
		check(err)
		break
	case "exercise":

		break
	default:

		break
	}


	check(err)
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
	w.Write(finalJson)
}

func getWorkouts(w http.ResponseWriter, r *http.Request) {
	userId, err := GetUserIdFromCookieJar(r)
	if err != nil {
		fmt.Println(err)
		return
	}
	workoutList := GetUserWorkoutList(userId)
	finalJson, err := json.Marshal(workoutList)
	check(err)
	w.Write(finalJson)
}



func getExercises(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("Id")
	title := r.URL.Query().Get("Title")
	requestType := RequestType{title, id}
	exList := GetUserWorkoutExercises(requestType)
	finalJson, err := json.Marshal(exList)
	check(err)
	w.Write(finalJson)
}


func getSets(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("Id")
	title := r.URL.Query().Get("Title")
	requestType := RequestType{title, id}
	setList := GetExerciseSets(requestType)
	finalJson, err := json.Marshal(setList)
	check(err)
	w.Write(finalJson)
}


func postWorkouts(w http.ResponseWriter, r *http.Request) {
	userId, err := GetUserIdFromCookieJar(r)
	postData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("1")
		fmt.Println(err)
		return
	}
	var workoutSubmitted SubmitType
	err = json.Unmarshal(postData, &workoutSubmitted)
	if err != nil {
		fmt.Println(err)
	}
	newWorkout := InsertUserWorkout(workoutSubmitted, userId)
	finalJson, err := json.Marshal(newWorkout)
	check(err)
	w.Write(finalJson)
}

func postExercises(w http.ResponseWriter, r *http.Request) {
	postData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	var exSubmitted SubmitType
	err = json.Unmarshal(postData, &exSubmitted)
	if err != nil {
		fmt.Println("JSON ERROR")
		fmt.Println(err)
	}
	newEx := InsertUserExercise(exSubmitted)
	finalJson, err := json.Marshal(newEx)
	check(err)
	w.Write(finalJson)
}

func postSets(w http.ResponseWriter, r *http.Request) {
	postData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	var submitSet SubmitSet
	err = json.Unmarshal(postData, &submitSet)
	if err != nil {
		fmt.Println(string(postData[:len(postData)]))
		fmt.Println("JSON ERROR")
		fmt.Println(err)
		return
	}
	newEx := InsertExerciseSet(submitSet)
	finalJson, err := json.Marshal(newEx)
	check(err)
	fmt.Println("Finished Set")
	w.Write(finalJson)
}

func deleteData(w http.ResponseWriter, r *http.Request) {
	postData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	var deleteRequest DeleteRequest
	err = json.Unmarshal(postData, &deleteRequest)
	if err != nil {
		fmt.Println(string(postData[:len(postData)]))
		fmt.Println("JSON ERROR")
		fmt.Println(err)
		return
	}
	//fmt.Println(deleteRequest)
	deleteRecord(deleteRequest)
}


func getLogout(w http.ResponseWriter, r *http.Request) {
	noCache(&w)
	c, err := r.Cookie("session")
	if err != nil {
		http.Redirect(w, r, "/login", 302)

	} else {
		delete(cookieJar, c.Value)
		c = &http.Cookie{
			Name:   "session",
			Value: "",
			MaxAge: -1}
		http.SetCookie(w, c)
		http.ServeFile(w, r, config.RootDir + "logout.html")
	}

}

func getLogin(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("session")
	if err != nil {
		http.ServeFile(w, r, config.RootDir + "login.html")
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
		c = createCookie(id)
		http.SetCookie(w, c)
		fmt.Println("Session Cookie Set")
		http.Redirect(w, r, "/", 302)
	}
}

func getSignUp(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, config.RootDir + "signup.html")
}

func postSignUp(w http.ResponseWriter, r *http.Request) {
	jsonData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		check(err)
		return
	}
	signupInfo := SignUpInfo{}
	err = json.Unmarshal(jsonData, &signupInfo)
	fmt.Println("SignupData")
	fmt.Println(signupInfo)
	defer r.Body.Close()
	if CreateUser(signupInfo) {
		fmt.Println("Successfully Created User:" + signupInfo.Username)
		//login := LoginInfo{Username: signupInfo.Username, Password: signupInfo.Password, Remember: false}
		//w.WriteHeader(201)
		//w.Write([]byte("Successfully Created User:" + signupInfo.Username))
	}
	fmt.Println(" Posted to /signup")
}

func getHome(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")
	if err != nil || c == nil {
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
	http.FileServer(http.Dir(config.RootDir + "")).ServeHTTP(w, r)
}

func getCSS(w http.ResponseWriter, r *http.Request) {
	value := mux.Vars(r)
	w.Header().Add("Content-Type", "text/css")
	fileName := config.RootDir + "css/" + value["stylesheet"]
	http.ServeFile(w, r, fileName)
	return
}

func getJS(w http.ResponseWriter, r *http.Request) {
	value := mux.Vars(r)
	fileName := config.RootDir + "js/" + value["script"]
	w.Header().Add("Content-Type", "application/javascript")
	http.ServeFile(w, r, fileName)
	return
}

func getJSPlugin(w http.ResponseWriter, r *http.Request) {
	value := mux.Vars(r)
	fileName := config.RootDir + "js/plugins/" + value["plugin"]
	w.Header().Add("Content-Type", "application/javascript")
	http.ServeFile(w, r, fileName)
}

func getImage(w http.ResponseWriter, r *http.Request) {
	value := mux.Vars(r)
	fileName := config.RootDir + "images/" + value["image"]
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