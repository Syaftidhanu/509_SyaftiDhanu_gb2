package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "310399"
	dbname   = "db-go"
)

var (
	db  *sql.DB
	err error
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database")

	CreateEmployee()
}

// membuat data employee baru dengan package database/sql
// struct employee
type Employee struct {
	ID        int
	Full_name string
	Email     string
	Age       int
	Division  string
}

// fungsi untuk create data
func CreateEmployee() {
	var employee = Employee{}

	//$1 dst merupakan representasi dari nilai-nilai yang akan dimasukkan
	//Returning * memiliki arti bahwa kita ingin mendapat nilai2 dari seluruh field yang berasa dari data yang baru dibuat
	sqlStatement := `
	INSERT INTO employees (full_name, email, age, division)
	VALUES ($1, $2, $3, $4)
	Returning *
	`

	err = db.QueryRow(sqlStatement, "Airell Jordan", "airell@gmail.com", 23, "IT").
		Scan(&employee.ID, &employee.Full_name, &employee.Email, &employee.Age, &employee.Division)

	if err != nil {
		panic(err)
	}

	fmt.Printf("New Employee Data : %+v\n", employee)

}
