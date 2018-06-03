package controller

import (
    "github.com/dazhenghu/ginApp/controller"
    "github.com/gin-gonic/gin"
    "github.com/dazhenghu/gueditor"
    "github.com/dazhenghu/util/fileutil"
    "net/http"
    "path/filepath"
    "syscall"
)

type ueditorController struct {
    controller.Controller
}

var ueditorInstance *ueditorController
var uedService *gueditor.Service

func init()  {
    ueditorInstance = &ueditorController{}
    ueditorInstance.Init(ueditorInstance)

    syscall.Umask(0)
    rootPath, _ := fileutil.GetCurrentDirectory()
    configFilePath := filepath.Join(rootPath, "config/ueditor.json") // 设置自定义配置文件路径
    rootPath = filepath.Join(rootPath, "../") // 设置项目根目录
    uedService, _ = gueditor.NewService(nil, nil, rootPath, configFilePath)

    ueditorInstance.PostAndGet("/ueditor", ueditorInstance.index)
}

func (ued *ueditorController) index(context *gin.Context) {
    action := context.Query("action")

    switch action {
    case "config":
        ued.config(context)
    case "uploadimage":
        ued.uploadImage(context)
    }

}

func (ued *ueditorController) config(context *gin.Context) {
    cnf := uedService.Config()
    context.JSON(http.StatusOK, cnf)
}

func (ued *ueditorController) uploadImage(context *gin.Context)  {
    res, _ := uedService.Uploadimage(context.Request)
    context.JSON(http.StatusOK, res)
}