package controller

import (
    "github.com/dazhenghu/ginApp/controller"
    "github.com/gin-gonic/gin"
    "net/http"
    "github.com/dazhenghu/ginCms/common/service"
)

type postController struct {
    controller.Controller
}

var postInstance *postController

func init()  {
    postInstance = &postController{}
    postInstance.Init(postInstance)

    postInstance.Get("/post", postInstance.index)
    postInstance.PostAndGet("/post/save", postInstance.save)
    //postInstance.Get("/save", postInstance.index)
}

func (post *postController) index(context *gin.Context)  {
    context.HTML(http.StatusOK, "post/index.html", gin.H{
        "pageTitle": "文章列表",
    })
}

func (post *postController) save(context *gin.Context)  {
    if context.Request.Method == http.MethodPost {
        // 保存或更新文章
    } else {
        postId, _ := context.GetQuery("post_id")
        if postId != "" {
            post := service.Post.FindPostById(postId)
            // 显示操作页面
            context.HTML(http.StatusOK, "post/save.html", gin.H{
                "pageTitle": "文章修改",
                "post":post,
            })
            return
        }
        context.HTML(http.StatusOK, "post/save.html", gin.H{
            "pageTitle": "文章添加",
        })
    }
}