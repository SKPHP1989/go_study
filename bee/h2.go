package main

import (
    "github.com/astaxie/beego"
)

type MainController struct {
    beego.Controller
}
type Main2Controller struct {
    beego.Controller
}

func (this *MainController) Get() {
    this.Ctx.WriteString("hello world")
}
func (this *Main2Controller) Get() {
    this.Ctx.WriteString("hello worldasdasd")
}

func main() {
    beego.Router("/", &MainController{})
    beego.Router("/asd", &Main2Controller{})
    beego.Run()
}