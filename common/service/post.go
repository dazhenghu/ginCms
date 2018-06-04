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
    db.Table("post").Where("post_id = ?", postId).Find(post)
    return
}

