package main

import (
    "github.com/dazhenghu/ginApp"
    "github.com/dazhenghu/util/fileutil"
    "path"
    "path/filepath"
    "github.com/dazhenghu/ginCms/admin/web"
    "github.com/dazhenghu/ginCms/common/service"
    _ "github.com/dazhenghu/ginCms/admin/controller"
)

func main()  {
    Run()
}

func Run()  {
    ginAppInstance := ginApp.Instance()
    // 保存全局app实例
    web.App = ginAppInstance
    // 获取挡墙文件夹
    currPath, _ := fileutil.GetCurrentDirectory()
    // 设置common文件夹位置
    web.App.SetCommonPath(path.Join(currPath, "../common"))
    // 设置根目录文件夹位置
    web.App.SetRootPath(currPath)
    // 读取默认位置的配置文件
    web.App.DefaultLoadConfig("")
    // static参数：relativePath，表示模板中的设定的读取资源文件的路径;root，表示真实的静态文件物理路径。
    // 相当于url中relativePath映射到root
    web.App.Engine().Static(web.App.AppConfig.Pulic.RelativePath, web.App.AppConfig.Pulic.Root)
    // 加载html
    htmlDir := filepath.Join(currPath, web.App.AppConfig.ViewBaseDir)
    web.App.Engine().LoadHTMLGlob(htmlDir)

    // service层初始化
    service.Init(web.App.AppConfig.Dblist)

    addr := web.App.AppConfig.Addr
    if addr == "" {
        addr = ":8001"
    }
    web.App.Run(web.App.AppConfig.Addr)
}