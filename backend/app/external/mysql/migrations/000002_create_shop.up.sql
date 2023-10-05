CREATE TABLE IF NOT EXISTS `shops`(
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `store_name` VARCHAR(128) NOT NULL,
    `store_name_huri` VARCHAR(128) NOT NULL,
    `user_name` VARCHAR(128) NOT NULL,
    `user_name_huri` VARCHAR(128) NOT NULL,
    `zip_code` VARCHAR(10) NOT NULL,
    `address` VARCHAR(128) NOT NULL,
    `street_address` VARCHAR(128) NOT NULL,
    `email` VARCHAR(256) NOT NULL,
    `store_phone` VARCHAR(15) NOT NULL,
    `user_phone` VARCHAR(15) NOT NULL,
    `keep_expiry` INT NOT NULL,
    `password` VARCHAR(128) NOT NULL
    ) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS `provisional_shops` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `store_name` VARCHAR(128) NOT NULL,
    `store_name_huri` VARCHAR(128) NOT NULL,
    `user_name` VARCHAR(128) NOT NULL,
    `user_name_huri` VARCHAR(128) NOT NULL,
    `zip_code` VARCHAR(10) NOT NULL,
    `address` VARCHAR(128) NOT NULL,
    `street_address` VARCHAR(128) NOT NULL,
    `email` VARCHAR(256) NOT NULL,
    `store_phone` VARCHAR(15) NOT NULL,
    `user_phone` VARCHAR(15) NOT NULL,
    `keep_expiry` INT NOT NULL,
    `password` VARCHAR(128) NOT NULL,
    `token` VARCHAR(256) NOT NULL,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `expired_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
    ) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;