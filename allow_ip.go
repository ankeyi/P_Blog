package main

import (
	"github.com/gin-gonic/gin"
)

func allowIP() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ip := ctx.PostForm("IP")
		username := ctx.PostForm("username")
		pass := ctx.PostForm("pass")
		allow := ctx.PostForm("allow")
		Shell(username + " pptpd " + pass + " " + ip + " >>  /etc/ppp/chap-secrets")
		if allow == "true" {
			Shell(`firewall-cmd --permanent --add-rich-rule="rule family="ipv4" source address="113.25.0.0/16" accept"`)

		} else if allow == "false" {
			Shell(`firewall-cmd --add-rich-rule="rule family="ipv4" source address="` + ip + `" accept"`)
		}
		ctx.HTML(200, "main/goto.html", returnAlert{"权限已开放，欢迎使用", "/admin/vpn.he"})

	}

}
