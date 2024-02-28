CREATE DATABASE dravel_trade;
USE dravel_trade;

-- 作品表
CREATE TABLE `travel_trade_work` (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
    `user_id`   bigint  NOT NULL    COMMENT '用户 ID',
    `name`  varchar(50)  NOT NULL    COMMENT '作品名称',
    `content`  varchar(512) NOT NULL COMMENT    '作品信息',
    `price` varchar(50) NOT NULL    COMMENT '作品价格',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `is_deleted` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否删除',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '作品表';

CREATE INDEX idx_user_id ON travel_trade_work(`user_id`);

-- 用户作品表
CREATE TABLE `travel_trade_user_work` (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
    `user_id`   bigint  NOT NULL    COMMENT '用户 ID',
    `work_id`  bigint  NOT NULL    COMMENT '作品 ID',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `is_deleted` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否删除',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '用户作品表';

CREATE INDEX idx_user_id ON travel_trade_user_work(`user_id`);