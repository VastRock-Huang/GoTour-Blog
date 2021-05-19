-- 创建数据库
CREATE
DATABASE IF NOT EXISTS blog_service  -- 创建表仅当表不存在时
    DEFAULT CHARACTER SET utf8mb4   -- 设定数据库默认字符集
    DEFAULT COLLATE utf8mb4_general_ci;     -- 设定字符集的默认校对规则
-- utf8mb4 即utf8 more bytes 4, 指的是兼容4字节Unicode的编码格式
-- utf8mb4_general_ci 为utf8mb4的校对规则,一般用于字符比较,
-- ci(case insensitive)大小写不敏感,
-- cs(case sensitive)大小写敏感
-- bin 二元, 二进制存储字符,大小写敏感

-- 使用当前数据库
use
blog_service;

-- 创建标签表
CREATE TABLE `blog_tag`
(
    `id`          int(10) unsigned NOT NULL AUTO_INCREMENT,
    `name`        varchar(100) DEFAULT '' COMMENT '标签名称',
    `created_on`  int(10) unsigned DEFAULT 0 COMMENT '创建时间',
    `created_by`  varchar(100) DEFAULT '' COMMENT '创建人',
    `modified_on` int(10) unsigned DEFAULT 0 COMMENT '修改时间',
    `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
    `deleted_on`  int(10) unsigned DEFAULT 0 COMMENT '删除时间',
    `is_del`      tinyint(1) unsigned DEFAULT 0 COMMENT '是否删除,0未删除,1已删除',
    `state`       tinyint(1) unsigned DEFAULT 1 COMMENT '状态,0禁用,1启用',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='标签管理';

-- 创建文章表
CREATE TABLE `blog_article`
(
    `id`              int(10) unsigned NOT NULL AUTO_INCREMENT,
    `title`           varchar(100) DEFAULT '' COMMENT '文章标题',
    `desc`            varchar(255) DEFAULT '' COMMENT '文章简述',
    `cover_image_url` varchar(255) DEFAULT '' COMMENT '封面图片地址',
    `content`         longtext COMMENT '文章内容',
    `created_on`      int(10) unsigned DEFAULT 0 COMMENT '创建时间',
    `created_by`      varchar(100) DEFAULT '' COMMENT '创建人',
    `modified_on`     int(10) unsigned DEFAULT 0 COMMENT '修改时间',
    `modified_by`     varchar(100) DEFAULT '' COMMENT '修改人',
    `deleted_on`      int(10) unsigned DEFAULT 0 COMMENT '删除时间',
    `is_del`          tinyint(1) unsigned DEFAULT 0 COMMENT '是否删除,0未删除,1已删除',
    `state`           tinyint(1) unsigned DEFAULT 1 COMMENT '状态,0禁用,1启用',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章管理';


-- 创建文章标签关联表
CREATE TABLE `blog_article_tag`
(
    `id`          int(10) unsigned NOT NULL AUTO_INCREMENT,
    `article_id`  int(11) unsigned NOT NULL COMMENT '文章ID',
    `tag_id`      int(10) unsigned NOT NULL DEFAULT '0' COMMENT '标签ID',
    `created_on`  int(10) unsigned DEFAULT 0 COMMENT '创建时间',
    `created_by`  varchar(100) DEFAULT '' COMMENT '创建人',
    `modified_on` int(10) unsigned DEFAULT 0 COMMENT '修改时间',
    `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
    `deleted_on`  int(10) unsigned DEFAULT 0 COMMENT '删除时间',
    `is_del`      tinyint(1) unsigned DEFAULT 0 COMMENT '是否删除,0未删除,1已删除',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章标签关联表';

-- 创建签发认证信息表
CREATE TABLE `blog_auth`
(
    `id`          int(10) unsigned NOT NULL AUTO_INCREMENT,
    `app_key`     varchar(20)  DEFAULT '' COMMENT 'Key',
    `app_secret`  varchar(50)  DEFAULT '' COMMENT 'Secret',
    `created_on`  int(10) unsigned DEFAULT 0 COMMENT '创建时间',
    `created_by`  varchar(100) DEFAULT '' COMMENT '创建人',
    `modified_on` int(10) unsigned DEFAULT 0 COMMENT '修改时间',
    `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
    `deleted_on`  int(10) unsigned DEFAULT 0 COMMENT '删除时间',
    `is_del`      tinyint(1) unsigned DEFAULT 0 COMMENT '是否删除,0未删除,1已删除',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='认证管理';

-- 新增一条认证信息
INSERT INTO `blog_service`.`blog_auth`
VALUES(1,'vastrock-huang','gotour-blogservice',0,'vastrock-huang',0,'',0,0);