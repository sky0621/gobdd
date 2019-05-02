# テーブル定義は適当
CREATE TABLE IF NOT EXISTS `notice` (
  `id` varchar(64) NOT NULL,
  `title` varchar(256) NOT NULL,
  `text` varchar(256) NOT NULL,
  `publish_from` int NULL,
  `publish_to` int NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE utf8_bin;
