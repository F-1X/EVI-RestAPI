CREATE TABLE IF NOT EXISTS ads (
    id SERIAL PRIMARY KEY,
    id_ad VARCHAR(255) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT ,
    price FLOAT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


INSERT INTO ads (name, id_ad, description, price) VALUES 
    ('test 1', 'test_id1','desciption test 1', 10.50),
    ('test 2','test_id2', 'desciption test 2', 20.75),
    ('test 3', 'test_id3','desciption test 3', 30.25);