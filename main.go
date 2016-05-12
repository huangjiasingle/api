package main

import (
	_ "api/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.BConfig.CopyRequestBody = true
	beego.Run()
}
