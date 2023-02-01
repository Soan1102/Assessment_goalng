package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DNS = "root:Tungusonali@1102@tcp(127.0.0.1:3306)/sk2?parseTime=true"

type employees struct {
	ID      int64   `json:"id"`
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employee employees
	json.NewDecoder(r.Body).Decode(&employee)
	DB.Create(&employee)
	json.NewEncoder(w).Encode(employee)
}
func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user employees
	DB.Delete(&user, params["id"])
	json.NewEncoder(w).Encode("The User is Deleted successfully")
}
func initializeRouter() {
	r := mux.NewRouter() //to set the router

	r.HandleFunc("/employees", CreateEmployee).Methods("POST")
	r.HandleFunc("/employees/{id}", DeleteEmployee).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4000", r))

}
func main() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Can't connect to DB")
	} else {
		fmt.Println("Connection Established")
	}

	DB.AutoMigrate(&employees{})
	initializeRouter()

}
