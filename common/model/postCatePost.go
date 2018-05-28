package model

import "time"
import "github.com/dazhenghu/ginApp/model"

type PostCatePost struct {
    model.BaseModel
    PostCatePostId int64 `gorm:"column:post_cate_post_id" json:"post_cate_post_id" form:"post_cate_post_id"`
    PostCatePostPostId int64 `gorm:"column:post_cate_post_post_id" json:"post_cate_post_post_id" form:"post_cate_post_post_id"`
    PostCatePostPostCateId int64 `gorm:"column:post_cate_post_post_cate_id" json:"post_cate_post_post_cate_id" form:"post_cate_post_post_cate_id"`
    PostCatePostCreateAt time.Time `gorm:"column:post_cate_post_create_at" json:"post_cate_post_create_at" form:"post_cate_post_create_at"`
    PostCatePostUpdateAt time.Time `gorm:"column:post_cate_post_update_at" json:"post_cate_post_update_at" form:"post_cate_post_update_at"`
}