\c warehouse;

CREATE TABLE stores (
                        id SERIAL PRIMARY KEY,
                        name VARCHAR(255) NOT NULL,
                        address TEXT NOT NULL
);

CREATE TABLE animals (
                         id SERIAL PRIMARY KEY,
                         name VARCHAR(255) NOT NULL,
                         type VARCHAR(255) NOT NULL,
                         color VARCHAR(255) NOT NULL,
                         store_id INT NOT NULL,
                         age INT NOT NULL,
                         price DECIMAL(10,2),
                         FOREIGN KEY (store_id) REFERENCES stores(id) ON DELETE CASCADE
);
