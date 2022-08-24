package routes

import (
	"github.com/gorilla/mux"
	"github.com/jyoti/controller"
)

var RegisterRoute = func(route *mux.Router) {

	route.HandleFunc("/", controller.Serverpage).Methods("GET")
	route.HandleFunc("/login/{id},{secretcode}", controller.Login).Methods("GET")
	route.HandleFunc("/register", controller.Register).Methods("POST")
	route.HandleFunc("/getAllStudent", controller.GetAllStudent).Methods("GET")
	route.HandleFunc("/viewReportCard/{id}", controller.ViewReportCard).Methods("GET")
	route.HandleFunc("/addGrade/{id}", controller.AddGrade).Methods("PUT")
	route.HandleFunc("/updateGrade/{id},{course}", controller.UpdateGrade).Methods("PUT")
	route.HandleFunc("/deleteGrade/{id}", controller.DeleteGrade).Methods("DELETE")
	route.HandleFunc("/filterStudentsByCourse/{course}", controller.FilterStudentsByCourse).Methods("GET")

}
