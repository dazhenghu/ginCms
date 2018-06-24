package model

import (
    "github.com/dazhenghu/ginApp/model"
    "time"
    "regexp"
    "github.com/dazhenghu/util/stringutil"
    "errors"
    "unicode/utf8"
)



type User struct {
    model.BaseModel
    UserId int64 `gorm:"column:user_id" json:"user_id"`
    UserName string `gorm:"column:user_name" json:"user_name"`
    UserPassword string `gorm:"column:user_password" json:"user_password"`
    UserAliasName string `gorm:"column:user_alias_name" json:"user_alias_name"`
    UserMail string `gorm:"column:user_mail" json:"user_mail"`
    UserAvatar string `gorm:"column:user_avatar" json:"user_avatar"`
    UserStatus int64 `gorm:"column:user_status" json:"user_status"`
    UserSalt string `gorm:"column:user_salt" json:"user_salt"`
    UserCreateAt time.Time `gorm:"column:user_create_at" json:"user_create_at"`
    UserUpdateAt time.Time `gorm:"column:user_update_at" json:"user_update_at"`
}

/**
校验数据是否有效
 */
func (u *User) IsValid() (err error) {
    if err = u.checkName(); err != nil {
        return
    }
    if err = u.checkPassword(); err != nil {
        return
    }
    if err = u.checkAliasName(); err != nil {
        return
    }
    if err = u.checkMail(); err != nil {
        return
    }
    return
}

/**
校验用户名
 */
func (u *User) checkName() (err error) {
    nameLen := len(u.UserName)
    if !(nameLen > 3 && nameLen < 21) {
        err = errors.New("用户名长度必须是4到20位")
        return
    }

    reg := regexp.MustCompile("[^0-9A-Za-z_]")
    if reg.MatchString(u.UserName) {
        err = errors.New("用户名只能由英文字母、数字、下划线组成")
        return
    }

    reg = regexp.MustCompile("^[A-Za-z]")
    if !reg.MatchString(u.UserName) {
        err = errors.New("用户名必须以英文字母开头")
        return
    }

    return
}

/**
校验密码
 */
func (u *User) checkPassword() (err error) {
    pswLen := len(u.UserPassword)
    if !(pswLen > 7 && pswLen < 21) {
        err = errors.New("密码长度必须是8到20位")
        return
    }

    if stringutil.ContainSpace(u.UserPassword) {
        err = errors.New("密码中不能包含空白字符，如空格、回车等")
        return
    }

    return
}

/**
校验昵称
 */
func (u *User) checkAliasName() (err error)  {
    aliasLen := utf8.RuneCountInString(u.UserAliasName)
    if !(aliasLen > 1 && aliasLen < 11) {
        err = errors.New("昵称长度必须是2到10位")
        return
    }

    if stringutil.ContainSpecialChar(u.UserAliasName) {
        err = errors.New("昵称中不能包含特殊字符")
        return
    }

    return
}

/**
校验邮箱
 */
func (u *User) checkMail() (err error) {
    if !stringutil.ValidMail(u.UserMail) {
        err = errors.New("无效的邮箱格式")
        return
    }
    return
}
