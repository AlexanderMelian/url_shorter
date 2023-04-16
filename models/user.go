// models/person.go
package models

type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Password string `json:"password"`
}
