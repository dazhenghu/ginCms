package controller

import (
    "github.com/dazhenghu/ginApp/controller"
    "sync"
    "github.com/gin-gonic/gin"
    "github.com/dazhenghu/util/dhutil"
    "net/http"
)

type indexController struct {
    controller.Controller
}

var indexInstace *indexController

func init()  {
    var once sync.Once
    once.Do(func() {
        indexInstace = &indexController{}
        indexInstace.Init(indexInstace)
    })

    indexInstace.Get("/", indexInstace.index)
    //indexInstace.Get("/index", indexInstace.index)
}

func (c *indexController)index(context *gin.Context)  {
    beginTime := dhutil.CurrTimeFormat(dhutil.TIME_FORMAT_MIDDLE_SPLIT)
    context.HTML(http.StatusOK, "index/index.html", gin.H{
        "pageTitle": "GINCMS-后台首页",
        "beginTime": beginTime,
        "endTime": dhutil.CurrTimeFormat(dhutil.TIME_FORMAT_MIDDLE_SPLIT),
    })
}
