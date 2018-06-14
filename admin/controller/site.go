package controller

import (
    "github.com/dazhenghu/ginApp/controller"
    "github.com/gin-gonic/gin"
    "net/http"
)

type siteController struct {
    controller.Controller
}

var siteInstance *siteController

func init()  {
    siteInstance = &siteController{}
    siteInstance.Init(siteInstance)


}

func (site *siteController) Login(context *gin.Context)  {
    if context.Request.Method == http.MethodGet {
        // get请求
        context.HTML(http.StatusOK, "site/login.html", gin.H{
            "pageTitle": "登录",
        })
    }
}
