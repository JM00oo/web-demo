
-- +migrate Up

CREATE TABLE `user` (
  `id` varchar(64) COLLATE utf8mb4_bin NOT NULL,
  `token` varchar(36) COLLATE utf8mb4_bin NOT NULL,
  `username` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `password` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `created_at` Datetime COLLATE utf8mb4_bin NOT NULL,
  UNIQUE (`username`),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC;

-- +migrate Down

DROP TABLE `user`;
