package main

import (
	"github.com/satori/go.uuid"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func getUserInfo(userId string) UserData {
	query := "SELECT username, full_name FROM users WHERE user_id=?"
	rows, err := db.Query(query, userId)
	check(err)
	userData := UserData{}
	if rows.Next() {
		rows.Scan(&userData.UserName, &userData.FullName)
	}
	return userData
}

func CreateUser(info SignUpInfo) bool {
	if (UsernameExist(info.Username)) {
		return false
	}
	if (EmailExist(info.Email)) {
		return false
	}
	userId, err := uuid.NewV4()
	check(err)
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(info.Password), bcrypt.DefaultCost)
	check(err)
	insertStatement := "INSERT INTO users " +
		"(user_id, username, email, pass_hash, full_name)" +
		" VALUES "
	insertStatement += fmt.Sprintf("(\"%s\",\"%s\",\"%s\",\"%s\",\"%s\")",
		userId, info.Username, info.Email, passwordHash, info.Name)
	db.Query(insertStatement)
	return UsernameExist(info.Username)
}

func InsertUserWorkout(workout SubmitType, userId string) SubItem {
	insertQuery := "INSERT INTO workouts (workout_id, user_id, workout_name, date_complete)" +
		"Values" +
		"(?, ?, ?, ?);"
	workoutId, err := uuid.NewV4()
	check(err)
	year, month, day := time.Now().Date()
	todaysDate := fmt.Sprintf("%d-%02d-%d", year, int(month), day)
	workoutString := "w" + workoutId.String()
	_, err = db.Query(insertQuery, workoutString, userId, workout.Name, todaysDate)
	check(err)
	return GetSingleWorkout(workoutString)
}

func InsertUserExercise(exercise SubmitType) SubItem {
	insertQuery := "INSERT INTO exercises (ex_id, workout_id, ex_name, ex_order)" +
		"Values" +
		"(?, ?, ?, ?);"
	exId, err := uuid.NewV4()
	check(err)
	exString := "e" + exId.String()
	orderNum := IncrementWorkoutNumExercise(exercise.Id)
	_, err = db.Query(insertQuery, exString, exercise.Id, exercise.Name, orderNum)
	check(err)
	return GetSingleExercise(exString)
}

func InsertExerciseSet(set SubmitSet) Set {
	insertQuery := "INSERT INTO exercise_sets (set_id, ex_id, set_order, reps, set_weight, done)" +
		"Values" +
		"(?, ?, ?, ?, ?, ?);"
	setId, err := uuid.NewV4()
	check(err)
	setString := "s" + setId.String()
	setOrder := IncrementExerciseNumSets(set.Id)
	_, err = db.Query(insertQuery, setString, set.Id, setOrder, set.Reps, set.Weight, set.Done)
	check(err)
	return GetSingleSet(setString)
}

func UpdateExerciseSet() Set {
	return Set{}
}

func deleteRecord(deleteRequest DeleteRequest) bool {
	switch deleteRequest.Id[0] {
		case 'w':
			_, err1 := db.Query("DELETE FROM workouts WHERE workout_id=?", deleteRequest.Id)
			check(err1)
			rows, err := db.Query("SELECT ex_id FROM exercises WHERE workout_id=?", deleteRequest.Id)
			check(err)
			for rows.Next() {
				var eid string
				rows.Scan(&eid)
				deleteRecord(DeleteRequest{eid})
			}
			break
		case 'e':
			_, err := db.Query("DELETE FROM exercises WHERE ex_id=?", deleteRequest.Id)
			check(err)
			_, err = db.Query("DELETE FROM exercise_sets WHERE ex_id=?", deleteRequest.Id)
			check(err)
			break
		case 's':
			_, err := db.Query("DELETE FROM exercise_sets WHERE set_id=?", deleteRequest.Id)
			check(err)
			break
		default:
			break
	}
	return true
}

func GetSingleSet(setId string) Set {
	getSingleQuery := "SELECT set_id, reps, set_weight, set_order, done FROM exercise_sets WHERE set_id=?;"
	rows, err2 := db.Query(getSingleQuery, setId)
	check(err2)
	var setItem Set
	if rows.Next() {
		rows.Scan(&setItem.Id, &setItem.Reps, &setItem.Weight, &setItem.Number, &setItem.Done)
	}
	fmt.Println(setItem)
	return setItem
}

func IncrementExerciseNumSets(exId string) int {
	updateQuery := "UPDATE exercises SET num_sets=num_sets+1 WHERE ex_id=?"
	_, err := db.Query(updateQuery, exId)
	check(err)
	selectQuery := "SELECT num_sets FROM exercises WHERE ex_id=?"
	row, err := db.Query(selectQuery, exId)
	check(err)
	var num int
	if row.Next() {
		row.Scan(&num)
	}
	return num
}

func IncrementWorkoutNumExercise(workoutId string) int {
	updateQuery := "UPDATE workouts SET num_ex=num_ex+1 WHERE workout_id=?"
	_, err := db.Query(updateQuery, workoutId)
	check(err)
	selectQuery := "SELECT num_ex FROM workouts WHERE workout_id=?"
	row, err := db.Query(selectQuery, workoutId)
	check(err)
	var num int
	if row.Next() {
		row.Scan(&num)
	}
	return num
}

func GetSingleExercise(exId string) SubItem {
	getSingleQuery := "SELECT ex_name, ex_id FROM exercises WHERE ex_id=?;"
	rows, err2 := db.Query(getSingleQuery, exId)
	check(err2)
	var exItem SubItem
	if rows.Next() {
		rows.Scan(&exItem.Title, &exItem.Value)
	}
	fmt.Println(exItem)
	return exItem
}

func GetSingleWorkout(workoutId string) SubItem {
	getSingleQuery := "SELECT workout_name, workout_id FROM workouts WHERE workout_id=?;"
	rows, err2 := db.Query(getSingleQuery, workoutId)
	check(err2)
	var workoutItem SubItem
	if rows.Next() {
		rows.Scan(&workoutItem.Title, &workoutItem.Value)
	}
	fmt.Println(workoutItem)
	return workoutItem
}

func GetUserWorkoutList(userId string) Item {
	query := "SELECT workout_name, workout_id FROM workouts WHERE user_id=? ORDER BY date_complete DESC"
	rows, err := db.Query(query, userId)
	check(err)
	workoutList := Item{}
	workoutList.List = make([]SubItem, 0)
	workoutList.Title = "Workout List"
	workoutList.Value = ""
	for rows.Next() {
		workout := SubItem{}
		rows.Scan(&workout.Title, &workout.Value)
		workoutList.add(workout)
	}
	return workoutList
}

func GetUserWorkoutExercises(workout RequestType) Item {
	query := "SELECT ex_name, ex_id FROM exercises WHERE workout_id=? ORDER BY ex_order ASC"
	rows, err := db.Query(query, workout.Id)
	check(err)
	workoutList := Item{}
	workoutList.List = make([]SubItem, 0)
	workoutList.Title = workout.Title
	workoutList.Value = workout.Id
	for rows.Next() {
		workout := SubItem{}
		rows.Scan(&workout.Title, &workout.Value)
		workoutList.add(workout)
	}
	return workoutList
}

func GetExerciseSets(exercise RequestType) SetList {
	query := "SELECT set_id, reps, set_weight, set_order, done FROM exercise_sets WHERE ex_id=? ORDER BY set_order ASC"
	rows, err := db.Query(query, exercise.Id)
	check(err)
	setList := SetList{}
	setList.List = make([]Set, 0)
	setList.Title = exercise.Title
	setList.Value = exercise.Id
	for rows.Next() {
		set := Set{}
		rows.Scan(&set.Id, &set.Reps, &set.Weight, &set.Number, &set.Done)
		setList.add(set)
	}
	return setList
}