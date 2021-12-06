package api

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/frame/g"
	"web/app/model"
)

var (
	// 用户API管理对象
	Field = new(fieldApi)
)

type fieldApi struct{}

//获取字段属性
func (a *fieldApi)GetFileds(r *ghttp.Request){
	var (
		request  *model.FieldApiGetReq
		page = r.GetInt("page")
		limit = r.GetInt("limit")
		m = g.DB().Model("fields")
	)
	if err := r.Parse(&request); err != nil {
		r.Response.WriteJsonPExit(g.Map{
			"err": err.Error(),
		})
	}
	result,errSql :=m.Where("field_table = ?", request.FieldTable).Limit((page-1)*limit, limit).All()
	total,errSqlCount :=g.DB().Model("fields").Where("field_table = ?", request.FieldTable).Count()
	if errSql != nil && errSqlCount !=nil{
		panic(errSql)
	}
	r.Response.WriteJsonPExit(g.Map{
		"code":   200,
		"message": "获取成功",
		"data":g.Map{
			"total": total,
			"data":result,
		},
	})
}

//添加字段属性
func (a *fieldApi)AddFileds(r *ghttp.Request) {
	var (
		request *model.FieldApiAddReq
		m = g.DB().Model("fields")
	)
	if err := r.Parse(&request); err != nil {
		r.Response.WriteJsonPExit(g.Map{
			"err": err.Error(),
		})
	}
	field := &model.Field{
		FieldTable:  request.FieldTable,
		FieldName: request.FieldName,
		FieldType: request.FieldType,
		FieldIsShow: request.FieldIsShow,
	}
	result,errSql := m.Data(field).Insert()
	if errSql != nil {
		r.Response.WriteJsonPExit(g.Map{
			"err": errSql.Error(),
		})
	}
	r.Response.WriteJsonPExit(g.Map{
		"code":    200,
		"message": "添加成功",
		"data":    result,
	})
}

//修改字段属性
func (a *fieldApi)EditFileds(r *ghttp.Request){

}


