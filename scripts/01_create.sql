CREATE TABLE IF NOT EXISTS products (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    price REAL NOT NULL,
    stock INT NOT NULL,
    min_stock INT NOT NULL,
    is_active INT 
);

CREATE INDEX IF NOT EXISTS idx_products_name ON products (name);
CREATE INDEX IF NOT EXISTS idx_products_min_stock ON products (min_stock);
CREATE INDEX IF NOT EXISTS idx_products_is_active ON products (is_active);
