package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

// 通过cookie简易判断是否登录用户
func indexCheckUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 通过url路径判断用户是否登录

		// path := strings.Split(ctx.Request.URL.Path, "/")
		// if path[0] != "user" {
		c_name, _ := ctx.Cookie("userName")
		if !findUser(c_name) {

			user := user{}
			file, err := os.ReadFile("json/users/default.json")
			if err != nil {
				// 如果不存在default用户, 跳转到注册页面
				ctx.HTML(200, "main/sign_up.html", returnAlert{"初始化博客", ctx.ClientIP()})
				ctx.Abort()
				return
			}

			err = json.Unmarshal(file, &user)
			if err != nil {
				ctx.HTML(500, "main/goto.html", returnAlert{"读取默认用户配置失败", "/"})
			}
			user.Url = "/"
			ctx.HTML(200, "main/index.html", user)
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

// 显示用户信息
func Index() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c_userInfo, _ := ctx.Cookie("userName")
		file, _ := os.ReadFile("json/users/private/" + c_userInfo + ".json")
		u_info := user{}
		json.Unmarshal(file, &u_info)

		read, _ := os.ReadFile("json/articles/private/" + c_userInfo + ".json")
		fmt.Println("string", string(read))
		article := []article{}

		err := json.Unmarshal(read, &article)
		if err != nil {
			ctx.HTML(500, "goto.html", returnAlert{err.Error(), "/"})
		}
		u_info.Articles = append(u_info.Articles, article...)

		json.Unmarshal(file, &article)

		u_info.Url = "/"
		ctx.HTML(200, "main/index.html", u_info)

	}
}
