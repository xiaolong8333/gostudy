drop table if exists go_articles;
create table go_articles(
id int(11) primary key auto_increment,
artcle_category tinyint(2) not null default 0,
article_title varchar(255) not null default '',
author_id  int(11) not null default 0,
article_content text,
created_at int(13) not null default 0,
updated_at int(13) not null default 0,
key article_key(article_title)
)engine=innodb default charset=utf8;