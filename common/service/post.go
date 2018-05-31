package service

import (
    "github.com/dazhenghu/ginCms/common/model"
    "fmt"
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
    fmt.Printf("db:%+v dbconfiglist:%+v\n", db, dbConfigList)
    db.Where("post_id = ?", postId).Find(post)
    //db.Model(post).Where("post_id = ?", postId).Row()
    return
}

