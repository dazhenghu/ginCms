package main

import (
    "github.com/dazhenghu/migrate"
    "flag"
    _ "github.com/go-sql-driver/mysql"
    "github.com/go-yaml/yaml"
    "github.com/dazhenghu/ginApp/config"
    "io/ioutil"
    "fmt"
)

type DbConfigs struct {
    Dblist map[string]config.DbConfg `yaml:"dblist"`
} 

// migration所在文件夹
var migrationPath string
var dbconfigs *DbConfigs

func init()  {
    migrationPath = "./migration"

    // 读取配置文件
    configFile, err := ioutil.ReadFile("./common/config/main-local.yaml")
    if err != nil {
        panic(fmt.Sprintf("migrate init err:%+v\n", err))
    }

    dbconfigs = &DbConfigs{}
    err = yaml.Unmarshal(configFile, dbconfigs)
    if err != nil {
        panic(err)
    }
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

    migrateDbConfs := make(map[string]*migrate.DbConf)

    for key, dbconfig := range dbconfigs.Dblist {
        migrateDbConfs[key] = &migrate.DbConf{
            Type:dbconfig.Type,
            Dsn:dbconfig.Dsn,
        }
    }

    // 创建执行migrate操作的对象，第二个参数显示指定migration文件所在路径
    migrateObj := migrate.New(migrationPath, migrateDbConfs)
    // 初始化，如果没有migration_log表，将会创建
    migrateObj.InitSelf()

    // 此句是执行migration文件中up所指定的sql
    migrateObj.ExecUp()
}
