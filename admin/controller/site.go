package controller

import (
    "github.com/dazhenghu/ginApp/controller"
    "github.com/gin-gonic/gin"
    "net/http"
    "github.com/dazhenghu/ginApp/session"
    "github.com/dazhenghu/ginCms/common/service"
    "github.com/gin-contrib/sessions"
    "github.com/dazhenghu/ginCms/common/consts"
    adminConsts "github.com/dazhenghu/ginCms/admin/consts"
    "github.com/dazhenghu/ginCms/admin/util"
)

type siteController struct {
    controller.Controller
}

var siteInstance *siteController

func init()  {
    siteInstance = &siteController{}
    siteInstance.Init(siteInstance)

    siteInstance.PostAndGet(adminConsts.URL_LOGIN, siteInstance.Login)
    siteInstance.PostAndGet(adminConsts.URL_REGISTER, siteInstance.Register)
}

/**
登录
 */
func (site *siteController) Login(context *gin.Context)  {
    token, _ := session.GenerateSessionToken(context, consts.SESSION_KEY_POST_TOKEN)

    if context.Request.Method == http.MethodGet {
        // get请求
        context.HTML(http.StatusOK, "site/login.html", gin.H{
            "pageTitle": "登录",
            "token": token,
        })
    } else if context.Request.Method == http.MethodPost {
        // post请求登录
        sessions.Default(context)
        tokenErr := session.CheckSessionToken(context, consts.SESSION_KEY_POST_TOKEN, token)
        if tokenErr != nil {
            context.JSON(http.StatusOK, map[string]string {
                "code":"error",
                "message":"令牌已过期，请刷新重试",
            })
            return
        }

        account := context.PostForm("account")
        password := context.PostForm("password")
        _, err := service.User.LoginWithSession(account, password, context)
        if err != nil {
            context.JSON(http.StatusOK, map[string]string {
                "code":consts.ERROR,
                "message":"登录失败，" + err.Error(),
            })
            return
        }

        redirectUrl := site.redirectUrl(context);
        context.JSON(http.StatusOK, map[string]string {
            "code":consts.SUCCESS,
            "message":"登录成功",
            "redirect":redirectUrl,
        })
        return
    }
}

/**
注册
 */
func (site *siteController) Register(context *gin.Context) {
    token, _ := session.GenerateSessionToken(context, consts.SESSION_KEY_POST_TOKEN)
    if context.Request.Method == http.MethodGet {
        // get请求
        context.HTML(http.StatusOK, "site/register.html", gin.H{
            "pageTitle": "登录",
            "token": token,
        })
    } else if context.Request.Method == http.MethodPost {
        userInfoMap := service.NewUserInfoMap()
        userInfoMap["user_name"] = context.PostForm("user_name")
        userInfoMap["alias_name"] = context.PostForm("alias_name")
        userInfoMap["mail"] = context.PostForm("mail")

        pwd := context.PostForm("user_password")
        pwdConfirm := context.PostForm("user_password_confirm")

        _, err := service.User.Register(userInfoMap, pwd, pwdConfirm)
        if err != nil {
            context.JSON(http.StatusOK, map[string]string {
                "code":consts.ERROR,
                "message":"注册失败，" + err.Error(),
            })
            return
        }

        redirectUrl := site.redirectUrl(context);
        context.Redirect(http.StatusFound, redirectUrl)
        //context.JSON(http.StatusOK, map[string]string {
        //    "code":consts.SUCCESS,
        //    "message":"注册成功",
        //})
        return
    }
}

/**
跳转url
 */
func (site *siteController) redirectUrl(context *gin.Context) (redirectUrl string) {
    referer := context.Request.Referer()

    isLoginPage, urlLoginErr := util.IsLoginPage(referer)
    isRegisterPage, urlRegisterErr := util.IsRegisterPage(referer)

    if len(referer) == 0 || urlLoginErr != nil || urlRegisterErr != nil || isLoginPage || isRegisterPage {
        // 没有referer则跳转到首页
        redirectUrl = "/"
        return
    }

    redirectUrl = referer
    return
}
