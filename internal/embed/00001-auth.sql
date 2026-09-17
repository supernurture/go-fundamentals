-- PostgreSQL

CREATE TABLE IF NOT EXISTS users (
    user_id     SERIAL PRIMARY KEY,
    name        VARCHAR(50),
    age         INTEGER
);

-- Used to add a new column.
ALTER TABLE users ADD       COLUMN place_of_birth TEXT;

-- Used to change the data type of a column.
ALTER TABLE users ALTER     COLUMN place_of_birth TYPE VARCHAR(100);

-- Used to delete a column.
ALTER TABLE users DROP      COLUMN place_of_birth;