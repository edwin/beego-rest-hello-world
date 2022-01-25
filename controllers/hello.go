package controllers

import "github.com/beego/beego/v2/server/web"

type HelloController struct {
	web.Controller
}

type helloResponse struct {
	Greeting string `json:"greeting"`
}

func (u *HelloController) Get() {

	user := helloResponse{}
	user.Greeting = "Hello World Edwin"

	u.Data["json"] = user
	u.ServeJSON()
}