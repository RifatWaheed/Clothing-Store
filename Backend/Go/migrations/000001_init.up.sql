CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE products (
    -- Base fields
    pkid BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    created_by BIGINT NOT NULL,
    modified_by BIGINT,
    created_date TIMESTAMP NOT NULL DEFAULT NOW(),
    modified_date TIMESTAMP,

    -- Product description
    description TEXT,

    -- Pricing
    price NUMERIC(10,2) NOT NULL,
    discount_amount NUMERIC(10,2) DEFAULT 0,
    discount_percent NUMERIC(5,2) DEFAULT 0,

    -- SKU
    sku_id BIGINT,
    sku_code TEXT UNIQUE,

    -- Color
    color_id BIGINT,
    color_name TEXT,

    -- Gender
    gender_id BIGINT,
    gender_name TEXT,

    -- Size
    size_id BIGINT,
    size_name TEXT,

    -- Stock
    stock_id BIGINT,
    stock_qty INT DEFAULT 0,

    -- Product type
    type_id BIGINT,
    type_name TEXT,

    -- Voucher
    voucher_id BIGINT,
    voucher_code TEXT
);
