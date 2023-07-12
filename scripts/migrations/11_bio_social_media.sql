CREATE TABLE `bio_interest`
(
    `bio_id` INTEGER UNSIGNED NOT NULL,
    `social_media_id` INTEGER UNSIGNED NOT NULL,
    `media_id`          INTEGER UNSIGNED NOT NULL,
    `created_at`     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at`     TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`bio_id`, `interest_id`)
);