CREATE TABLE IF NOT EXISTS users (
    username TEXT PRIMARY KEY NOT NULL,
    password varchar(150) NOT NULL,
    firstname TEXT NOT NULL,
    lastname TEXT NOT NULL,
    surname TEXT,
    user_group TEXT,
    role TEXT NOT NULL,
    balance INTEGER NOT NULL
);