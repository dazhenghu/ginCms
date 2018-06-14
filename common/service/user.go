package service

import (
    "github.com/dazhenghu/ginCms/common/model"
    "github.com/dazhenghu/util/encryptutil"
    "errors"
)

type user struct {

} 

var User *user

func init()  {
    User = &user{}
}

func (u *user) NewUser() *model.User  {
    return &model.User{}
}

/**
用户注册
 */
func (u *user) Register(userObj *model.User) (registerResult bool, err error) {

    if err = userObj.IsValid(); err != nil {
        return
    }
    // 生成盐
    salt := encryptutil.GenerateSalt()
    userObj.UserSalt = salt
    // 密码加密
    pswdEncrypt := encryptutil.EncryptWithSalt(userObj.UserPassword, salt)
    userObj.UserPassword = pswdEncrypt

    err = db.Create(userObj).Error
    registerResult = db.NewRecord(userObj)
    return
}

/**
用户登录
 */
func (u *user) Login(account, plainPassword string) (userObj *model.User, err error)  {
    userObj = &model.User{}
    db.Where("user_name = ?", account).Find(userObj)
    if userObj.UserId > 0 {
        // 存在用户，校验密码
        pswdEncrypt := encryptutil.EncryptWithSalt(plainPassword, userObj.UserSalt)
        if pswdEncrypt != userObj.UserPassword {
            userObj = nil
            err = errors.New("账号或密码错误")
            return
        }
    } else {
        userObj = nil
        err = errors.New("账号或密码错误")
    }
    return
}


