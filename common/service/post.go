package service

import (
    "github.com/dazhenghu/ginCms/common/model"
)

type post struct {

}

var Post *post

func init()  {
    Post = &post{}
}

/**
根据postid获取文章信息
 */
func (p *post) FindPostById(postId string) (post *model.Post) {
    post = &model.Post{}
    db.Where("post_id = ?", postId).Find(post)
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

