drop table if exists stocks;

CREATE TABLE `stocks` (
  `id` int NOT NULL AUTO_INCREMENT,
  `companynamec` varchar(255) CHARACTER SET utf8 NOT NULL DEFAULT '' COMMENT '公司名称',
  `companynamee` varchar(255) CHARACTER SET utf8 NOT NULL DEFAULT '' COMMENT '英文名称',
  `stockcode` varchar(255) CHARACTER SET utf8 NOT NULL DEFAULT '' COMMENT 'A股代码',
  `stockname` varchar(255) CHARACTER SET utf8 NOT NULL DEFAULT '' COMMENT 'A股股票简称',
  `stockcate` varchar(255) CHARACTER SET utf8 NOT NULL DEFAULT '' COMMENT '证券类别',
  `stockindustry` varchar(255) CHARACTER SET utf8 NOT NULL DEFAULT '' COMMENT '所属行业',
  `listedcate` varchar(255) CHARACTER SET utf8 NOT NULL DEFAULT '' COMMENT '上市交易所',
  `regulator` text CHARACTER SET utf8 COMMENT '所属证监会行业',
  PRIMARY KEY (`id`),
  UNIQUE KEY `stockcode` (`stockcode`),
  UNIQUE KEY `stockname` (`stockname`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci  