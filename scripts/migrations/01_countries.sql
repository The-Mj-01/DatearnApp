CREATE TABLE `countries`
(
    `id`         INTEGER UNSIGNED NOT NULL AUTO_INCREMENT,
    `name`       VARCHAR(255)     NOT NULL,
    `created_at`     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at`     TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);