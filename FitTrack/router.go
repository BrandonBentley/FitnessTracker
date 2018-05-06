package main

import "github.com/gorilla/mux"

func GetRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", getHome).Methods("GET")
	router.HandleFunc("/js/{script}", getJS).Methods("GET")
	router.HandleFunc("/js/plugins/{plugin}", getJSPlugin).Methods("GET")
	router.HandleFunc("/css/{stylesheet}", getCSS).Methods("GET")
	router.HandleFunc("/images/{image}", getImage).Methods("GET")
	router.HandleFunc("/json", JSONRequestHandler).Methods("GET")
	router.HandleFunc("/login", getLogin).Methods("GET")
	router.HandleFunc("/logout", getLogout).Methods("GET")
	router.HandleFunc("/signup", getSignUp).Methods("GET")
	router.HandleFunc("/json", JSONPostHandler).Methods("POST")
	router.HandleFunc("/login", postLogin).Methods("POST")
	router.HandleFunc("/signup", postSignUp).Methods("POST")
	router.HandleFunc("/workouts", getWorkouts).Methods("GET")
	router.HandleFunc("/exercises", getExercises).Methods("GET")
	router.HandleFunc("/sets", getSets).Methods("GET")
	router.HandleFunc("/workouts", postWorkouts).Methods("POST")
	router.HandleFunc("/exercises", postExercises).Methods("POST")
	router.HandleFunc("/sets", postSets).Methods("POST")
	router.HandleFunc("/userData", getUserData).Methods("GET")
	router.HandleFunc("/delete", deleteData).Methods("POST")
	return router
}