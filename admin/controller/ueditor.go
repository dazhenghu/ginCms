package controller

import (
    "github.com/dazhenghu/ginApp/controller"
    "github.com/gin-gonic/gin"
    "github.com/dazhenghu/gueditor"
    "github.com/dazhenghu/util/fileutil"
    "net/http"
    "fmt"
)

type ueditorController struct {
    controller.Controller
}

var ueditorInstance *ueditorController
var uedService *gueditor.Service

func init()  {
    ueditorInstance = &ueditorController{}
    ueditorInstance.Init(ueditorInstance)

    rootPath, _ := fileutil.GetCurrentDirectory()
    uedService, _ = gueditor.NewService(nil, nil, rootPath, "")

    ueditorInstance.PostAndGet("/ueditor", ueditorInstance.index)
}

func (ued *ueditorController) index(context *gin.Context) {
    action := context.Query("action")

    if action == "config" {
        ued.config(context)
    }

}

func (ued *ueditorController) config(context *gin.Context) {
    cnf := uedService.Config()
    context.JSON(http.StatusOK, cnf)
}
