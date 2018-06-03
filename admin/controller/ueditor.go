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

func init() {
    ueditorInstance = &ueditorController{}
    ueditorInstance.Init(ueditorInstance)

    syscall.Umask(0)
    rootPath, _    := fileutil.GetCurrentDirectory()
    configFilePath := filepath.Join(rootPath, "config/ueditor.json") // 设置自定义配置文件路径

    rootPath      = filepath.Join(rootPath, "../") // 设置项目根目录
    uedService, _ = gueditor.NewService(nil, nil, rootPath, configFilePath)

    ueditorInstance.PostAndGet("/ueditor", ueditorInstance.index)
}

func (ued *ueditorController) index(context *gin.Context) {
    action := context.Query("action")

    switch action {
    case "config":
        // config接口
        ued.config(context)
    case "uploadimage":
        // 上传图片
        ued.uploadImage(context)
    case "uploadscrawl":
        // 上传涂鸦
        ued.uploadScrawl(context)
    case "uploadvideo":
        // 上传视频
        ued.uploadVideo(context)
    case "uploadfile":
        // 上传附件
        ued.uploadfile(context)
    case "listfile":
        // 查询上传的文件列表
        ued.listFile(context)
    case "listimage":
        // 查询上传的图片列表
        ued.listImage(context)
    }

}

func (ued *ueditorController) config(context *gin.Context) {
    cnf := uedService.Config()
    context.JSON(http.StatusOK, cnf)
}

func (ued *ueditorController) uploadImage(context *gin.Context) {
    res, _ := uedService.Uploadimage(context.Request)
    context.JSON(http.StatusOK, res)
}

func (ued *ueditorController) uploadScrawl(context *gin.Context)  {
    res, _ := uedService.UploadScrawl(context.Request)
    context.JSON(http.StatusOK, res)
}

func (ued *ueditorController) uploadVideo(context *gin.Context)  {
    res, _ := uedService.UploadVideo(context.Request)
    context.JSON(http.StatusOK, res)
}

func (ued *ueditorController) uploadfile(context *gin.Context)  {
    res, _ := uedService.UploadFile(context.Request)
    context.JSON(http.StatusOK, res)
}

func (ued *ueditorController) listFile(context *gin.Context) {
    listFileItem := &gueditor.ListFileItem{}
    uedService.Listfile(listFileItem, 0, 10)
    context.JSON(http.StatusOK, listFileItem)
}

func (ued *ueditorController) listImage(context *gin.Context)  {
    listFileItem := &gueditor.ListFileItem{}
    uedService.ListImage(listFileItem, 0, 10)
    context.JSON(http.StatusOK, listFileItem)
}