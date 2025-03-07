CREATE DATABASE dravel_data;
USE dravel_data;

-- 用户行为表
CREATE TABLE `travel_data_behavior` (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
    `user_id`   bigint  NOT NULL    COMMENT '用户 ID',
    `behavior_item_type`    int NOT NULL    COMMENT '行为对象类型',
    `behavior_item_id`  bigint  NOT NULL    COMMENT '行为对象 ID',
    `behavior_type` int NOT NULL    COMMENT '行为类型',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `is_deleted` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否删除',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '用户行为表';

CREATE INDEX idx_user_id ON travel_data_behavior(`user_id`);

-- 内容标签表
CREATE TABLE `travel_data_content_tag` (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
    `name`   varchar(255)  NOT NULL    COMMENT '标签名',
    `item_type`   int   NOT NULL    COMMENT '对象类型',
    `item_id`     bigint    NOT NULL    COMMENT '对象 ID',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `is_deleted` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否删除',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '内容标签表';

-- 用户标签表
CREATE TABLE `travel_data_user_tag` (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
    `user_id`   bigint  NOT NULL COMMENT    '用户 ID',
    `tag`   varchar(2048)  NOT NULL    COMMENT '标签数组',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `is_deleted` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否删除',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '用户标签表';

CREATE INDEX idx_user_id ON travel_data_user_tag(`user_id`);