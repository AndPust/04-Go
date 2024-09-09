CREATE TABLE `cart_items` (`id` integer PRIMARY KEY AUTOINCREMENT,`created_at` datetime,`updated_at` datetime,`deleted_at` datetime,`item_id` integer, `amount` integer);
CREATE INDEX `idx_carts_deleted_at` ON `cart_items`(`deleted_at`);
CREATE TABLE `products` (`id` integer PRIMARY KEY AUTOINCREMENT,`created_at` datetime,`updated_at` datetime,`deleted_at` datetime,`name` text,`price` integer,`category_id` integer);
CREATE INDEX `idx_products_deleted_at` ON `products`(`deleted_at`);
CREATE TABLE `categories` (`id` integer PRIMARY KEY AUTOINCREMENT,`created_at` datetime,`updated_at` datetime,`deleted_at` datetime,`name` text);
CREATE INDEX `idx_categories_deleted_at` ON `categories`(`deleted_at`);


INSERT INTO categories(name) VALUES 
    ("Headphones"),
    ("Keyboards"),
    ("Mice");


INSERT INTO products(name, price, category_id) VALUES 
    ("Bose", 129, 1),
    ("Sony", 149, 1),
    ("Sennheiser", 219, 1),
    ("Audio-Technica", 79, 1),
    ("CherryMX", 250, 2),
    ("Corsair", 89, 2),
    ("Logitech", 25, 2),
    ("Ducky", 50, 2),
    ("Razer", 29, 3),
    ("Steelseries", 40, 3),
    ("Logitech", 12, 3);