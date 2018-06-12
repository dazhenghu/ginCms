package service

import (
    "github.com/dazhenghu/ginCms/common/model"
    "github.com/dazhenghu/ginCms/admin/web"
    "github.com/dazhenghu/util/encryptutil"
    "crypto/md5"
    "encoding/hex"
    "strings"
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

func (u *user) Register(userObj *model.User)  {



    salt := encryptutil.GenerateSalt()
    pswdEncrypt := encryptutil.EncryptWithSalt(userObj.UserPassword, salt)
    userObj.UserPassword = pswdEncrypt
    userObj.UserSalt = salt

}

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


