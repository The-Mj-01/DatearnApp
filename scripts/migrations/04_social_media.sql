CREATE TABLE `social_media`
(
    `id`                    INTEGER UNSIGNED NOT NULL AUTO_INCREMENT,
    `title`                 VARCHAR(255),
    `social_media_id`       BIT  NOT NULL,
    `tiktok`                VARCHAR(255),
    `created_at`            TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at`            TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `delete_at`             TIMESTAMP,
    PRIMARY KEY (`id`),
);