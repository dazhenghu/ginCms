package main

import (
    "github.com/dazhenghu/ginCms/admin/web" // web包引用要尽量放在前面，因为这个包内会初始化App
    _ "github.com/dazhenghu/ginCms/admin/controller" // 执行controller包中的init方法
)

func main()  {
    Run()
}

func Run()  {
    addr := web.App.AppConfig.Addr
    if addr == "" {
        addr = ":8001"
    }
    web.App.Run(web.App.AppConfig.Addr)
}