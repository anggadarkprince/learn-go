ALTER TABLE categories
    ADD COLUMN description TEXT AFTER name,
    ADD COLUMN created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP;