SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;


DROP DATABASE IF EXISTS `machine`;
CREATE DATABASE `machine`;
use `machine`;


DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (  
    `id` bigint AUTO_INCREMENT COMMENT 'id',
    `name` varchar(63) NOT NULL DEFAULT '' COMMENT 'name',
    `avatar` varchar(255) NOT NULL DEFAULT '' COMMENT 'avatar',

    `username` varchar(63) NOT NULL DEFAULT '' COMMENT 'username',
    `phone` varchar(63) NOT NULL DEFAULT '' COMMENT 'phone',
    `openid` varchar(63) NOT NULL DEFAULT '' COMMENT 'openid',
    `unionid` varchar(63) NOT NULL DEFAULT '' COMMENT 'unionid',

    `password` varchar(63) NOT NULL DEFAULT '' COMMENT 'password',

    `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'create time',
    `update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'update time',
    PRIMARY KEY(id)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4 COMMENT='用户表';

alter table user AUTO_INCREMENT=100000;
