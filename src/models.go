package main

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string
	Locks    []Lock `gorm:"many2many:subscriptions;"`
}

type Lock struct {
	gorm.Model

	Serial string
	Users  []User `gorm:"many2many:subscriptions;"`
}
