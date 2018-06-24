package controller

import (
    "github.com/dazhenghu/ginApp/controller"
    "github.com/gin-gonic/gin"
    "github.com/dazhenghu/util/dhutil"
    "net/http"
    "github.com/dazhenghu/ginCms/common/service"
    adminConsts "github.com/dazhenghu/ginCms/admin/consts"
    "github.com/kataras/iris/core/errors"
)

type indexController struct {
    controller.Controller
}

var indexInstace *indexController

func init()  {
    indexInstace = &indexController{}
    indexInstace.Init(indexInstace)

    indexInstace.Get("/", indexInstace.index)
    //indexInstace.Get("/index", indexInstace.index)
}

/**
action调用前回调
 */
func (c *indexController) beforeAction(context *gin.Context) error {
    isLogin := service.User.IsLogin(context)
    if !isLogin {
        context.Redirect(http.StatusFound, adminConsts.URL_LOGIN)
        return errors.New("未登录")
    }

    return nil
}

func (c *indexController)index(context *gin.Context)  {
    beginTime := dhutil.CurrTimeFormat(dhutil.TIME_FORMAT_MIDDLE_SPLIT)
    context.HTML(http.StatusOK, "index/index.html", gin.H{
        "pageTitle": "GINCMS-后台首页",
        "beginTime": beginTime,
        "endTime": dhutil.CurrTimeFormat(dhutil.TIME_FORMAT_MIDDLE_SPLIT),
    })
}
