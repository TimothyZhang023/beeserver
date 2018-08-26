package controllers

import (
	"github.com/TimothyZhang023/beeserver/core/mipush"
	"github.com/astaxie/beego"
)

type MiPushController struct {
	beego.Controller
}

func (c *MiPushController) Get() {
	releases := mipush.GetReleases()
	c.Data["releases"] = releases
	c.TplName = "mipush.tpl"
}
