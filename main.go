package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

)

var db *gorm.DB
var err error

type employees struct {
	ID      int64   `json:"id"`
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}

func createNewEmployee(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var employee employees
	json.Unmarshal(reqBody, &employee)
	db.Create(&employee)
	fmt.Println("Create New Employee")
	json.NewEncoder(w).Encode(employee)
}
func deleteEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p := mux.Vars(r)
	var employee employees
	db.Delete(&employee, p["id"])
	json.NewEncoder(w).Encode("The Employee is deleted successfully")

}

func main() {

	db, err = gorm.Open("mysql", "root:Tungusonali@1102@tcp(127.0.0.1:3306)/sk2?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("Connection failed to open!")
	} else {
		fmt.Println("Connection Established!")
	}
	defer db.Close()
	db.AutoMigrate(&employees{})
	router := mux.NewRouter()
	router.HandleFunc("/employee", createNewEmployee).Methods("POST")
	http.ListenAndServe(":9000", router)
	router.HandleFunc("/employee/{id}", createNewEmployee).Methods("DELETE")

	
}