package service

import (
    "github.com/dazhenghu/ginApp/config"
    "github.com/jinzhu/gorm"
    _ "github.com/go-sql-driver/mysql"
)

var dbConfigList map[string]config.DbConfg

var db *gorm.DB

func Init(dbList map[string]config.DbConfg)  {
    dbConfigList = dbList
    // 获取默认db配置
    defaultDbConf, ok := dbConfigList["db"]
    if ok {
        defaultDb, err := gorm.Open(defaultDbConf.Type, defaultDbConf.Dsn)
        if err != nil {
            panic(err)
        }
        db = defaultDb
    }
}
