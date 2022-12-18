package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// 需要增加邮件功能注册

// 注册用户，写入个性签名以及上传头像
// 使用函数不定式
func sign_up() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		// post data
		imgInfo, _ := ctx.FormFile("head_portrait")
		userName := ctx.PostForm("userName")
		password := ctx.PostForm("password")
		sign := ctx.PostForm("sign")
		email := ctx.PostForm("email")

		pass_hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		ext := filepath.Ext(imgInfo.Filename)
		// 验证上传是否是图片
		if ext == ".jpg" || ext == ".jpeg" || ext == ".png" {

		} else {
			ctx.HTML(500, "main/goto.html", returnAlert{"允许上传图片格式:\njpg、png、jpeg", "/sign_up.he"})
			return
		}
		// 应该实现随机名字上传
		hash_img_name := getMd5Hash(time.Now().String(), imgInfo.Filename)

		if _, err := os.ReadDir("static/touxiang/"); err != nil {
			os.Mkdir("static/touxiang/", 0700)
		}
		img_path := "static/touxiang/" + hash_img_name + ext
		ctx.SaveUploadedFile(imgInfo, img_path)
		os.Chmod(img_path, 0600)
		// 结构体存储用户信息
		user := user{
			user_info{
				userName,
				string(pass_hash),
				email,
				""},

			user_config{
				img_path,
				sign,
				"",
				nil,
			},
		}

		var file *os.File
		var err error
		// 创建用户文件
		if _, err := os.ReadFile("json/users/default.json"); err != nil {
			file, _ = os.Create("json/users/default.json")

		} else {
			// 判断是否存在用户
			if findUser(userName) {
				ctx.HTML(500, "main/goto.html", returnAlert{"用户已存在", "/sign_up.he"})
				defer file.Close()
				return
			}

			file, _ = os.Create("json/users/private/" + userName + ".json")
		}

		data, err := json.MarshalIndent(user, "", " ")

		file.WriteString(string(data))
		if err != nil {
			ctx.HTML(500, "main/goto.html", returnAlert{err.Error(), "/sign_up.he"})
			defer file.Close()
			return
		}

		ctx.HTML(200, "main/goto.html", returnAlert{"注册成功，请登录", "/"})
		defer file.Close()

	}

}
