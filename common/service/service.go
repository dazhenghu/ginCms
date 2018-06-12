package service

import (
    "github.com/dazhenghu/ginApp/config"
    "github.com/jinzhu/gorm"
    _ "github.com/go-sql-driver/mysql"
    "github.com/dazhenghu/ginCms/admin/web"
    "errors"
)

var dbConfigList map[string]config.DbConfg

var db *gorm.DB

func init()  {
    dbConfigList = web.App.AppConfig.Dblist
    // 获取默认db配置i
    defaultDbConf, ok := dbConfigList["db"]
    if ok {
        defaultDb, err := gorm.Open(defaultDbConf.Type, defaultDbConf.Dsn)
        if err != nil {
            panic(err)
        }
        db = defaultDb
        db.SingularTable(true) // 全局禁用表名复数
    } else {
        panic(errors.New("default db is null"))
    }
}
