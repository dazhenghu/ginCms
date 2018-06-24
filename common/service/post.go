package service

import (
    "github.com/dazhenghu/ginCms/common/model"
    "github.com/gin-gonic/gin"
)

type post struct {

}

var Post *post

func init()  {
    Post = &post{}
}

/**
更新文章
 */
func (p *post) UpdateFromForm(context *gin.Context) (err error) {
    post, err := p.FindPostById(context.PostForm("post_id"))
    post.PostTitle = context.PostForm("post_title")
    post.PostKey = context.PostForm("post_key")
    post.PostContent = context.PostForm("post_content")

    err = db.Save(post).Error
    return
}

/**
根据postid获取文章信息
 */
func (p *post) FindPostById(postId string) (post *model.Post, err error) {
    post = &model.Post{}
    err = db.Where("post_id = ?", postId).Find(post).Error
    return
}

/**
更新数据
 */
func (p *post) UpdatePost(post *model.Post) (err error) {
    err = db.Save(post).Update().Error
    return
}

/**
增加文章
 */
func (p *post) AddPost(postTitle, postKey, postContent string) bool {
    postObj := &model.Post{}
    postObj.PostTitle = postTitle
    postObj.PostKey = postKey
    postObj.PostContent = postContent
    postObj.PostStatus = model.POST_STATUS_PASS
    postObj.PostShowTimes = 1
    postObj.PostLikeTimes = 0
    postObj.PostCreateUserId = 0
    postObj.PostCreateUserName = ""
    postObj.PostUpdateUserId = 0
    postObj.PostUpdateUserName = ""
    db.Create(postObj)
    return db.NewRecord(postObj)
}

