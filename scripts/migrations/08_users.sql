CREATE TABLE `users`
(
    `id`            INTEGER UNSIGNED AUTO_INCREMENT NOT NULL,
    `uuid`          VARCHAR(255)                    NOT NULL,
    `username`      varchar(255)                    NOT NULL,
    `first_name`    varchar(255),
    `last_name`     varchar(255),
--     `bio`           INTEGER UNSIGNED                NOT NULL,
    `born`          TIMESTAMP                       NOT NULL,
    `country`       INTEGER UNSIGNED                NOT NULL,
    `city`          INTEGER UNSIGNED                NOT NULL,
    `sex`           INTEGER UNSIGNED                NOT NULL,
    `email`         VARCHAR(255)                    NOT NULL,
    `password`      VARCHAR(255)                    NOT NULL,
    `last_login_at` TIMESTAMP,
    `created_at`    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at`    TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `delete_at` TIMESTAMP,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`city`) REFERENCES `cities` (`id`),
    FOREIGN KEY (`country`) REFERENCES `countries` (`id`),
    FOREIGN KEY (`sex`) REFERENCES `sexs` (`id`),
    FOREIGN KEY (`bio`) REFERENCES `bio` (`id`),


);

CREATE UNIQUE INDEX users_email_unique_index ON `users` (`email`);
CREATE UNIQUE INDEX users_uuid_unique_index ON `users` (`uuid`);