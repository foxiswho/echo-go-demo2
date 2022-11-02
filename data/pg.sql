


CREATE TABLE IF NOT EXISTS `iam_account` (
  `id` bigint(20) NOT NULL,
  `create_time` datetime DEFAULT current_timestamp() COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL ON UPDATE current_timestamp() COMMENT '更新时间',
  `create_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建人',
  `update_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '更新人',
  `tenant_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '租户',
  `org_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '组织id',
  `account` char(255) DEFAULT NULL COMMENT '账户',
  `account_md5` char(32) DEFAULT NULL COMMENT '账户md5',
  `password` char(32) DEFAULT NULL COMMENT '密码',
  `salt` char(32) DEFAULT NULL COMMENT '组织id',
  `phone` varchar(50) DEFAULT NULL COMMENT '手机号',
  `mail` varchar(255) DEFAULT NULL COMMENT '邮箱',
  `mail_md5` char(32) DEFAULT NULL,
  `code` char(70) DEFAULT NULL COMMENT '编码',
  `mail_verify` tinyint(1) NOT NULL DEFAULT 2 COMMENT '邮箱验证1是2否',
  `phone_verify` tinyint(1) NOT NULL DEFAULT 2 COMMENT '手机验证1是2否',
  `state` tinyint(1) DEFAULT 1 COMMENT '启用1是2否',
  `register_time` datetime NOT NULL DEFAULT '0001-01-01 00:00:00' COMMENT '注册时间',
  `register_ip` char(100) DEFAULT NULL COMMENT '注册ip',
  `login_time` datetime DEFAULT '0001-01-01 00:00:00' COMMENT '登陆时间',
  `role_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '角色id',
  `department_id_main` bigint(20) NOT NULL DEFAULT 0 COMMENT '主部门id',
  `department_ids` text DEFAULT NULL COMMENT '部门',
  `team_ids` text DEFAULT NULL COMMENT '团队',
  `org_ids` text DEFAULT NULL COMMENT '团队',
  `type_domain` int(11) NOT NULL DEFAULT 0 COMMENT '域类型',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `tenant_id` (`tenant_id`),
  KEY `create_time` (`create_time`),
  KEY `create_id` (`create_id`),
  KEY `organization_id` (`org_id`),
  KEY `account` (`account`),
  KEY `phone` (`phone`),
  KEY `mail` (`mail`),
  KEY `code` (`code`),
  KEY `state` (`state`),
  KEY `type_domain` (`type_domain`),
  KEY `department_id_main` (`department_id_main`),
  KEY `role_id` (`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


CREATE TABLE `iam_account_login_log` (
  `id` bigint(20) NOT NULL,
  `create_time` datetime DEFAULT current_timestamp() COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL ON UPDATE current_timestamp() COMMENT '更新时间',
  `create_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建人',
  `update_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '更新人',
  `tenant_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '租户',
  `org_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '组织id',
  `login_source` bigint(20) NOT NULL DEFAULT 0 COMMENT '登陆来源',
  PRIMARY KEY (`id`),
  KEY `tenant_id` (`tenant_id`),
  KEY `create_time` (`create_time`),
  KEY `create_id` (`create_id`),
  KEY `organization_id` (`org_id`),
  KEY `account` (`account`),
  KEY `account` (`account`),
  KEY `phone` (`phone`),
  KEY `mail` (`mail`),
  KEY `code` (`code`),
  KEY `department_id_main` (`department_id_main`),
  KEY `type_domain` (`type_domain`),
  KEY `role_id` (`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;