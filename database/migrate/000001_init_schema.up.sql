CREATE TABLE IF NOT EXISTS tasks (
    id serial PRIMARY KEY,
    quantity INTEGER NOT NULL,
    heights  INTEGER[] NOT NULL,
    price INTEGER NOT NULL,
    answer INTEGER NOT NULL
);

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

CREATE TABLE IF NOT EXISTS userOrders (
    username TEXT REFERENCES users(username),
    orderId INTEGER REFERENCES tasks(id)
);

INSERT INTO users
SELECT 'root','root','root','root','root','admin-group','admin',5000
WHERE NOT EXISTS (
    SELECT 1
    FROM users
    WHERE username = 'root'
);