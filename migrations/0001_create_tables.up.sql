-- Пользователи
CREATE TABLE IF NOT EXISTS users (
    user_id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    salt VARCHAR(255) NOT NULL,
    role TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- pvz
CREATE TABLE IF NOT EXISTS pvz (
    pvz SERIAL PRIMARY KEY,
    id UUID NOT NULL UNIQUE,  
    creator_id INT, 
    registration_date TIMESTAMP NOT NULL,
    city VARCHAR(100) NOT NULL,
    FOREIGN KEY (creator_id) REFERENCES users(user_id)
);

-- Приёмки товара
CREATE TABLE IF NOT EXISTS receptions (
    reception_serial_id SERIAL PRIMARY KEY,
    reception_id UUID NOT NULL DEFAULT gen_random_uuid() UNIQUE,
    reception_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    pvz_id UUID NOT NULL,
    status TEXT NOT NULL CHECK (status IN ('in_progress', 'close')),
    FOREIGN KEY (pvz_id) REFERENCES pvz(id)
);

-- Товары
CREATE TABLE IF NOT EXISTS products (
    product_serial_id SERIAL PRIMARY KEY,
    product_id UUID NOT NULL DEFAULT gen_random_uuid() UNIQUE,  
    reception_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    type TEXT NOT NULL CHECK (type IN ('electronics', 'clothing', 'footwear')), 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    is_active BOOLEAN DEFAULT TRUE
);

-- Принятые товары в рамках приёмки
CREATE TABLE IF NOT EXISTS reception_products (
    id SERIAL PRIMARY KEY,
    reception_id UUID NOT NULL,
    product_id UUID NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    FOREIGN KEY (reception_id) REFERENCES receptions(reception_id),
    FOREIGN KEY (product_id) REFERENCES products(product_id)
);

CREATE INDEX IF NOT EXISTS idx_receptions_pvz_id ON receptions (pvz_id);
CREATE INDEX IF NOT EXISTS idx_receptions_reception_time ON receptions (reception_time);
CREATE INDEX IF NOT EXISTS idx_receptions_status ON receptions (status);
CREATE INDEX IF NOT EXISTS idx_reception_products_product_id ON reception_products (product_id);
CREATE INDEX IF NOT EXISTS idx_reception_products_reception_id ON reception_products (reception_id);
CREATE INDEX IF NOT EXISTS idx_products_product_id ON products (product_id);
CREATE INDEX IF NOT EXISTS idx_products_is_active ON products (is_active);
CREATE INDEX IF NOT EXISTS idx_pvz_pvz_id ON pvz (id);
CREATE INDEX IF NOT EXISTS idx_pvz_registration_date ON pvz (registration_date);
CREATE INDEX IF NOT EXISTS idx_reception_products_is_active ON reception_products (is_active);
CREATE INDEX IF NOT EXISTS idx_receptions_pvz_reception_time ON receptions (pvz_id, reception_time);
CREATE INDEX IF NOT EXISTS idx_reception_products_product_id_is_active ON reception_products (product_id, is_active);
CREATE INDEX IF NOT EXISTS idx_products_product_id_is_active ON products (product_id, is_active);