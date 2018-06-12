package web

import (
    "github.com/dazhenghu/ginApp"
    "github.com/dazhenghu/util/fileutil"
    "path"
    "path/filepath"
    "html/template"
)

var App *ginApp.GinApp

func init()  {
    ginAppInstance := ginApp.Instance()
    // 保存全局app实例
    App = ginAppInstance
    // 获取挡墙文件夹
    currPath, _ := fileutil.GetCurrentDirectory()
    // 设置common文件夹位置
    App.SetCommonPath(path.Join(currPath, "../common"))
    // 设置根目录文件夹位置
    App.SetRootPath(currPath)
    // 读取默认位置的配置文件
    App.DefaultLoadConfig("")
    // 初始化session配置
    App.InitSession()

    // static参数：relativePath，表示模板中的设定的读取资源文件的路径;root，表示真实的静态文件物理路径。
    // 相当于url中relativePath映射到root
    App.Engine().Static(App.AppConfig.Pulic.RelativePath, App.AppConfig.Pulic.Root)
    // 加载html
    htmlDir := filepath.Join(currPath, App.AppConfig.ViewBaseDir)

    App.Engine().SetFuncMap(template.FuncMap{
        "unescape2": func(htmlStr string) template.HTML {
            return template.HTML(htmlStr)
        },
    })
    App.Engine().LoadHTMLGlob(htmlDir)
}
