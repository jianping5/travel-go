CREATE DATABASE dravel_user;
USE dravel_user;

-- 用户表
CREATE TABLE `travel_user_info` (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
    `account` varchar(50)   NOT NULL   COMMENT '账户',
    `password`   varchar(50)     NULL    COMMENT '密码',
    `name`  varchar(50) NULL    COMMENT '昵称',
    `avatar`  varchar(255) NULL    COMMENT '头像',
    `signature`  varchar(255) NULL    COMMENT '个性签名',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `is_deleted` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否删除',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '用户表';

CREATE INDEX idx_account ON travel_user_info(`account`);

-- 关注表
CREATE TABLE `travel_user_follow` (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
    `user_id`   bigint  NOT NULL    COMMENT '用户 ID',
    `follow_user_id`   bigint  NOT NULL    COMMENT '关注用户 ID',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `is_deleted` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否删除',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '关注表';

CREATE INDEX idx_user_id ON travel_user_follow(`user_id`);

-- 个人动态表
CREATE TABLE `travel_user_dynamic` (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
    `user_id`   bigint  NOT NULL    COMMENT '用户 ID',
    `content`   text   NULL    COMMENT '动态内容',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `is_deleted` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否删除',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '个人动态表';

CREATE INDEX idx_user_id ON travel_user_dynamic(`user_id`);