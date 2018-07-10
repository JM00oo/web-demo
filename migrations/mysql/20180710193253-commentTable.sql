
-- +migrate Up

CREATE TABLE `comment` (
  `id` varchar(64) COLLATE utf8mb4_bin NOT NULL,
  `comment` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `post_id` varchar(64) COLLATE utf8mb4_bin NOT NULL,
  `created_at` Datetime COLLATE utf8mb4_bin NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC;

-- +migrate Down

DROP TABLE `comment`;
