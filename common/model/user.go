package model

import (
    "github.com/jinzhu/gorm"
    "github.com/dazhenghu/ginApp/model"
)

type User struct {
    model.BaseModel
    UserName string `gorm:"size:64"`
    UserPassword string `gorm:"size:64"`
}

