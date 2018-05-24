package web

import (
    "github.com/dazhenghu/ginApp"
    _ "github.com/dazhenghu/ginCms/admin/controller"
    "path/filepath"
    "fmt"
)

var App *ginApp.GinApp

func Init()  {
    App = ginApp.Instance()
    App.Engine().Static("../public", "")
    htmlDir := filepath.Join("", "./views/**/*")
    fmt.Printf("htmlDir:%s\n", htmlDir)
    App.Engine().LoadHTMLGlob(htmlDir)
}
