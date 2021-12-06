drop table if exists go_friends_group;
create table go_friends_group(
id int(11) primary key auto_increment,
user_id int(11) not null default 0,
group_name char(50) not null default '',
sort tinyint(2) not null default 0,
key friends_group_key(user_id)
)engine=innodb default charset=utf8;  
