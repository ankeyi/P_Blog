package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"math/rand"
	"os"
	"os/exec"
	"reflect"

	"github.com/gin-gonic/gin"
)

// 验证用户具体信息，出错次数过多封禁该IP
func isAuthUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c_userInfo, _ := ctx.Cookie("userInfo")
		c_name, _ := ctx.Cookie("userName")
		myClaims, err := ParseToken(c_userInfo)
		if err != nil {
			ctx.HTML(500, "main/goto.html", returnAlert{"认证失败，请登录", "/sign_in.he"})
			return
		}
		file, _ := os.ReadFile("json/users/private/" + c_name + ".json")
		u_Info := user_info{}
		json.Unmarshal(file, &u_Info)

		if !reflect.DeepEqual(u_Info, myClaims.user_info) {
			ctx.HTML(500, "main/goto.html", returnAlert{"认证失败,请重新登录", "/login.he"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

// 获取文章，文章需要自带后缀名
func get_Article(authority, user_file_name string) []article {
	file, err := os.ReadFile("json/articles/" + authority + "/" + user_file_name)
	if err != nil {
		os.Create("json/articles/" + authority + "/" + user_file_name)
	}
	articles := []article{}
	json.Unmarshal(file, &articles)

	return articles
}

// 加载
func loadNoParHtml(webPath, filePath string) {
	router.GET(webPath, func(ctx *gin.Context) {
		// ctx.SetCookie("gin_cookie", "testvalue", 3600, "/", "localhost", false, true)
		ctx.HTML(200, filePath, ctx.ClientIP())
	})
}

func getMd5Hash(key, data string) string {
	hash := hmac.New(md5.New, []byte(key)) //创建对应的md5哈希加密算法
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum([]byte("")))

}

// Shell
const bash = "/usr/bin/bash"

func Shell(command string) (out string, outerr string, errors error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(bash, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}

// 接受用户文件名
func findUser(user string) bool {

	file, err := os.ReadDir("json/users/private/")
	if err != nil {
		if os.Mkdir("json/users/private/", 0700); err != nil {
			return false
		}
	}
	for _, f := range file {
		if user+".json" == f.Name() {
			return true
		}
	}

	return false
}

// 随机生成字符串
var defaultLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandomString(n int, allowedChars ...[]rune) string {
	var letters []rune

	if len(allowedChars) == 0 {
		letters = defaultLetters
	} else {
		letters = allowedChars[0]
	}

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}
