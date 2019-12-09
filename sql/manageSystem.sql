CREATE DATABASE `longmarch01_managesystem` /*!40100 DEFAULT CHARACTER SET utf8 */

CREATE TABLE `user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8 NOT NULL COMMENT '用户名称',
  `password` varchar(255) CHARACTER SET utf8 NOT NULL COMMENT '用户密码',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '0-未知,1-正常,2-冻结',
  `role` tinyint(4) NOT NULL DEFAULT '0' COMMENT '0-未知,1-管理员,2-普通用户',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='用户表'

CREATE TABLE `goods` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8 NOT NULL COMMENT '商品名称',
  `price` int(11) NOT NULL DEFAULT '0' COMMENT '商品价格，单位：分',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '0-未知,1-正常,2-冻结',
  `stock` int(11) NOT NULL DEFAULT '0' COMMENT '商品库存',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='商品表'