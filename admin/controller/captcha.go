package controller

import (
    "github.com/dazhenghu/ginApp/controller"
    "github.com/gin-gonic/gin"
    "github.com/dazhenghu/ginApp/identify"
    "fmt"
    "github.com/dazhenghu/ginApp/logs"
)

type captchaController struct {
    controller.Controller
}

var captchaControllerInstance *captchaController

func init()  {
    captchaControllerInstance = &captchaController{}
    captchaControllerInstance.Init(captchaControllerInstance)

    captchaControllerInstance.Get("captcha/:name", captchaControllerInstance.captchaCreate)
}

func (cc *captchaController) captchaCreate(context *gin.Context) {
    name := context.Param("name")
    logs.Debug(fmt.Sprintf("captcha name:%+v\n", name))
    // 生成验证码信息
    captcha := identify.CaptchaNew(240, 80)
    err := captcha.Handle(context)
    logs.Error(fmt.Sprintf("captcha err:%+v\n", err))
    return
}