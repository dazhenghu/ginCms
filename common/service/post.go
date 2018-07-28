package service

import (
    "github.com/dazhenghu/ginCms/common/model"
    "github.com/gin-gonic/gin"
    "time"
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

/**
按分类查找文章
 */
func (p *post) GetPostList(cateId int) (list []model.Post) {
    if cateId < 1 {
        db.Find(&list)
    } else {
        db.Joins("inner join post on post_id=post_cate_post_post_id").Where("post_cate_post_post_cate_id = ?", cateId).Find(&list)
    }
    return
}

/**
获取分类列表
 */
func (p *post) GetCateList() (cates []model.PostCate)  {
    db.Find(&cates)
    return
}

func (p *post) FindCate(cateId string) (cate *model.PostCate, err error)  {
    cate = &model.PostCate{}
    err = db.Where("post_cate_id = ?", cateId).Find(cate).Error
    return
}

/**
添加分类
 */
func (p *post) AddPostCate(context *gin.Context) bool {
    postcateObj := &model.PostCate{}
    postcateObj.PostCateName = context.PostForm("post_cate_name")
    postcateObj.PostCateDes = context.PostForm("post_cate_des")
    postcateObj.PostCateCreateAt = time.Now()

    db.Create(postcateObj)
    return db.NewRecord(postcateObj)
}

/**
更新类别
 */
func (p *post) UpdatePostCate(context *gin.Context) (err error) {
    postcate, err := p.FindCate(context.PostForm("post_cate_id"))
    postcate.PostCateName = context.PostForm("post_cate_name")
    postcate.PostCateDes = context.PostForm("post_cate_des")
    postcate.PostCateUpdateAt = time.Now()

    err = db.Save(postcate).Error
    return
}
