CREATE TABLE IF NOT EXISTS accounts (
    id serial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    username varchar(50) NOT NULL,
    firstname varchar(50) NOT NULL,
    lastname varchar(50) NOT NULL,
    password varchar(255) NOT NULL,
    email varchar(255) UNIQUE NOT NULL
);
