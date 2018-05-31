package service

import "github.com/dazhenghu/ginCms/common/model"

const (
    // 菜单类型
    MENU_TYPE_FRONT   string = "front"   // 前台
    MENU_TYPE_BACKEND string = "backend" // 后台
)

type MenuTree struct {
    parent *model.Menu
    children []*model.Menu
}

func (m *MenuTree)GetBackendMenuTree() *MenuTree  {
    return nil
}