package model

import "time"
import "github.com/dazhenghu/ginApp/model"

const (
    POST_STATUS_EXAMING int64 = 0 // 审核中
    POST_STATUS_PASS    int64 = 1 // 审核通过
    POST_STATUS_REJECT  int64 = 2 // 审核拒绝
)

type Post struct {
    model.BaseModel
    PostId int64 `gorm:"column:post_id" json:"post_id" form:"post_id"`
    PostTitle string `gorm:"column:post_title" json:"post_title" form:"post_title"`
    PostKey string `gorm:"column:post_key" json:"post_key" form:"post_key"`
    PostContent string `gorm:"column:post_content" json:"post_content" form:"post_content"`
    PostStatus int64 `gorm:"column:post_status" json:"post_status" form:"post_status"`
    PostStatusDes string `gorm:"column:post_status_des" json:"post_status_des" form:"post_status_des"`
    PostShowTimes int64 `gorm:"column:post_show_times" json:"post_show_times" form:"post_show_times"`
    PostLikeTimes int64 `gorm:"column:post_like_times" json:"post_like_times" form:"post_like_times"`
    PostCreateUserId int64 `gorm:"column:post_create_user_id" json:"post_create_user_id" form:"post_create_user_id"`
    PostCreateUserName string `gorm:"column:post_create_user_name" json:"post_create_user_name" form:"post_create_user_name"`
    PostUpdateUserId int64 `gorm:"column:post_update_user_id" json:"post_update_user_id" form:"post_update_user_id"`
    PostUpdateUserName string `gorm:"column:post_update_user_name" json:"post_update_user_name" form:"post_update_user_name"`
    PostCreateAt time.Time `gorm:"column:post_create_at" json:"post_create_at" form:"post_create_at"`
    PostUpdateAt time.Time `gorm:"column:post_update_at" json:"post_update_at" form:"post_update_at"`
}
