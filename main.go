package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jyoti/controller"
	"github.com/jyoti/models"
	"github.com/jyoti/routes"
)

func main() {
	fmt.Println("Server running at port 5000...")

	r := mux.NewRouter()
	routes.RegisterRoute(r)

	controller.NewUser = append(controller.NewUser, models.User{ID: "1", Name: "sam", Email: "sam@gmail.com", UserType: "teacher", SecretCode: "123456", Course: &models.Course{Name: "java", Grade: &models.Grade{"C"}}})

	controller.NewUser = append(controller.NewUser, models.User{ID: "2", Name: "joen", Email: "joen@gmail.com", ListOfCourse: []string{"js", "java", "html"}, UserType: "student", SecretCode: "123456"})
	// add Course
	controller.NewCourse = append(controller.NewCourse, models.Course{Name: "go", Grade: &models.Grade{GradeName: "B"}})
	controller.NewCourse = append(controller.NewCourse, models.Course{Name: "html", Grade: &models.Grade{GradeName: "C"}})

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":5000", r))
}
