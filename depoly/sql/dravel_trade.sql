CREATE DATABASE dravel_trade;
USE dravel_trade;

-- 作品表
CREATE TABLE `travel_trade_work` (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
    `user_id`   bigint  NOT NULL    COMMENT '用户 ID',
    `copyright_id` bigint  NOT NULL    COMMENT '版权 ID',
    `price` varchar(50) NOT NULL    COMMENT '作品价格',
    `status`    int DEFAULT 0   NOT NULL    COMMENT '商品状态',
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

-- 交易记录表
CREATE TABLE `travel_trade_record` (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
    `work_id`  bigint  NOT NULL    COMMENT '作品 ID',
    `old_user_id`   bigint  NOT NULL    COMMENT '原用户 ID',
    `new_user_id`   bigint  NOT NULL    COMMENT '新用户 ID',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `is_deleted` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否删除',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '交易记录表';

CREATE INDEX idx_work_id ON travel_trade_record(`work_id`);