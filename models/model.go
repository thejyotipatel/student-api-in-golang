package models

type User struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Email        string   `json:"email"`
	UserType     string   `json:"userType"`
	ListOfCourse []string `json:"listofcourse,omitempty"`
	SecretCode   string   `json:"secretcode"`
	Course       *Course  `json:"course,omitempty"`
}

type Course struct {
	Name  string `json:"name"`
	Grade *Grade `json:"grade"`
}

type Grade struct {
	GradeName string `json:"grade"`
}
