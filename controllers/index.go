package controllers

import (
	"github.com/astaxie/beego"
	"easylog/consts"
	"easylog/models"
)

var json consts.Json

type IndexController struct {
	beego.Controller
}


func (this *IndexController) Index() {
	this.Data["json"] = "ok"
	this.ServeJSON()
}

func (this *IndexController) Log() {
	content := this.GetString("content")
	appname := this.GetString("appname")

	if appname == "" || content == "" {
		this.Data["json"] = json.VendorError(500,"miss appname or content")
		this.ServeJSON()
	}
	go func() {
		log := new(models.Log)

		log.AppName = appname
		log.Content = content
		log.Save()

	}()

	json.SetData([...]string{content,appname})
	this.Data["json"] = json.VendorOk()
	this.ServeJSON()
}