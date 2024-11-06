
-- Categories
INSERT INTO categories (id, name, picture_url, description, sub_category_id) VALUES
(1, 'Electronics', 'https://example.com/electronics.jpg', 'Electronic devices and accessories', NULL),
(2, 'Smartphones', 'https://example.com/smartphones.jpg', 'Mobile phones and accessories', 1),
(3, 'Laptops', 'https://example.com/laptops.jpg', 'Portable computers', 1),
(4, 'Clothing', 'https://example.com/clothing.jpg', 'Apparel and fashion items', NULL),
(5, 'T-shirts', 'https://example.com/tshirts.jpg', 'Casual and formal t-shirts', 4),
(6, 'Jeans', 'https://example.com/jeans.jpg', 'Denim pants for all occasions', 4);

-- Brands
INSERT INTO brands (name) VALUES
('Apple'),
('Samsung'),
('Dell'),
('Levi''s'),
('Nike'),
('Adidas');

-- Products
INSERT INTO products (title, description, category_id, brand_id, price, discount, colors, sizes, picture_url) VALUES
('iPhone 12', 'Latest iPhone model', 2, 1, 999.99, 0, 'Black,White,Red', '64GB,128GB,256GB', 'https://example.com/iphone12.jpg'),
('MacBook Pro', '13-inch MacBook Pro with M1 chip', 3, 1, 1299.99, 1199,99, 'Silver,Space Gray', '8GB,16GB', 'https://example.com/macbookpro.jpg'),
('Samsung Galaxy S21', 'Flagship Android smartphone', 2, 2, 799.99, 0, 'Phantom Black,Phantom Silver,Phantom Violet', '128GB,256GB', 'https://example.com/galaxys21.jpg'),
('Dell XPS 13', 'Ultra-thin and light laptop', 3, 3, 999.99, 0, 'Platinum Silver,Frost White', '8GB,16GB,32GB', 'https://example.com/dellxps13.jpg'),
('Classic Cotton T-Shirt', 'Comfortable everyday t-shirt', 5, 5, 19.99, 0, 'White,Black,Navy,Gray', 'S,M,L,XL', 'https://example.com/cottontshirt.jpg'),
('Slim Fit Jeans', 'Modern slim fit denim jeans', 6, 4, 59.99, 0, 'Blue,Black,Gray', '28,30,32,34,36', 'https://example.com/slimfitjeans.jpg');
