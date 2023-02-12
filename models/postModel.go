package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string 
	Password string
}

type Post struct {
	gorm.Model
	Title string
	Body  string
}
