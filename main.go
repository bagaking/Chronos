package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

const (
	rdsHOST, rdsPORT, rdsNUM = "localhost", ":6379", 0
	svrPORT                  = ":8020"
	confPATH                 = "./conf.json"
)

var _rds *redis.Client

func main() {

	conf := loadConfig(confPATH)
	fmt.Printf("Conf loaded %s\n", conf)
	//run hub
	fmt.Println("Hub started")
	workerhub := &Hub{}
	for _, workerCfg := range conf.Workers {
		workerhub.Insert(workerCfg.Workername, workerCfg.Timespan, workerCfg.Srcpath)
	}
	workerhub.Start()

	//prepare db
	_rds = redis.NewClient(&redis.Options{
		Addr:     rdsHOST + rdsPORT,
		Password: "",
		DB:       rdsNUM,
	})

	pong, err := _rds.Ping().Result()
	fmt.Println(pong, err)

	router := gin.Default()

	router.HTMLRender = gintemplate.New(gintemplate.TemplateConfig{
		Root:      "views",
		Extension: ".tpl",
		Master:    "layout/master",
		Partials:  []string{"ad"},
		Funcs: template.FuncMap{
			"sub": func(a, b int) int {
				return a - b
			},
			"copy": func() string {
				return time.Now().Format("2006")
			},
		},
		DisableCache: true,
	})

	router.GET("/page", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index", gin.H{
			"title": "Index title!",
			"add": func(a int, b int) int {
				return a + b
			},
		})
	})

	router.Run(svrPORT)

}
