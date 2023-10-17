package main

import (
	"database/sql"
	"fmt"
	entity "golang-transaction/student"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "stanners2020"
	dbname   = "enigmacamp"
)

var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

func main() {

}

func enrollSubject(studentEnrollment entity.StudentEnrollment) {
	db := connectDb()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

}

func insertStudentEnrollment(studentEnrollment entity.StudentEnrollment, tx *sql.Tx) {

	insertStudentEnrollment := "INSERT INTO tx_student_enrollment (id, student_id, subject, credit) VALUES($1, $2, $3, $4)"

	_, err := tx.Exec(insertStudentEnrollment, studentEnrollment.Id, studentEnrollment.Student_Id, studentEnrollment.Subject, studentEnrollment.Credit)

}

func connectDb() *sql.DB {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully connected!")
	}

	return db

}
