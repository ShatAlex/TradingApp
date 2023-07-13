CREATE TABLE users
(
    id serial not null unique,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password_hash varchar(255) not null,
    balance int default 0,
    trading_status varchar(255) default 'beginner'
);

CREATE TABLE types
(
    id serial not null unique,
    user_id int references users(id) on delete cascade not null,
    trade_type varchar(255) unique
);

CREATE TABLE trades
(
    id serial not null unique,
    ticker varchar(255) not null unique,
    user_id int references users(id) on delete cascade not null,
    type_id int references types(id) on delete cascade not null,
    price int,
    amount int
);