package main

import (
	"os"
	"fmt"
	"time"
	"net/http"
	"encoding/json"
	"github.com/astaxie/goredis"
	"github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
)

const (
	DB_RDS_HOST, DB_RDS_PORT, DB_RDS_NUM = "localhost", ":6379", 0
	SVR_PORT = ":2333"
	CONF_SH = "conf.json"
)

var _rds *redis.Client
 

func main() {
	_rds = redis.NewClient(&redis.Options{
		Addr:	DB_RDS_HOST + DB_RDS_PORT,
		Password: "",
		DB: DB_RDS_NUM
	})

	pong, err := _rds.Ping().Result()
	fmt.Println(pong, err)

	router := gin.Default()

	router.HTMLRender = gintemplate.New(gintemplate.TemplateConfig{
		Root:      "views",
		Extension: ".tpl",
		Master:    "layouts/master",
		Partials:  []string{"partials/ad"},
		Funcs: template.FuncMap{
			"sub": func(a, b int) int {
				return a - b
			},
			"copy": func() string{
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


	router.Run(SVR_PORT)
}