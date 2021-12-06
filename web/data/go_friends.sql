drop table if exists go_friends;
create table go_friends(
id int(11) primary key auto_increment,
user_id int(11) not null default 0,
friends_id int(11) not null default 0,
friends_group int(11) not null default 0,
key friends_key(user_id)
)engine=innodb default charset=utf8;  
