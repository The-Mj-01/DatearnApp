CREATE TABLE `images`
(
    `id`             INTEGER UNSIGNED AUTO_INCREMENT NOT NULL,
    `imageable_id`   INTEGER UNSIGNED                NOT NULL,
    `imageable_type` VARCHAR(255)                    NOT NULL,
    `name`           VARCHAR(255)                    NOT NULL,
    `path`           VARCHAR(255)                    NOT NULL,
    `created_at`     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at`     TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `delete_at` TIMESTAMP,
    PRIMARY KEY (`id`)
);