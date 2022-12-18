package main

import (
	"encoding/json"
	"os"

	"github.com/gin-gonic/gin"
)

// Add Article
func add_article() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		title := ctx.PostForm("Title")
		lable := ctx.PostForm("Lable")
		content := ctx.PostForm("Content")
		userName, _ := ctx.Cookie("userName")
		userName = userName + ".json"
		newArticle := article{
			Title:   title,
			Lable:   lable,
			Content: content,
		}

		if title == "" || lable == "" || content == "" {
			ctx.HTML(500, "main/goto.html", returnAlert{"内容缺失，请补全", "/add_article.he"})
			return
		}

		if newArticle.Lable == "public" {

			article := get_Article("public", "public.json")
			article = append(article, newArticle)
			p1, _ := json.MarshalIndent(article, "", "	")
			wfile, _ := os.Create("json/articles/public/public.json")
			wfile.WriteString(string(p1))
			defer wfile.Close()
			ctx.HTML(200, "main/goto.html", returnAlert{"新增文章成功", "/"})
			return
		}
		if newArticle.Lable == "private" {

			article := get_Article("private", userName)

			article = append(article, newArticle)
			p1, _ := json.MarshalIndent(article, "", "	")
			wfile, _ := os.Create("json/articles/private/" + userName)
			wfile.WriteString(string(p1))
			defer wfile.Close()
			ctx.HTML(200, "main/goto.html", returnAlert{"新增文章成功", "/"})
			return
		}
		if newArticle.Lable == "protected" {
			article := get_Article("protected", userName)

			article = append(article, newArticle)
			p1, _ := json.MarshalIndent(article, "", "	")
			wfile, _ := os.Create("json/articles/protected/" + userName)
			wfile.WriteString(string(p1))
			defer wfile.Close()
			ctx.HTML(200, "main/goto.html", returnAlert{"新增文章成功", "/"})
			return

		}
		ctx.HTML(500, "main/goto.html", returnAlert{"你干嘛哎呦", "/article.he"})

	}
}
