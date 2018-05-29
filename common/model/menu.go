package model

import "time"

type Menu struct {
    MenuId int64 `gorm:"column:menu_id" json:"menu_id" form:"menu_id"`
    MenuType string `gorm:"column:menu_type" json:"menu_type" form:"menu_type"`
    MenuName string `gorm:"column:menu_name" json:"menu_name" form:"menu_name"`
    MenuParentId int64 `gorm:"column:menu_parent_id" json:"menu_parent_id" form:"menu_parent_id"`
    MenuRoute string `gorm:"column:menu_route" json:"menu_route" form:"menu_route"`
    MenuOrder int64 `gorm:"column:menu_order" json:"menu_order" form:"menu_order"`
    MenuData string `gorm:"column:menu_data" json:"menu_data" form:"menu_data"`
    MenuCreateAt time.Time `gorm:"column:menu_create_at" json:"menu_create_at" form:"menu_create_at"`
    MenuUpdteAt time.Time `gorm:"column:menu_updte_at" json:"menu_updte_at" form:"menu_updte_at"`
}
