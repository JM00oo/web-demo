
-- +migrate Up

CREATE TABLE `comment` (
  `id` varchar(64) COLLATE utf8mb4_bin NOT NULL,
  `comment` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `post_id` varchar(64) COLLATE utf8mb4_bin NOT NULL,
  `created_at` Datetime COLLATE utf8mb4_bin NOT NULL,
  `owner_name` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `comment_FK_post` FOREIGN KEY (`post_id`) REFERENCES `post` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `comment_FK_user` FOREIGN KEY (`owner_name`) REFERENCES `user` (`username`) ON DELETE CASCADE ON UPDATE CASCADE

) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC;

-- +migrate Down

DROP TABLE `comment`;
