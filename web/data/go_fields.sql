drop table if exists go_fields;
create table go_fields(
id int(11) primary key auto_increment,
field_table char(50) not null default '' comment '字段表',
field_name char(50) not null default '' comment '字段名',
field_type char(50) not null default '' comment '字段类型',
field_length middleint(5) not null default 0 comment '字段长度',
field_sort tinyint(2) not null default 0 comment '字段排序',
field_form_type char(50) not null default '' comment '字段表单类型',
field_search_type char(50) not null default '' comment '字段查询类型',
field_setting text comment '字段查询类型',
field_validation varchar(500) not null default '' comment '字段提交验证类型',
field_is_display tinyint(1) not null default 1 comment '字段是否展示',
field_can_delete tinyint(1) not null default 0 comment '字段是否可删除',
field_associate_table varchar(255) not null default '' comment '字段关联其他表',
unique key field_table_name(field_table,field_name)
)engine=innodb default charset=utf8;


  
