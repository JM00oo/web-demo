
-- +migrate Up

CREATE TABLE `post` (
  `id` varchar(64) COLLATE utf8mb4_bin NOT NULL,
  `content` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `title` varchar(16) COLLATE utf8mb4_bin NOT NULL,
  `owner_id` varchar(64) COLLATE utf8mb4_bin NOT NULL,
  `created_at` Datetime COLLATE utf8mb4_bin NOT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `Post_FK_Owner` FOREIGN KEY (`owner_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC;

-- +migrate Down

DROP TABLE `post`;
