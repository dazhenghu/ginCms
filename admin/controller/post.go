package controller

import (
    "github.com/dazhenghu/ginApp/controller"
    "github.com/gin-gonic/gin"
    "net/http"
    "github.com/dazhenghu/ginCms/common/service"
    "github.com/dazhenghu/ginApp/session"
    "github.com/gin-contrib/sessions"
    "github.com/dazhenghu/ginCms/common/consts"
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
                "message":"令牌已过期，请刷新重试",
            })
            return
        }

        postToken, _ := session.GenerateSessionToken(context, service.POST_TOKEN_KEY)

        if postId != "" {
            // 更新文章
            updateErr := service.Post.UpdateFromForm(context)
            if updateErr != nil {
                context.JSON(http.StatusOK, map[string]string{
                    "code":consts.ERROR,
                    "message":updateErr.Error(),
                    "token":postToken,
                })
            } else {
                context.JSON(http.StatusOK, map[string]string{
                    "code":consts.SUCCESS,
                    "message":"更新成功",
                    "token":postToken,
                })
            }
        } else {
            // 新增文章
            ok := service.Post.AddPost(
                context.PostForm("post_title"),
                context.PostForm("post_key"),
                context.PostForm("post_content"),
            )

            if ok {
                context.JSON(http.StatusOK, map[string]string {
                    "code":consts.SUCCESS,
                    "message":"成功",
                    "token":postToken,
                })
            } else {
                context.JSON(http.StatusOK, map[string]string{
                    "code":consts.ERROR,
                    "message":"失败",
                    "token":postToken,
                })
            }
        }
    } else {
        postId, _ := context.GetQuery("post_id")
        postToken, _ := session.GenerateSessionToken(context, service.POST_TOKEN_KEY)
        if postId != "" {
            post, _ := service.Post.FindPostById(postId)
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