package ginApp

import "github.com/gin-gonic/gin"

type GinApp struct {
    *gin.Engine
}

func Run() *GinApp {
    app := gin.Default()
    ginApp := &GinApp{
        app,
    }

    ginApp.Run()
    return ginApp
}
