package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Jwt初始化
func init() {
	_, err := os.ReadFile("key")
	if err != nil {
		time := time.Now()
		os.WriteFile("key", []byte(time.String()), 0400)
	}
}

// 初始化文件夹
func init() {
	os.Mkdir("json/", 0700)
	os.Mkdir("json/articles/", 0700)
	os.Mkdir("json/users/", 0700)

	os.Mkdir("static/touxiang", 0700)

	os.Mkdir("json/articles/public/", 0700)
	os.Mkdir("json/articles/private/", 0700)
	os.Mkdir("json/articles/protected", 0700)

	os.Mkdir("json/users/private/", 0700)

	os.Mkdir("files", 0700)

}

var router = gin.Default()

func main() {

	// analysis(解析) command parament
	var port int
	flag.IntVar(&port, "p", 8080, "指定服务器端口\t示例:\n ./p_blog -p 8080\n")
	flag.Parse()

	// disable debug
	// gin.SetMode(gin.ReleaseMode)
	if port == 80 || port == 443 {
		gin.SetMode(gin.ReleaseMode)
	}

	// no display files list
	router.Static("/static", "static")

	// display files list
	// router.StaticFS("/files", http.Dir("files/"))

	// load templates/*/* all templates
	router.LoadHTMLGlob("templates/**/*.html")

	// index page
	router.GET("/", indexCheckUser(), isAuthUser(), Index())

	// sign_up.he page
	router.GET("/sign_up.he", func(ctx *gin.Context) {
		ctx.HTML(200, "main/sign_up.html", returnAlert{"", ctx.ClientIP()})
	})

	admin := router.Group("/admin", isAuthUser())
	{
		admin.GET("/lianai.he", func(ctx *gin.Context) {
			u_c := user{}
			file, err := os.ReadFile("json/users/default.json")
			if err != nil {
				ctx.HTML(200, "main/goto.html", returnAlert{"", "/"})
				return

			}

			json.Unmarshal(file, &u_c)

			u_c.Url = "/lianai.he"

			ctx.HTML(200, "admin/nav.html", u_c)

		})

		admin.GET("/vpn.he", func(ctx *gin.Context) {

			user := user{}
			user.IP = ctx.ClientIP()
			user.Url = "/vpn.he"
			ctx.HTML(200, "admin/nav.html", user)
		})

		admin.POST("/allow_IP", isAuthUser(), allowIP())

	}

	// About page
	router.GET("/about.he", func(ctx *gin.Context) {
		u_c := user{}
		file, err := os.ReadFile("json/users/default.json")
		if err != nil {
			ctx.HTML(200, "main/admin.html", nil)

		}

		json.Unmarshal(file, &u_c)

		u_c.Url = "/about.he"

		ctx.HTML(200, "main/about.html", u_c)
	})

	// album Page
	album()

	// allow Files
	router.GET("/files", isAuthUser(), func(ctx *gin.Context) {
		files, _ := os.ReadDir("files")
		var fs []string
		for _, file := range files {
			f := file.Name()
			fs = append(fs, f)
		}
		ctx.HTML(200, "main/files", fs)
	})

	router.POST("/sign_up", sign_up())

	//sign_in.he page
	router.POST("/sign_in", sign_in())

	// router.POST("/add_article", isAuthUser(), add_article_1())
	loadNoParHtml("/sign_in.he", "main/sign_in.html")

	loadNoParHtml("/lianai_time.he", "main/lianai_time.html")

	// add_article.he page
	router.GET("/add_article.he", func(ctx *gin.Context) {
		ctx.HTML(200, "main/add_article.html", nil)
	})

	// add_article.he handlers
	router.POST("/add_article", isAuthUser(), add_article())

	// start
	err := router.Run(":" + strconv.Itoa(port))
	if err != nil {
		if port <= 1023 {
			fmt.Println("使用周知端口需要root权限,Windows需要以管理员权限运行")
		} else {
			fmt.Println(strconv.Itoa(port) + "端口已被占用，请更换端口")
		}

	}

}
