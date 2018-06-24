package service

import (
    "github.com/dazhenghu/ginCms/common/model"
    "github.com/dazhenghu/util/encryptutil"
    "errors"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/sessions"
    "github.com/dazhenghu/ginCms/common/consts"
    "encoding/gob"
)

type user struct {

}

// 用户信息
type UserInfoMap map[string]interface{}

func NewUserInfoMap() UserInfoMap {
    return make(UserInfoMap)
}

func (um UserInfoMap) GetUserId() int64  {
    if um["user_id"] == nil {
        return 0
    }
    return um["user_id"].(int64)
}

func (um UserInfoMap) GetUserName() string {
    if um["user_name"] == nil {
        return ""
    }
    return um["user_name"].(string)
}

func (um UserInfoMap) GetAliasName() string  {
    if um["alias_name"] == nil {
        return ""
    }
    return um["alias_name"].(string)
}

func (um UserInfoMap) GetMail() string  {
    if um["mail"] == nil {
        return ""
    }
    return um["mail"].(string)
}

func (um UserInfoMap) GetAvatar() string  {
    if um["avatar"] == nil {
        return ""
    }
    return um["avatar"].(string)
}
/**
将Map数据转换成user model
 */
func (um UserInfoMap) ConvertToUserModel() (userObj *model.User)  {
    userObj = &model.User{}
    userObj.UserId = um.GetUserId()
    userObj.UserName = um.GetUserName()
    userObj.UserAliasName = um.GetAliasName()
    userObj.UserMail = um.GetMail()
    userObj.UserAvatar = um.GetAvatar()
    return
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
func (u *user) Register(userInfoMap UserInfoMap, pwd, pwdConfirm string) (registerResult bool, err error) {
    if pwd != pwdConfirm {
        err = errors.New("两次输入的密码不同")
        return
    }

    userObj := userInfoMap.ConvertToUserModel()
    userObj.UserPassword = pwd
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

/**
用户登录校验，并将用户登录信息保存在session中
 */
func (u *user) LoginWithSession(account, plainPassword string, context *gin.Context) (userInfo UserInfoMap, err error) {
    userObj, err := u.Login(account, plainPassword)
    if err != nil {
        return
    }

    gob.RegisterName("UserInfoMap", UserInfoMap{})
    userInfo = make(map[string]interface{})
    userInfo["user_id"] = userObj.UserId
    userInfo["user_name"] = userObj.UserName
    userInfo["alias_name"] = userObj.UserAliasName
    userInfo["mail"] = userObj.UserMail
    userInfo["avatar"] = userObj.UserAvatar

    sess := sessions.Default(context)
    sess.Set(consts.SESSION_KEY_USER, userInfo)
    err = sess.Save()
    return
}

/**
判断用户是否登录
 */
func (u *user) IsLogin(context *gin.Context) bool  {
    sess := sessions.Default(context)
    userInfo := sess.Get(consts.SESSION_KEY_USER)
    if userInfo == nil {
        return false
    }

    return true
}

/**
退出登录，删除session中保留的登录状态信息
 */
func (u *user) LogOut(context *gin.Context) bool {
    sess := sessions.Default(context)
    sess.Delete(consts.SESSION_KEY_USER)
    return true
}

/**
获取当前登录用户信息
 */
func (u *user) GetCurrLoginUserInfo(context *gin.Context, key string) (value interface{}, ok bool) {
    sess := sessions.Default(context)
    value = sess.Get(key)
    if value == nil {
        ok = false
    } else {
        ok = true
    }

    return
}
