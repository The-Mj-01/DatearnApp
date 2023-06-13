CREATE TABLE `likes`
(
    `id`            INTEGER UNSIGNED NOT NULL AUTO_INCREMENT,
    `liker`         INTEGER UNSIGNED NOT NULL,
    `liked`         INTEGER UNSIGNED NOT NULL,
    `count`         INTEGER UNSIGNED NOT NULL,
    `cost`          INTEGER UNSIGNED NOT NULL,
    `created_at`    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at`    TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `delete_at` TIMESTAMP,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`liker`) REFERENCES `users` (`id`),
    FOREIGN KEY (`liked`) REFERENCES `users` (`id`),
);