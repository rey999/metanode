package entity

import (
	"gorm.io/gorm"
)

type Log struct {
	gorm.Model
	UserID       uint
	InParams     string
	OutParameter string
	Url          string
	Time         string
}
