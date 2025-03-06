-- Active: 1712051663122@@127.0.0.1@3306@blog_service

-- 创建数据库
CREATE DATABASE
IF
    NOT EXISTS blog_service DEFAULT CHARACTER
    SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;

-- 标签表
CREATE Table `blog_tag`(
    `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(100) DEFAULT '' COMMENT '标签名称',
    -- 公共字段
    `created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
    `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
    `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
    `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
    `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
    `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
    -- 公共字段
    `state` TINYINT(3) UNSIGNED DEFAULT '1' COMMENT '状态 0 为禁用、1 为启用',
    PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='标签管理';

-- 文章表
CREATE Table `blog_article`(
    `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `title` VARCHAR(100) DEFAULT '' COMMENT '文章标题',
    `desc` VARCHAR(255) DEFAULT '' COMMENT '文章简述',
    `cover_image_url` VARCHAR(255) DEFAULT '' COMMENT '封面图片地址',
     -- 公共字段
    `created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
    `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
    `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
    `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
    `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
    `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
    -- 公共字段
    `state` TINYINT(3) UNSIGNED DEFAULT '1' COMMENT '状态 0 为禁用、1 为启用',
    PRIMARY KEY(`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章管理';

-- 文章标签关联表
CREATE Table `blog_acticle_tag`(
    `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `article_id` INT(11) NOT NULL COMMENT '文章 ID',
    `tag_id` INT(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '标签 ID',
     -- 公共字段
    `created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
    `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
    `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
    `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
    `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
    `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
    -- 公共字段
    PRIMARY KEY(`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章标签关联';

# 认证
CREATE TABLE `blog_auth`(
    `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `app_key` VARCHAR(20) DEFAULT '' COMMENT 'Key',
    `app_secret` VARCHAR(50) DEFAULT '' COMMENT 'Secret',
         -- 公共字段
    `created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
    `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
    `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
    `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
    `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
    `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
    -- 公共字段
    PRIMARY KEY(`id`) USING BTREE
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='认证管理';

INSERT INTO `blog_service`.`blog_auth`(`id`, `app_key`, `app_secret`, `created_on`, `created_by`, `modified_on`, `modified_by`, `deleted_on`, `is_del`) 
VALUES (1, 'eddycjy', 'go-programming-tour-book', 0, 'eddycjy', 0, '', 0, 0);

 SELECT * FROM `blog_auth` WHERE app_key = 'eddycjy' AND app_secret = 'go-programming-tour-book' AND is_del = 0 ORDER BY `blog_auth`.`id` LIMIT 1;