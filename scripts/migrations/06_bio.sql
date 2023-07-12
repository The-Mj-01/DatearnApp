CREATE TABLE `bio`
(
    `id`                    INTEGER UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id`               INTEGER UNSIGNED UNIQUE NOT NULL,
    `social_media`          INTEGER UNSIGNED NOT NULL,
    `description`           Text,
    `photos`                INTEGER UNSIGNED NOT NULL,
    `primary_photos`        INTEGER UNSIGNED NOT NULL,
    `created_at`    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at`    TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `delete_at` TIMESTAMP,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`social_media`) REFERENCES `social_media` (`id`),
    FOREIGN KEY (`photos`) REFERENCES `images` (`id`),
    FOREIGN KEY (`primary_photos`) REFERENCES `images` (`id`),
);