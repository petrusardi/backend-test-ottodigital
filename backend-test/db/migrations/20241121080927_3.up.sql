ALTER TABLE transactions
DROP COLUMN IF EXISTS total_points,
DROP COLUMN IF EXISTS total_value;

ALTER TABLE redemptions
DROP COLUMN IF EXISTS redeemed_points,
DROP COLUMN IF EXISTS redeemed_value;
