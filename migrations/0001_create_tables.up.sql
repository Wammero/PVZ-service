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
