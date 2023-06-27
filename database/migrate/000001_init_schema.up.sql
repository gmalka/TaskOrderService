CREATE TABLE IF NOT EXISTS tasks (
    id serial PRIMARY KEY,
    quantity INTEGER NOT NULL CHECK(quantity >= 0),
    heights  INTEGER[] NOT NULL,
    price INTEGER NOT NULL CHECK(price > 0),
    answer INTEGER NOT NULL CHECK(answer >= 0)
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

INSERT INTO users (username,password,firstname,lastname,surname,user_group,role,balance)
SELECT 'root','$2a$10$arAEdUAW2TewfoFX8EC09uh.6HoJyXGXX9fJQuw2iygDf1xkEoTES','root','root','root','admin-group','admin',5000
WHERE NOT EXISTS (
    SELECT 1
    FROM users
    WHERE username = 'root'
);