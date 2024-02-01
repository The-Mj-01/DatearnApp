CREATE TABLE `interests`
(
    `id` INTEGERUNSIGNED AUTO_INCREMENT NOT NULL,
    `name`  VARCHAR(255)                NOT NULL,
    `deleted_at` TIMESTAMP,
    PRIMARY KEY (`id`)
);