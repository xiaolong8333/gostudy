drop table if exists stocksdetail;

create table stocksdetail(
id int(11) primary key auto_increment,
stockcode  varchar(255) not null default '',
v1 varchar(255) not null default '' comment '最高',
v2 varchar(255) not null default ''comment '今开',
v3 varchar(255) not null default '' comment '涨停',
v4 varchar(255) not null default '' comment '成交量',
v5 varchar(255) not null default '' comment '最低',
v6 varchar(255) not null default '' comment '昨收',
v7 varchar(255) not null default '' comment '跌停',
v8 varchar(255) not null default '' comment '成交额',
v9 varchar(255) not null default '' comment '量比',
v10 varchar(255) not null default '' comment '换手',
v11 varchar(255) not null default '' comment '市盈率(动)',
v12 varchar(255) not null default '' comment '市盈率(TTM)',
v13 varchar(255) not null default '' comment '委比',
v14 varchar(255) not null default '' comment '振幅',
v15 varchar(255) not null default '' comment '市盈率(静)',
v16 varchar(255) not null default '' comment '市净率',
v17 varchar(255) not null default '' comment '每股收益',
v18 varchar(255) not null default '' comment '股息(TTM)',
v19 varchar(255) not null default '' comment '总股本',
v20 varchar(255) not null default '' comment '总市值',
v21 varchar(255) not null default '' comment '每股净资产',
v22 varchar(255) not null default '' comment '股息率(TTM)',
v23 varchar(255) not null default '' comment '流通股',
v24 varchar(255) not null default '' comment '流通值',
v25 varchar(255) not null default '' comment '52周最高',
v26 varchar(255) not null default '' comment '52周最低',
v27 varchar(255) not null default '' comment '货币单位',
unique key (stockcode),
key (v11),
key (v12),
key (v15),
key (v16)
)engine=innodb default charset=utf8; 
