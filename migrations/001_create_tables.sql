CREATE TABLE users (
    id UUID PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100) UNIQUE
);

CREATE TABLE meals (
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES users(id),  
    name VARCHAR(100),
    description TEXT,
    date TIMESTAMP,
    is_diet BOOLEAN
);