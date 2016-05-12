package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

var ControllerLog = logs.NewLogger(10000)

func init() {
	ControllerLog.SetLogger("console", "")
	ControllerLog.EnableFuncCallDepth(true)
}

type APIController struct {
	beego.Controller
	version interface{}
	err     error
	data    interface{}
}

// 函数结束时,组装成json结果返回
func (this *APIController) Finish() {
	r := struct {
		Version interface{} `json:"version"`
		Error   interface{} `json:"error"`
		Data    interface{} `json:"data"`
	}{}
	r.Version = this.version
	if this.err != nil {
		r.Error = this.err.Error()
	}
	r.Error = ""
	r.Data = this.data
	this.Data["json"] = r
	this.ServeJSON()
}

// 如果请求的参数不存在,就直接 error返回
func (this *APIController) MustString(key string) string {
	v := this.GetString(key)
	if v == "" {
		this.Data["json"] = map[string]string{
			"version": beego.AppConfig.String("version"),
			"error":   fmt.Sprintf("require filed: %s", key),
			"data":    "orz!!",
		}
		this.ServeJSON()
		this.StopRun()
	}
	return v
}

// 其他的函数跟它累似,就不写了
