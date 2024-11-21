ALTER TABLE transactions
ADD COLUMN total_points INT NOT NULL,
ADD COLUMN total_value DECIMAL(10, 2) NOT NULL;

ALTER TABLE redemptions
ADD COLUMN redeemed_points INT NOT NULL,
ADD COLUMN redeemed_value DECIMAL(10, 2) NOT NULL;
