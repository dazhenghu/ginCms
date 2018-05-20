package web

import (
    "github.com/dazhenghu/ginApp"
    "sync"
)

var App *ginApp.GinApp
var indexOnce sync.Once

func init()  {
    indexOnce.Do(func() {
        App = ginApp.Run()
    })
}
