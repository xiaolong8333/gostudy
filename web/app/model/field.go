package model

//获取字段属性列表
type FieldApiGetReq struct {
	FieldTable  string  `v:"required#字段表不能为空"`
}
// 添加字段属性查询
type FieldApiAddReq struct {
	FieldTable  string  `v:"required#字段表不能为空"`
	FieldIsShow string  `v:"required#字段是否显示不能为空"`
	FieldName   string  `v:"required#字段是否显示不能为空"`
	FieldType   string  `v:"required#字段是否显示不能为空"`
}

type Field struct {
	FieldTable  string    `orm:"field_table"`
	FieldName   string    `orm:"field_name"`
	FieldType   string    `orm:"field_type"`
	FieldIsShow string 	  `orm:"field_is_display"`
}
