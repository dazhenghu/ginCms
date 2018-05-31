package web

import (
    "github.com/dazhenghu/ginApp"
    _ "github.com/dazhenghu/ginCms/admin/controller"
    "path/filepath"
    "github.com/dazhenghu/util/fileutil"
    "path"
    "github.com/dazhenghu/ginCms/common/service"
)

var App *ginApp.GinApp

func Run()  {
    App = ginApp.Instance()
    // 获取挡墙文件夹
    currPath, _ := fileutil.GetCurrentDirectory()
    // 设置common文件夹位置
    App.SetCommonPath(path.Join(currPath, "../common"))
    // 设置根目录文件夹位置
    App.SetRootPath(currPath)
    // 读取默认位置的配置文件
    App.DefaultLoadConfig("")
    // static参数：relativePath，表示模板中的设定的读取资源文件的路径;root，表示真实的静态文件物理路径。
    // 相当于url中relativePath映射到root
    App.Engine().Static(App.AppConfig.Pulic.RelativePath, App.AppConfig.Pulic.Root)
    // 加载html
    htmlDir := filepath.Join(currPath, App.AppConfig.ViewBaseDir)
    App.Engine().LoadHTMLGlob(htmlDir)

    // service层初始化
    service.Init(App.AppConfig.Dblist)

    addr := App.AppConfig.Addr
    if addr == "" {
        addr = ":8001"
    }
    App.Run(App.AppConfig.Addr)
}
