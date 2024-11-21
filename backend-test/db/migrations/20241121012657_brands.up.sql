CREATE TABLE brands (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE vouchers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    cost INTEGER NOT NULL CHECK (cost >= 50000 AND cost <= 100000),
    brand_id INT NOT NULL REFERENCES brands(id)
);

CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    customer_name VARCHAR(255) NOT NULL,
    voucher_id INT NOT NULL REFERENCES vouchers(id),
    redemption_date TIMESTAMP DEFAULT NOW()
);