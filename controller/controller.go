package controller

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jyoti/models"
)

var NewUser []models.User
var NewCourse []models.Course

// display on browser
func Serverpage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
	<h1> welcome to Student Ledger Backend API</h1>
	`))
}

// Login
func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login")

	w.Header().Set("Contant-Type", "appication/json")

	param := mux.Vars(r)

	for _, item := range NewUser {
		if param["id"] == item.ID && param["secretcode"] == item.SecretCode {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

}

// get all student
func GetAllStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all student")

	w.Header().Set("Contant-Type", "appication/json")

	for _, item := range NewUser {
		if item.UserType == "student" {
			json.NewEncoder(w).Encode(item)

		}
	}

}

// ViewReportCard
func ViewReportCard(w http.ResponseWriter, r *http.Request) {
	fmt.Println("View ReportCard")

	w.Header().Set("Contant-Type", "appication/json")

	param := mux.Vars(r)

	for _, item := range NewUser {
		if param["id"] == item.ID && item.UserType == "student" {

			json.NewEncoder(w).Encode(item)
			break
		}

	}
}

// Register
func Register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create user")

	w.Header().Set("Contant-Type", "appication/json")

	if r.Body == nil {
		http.Error(w, "please send some date", http.StatusBadRequest)
		return
	}

	var u models.User
	_ = json.NewDecoder(r.Body).Decode(&u)

	rand.Seed(time.Now().UnixNano())
	u.ID = string(rand.Intn(100))

	NewUser = append(NewUser, u)
	json.NewEncoder(w).Encode(u)
}

// add grade
func AddGrade(w http.ResponseWriter, r *http.Request) {
	fmt.Println("add grade")

	w.Header().Set("Contant-Type", "appication/json")

	param := mux.Vars(r)

	for index, item := range NewUser {
		if param["id"] == item.ID && item.UserType == "student" {
			NewUser = append(NewUser[:index], NewUser[index+1:]...)
			var u models.User
			_ = json.NewDecoder(r.Body).Decode(&u)
			u.ID = param["id"]

			NewUser = append(NewUser, u)
			json.NewEncoder(w).Encode(u)
			return
		}
	}

}

// update grade
func UpdateGrade(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update grade")

	w.Header().Set("Contant-Type", "appication/json")

	param := mux.Vars(r)
	// fmt.Println(param)
	for index, item := range NewUser {
		if param["id"] == item.ID && item.UserType == "student" && param["course"] == item.Course.Name {

			NewUser = append(NewUser[:index], NewUser[index+1:]...)

			var u models.User
			_ = json.NewDecoder(r.Body).Decode(&u)
			u.ID = param["id"]

			NewUser = append(NewUser, u)
			json.NewEncoder(w).Encode(u)
			return
		}
	}

}

// DeleteGrade
func DeleteGrade(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete grade")

	w.Header().Set("Contant-Type", "appication/json")

	param := mux.Vars(r)

	for index, u := range NewUser {
		if u.ID == param["id"] {

			NewUser = append(NewUser[:index], NewUser[index+1:]...)

			var u models.User
			_ = json.NewDecoder(r.Body).Decode(&u)
			u.ID = param["id"]
			u.Course.Grade = nil

			NewUser = append(NewUser, u)
			json.NewEncoder(w).Encode(u)
			return
		}
	}
	json.NewEncoder(w).Encode(&NewUser)

}

// FilterStudentsByCourse
func FilterStudentsByCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Filter Students By Course")

	w.Header().Set("Contant-Type", "appication/json")

	param := mux.Vars(r)

	for _, item := range NewUser {
		if item.Course.Name == param["course"] {
			json.NewEncoder(w).Encode(item)
		}
	}

	json.NewEncoder(w).Encode(&NewUser)
}
