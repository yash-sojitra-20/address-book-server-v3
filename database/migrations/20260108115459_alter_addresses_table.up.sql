ALTER TABLE addresses
ADD COLUMN first_name VARCHAR(100) NOT NULL AFTER user_id,
ADD COLUMN last_name VARCHAR(100) AFTER first_name;
