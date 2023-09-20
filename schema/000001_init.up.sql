CREATE TABLE users
(
    id serial not null unique,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password_hash varchar(255) not null,
    is_superuser boolean default false
);

CREATE TABLE types
(
    id serial not null unique,
    trade_type varchar(255) not null unique
);

INSERT INTO types (trade_type) VALUES ('Покупка ценных бумаг');

INSERT INTO types (trade_type) VALUES ('Продажа ценных бумаг');

CREATE TABLE trades
(
    id serial not null unique,
    ticker varchar(255) not null,
    user_id int references users(id) on delete cascade not null,
    type_id int references types(id) on delete cascade not null,
    price float,
    amount int
);

CREATE TABLE portfolio
(
    id serial not null unique,
    ticker varchar(255) not null,
    user_id int references users(id) on delete cascade not null,
    amount int not null
);
