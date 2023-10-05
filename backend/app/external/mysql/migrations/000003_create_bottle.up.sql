CREATE TABLE IF NOT EXISTS `bottles` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `user_id` INT NOT NULL,
    `shop_id` INT NOT NULL,
    `kind` INT NOT NULL,
    `status` INT NOT NULL,
    `opened_at` DATETIME NOT NULL,
    `name` VARCHAR(255),
    FOREIGN KEY (`user_id`) REFERENCES users(id),
    FOREIGN KEY (`shop_id`) REFERENCES shops(id)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;