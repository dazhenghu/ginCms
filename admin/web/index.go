package web

import (
    "github.com/dazhenghu/ginApp"
    _ "github.com/dazhenghu/ginCms/admin/controller"
    "path/filepath"
    "fmt"
    "github.com/dazhenghu/util/fileutil"
)

var App *ginApp.GinApp

func Run()  {
    App = ginApp.Instance()
    currPath, _ := fileutil.GetCurrentDirectory()
    App.SetRootPath(currPath)
    App.DefaultLoadConfig("")
    // static参数：relativePath，表示模板中的设定的读取资源文件的路径;root，表示真实的静态文件物理路径。
    // 相当于url中relativePath映射到root
    fmt.Printf("Public:%+v\n", App.AppConfig.Pulic)
    App.Engine().Static(App.AppConfig.Pulic.RelativePath, App.AppConfig.Pulic.Root)
    htmlDir := filepath.Join(currPath, App.AppConfig.ViewBaseDir)
    fmt.Printf("htmlDir:%s\n", htmlDir)
    App.Engine().LoadHTMLGlob(htmlDir)

    addr := App.AppConfig.Addr
    if addr == "" {
        addr = ":8001"
    }
    App.Run(App.AppConfig.Addr)
}
