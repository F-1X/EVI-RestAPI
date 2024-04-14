CREATE TABLE IF NOT EXISTS ads (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price FLOAT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


INSERT INTO ads (name, description, price) VALUES 
    ('test 1', 'desciption test 1', 10.50),
    ('test 2', 'desciption test 2', 20.75),
    ('test 3', 'desciption test 3', 30.25);