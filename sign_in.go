package main

import (
	"encoding/json"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// 登录功能
func sign_in() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		userName := ctx.PostForm("userName")
		password := ctx.PostForm("password")
		ip := ctx.ClientIP()

		user := user{}

		// 没有设置验证码，可能存在爆破行为
		if !findUser(userName) {
			ctx.HTML(200, "main/goto.html", returnAlert{"不存在该用户", "/sign_in.he"})
			return
		}
		path := "json/users/private/" + userName + ".json"
		file, _ := os.ReadFile(path)

		err := json.Unmarshal(file, &user)
		user.IP = ip
		data, _ := json.MarshalIndent(&user, "", " ")
		os.WriteFile(path, data, 0600)
		// 如果解析JSON报错，返回错误
		if err != nil {
			ctx.HTML(200, "main/goto.html", returnAlert{err.Error(), "/sign_in.he"})
			return
		}

		// 比较密码和HASH Value是否相等
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err != nil {
			ctx.HTML(500, "main/goto.html", returnAlert{"密码错误", "/sign_in.he"})
			return
		}

		// 生成Cookie密钥
		tokenString, err := MakeToken(user.user_info)
		if err != nil {
			ctx.HTML(500, "main/goto.html", returnAlert{"身份凭证生成失败", "/sign_in.he"})
			return
		}

		// 设置主域名
		ctx.SetCookie("userName", userName, 3600*24, "/", "localhost", false, true)
		ctx.SetCookie("userInfo", tokenString, 3600*24, "/", "localhost", false, true)

		ctx.HTML(200, "main/goto.html", returnAlert{"登录成功", "/"})
	}

}
