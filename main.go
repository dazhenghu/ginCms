package main

import (
	"github.com/gin-gonic/gin"
	"context"
	"flag"
	"fmt"
	"github.com/dazhenghu/migrate"
	"gopkg.in/yaml.v2"
	"container/list"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/dazhenghu/ginCms/admin/controller"
)

var DB = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()
	r.Use()

	flag.Parse()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
		c.Handler()
		c.HandlerName()
		c.IsWebsocket()
		c.MultipartForm()
	})

	var tmp context.Context
	tmp = context.Background()
	tmp.Done()

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := DB[user]
		if ok {
			c.JSON(200, gin.H{"user": user, "value": value})
		} else {
			c.JSON(200, gin.H{"user": user, "status": "no value"})
		}
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			DB[user] = json.Value
			c.JSON(200, gin.H{"status": "ok"})
		}
	})

	return r
}

type T struct {
	blog    string
	Authors []string `yaml:"best_authors,flow"`
	Desc    struct {
		Counter int   `yaml:"Counter"`
		Plist   []int `yaml:",flow"`
	}
}

func main() {
	list.New()
	//sqls := migrate.MigrateSqls{
	//	UpList: make([]string, 0 , 10),
	//	DownList: make([]string, 0 , 10),
	//}
	//sqls.DownList = append(sqls.DownList, "asd")
	//str, _ := yaml.Marshal(&sqls)
	//fmt.Printf("--- t dump:\n%s\n\n", string(str))


	t := T{}
	//修改struct里面的记录
	//t.Blog = "this is Blog"
	t.Authors = append(t.Authors, "myself")
	t.Desc.Counter = 99
	fmt.Printf("--- t:\n%v\n\n", t)
	//转换成yaml字符串类型
	d, err := yaml.Marshal(&t)
	if err != nil {
	}
	fmt.Printf("--- t dump:\n%s\n\n", string(d))

	db, err := gorm.Open("mysql", "root:qsqfrms@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	fmt.Printf("db:%+v\n", db)

	flag.String("d", "", "dd")

	flag.Parse()
	cmd := flag.Arg(0)
	fmt.Println("cmd:"+cmd)



	//migrateInstance := migrate.New(db, "./migration")
	//err = migrateInstance.ExecUp()
	//fmt.Printf("err:%+v", err)

	//r := setupRouter()
	//// Listen and Server in 0.0.0.0:8080
	//r.Run(":8080")
}
