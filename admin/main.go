package main

import "github.com/dazhenghu/ginCms/admin/web"

func main()  {
    web.Init()
    web.App.Run(":8001")
}
