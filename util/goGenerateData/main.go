package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

var db sql.DB

func main() {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/fitnessdb")
	check(err)
	defer db.Close()

	theQuery := generateUserDataQuery()
	fmt.Println(theQuery)

	a, err := db.Query(theQuery)
	fmt.Println(a)
	check(err)
}

func generateUserDataQuery() string {
	usernames := []string{"johnnyboy", "jim_h", "tjif", "bhumble", "lexus", "blackbriar", "jilloutwouldya", "sarbear", "thetrain"}
	fnames := []string{"john", "jim", "thomas", "braden", "lexy", "sara", "jill", "sara", "thomas"}
	lnames := []string{"jefferson", "halpert", "jefferson", "crist", "luther", "clark", "livingston", "madison", "stott"}
	passwords := []string{"password", "office", "old", "friend", "superman", "neighbor", "jack", "choco", "cousin"}
	inserts := "INSERT INTO users (user_id, username, email, pass_hash, first_name, last_name) VALUES "
	for i := 0; i < 9; i++ {
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
		inserts += fmt.Sprintf("(\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\")", "u"+id.String(), usernames[i], email, passwordHash, fnames[i], lnames[i])
		if i < 8 {
			inserts += ", \n"
		} else {
			inserts += " \n"
		}
		fmt.Println(len("u" + id.String()))
	}
	return inserts
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
