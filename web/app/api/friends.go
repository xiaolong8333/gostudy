package api

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var (
	// 用户API管理对象
	Friends = new(friendsApi)
)

type friendsApi struct{}

func (f *friendsApi) List(r *ghttp.Request) {
	token := r.GetHeader("token")
	userId, _ := g.Redis().DoVar("HGET", token, "id")
	result, errSql := g.DB().
		Table("friends f").
		LeftJoin("users u", "u.id=f.user_id").
		LeftJoin("friends_group g", "g.id=f.friends_group").
		Fields("f.*,u.name,g.group_name").
		Where("f.user_id = ?", userId.Int()).All()
	if errSql == nil {
		panic(errSql)
	}
	r.Response.WriteJsonPExit(g.Map{
		"code":    200,
		"message": "获取成功",
		"data":    result,
	})
}
