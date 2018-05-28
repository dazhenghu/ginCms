package model

import (
    "github.com/dazhenghu/ginApp/model"
    "time"
)

type User struct {
    model.BaseModel
    UserId int64 `gorm:"column:user_id" json:"user_id" form:"user_id"`
    UserName string `gorm:"column:user_name" json:"user_name" form:"user_name"`
    UserPassword string `gorm:"column:user_password" json:"user_password" form:"user_password"`
    UserAliasName string `gorm:"column:user_alias_name" json:"user_alias_name" form:"user_alias_name"`
    UserMail string `gorm:"column:user_mail" json:"user_mail" form:"user_mail"`
    UserFullname string `gorm:"column:user_fullname" json:"user_fullname" form:"user_fullname"`
    UserAvatar string `gorm:"column:user_avatar" json:"user_avatar" form:"user_avatar"`
    UserStatus int64 `gorm:"column:user_status" json:"user_status" form:"user_status"`
    UserCreateAt time.Time `gorm:"column:user_create_at" json:"user_create_at" form:"user_create_at"`
    UserUpdateAt time.Time `gorm:"column:user_update_at" json:"user_update_at" form:"user_update_at"`
}


