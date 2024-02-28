CREATE DATABASE dravel_intelligence;
USE dravel_intelligence;

-- 对话表
CREATE TABLE `travel_intelligence_conversation` (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
    `user_id`   bigint  NOT NULL    COMMENT '用户 ID',
    `content`  varchar(1024)  NOT NULL    COMMENT '对话内容',
    `type`  int NOT NULL COMMENT    '对话类型',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `is_deleted` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否删除',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '对话表';

CREATE INDEX idx_user_id ON travel_intelligence_conversation(`user_id`);

-- 攻略表
CREATE TABLE `travel_intelligence_strategy` (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
    `user_id`   bigint  NOT NULL    COMMENT '用户 ID',
    `ask_content`  varchar(1024)  NOT NULL    COMMENT '询问内容',
    `strategy_content`  varchar(1024)  NOT NULL    COMMENT '攻略内容',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `is_deleted` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否删除',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '攻略表';

CREATE INDEX idx_user_id ON travel_intelligence_strategy(`user_id`);