ALTER TABLE vouchers
ADD COLUMN points_required INT NOT NULL;

ALTER TABLE transactions
ADD COLUMN voucher_id INT REFERENCES vouchers(id) ON DELETE CASCADE;