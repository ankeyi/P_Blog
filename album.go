package main

import (
	"encoding/json"
	"os"

	"github.com/gin-gonic/gin"
)

type allImg struct {
	user
	Path []string
}

// Display img page
func album() {
	router.GET("/album.he", func(ctx *gin.Context) {

		u_c := user{}
		var paths []string

		// img_path, _ := os.ReadDir("files/album/")

		img_path, _ := os.ReadDir("static/touxiang/")
		file, err := os.ReadFile("json/users/default.json")
		if err != nil {
			ctx.HTML(200, "main/index.html", nil)

		}

		json.Unmarshal(file, &u_c)

		for i := 0; i < len(img_path)-1; i++ {
			// paths = append(paths, "/files/album/"+img_path[i].Name())
			paths = append(paths, "/static/touxiang/"+img_path[i].Name())

		}
		u_c.Url = "/album.he"
		u_c.IP = ctx.ClientIP()
		q := allImg{u_c, paths}

		ctx.HTML(200, "main/album.html", q)
	})

}
