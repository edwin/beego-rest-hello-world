package main

import (
	_ "HelloBeego/routers"
	"github.com/beego/beego/v2/server/web"
)

func main() {
	web.Run()
}