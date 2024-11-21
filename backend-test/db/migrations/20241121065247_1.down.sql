DROP TABLE IF EXISTS redemptions CASCADE;

ALTER TABLE vouchers 
DROP COLUMN IF EXISTS code, 
DROP COLUMN IF EXISTS description, 
DROP COLUMN IF EXISTS points_required, 
DROP COLUMN IF EXISTS value, 
DROP COLUMN IF EXISTS created_at;

ALTER TABLE transactions 
DROP COLUMN IF EXISTS total_points, 
DROP COLUMN IF EXISTS total_value, 
DROP COLUMN IF EXISTS created_at;