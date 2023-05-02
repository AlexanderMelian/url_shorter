// models/person.go
package models

type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Username string `json:"username" gorm:"unique;type:varchar(30)"`
	Email    string `json:"email" gorm:"unique;type:varchar(30)"`
	Name     string `json:"name" gorm:"type:varchar(20)"`
	LastName string `json:"last_name" gorm:"type:varchar(20)"`
	Password string `json:"password" gorm:"type:varchar(100)"`
}
