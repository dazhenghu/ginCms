package main

import (
    "github.com/dazhenghu/migrate"
    "flag"
    "github.com/jinzhu/gorm"
    _ "github.com/go-sql-driver/mysql"
)

var migrationPath string

func init()  {
    migrationPath = "./migration"
}

func main()  {
    flag.Parse()
    cmd := flag.Arg(0)

    if cmd == "create" {
        createMigrateFile()
    } else if cmd == "up" {
        up()
    }
}

func createMigrateFile()  {
    migrate.CreateMigrationFile(migrationPath)
}

func up()  {
    // 初始化数据库链接，执行sql的时候会用到
    db, err := gorm.Open("mysql", "root:qsqfrms@tcp(127.0.0.1:3306)/gincms?charset=utf8mb4&parseTime=True&loc=Local")
    if err != nil {
        panic(err)
    }
    // 释放链接
    defer db.Close()

    migrateObj := migrate.New(db, migrationPath)
    // 初始化，如果没有migration_log表，将会创建
    migrateObj.InitSelf()
    // 创建执行migrate操作的对象，第二个参数显示指定migration文件所在路径

    // 此句是执行migration文件中up所指定的sql
    err = migrateObj.ExecUp()
}
