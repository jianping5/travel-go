CREATE DATABASE dravel_social;
USE dravel_social;

-- 内容表
CREATE TABLE `travel_social_content` (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
    `user_id`   bigint  NOT NULL    COMMENT '用户 ID',
    `item_type` int NOT NULL DEFAULT 0 COMMENT '内容类型',
    `title` varchar(255)    NOT NULL    COMMENT '标题',
    `cover_url` varchar(255)    NOT NULL COMMENT    '封面',
    `content`   varchar(255)    NOT NULL    COMMENT '内容 URL',
    `description`   varchar(2048)   NULL    COMMENT '简介',
    `tag`    varchar(1024)  NULL    COMMENT '标签',
    `like_count`    int NOT NULL DEFAULT 0  COMMENT '点赞量',
    `comment_count` int NOT NULL DEFAULT 0  COMMENT '评论量',
    `favor_count`   int NOT NULL DEFAULT 0  COMMENT '收藏量',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `is_deleted` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否删除',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '内容表';

CREATE INDEX idx_user_id ON travel_social_content(`user_id`);

-- 评论表
CREATE TABLE `travel_social_comment` (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
    `user_id`   bigint  NOT NULL    COMMENT '用户 ID',
    `comment_item_type` int NOT NULL    COMMENT '评论类型',
    `comment_item_id`   bigint  NOT NULL    COMMENT '评论对象 ID',
    `parent_user_id`    bigint  NOT NULL DEFAULT 0   COMMENT '父评论用户 ID',
    `top_id`    bigint  NOT NULL DEFAULT 0   COMMENT '顶级评论 ID',
    `content`   varchar(2048)   NOT NULL    COMMENT '评论内容',
    `like_count`    int NOT NULL DEFAULT 0  COMMENT '点赞量',
    `reply_count` int NOT NULL DEFAULT 0  COMMENT '回复量',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `is_deleted` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否删除',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '评论表';

CREATE INDEX idx_user_id ON travel_social_comment(`user_id`);

-- 版权表
CREATE TABLE `travel_social_copyright` (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
    `user_id`   bigint  NOT NULL    COMMENT '用户 ID',
    `item_type` int NOT NULL    COMMENT '对象类型',
    `item_id`   bigint  NOT NULL    COMMENT '对象 ID',
    `metadata`   varchar(3072)   NOT NULL    COMMENT '作品元数据',
    `trade_hash`    varchar(1024)   NULL    COMMENT '交易哈希',
    `address`   varchar(1024)   NULL    COMMENT '区块链地址',
    `status`    int NULL    COMMENT '版权状态',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `is_deleted` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否删除',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '版权表';

CREATE INDEX idx_user_id ON travel_social_copyright(`user_id`);
CREATE INDEX idx_item_type ON travel_social_copyright(`item_type`);

-- 社区表
CREATE TABLE `travel_social_community` (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
    `user_id`   bigint  NOT NULL    COMMENT '用户 ID',
    `name`   varchar(255)   NOT NULL    COMMENT '社区名称',
    `description` varchar(1024) NOT NULL    COMMENT '社区简介',
    `avatar`   varchar(255)  NOT NULL    COMMENT '社区头像',
    `status`   int NOT NULL DEFAULT 0 COMMENT '社区状态',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `is_deleted` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否删除',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '社区表';

CREATE INDEX idx_user_id ON travel_social_copyright(`user_id`);
CREATE INDEX idx_item_type ON travel_social_copyright(`item_type`);

-- 社区动态表
CREATE TABLE `travel_social_dynamic` (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
    `user_id`   bigint  NOT NULL    COMMENT '用户 ID',
    `community_id`  bigint  NOT NULL    COMMENT '社区 ID',
    `title` varchar(255)    NULL    COMMENT '标题',
    `description`   varchar(2048)   NULL COMMENT '描述',
    `content` varchar(2048) NOT NULL    COMMENT '动态内容',
    `file_type` int NOT NULL     COMMENT    '动态类型',
    `like_count`    int NOT NULL DEFAULT 0  COMMENT '点赞量',
    `comment_count` int NOT NULL DEFAULT 0  COMMENT '评论量',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `is_deleted` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否删除',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '社区动态表';

CREATE INDEX idx_user_id ON travel_social_dynamic(`user_id`);

-- 用户社区表
CREATE TABLE `travel_social_user_community` (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
    `user_id`   bigint  NOT NULL    COMMENT '用户 ID',
    `community_id`  bigint  NOT NULL    COMMENT '社区 ID',
    `role`  int NOT NULL COMMENT    '角色',
    `member_count` int NOT NULL DEFAULT 0  COMMENT '成员量',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `is_deleted` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否删除',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '用户社区表';

CREATE INDEX idx_user_id ON travel_social_user_community(`user_id`);

-- 收藏表
CREATE TABLE `travel_social_favor` (
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

CREATE INDEX idx_user_id ON travel_social_favor(`user_id`);
CREATE INDEX idx_favorite_id ON travel_social_favor(`favorite_id`);
CREATE INDEX idx_item_type ON travel_social_favor(`item_type`);

-- 收藏夹表
CREATE TABLE `travel_social_favorite` (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
    `user_id`   bigint  NOT NULL    COMMENT '用户 ID',
    `name`   varchar(50)  NOT NULL    COMMENT '收藏夹名称',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `is_deleted` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否删除',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '收藏夹表';

CREATE INDEX idx_user_id ON travel_social_favorite(`user_id`);

-- 点赞表
CREATE TABLE `travel_social_like` (
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

CREATE INDEX idx_user_id ON travel_social_like(`user_id`);
CREATE INDEX idx_item_type ON travel_social_like(`item_type`);

-- 历史记录表
CREATE TABLE `travel_social_history` (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
    `user_id`   bigint  NOT NULL    COMMENT '用户 ID',
    `item_type` int NOT NULL    COMMENT '对象类型',
    `item_id`   bigint  NOT NULL    COMMENT '对象 ID',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `is_deleted` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否删除',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '历史记录表';

CREATE INDEX idx_user_id ON travel_social_history(`user_id`);
CREATE INDEX idx_item_type ON travel_social_history(`item_type`);

-- 消息表
CREATE TABLE `travel_social_message` (
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

CREATE INDEX idx_user_id ON travel_social_message(`user_id`);
CREATE INDEX idx_item_type ON travel_social_message(`item_type`);