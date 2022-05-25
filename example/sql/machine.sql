CREATE TABLE `user` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
  `name` varchar(63) NOT NULL DEFAULT '' COMMENT 'name',
  `avatar` varchar(255) NOT NULL DEFAULT '' COMMENT 'avatar',
  `username` varchar(63) NOT NULL DEFAULT '' COMMENT 'username',
  `phone` varchar(63) NOT NULL DEFAULT '' COMMENT 'phone',
  `password` varchar(63) NOT NULL DEFAULT '' COMMENT 'password',
  `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'create time',
  `update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'update time',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=100015 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户表';


CREATE TABLE `group` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
  `name` varchar(63) NOT NULL DEFAULT '' COMMENT 'name',
  `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'create time',
  `update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'update time',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='群组表';
