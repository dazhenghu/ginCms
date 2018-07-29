package controller

import (
    "github.com/dazhenghu/ginApp/controller"
    "github.com/gin-gonic/gin"
    "github.com/dazhenghu/ginCms/common/service"
    "net/http"
    "github.com/dazhenghu/ginApp/session"
    "github.com/dazhenghu/ginCms/common/consts"
    "github.com/gin-contrib/sessions"
)

type postcateController struct {
    controller.Controller
} 

var postcateInstance *postcateController

func init()  {
    postcateInstance = &postcateController{}
    postcateInstance.Init(postcateInstance)

    postcateInstance.Get("/postcate", postcateInstance.index)
    postcateInstance.PostAndGet("/postcate/save", postcateInstance.save)
}

func (pc *postcateController) index(context *gin.Context)  {
    cateId := context.GetInt("cate_id")
    if cateId < 1 {
        cateList := service.Post.GetCateList()
        context.HTML(http.StatusOK, "postcate/index.html", gin.H{
            "pageTitle": "分类列表",
            "cateList": cateList,
        })
    }
}

func (pc *postcateController) save(context *gin.Context) {
    if context.Request.Method == http.MethodPost {
        postcateId := context.PostForm("post_cate_id")
        token := context.PostForm("token")
        sessions.Default(context)
        tokenErr := session.CheckSessionToken(context, consts.SESSION_KEY_POST_TOKEN, token)
        if tokenErr != nil {
            context.JSON(http.StatusOK, map[string]string {
                "code":consts.ERROR,
                "message":"令牌已过期，请刷新重试",
            })
            return
        }

        token, _ = session.GenerateSessionToken(context, consts.SESSION_KEY_POST_TOKEN)

        if postcateId != "" {
            // 更新文章
            updateErr := service.Post.UpdatePostCate(context)
            if updateErr != nil {
                context.JSON(http.StatusOK, map[string]string{
                    "code":consts.ERROR,
                    "message":updateErr.Error(),
                    "token":token,
                })
            } else {
                context.JSON(http.StatusOK, map[string]string{
                    "code":consts.SUCCESS,
                    "message":"更新成功",
                    "token":token,
                })
            }
        } else {
            // 新增文章
            ok := service.Post.AddPostCate(context)

            if ok {
                context.JSON(http.StatusOK, map[string]string {
                    "code":consts.SUCCESS,
                    "message":"成功",
                    "token":token,
                })
            } else {
                context.JSON(http.StatusOK, map[string]string{
                    "code":consts.ERROR,
                    "message":"失败",
                    "token":token,
                })
            }
        }

    } else {
        cateId, _ := context.GetQuery("post_cate_id")
        token, _ := session.GenerateSessionToken(context, consts.SESSION_KEY_POST_TOKEN)
        if cateId == "0" || cateId == "" {
            // 没有类别id，说明是新增
            context.HTML(http.StatusOK, "postcate/save.html", gin.H{
                "pageTitle": "添加分类",
                "token": token,
            })
        } else {
            cate, _ := service.Post.FindCate(cateId)
            context.HTML(http.StatusOK, "postcate/save.html", gin.H{
                "pageTitle": "修改分类",
                "postcate": cate,
                "token": token,
            })
        }
    }
}
