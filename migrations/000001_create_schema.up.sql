CREATE TABLE brands (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    description TEXT
);

CREATE TABLE vouchers (
    id SERIAL PRIMARY KEY,
    brand_id INT REFERENCES brands(id),
    cost_in_points FLOAT,
    voucher_code VARCHAR(255) UNIQUE
);

CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    customer_id INT,
    total_cost_in_points FLOAT
);

CREATE TABLE transaction_vouchers (
    transaction_id INT REFERENCES transactions(id),
    voucher_id INT REFERENCES vouchers(id)
);

