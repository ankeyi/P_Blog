package main

// import (
// 	"encoding/json"
// 	"os"

// 	"github.com/gin-gonic/gin"
// )

// func about() {
// 	router.GET("/about.he", func(ctx *gin.Context) {
// 		u_c := user_config{}
// 		rfile, err := os.ReadFile("json/users/default.json")

// 		if err != nil {
// 			file, _ := os.Create("json/users/default.json")
// 			u_c := user_config{
// 				"安可以为善",
// 				"",
// 				"/static/img/1.jpg",
// 				"如同闪电一般",
// 				"test@gmail.com",
// 				nil,
// 				"",
// 			}
// 			json, _ := json.MarshalIndent(u_c, "", "	")
// 			file.WriteString(string(json))
// 			rfile, _ = os.ReadFile("json/users/default.json")

// 		}

// 		json.Unmarshal(rfile, &u_c)

// 		u_c.Articles = get_Article("public", "public")

// 		// get url 区分nav页面
// 		u_c.Url = ctx.Request.URL.Path

// 		ctx.HTML(200, "main/about.html", u_c)
// 	})
// }
