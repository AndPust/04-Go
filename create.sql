CREATE TABLE `cart_items` (`id` integer PRIMARY KEY AUTOINCREMENT,`created_at` datetime,`updated_at` datetime,`deleted_at` datetime,`item_id` integer, `amount` integer);
CREATE INDEX `idx_carts_deleted_at` ON `cart_items`(`deleted_at`);
CREATE TABLE `products` (`id` integer PRIMARY KEY AUTOINCREMENT,`created_at` datetime,`updated_at` datetime,`deleted_at` datetime,`name` text,`price` integer,`category_id` integer);
CREATE INDEX `idx_products_deleted_at` ON `products`(`deleted_at`);
CREATE TABLE `categories` (`id` integer PRIMARY KEY AUTOINCREMENT,`created_at` datetime,`updated_at` datetime,`deleted_at` datetime,`name` text);
CREATE INDEX `idx_categories_deleted_at` ON `categories`(`deleted_at`);


-- INSERT INTO Category VALUES (1, "Headphones");
-- INSERT INTO Category VALUES (2, "Keyboards");
-- INSERT INTO Category VALUES (3, "Mice");

-- INSERT INTO products VALUES (1,  "Bose", 129, 1);
-- INSERT INTO products VALUES (2,  "Sony", 149, 1);
-- INSERT INTO products VALUES (3,  "Sennheiser", 219, 1);
-- INSERT INTO products VALUES (4,  "Audio-Technica", 79, 1);

-- INSERT INTO products VALUES (5,  "CherryMX", 250, 2);
-- INSERT INTO products VALUES (6,  "Corsair", 89, 2);
-- INSERT INTO products VALUES (7,  "Logitech", 25, 2);
-- INSERT INTO products VALUES (8,  "Ducky", 50, 2);

-- INSERT INTO products VALUES (9,  "Razer", 29, 3);
-- INSERT INTO products VALUES (10, "Steelseries", 40, 3);
-- INSERT INTO products VALUES (11, "Logitech", 12, 3);

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