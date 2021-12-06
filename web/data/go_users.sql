drop table if exists go_users;
create table go_users(
id int(11) primary key auto_increment,
name varchar(255) not null default '',
email varchar(255) not null default '',
role char(50) not null default '',
password varchar(255) not null default '',
created_at timestamp NULL DEFAULT NULL,
updated_at timestamp NULL DEFAULT NULL,
)engine=innodb default charset=utf8;