package main

// 用户全部信息
type user struct {
	user_info
	user_config
}

// 用户信息
type user_info struct {
	UserName string
	Password string
	Email    string
	IP       string
}

// 用户配置
type user_config struct {
	Head_portrait_path string
	Sign               string
	Url                string
	Articles           []article
}

// 文章信息
type article struct {
	Title   string
	Lable   string
	Content string
}

// Goto page use
type returnAlert struct {
	Alert    string
	GotoPath string
}
