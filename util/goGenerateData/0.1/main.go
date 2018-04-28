package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

var db sql.DB
var userId []string
var workoutId []string
var exerciseId []string
var setId []string
var numUsers int
var numWorkouts int
var numExercises int
var numSets int

func main() {
	numUsers = 9
	numWorkouts = 9
	numExercises = 9
	numSets = 9
	userId = newUserIdSlice()
	workoutId = newWorkoutIdSlice()
	exerciseId = newExerciseIdSlice()
	setId = newSetIdSlice()

	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/fitnessdb")
	check(err)
	defer db.Close()

	theQuery := generateUserDataQuery()
	fmt.Println(theQuery)

	// a, err := db.Query(theQuery)
	// fmt.Println(a)

	// theQuery = generateWorkouts()
	// fmt.Println(theQuery)

	// a, err = db.Query(theQuery)
	//fmt.Println(a)
	//check(err)
}

func generateUserDataQuery() string {
	usernames := []string{"johnnyboy", "jim_h", "tjif", "bhumble", "lexus", "blackbriar", "jilloutwouldya", "sarbear", "thetrain"}
	fnames := []string{"john", "jim", "thomas", "braden", "lexy", "sara", "jill", "sara", "thomas"}
	lnames := []string{"jefferson", "halpert", "jefferson", "crist", "luther", "clark", "livingston", "madison", "stott"}
	passwords := []string{"password", "office", "old", "friend", "superman", "neighbor", "jack", "choco", "cousin"}
	inserts := "INSERT INTO users (user_id, username, email, pass_hash, first_name, last_name) VALUES "
	for i := 0; i < numUsers; i++ {
		email := usernames[i]
		if i%3 == 0 {
			email += "@yahoo.com"
		} else {
			email += "@gmail.com"
		}
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(passwords[i]), bcrypt.DefaultCost)
		check(err)
		id, err := uuid.NewV4()
		check(err)
		inserts += fmt.Sprintf("(\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\")", userId[i], usernames[i], email, passwordHash, fnames[i], lnames[i])
		if i < 8 {
			inserts += ", \n"
		} else {
			inserts += " \n"
		}
		fmt.Println(len("u" + id.String()))
	}
	return inserts
}

func generateWorkouts() string {
	workoutName := []string{"john", "jim_h", "tjif", "bhumble", "lexus", "blackbriar", "jilloutwouldya", "sarbear", "thetrain"}
	inserts := "INSERT INTO users (workout_id, user_id, workout_name, begin_date, end_date) VALUES "
	for i := 0; i < numWorkouts; i++ {

		inserts += fmt.Sprintf("(\"%s\",\"%s\",\"%s\",\"%s\",\"%s\")", workoutId[i], userId[i], workoutName[i]+" Workout", "TO_DATE('17/12/2015', 'DD/MM/YYYY')", "TO_DATE('17/02/2016', 'DD/MM/YYYY')")
		if i < 8 {
			inserts += ", \n"
		} else {
			inserts += " \n"
		}
	}
	return inserts
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func newUserIdSlice() []string {
	s := make([]string, 9)
	for i := 0; i < 9; i++ {
		id, err := uuid.NewV4()
		check(err)
		s[i] = "u" + id.String()
	}
	return s
}

func newWorkoutIdSlice() []string {
	s := make([]string, 9)
	for i := 0; i < 9; i++ {
		id, err := uuid.NewV4()
		check(err)
		s[i] = "w" + id.String()
	}
	return s
}

func newExerciseIdSlice() []string {
	s := make([]string, 9)
	for i := 0; i < 9; i++ {
		id, err := uuid.NewV4()
		check(err)
		s[i] = "e" + id.String()
	}
	return s
}

func newSetIdSlice() []string {
	s := make([]string, 9)
	for i := 0; i < 9; i++ {
		id, err := uuid.NewV4()
		check(err)
		s[i] = "s" + id.String()
	}
	return s
}
