package api

import (
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	//"fmt"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"strings"
	"web/app/model"
)
var (
	// 用户API管理对象
	Chat = new(chatApi)
	UsersList = "user_list"
)

type chatApi struct{}

func (c *chatApi) MessageSend(r *ghttp.Request) {

	// 初始化WebSocket请求
	var (
		msg = &model.ChatMsg{}
		ws  *ghttp.WebSocket
		err error
	)
	ws, err = r.WebSocket()
	if err != nil {
		glog.Error(err)
		r.Exit()
	}
	for{
		// 阻塞读取WS数据
		_, msgByte, err := ws.ReadMessage()
		if err != nil {
			// 如果失败，那么表示断开，这里清除用户信息
			// 为简化演示，这里不实现失败重连机制
			// 通知所有客户端当前用户已下线
			break
		}
		// JSON参数解析
		if err := gjson.DecodeTo(msgByte, msg); err != nil {
			c.write(ws, model.ChatMsg{
				Type: "error",
				Data: "消息格式不正确: " + err.Error(),
				From: "",
			})
			continue
		}
		// 数据校验
		if err := g.Validator().Ctx(r.Context()).CheckStruct(msg); err != nil {
			c.write(ws, model.ChatMsg{
				Type: "error",
				Data: gerror.Current(err).Error(),
				From: "",
			})
			continue
		}
		// WS操作类型
		//ghttp.WS_MSG_TEXT
		switch msg.Type {
			case "send":
				list:=strings.Split(msg.From,"_")
				for _,v :=range list{
					user_fd, _ := g.Redis().DoVar("HGET", UsersList,v)
					msgData, _ := gjson.Encode(msg)
					ws.WriteMessage(user_fd.Int(),msgData)
				}
			case "login":
				c.login(msg.From,ghttp.WS_MSG_TEXT)
		default :

		}
	}
}
// 向客户端写入消息。
// 内部方法不会自动注册到路由中。
func (c *chatApi) write(ws *ghttp.WebSocket, msg model.ChatMsg) error {
	msgBytes, err := gjson.Encode(msg)
	if err != nil {
		return err
	}
	return ws.WriteMessage(ghttp.WS_MSG_TEXT, msgBytes)
}

//用户websocket登录
func (c *chatApi) login(users string,users_fd int)(string,string){
	g.Redis().Do("HSET", UsersList, users, users_fd)
	return "login","success"
}

