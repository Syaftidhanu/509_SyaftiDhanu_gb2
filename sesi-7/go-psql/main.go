package main

import (
	"database/sql"
	"fmt"

	//install modulnya dulu menggunakan go get
	_ "github.com/lib/pq"
)

// variable golabl yang digunakan untuk menyimpan seluruh info dari db
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "310399"
	dbname   = "db-go"
)

var (
	//var db memiliki tipe data pointer dari struct sql.db yang akan mereassign dgn obejct dari sql
	db  *sql.DB
	err error
)

func main() {
	//manggabungkan info dari Psotgresql
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	//funsi open yang berasal dari package database/sql
	//fungsi open hanya berfungsi untuk memvalidasi arg
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//metod ping berguna untuk membangun koneksi ke databse
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database")

	// CreateEmployee()
	// GetEmployess()
	//UpdateEmployess()
	DeleteEmployess()
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

func GetEmployess() {
	var results = []Employee{}

	sqlStatement := `SELECT * from employees`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var employee = Employee{}

		err = rows.Scan(&employee.ID, &employee.Full_name, &employee.Email, &employee.Age, &employee.Division)

		if err != nil {
			panic(err)
		}

		results = append(results, employee)
	}

	fmt.Println("Employee datas:", results)
}

func UpdateEmployess() {
	sqlStatement := `
	UPDATE employees
	SET full_name = $2, email = $3, division = $4, age = $5
	WHERE id = $1`

	res, err := db.Exec(sqlStatement, 1, "Airell Jordan Hidayat", "airellhidayat@yahoo.com", "CurDevs", 24)

	if err != nil {
		panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("Update Data Amount:", count)
}

func DeleteEmployess() {
	sqlStatement := `
	DELETe from employees
	WHERE id = $1;
	`

	res, err := db.Exec(sqlStatement, 1)

	if err != nil {
		panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("Deleted Data Amount:", count)
}
