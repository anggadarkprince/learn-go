CREATE TABLE customers
(
    id VARCHAR(255) NOT NULL,
    name VARCHAR(100) NOT NULL,
    PRIMARY KEY (id)
) ENGINE = InnoDB;

ALTER TABLE customers
    ADD COLUMN email VARCHAR(100) NOT NULL,
    ADD COLUMN balance DECIMAL(10, 2) NOT NULL DEFAULT 0.00,
    ADD COLUMN rating DOUBLE NOT NULL DEFAULT 0.0,
    ADD COLUMN created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    ADD COLUMN birth_date DATE,
    ADD COLUMN married BOOLEAN NOT NULL DEFAULT FALSE;

DESC customers;

INSERT INTO customers (id, name, email, balance, rating, birth_date, married)
VALUES ("angga", "Angga Ari", "angga@mail.com", 1000000, 4.3, '1992-01-01', true);

INSERT INTO customers (id, name, email, balance, rating, birth_date, married)
VALUES ("keenan", "Keenan Evander", "keenan@mail.com", 500000, 5.0, '1999-01-01', false); 

CREATE TABLE users(
    username VARCHAR(50) NOT NULL,
    password VARCHAR(255) NOT NULL,
    PRIMARY KEY (username)
) ENGINE = InnoDB;

INSERT INTO users (username, password) VALUES("admin", "admin");

CREATE TABLE comments(
    id INT NOT NULL AUTO_INCREMENT,
    email VARCHAR(255) NOT NULL,
    comment TEXT NOT NULL,
    PRIMARY KEY (id),
) ENGINE = InnoDB;