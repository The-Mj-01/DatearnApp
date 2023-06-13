CREATE TABLE `social_media`
(
    `id`                    INTEGER UNSIGNED NOT NULL AUTO_INCREMENT,
    `instagram`          VARCHAR(255),
    `twitter`          VARCHAR(255),
    `telegram`          VARCHAR(255),
    `spotify`          VARCHAR(255),
    `facebook`          VARCHAR(255),
    `tiktok`          VARCHAR(255),
    `created_at`    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at`    TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `delete_at` TIMESTAMP,
    PRIMARY KEY (`id`),
);