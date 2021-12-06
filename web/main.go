package main

import (
	"github.com/gogf/gf/frame/g"
	_ "web/boot"
	_ "web/router"
)

// @title       `gf-demo`示例服务API
// @version     1.0
// @description `GoFrame`基础开发框架示例服务API接口文档。
// @schemes     http
func main() {
	g.Log().Debug("[default]Debug")
	g.Log().Info("[default]Info")
	g.Log().Warning("[default]Warning")
	g.Log().Error("[default]Error")
	g.Server().Run()
}
