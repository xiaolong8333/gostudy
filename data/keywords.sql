drop table if exists keywords;
create table keywords(
id int(11) primary key auto_increment,
keywords varchar(500) not null default '',
unique key (keywords)
)engine=innodb default charset=utf8; 