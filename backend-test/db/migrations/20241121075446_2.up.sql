ALTER TABLE vouchers
DROP COLUMN IF EXISTS points_required;

ALTER TABLE transactions
DROP COLUMN IF EXISTS voucher_id;