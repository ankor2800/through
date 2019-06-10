-- +migrate Up
CREATE TABLE IF NOT EXISTS `auto_orders` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS `through` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `date` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS `orders_1` (
  `id` int(10) unsigned NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS `orders_2` (
  `id` int(10) unsigned NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS `orders_3` (
  `id` int(10) unsigned NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL
);

-- +migrate Down
DROP TABLE IF EXISTS `auto_orders`;
DROP TABLE IF EXISTS `through`;
DROP TABLE IF EXISTS `orders_1`;
DROP TABLE IF EXISTS `orders_2`;
DROP TABLE IF EXISTS `orders_3`;
