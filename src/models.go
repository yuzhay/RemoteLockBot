package main

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UserName  string
	FirstName string
	LastName  string
	//Locks         []Lock `gorm:"many2many:subscriptions;"`
	Subscriptions []Subscription
}

type Lock struct {
	gorm.Model
	Serial string `sql:"not null;unique"`
	//Users         []User `gorm:"many2many:subscriptions;"`
	Subscriptions []Subscription
}

type Subscription struct {
	gorm.Model
	User User
	Lock Lock
}
