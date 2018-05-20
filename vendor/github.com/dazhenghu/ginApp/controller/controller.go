package controller

import (
    "github.com/gin-gonic/gin"
    "github.com/dazhenghu/ginApp"
)

type Controller struct {
    App *ginApp.GinApp
}

/**
controller初始化
 */
func (c *Controller) Init() error {
    // 通过反射注册action方法
    //cReflect := reflect.ValueOf(c)

    // 获取方法数量
    //numMethod := cReflect.NumMethod()

    return nil
}

/**
action调用前回调
 */
func (c *Controller) beforeAction(context *gin.Context) error {
    return nil
}

/**
action调用后回调
 */
func (c *Controller) afterAction(context *gin.Context) error  {
    return nil
}

/**
GET方法路由handle设置
 */
func (c *Controller) Get(relativePath string, handler gin.HandlerFunc) {
    c.App.GET(relativePath, c.hook(handler))
}

/**
POST方法路由handle设置
 */
func (c *Controller) Post(relativePath string, handler gin.HandlerFunc)  {
    c.App.POST(relativePath, c.hook(handler))
}

func (c *Controller) hook(handler gin.HandlerFunc) func(context *gin.Context)  {
    return func(context *gin.Context) {
        // 执行handler之前执行beforeAction
        berforeErr := c.beforeAction(context)
        if berforeErr != nil {
            panic(berforeErr)
        }
        handler(context)
        // 执行handler之后执行beforeAction
        afterErr := c.afterAction(context)
        if afterErr != nil {
            panic(afterErr)
        }
    }
}