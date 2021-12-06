package api

import (
	"crypto/rand"
	"fmt"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtime"
	"log"
	"time"
	"web/app/model"
)

var (
	// 用户API管理对象
	User = new(userApi)
)

type userApi struct{}

func (a *userApi) Login(r *ghttp.Request) {
	var request *model.UserApiSignInReq
	if err := r.Parse(&request); err != nil {
		r.Response.WriteJsonPExit(g.Map{
			"err": err.Error(),
		})
	}
	passwordMd5, errPwd := gmd5.Encrypt(request.Password)
	if errPwd != nil {
		glog.Error("md5加密异常", errPwd)
		r.Response.WriteJsonPExit(g.Map{
			"err": errPwd,
		})
	}
	one, _ := g.DB().GetOne("select * from go_users where name='" + request.Name + "' && password='" + passwordMd5 + "'")
	if one == nil {
		r.Response.WriteJsonPExit(g.Map{
			"code":    401,
			"message": "用户名或密码错误！",
		})
	}
	token := getToken()
	_, errSet := g.Redis().Do("HSET", token, "id", one["id"])
	_, errSet = g.Redis().Do("HSET", token, "name", one["name"])
	if errSet != nil {
		panic(errSet)
	}
	r.Response.WriteJsonPExit(g.Map{
		"code":    200,
		"message": "登录成功",
		"data":    one,
		"token":   token,
	})
}

func (a *userApi) Register(r *ghttp.Request) {
	var request *model.UserApiSignUpReq
	if err := r.Parse(&request); err != nil {
		r.Response.WriteJsonPExit(g.Map{
			"err": err.Error(),
		})
	}
	//验证用户名
	one, _ := g.DB().GetOne("select * from go_users where name='" + request.Name + "'")
	if one != nil {
		r.Response.WriteJsonPExit(g.Map{
			"code":    401,
			"message": "用户名已存在！",
		})
	}
	email, _ := g.DB().GetOne("select * from go_users where email='" + request.Email + "'")
	if email != nil {
		r.Response.WriteJsonPExit(g.Map{
			"code":    401,
			"message": "用户名已存在！",
		})
	}

	//加密 密码
	passwordMd5, errPwd := gmd5.Encrypt(request.Password)
	if errPwd != nil {
		glog.Error("md5加密异常", errPwd)
		r.Response.WriteJsonPExit(g.Map{
			"err": errPwd,
		})
	}
	//写入数据库
	_, err := g.DB().Insert("go_users", g.Map{"name": request.Name, "password": passwordMd5, "email": request.Email, "role": request.Role})
	if err != nil {
		r.Response.WriteJsonPExit(g.Map{
			"err": err,
		})
	}
	r.Response.WriteJsonPExit(g.Map{
		"code":     200,
		"message":  "注册成功",
		"username": request.Name,
	})
}

//发送邮件忘记密码
func (a *userApi) Forgotpassword(r *ghttp.Request) {
	var (
		request  *model.UserApiCheckEmail
		userInfo *model.UserApiCheckUser
	)
	if err := r.Parse(&request); err != nil {
		r.Response.WriteJsonPExit(g.Map{
			"err": err.Error(),
		})
	}
	user, _ := g.DB().GetOne("select * from go_users where email='" + request.Email + "'")
	if user == nil {
		r.Response.WriteJsonPExit(g.Map{
			"code":    401,
			"message": "邮箱不存在！",
		})
	}
	token := getToken()
	/*	userInfo.Id = user["id"]
		userInfo.Email = user["email"]
		userInfo.Name = user["name"]*/
	_, errSet := g.Redis().Do("HSET", token, "id", userInfo, 1000*time.Millisecond)
	if errSet != nil {
		panic(errSet)
	}
	gtime.SetTimeZone("Asia/Tokyo")
	t1 := gtime.Datetime()
	t2 := gtime.Now().Add(time.Minute * 30).Format("Y-m-d H:i:s")
	r.Response.WriteJsonPExit(g.Map{
		"code":    200,
		"message": "发送邮件成功",
		"data": g.Map{
			"created_at":          t1,
			"email":               request.Email,
			"name":                user["name"],
			"id":                  user["id"],
			"role":                user["role"],
			"resetPasswordExpire": t2,
			"resetPasswordToken":  token,
		}})
}

//重置密码
func (a *userApi) Resetpassword(r *ghttp.Request) {
	var request *model.UserApiCheckPassword
	if err := r.Parse(&request); err != nil {
		r.Response.WriteJsonPExit(g.Map{
			"err": err.Error(),
		})
	}
	user, _ := g.Redis().DoVar("HGET", request.Token, "id")
	if user == nil {
		r.Response.WriteJsonPExit(g.Map{
			"code":    401,
			"message": "token已过期",
		})
	}
	password, errPwd := gmd5.Encrypt(request.Password)
	if errPwd != nil {
		glog.Error("md5加密异常", errPwd)
		r.Response.WriteJsonPExit(g.Map{
			"err": errPwd,
		})
	}
	updateData := g.Map{
		"password": password,
	}

	// UPDATE `article` SET `views`=`views`+1 WHERE `id`=1
	_, err := g.DB().Update("users", updateData, "id", user.String())
	token := getToken()
	if err == nil {
		r.Response.WriteJsonPExit(g.Map{
			"code":    200,
			"message": "重置密码成功",
			"data": g.Map{
				"resetPasswordToken": token,
			}})
	}
}

// 获取token
func getToken() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid
}
