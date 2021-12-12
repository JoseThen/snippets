-- Adminer 4.7.8 MySQL dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

SET NAMES utf8mb4;

DROP TABLE IF EXISTS `snippets`;
CREATE TABLE `snippets` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(100) NOT NULL,
  `content` text NOT NULL,
  `created` datetime NOT NULL,
  `expires` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_snippets` (`created`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `snippets` (`id`, `title`, `content`, `created`, `expires`) VALUES
(1,	'An old silent pond',	'An old silent pond...\nA frog jumps into the pond,\nsplash! Silence again.\n\n Matsuo Bash',	'2021-11-24 01:58:17',	'2022-11-24 01:58:17'),
(2,	'Over the wintry forest',	'Over the wintry\nforest, winds howl in rage\nwith no leaves to blow.\n\n Natsume Soseki',	'2021-11-24 01:58:25',	'2022-11-24 01:58:25'),
(3,	'First autumn morning',	'First autumn morning\nthe mirror I stare into\nshows my father\'s face.\n\n Murakami Kijo',	'2021-11-24 01:58:34',	'2021-12-01 01:58:34'),
(4,	'0 snail',	'O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa',	'2021-11-24 02:45:42',	'2021-12-01 02:45:42');

-- 2021-12-12 16:27:09
