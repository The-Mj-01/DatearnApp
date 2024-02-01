CREATE TABLE `likes`
(
    `id`            INTEGER UNSIGNED NOT NULL AUTO_INCREMENT,
    `disliker`         INTEGER UNSIGNED NOT NULL,
    `disliked`         INTEGER UNSIGNED NOT NULL,
    `count`         INTEGER UNSIGNED NOT NULL,
    `created_at`    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at`    TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`disliker`) REFERENCES `users` (`id`),
    FOREIGN KEY (`disliked`) REFERENCES `users` (`id`),
);