package model

import "time"
import "github.com/dazhenghu/ginApp/model"

type PostCate struct {
    model.BaseModel
    PostCateId int64 `gorm:"primary_key;column:post_cate_id" json:"post_cate_id" form:"post_cate_id"`
    PostCateName string `gorm:"column:post_cate_name" json:"post_cate_name" form:"post_cate_name"`
    PostCateDes string `gorm:"column:post_cate_des" json:"post_cate_des" form:"post_cate_des"`
    PostCateCreateAt time.Time `gorm:"column:post_cate_create_at" json:"post_cate_create_at" form:"post_cate_create_at"`
    PostCateUpdateAt time.Time `gorm:"column:post_cate_update_at" json:"post_cate_update_at" form:"post_cate_update_at"`
}