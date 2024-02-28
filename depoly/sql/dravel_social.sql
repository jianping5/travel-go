CREATE DATABASE dravel_social;
USE dravel_social;

-- 文章表
CREATE TABLE `travel_social_article` (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
    `user_id`   bigint  NOT NULL    COMMENT '用户 ID',
    `title` varchar(255)    NOT NULL    COMMENT '标题',
    `cover_url` varchar(255)    NOT NULL COMMENT    '封面',
    `content`   longtext    NOT NULL    COMMENT '内容',
    `like_count`    int NOT NULL DEFAULT 0  COMMENT '点赞量',
    `comment_count` int NOT NULL DEFAULT 0  COMMENT '评论量',
    `favor_count`   int NOT NULL DEFAULT 0  COMMENT '收藏量',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `is_deleted` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否删除',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '文章表';

CREATE INDEX idx_user_id ON travel_social_article(`user_id`);

-- 视频表
CREATE TABLE `travel_social_video` (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
    `user_id`   bigint  NOT NULL    COMMENT '用户 ID',
    `title` varchar(255)    NOT NULL    COMMENT '标题',
    `cover_url` varchar(255)    NOT NULL COMMENT    '封面',
    `video_url`   varchar(255)    NOT NULL    COMMENT '视频 URL',
    `like_count`    int NOT NULL DEFAULT 0  COMMENT '点赞量',
    `comment_count` int NOT NULL DEFAULT 0  COMMENT '评论量',
    `favor_count`   int NOT NULL DEFAULT 0  COMMENT '收藏量',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `is_deleted` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否删除',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '视频表';

CREATE INDEX idx_user_id ON travel_social_video(`user_id`);

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
    `content`   varchar(2048)   NOT NULL    COMMENT '版权信息',
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
    `content` varchar(2048) NOT NULL    COMMENT '动态内容',
    `file_type` int NOT NULL     COMMENT    '动态类型',
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
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `is_deleted` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否删除',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '用户社区表';

CREATE INDEX idx_user_id ON travel_social_user_community(`user_id`);