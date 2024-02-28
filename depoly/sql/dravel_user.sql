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

-- 收藏表
CREATE TABLE `travel_user_favor` (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
    `user_id`   bigint  NOT NULL    COMMENT '用户 ID',
    `favorite_id`   bigint  NULL    COMMENT '收藏夹 ID',
    `item_type` int NOT NULL    COMMENT '对象类型',
    `item_id`   bigint  NOT NULL    COMMENT '对象 ID',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `is_deleted` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否删除',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '收藏表';

CREATE INDEX idx_user_id ON travel_user_favor(`user_id`);
CREATE INDEX idx_favorite_id ON travel_user_favor(`favorite_id`);
CREATE INDEX idx_item_type ON travel_user_favor(`item_type`);

-- 收藏夹表
CREATE TABLE `travel_user_favorite` (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
    `user_id`   bigint  NOT NULL    COMMENT '用户 ID',
    `name`   varchar(50)  NOT NULL    COMMENT '收藏夹名称',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `is_deleted` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否删除',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '收藏夹表';

CREATE INDEX idx_user_id ON travel_user_favorite(`user_id`);

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

-- 点赞表
CREATE TABLE `travel_user_like` (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
    `user_id`   bigint  NOT NULL    COMMENT '用户 ID',
    `item_type` int NOT NULL    COMMENT '对象类型',
    `item_id`   bigint  NOT NULL    COMMENT '对象 ID',
    `liked_status`   bit(1)  NOT NULL DEFAULT b'0'   COMMENT '点赞状态',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `is_deleted` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否删除',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '点赞表';

CREATE INDEX idx_user_id ON travel_user_like(`user_id`);
CREATE INDEX idx_item_type ON travel_user_like(`item_type`);

-- 历史记录表
CREATE TABLE `travel_user_history` (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
    `user_id`   bigint  NOT NULL    COMMENT '用户 ID',
    `item_type` int NOT NULL    COMMENT '对象类型',
    `item_id`   bigint  NOT NULL    COMMENT '对象 ID',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `is_deleted` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否删除',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '历史记录表';

CREATE INDEX idx_user_id ON travel_user_history(`user_id`);
CREATE INDEX idx_item_type ON travel_user_history(`item_type`);

-- 消息表
CREATE TABLE `travel_user_message` (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
    `user_id`   bigint  NOT NULL    COMMENT '用户 ID',
    `item_type` int  NULL    COMMENT '对象类型',
    `item_id`   bigint   NULL    COMMENT '对象 ID',
    `message_user_id`   bigint  NULL    COMMENT '消息关联用户 ID',
    `message_type`   int  NULL    COMMENT '消息类型',
    `message_status`    bit(1)  NOT NULL DEFAULT b'0'   COMMENT '消息状态',
    `content`   varchar(512)    NULL    COMMENT '消息内容',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `is_deleted` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否删除',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '消息表';

CREATE INDEX idx_user_id ON travel_user_message(`user_id`);
CREATE INDEX idx_item_type ON travel_user_message(`item_type`);

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