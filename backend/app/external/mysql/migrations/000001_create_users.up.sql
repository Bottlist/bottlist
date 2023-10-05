CREATE TABLE IF NOT EXISTS `users`(
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `email` VARCHAR(128) UNIQUE NOT NULL,
    `first_name` VARCHAR(128) NOT NULL,
    `last_name` VARCHAR(128) NOT NULL,
    `first_name_hira` VARCHAR(128) NOT NULL,
    `last_name_hira` VARCHAR(128) NOT NULL,
    `screen_name` VARCHAR(128) NOT NULL,
    `birthday` DATE NOT NULL,
    `password` VARCHAR(128) NOT NULL,
    `image` VARCHAR(12) NOT NULL
    ) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS `provisional_users` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `email` VARCHAR(128) UNIQUE NOT NULL,
    `first_name` VARCHAR(128) NOT NULL,
    `last_name` VARCHAR(128) NOT NULL,
    `first_name_hira` VARCHAR(128) NOT NULL,
    `last_name_hira` VARCHAR(128) NOT NULL,
    `screen_name` VARCHAR(128) NOT NULL,
    `birthday` DATE NOT NULL,
    `password` VARCHAR(128) NOT NULL,
    `token` VARCHAR(256) NOT NULL,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `expired_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
    ) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;