package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gproc"
	"time"
	"web/app/api"
	"web/app/service"
)

func init() {

	s := g.Server()
	// 分组路由注册方式
	s.Group("/api", func(group *ghttp.RouterGroup) {
		group.Middleware(
			service.Middleware.Ctx,
			service.Middleware.CORS,
		)
		group.ALL("/chat", api.Chat, "MessageSend")
		//登录之后可访问
		group.Group("/", func(group *ghttp.RouterGroup) {
			group.Middleware(service.Middleware.Auth)
			group.POST("/mid", service.Middleware, "Auth")
			group.POST("/user/login", api.User, "Login")
			group.POST("/user/register", api.User, "Register")
			group.POST("/user/test", api.User, "Forgotpassword")
			group.PUT("user/resetpassword", api.User, "Resetpassword")
			group.GET("/fields", api.Field, "GetFileds")
			group.POST("/fields", api.Field, "AddFileds")
			group.GET("/friends", api.Friends, "List")
		})
	})

	//平滑重启
	s.BindHandler("/pid", func(r *ghttp.Request) {
		r.Response.Writeln(gproc.Pid())
	})
	s.BindHandler("/sleep", func(r *ghttp.Request) {
		r.Response.Writeln(gproc.Pid())
		time.Sleep(10 * time.Second)
		r.Response.Writeln(gproc.Pid())
	})
	// 分组路由注册方式
}
