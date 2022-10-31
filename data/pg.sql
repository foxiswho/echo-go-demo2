

CREATE TABLE IF EXISTS `iam_account` (
  `id` bigint(20) NOT NULL,
  `account` varchar(255) DEFAULT NULL COMMENT '账户',
  `account_md5` varchar(32) DEFAULT NULL COMMENT '账户md5',
  `password` varchar(32) DEFAULT NULL COMMENT '密码',
  `salt` varchar(20) DEFAULT NULL COMMENT '干扰码',
  `phone` varchar(20) DEFAULT NULL,
  `mail` varchar(255) DEFAULT NULL,
  `mail_md5` varchar(32) DEFAULT NULL,
  `mail_verify` tinyint(1) NOT NULL DEFAULT 2 COMMENT '邮箱验证1是2否',
  `phone_verify` tinyint(1) NOT NULL DEFAULT 2 COMMENT '手机验证1是2否',
  `state` tinyint(1) DEFAULT 1 COMMENT '启用1是2否',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;