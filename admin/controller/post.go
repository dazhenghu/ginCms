package controller

import (
    "github.com/dazhenghu/ginApp/controller"
    "github.com/gin-gonic/gin"
    "net/http"
    "github.com/dazhenghu/ginCms/common/service"
    "github.com/dazhenghu/ginApp/session"
    "github.com/gin-contrib/sessions"
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
        postId := context.PostForm("post_id")
        token := context.PostForm("token")
        sessions.Default(context)
        tokenErr := session.CheckSessionToken(context, service.POST_TOKEN_KEY, token)
        if tokenErr != nil {
            context.JSON(http.StatusOK, map[string]string {
                "code":"error",
                "message":tokenErr.Error(),
            })
            return
        }

        if postId != "" {
            // 更新文章
        } else {
            // 新增文章
            ok := service.Post.AddPost(
                context.PostForm("post_title"),
                context.PostForm("post_key"),
                context.PostForm("post_content"),
            )

            if ok {
                context.JSON(http.StatusOK, map[string]string {
                    "code":"success",
                    "message":"成功",
                })
            } else {
                context.JSON(http.StatusOK, map[string]string{
                    "code":"error",
                    "message":"失败",
                })
            }
        }
    } else {
        postId, _ := context.GetQuery("post_id")
        postToken, _ := session.GenerateSessionToken(context, service.POST_TOKEN_KEY)
        if postId != "" {
            post := service.Post.FindPostById(postId)
            // 显示操作页面
            context.HTML(http.StatusOK, "post/save.html", gin.H{
                "pageTitle": "文章修改",
                "post":post,
                "token": postToken,
            })
            return
        }
        context.HTML(http.StatusOK, "post/save.html", gin.H{
            "pageTitle": "文章添加",
            "token": postToken,
        })
    }
}