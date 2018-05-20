package controller

import (
    "github.com/dazhenghu/ginApp/controller"
    "sync"
    "github.com/dazhenghu/ginCms/admin/web"
    "github.com/gin-gonic/gin"
)

type indexController struct {
    controller.Controller
}

var indexInstace *indexController

func init()  {
    var once sync.Once
    once.Do(func() {
        indexInstace = &indexController{}
        indexInstace.App = web.App
    })

    indexInstace.Get("/", indexInstace.index)
}

func (c *indexController)index(context *gin.Context)  {

}
