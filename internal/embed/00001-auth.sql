CREATE TABLE IF NOT EXISTS users (
    user_id     SERIAL PRIMARY KEY,
    name        VARCHAR(50),
    age         INTEGER
);