package model

import "github.com/jinzhu/gorm"

type User struct {
    gorm.Model
    UserName string `gorm:"size:128"`
}
