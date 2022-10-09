CREATE TABLE IF NOT EXISTS users (
                       id bigserial not null primary key,
                       username varchar not null unique,
                       first_name varchar,
                       last_name varchar,
                       avatar varchar,
                       email varchar not null unique,
                       password varchar not null,
                       phone text unique,
                       is_author boolean not null,
                       about text
);

CREATE TABLE IF NOT EXISTS posts (
    post_id bigserial not null primary key,
    user_id bigserial not null references users(id),
    img text not null,
    title varchar(128) not null,
    text text not null
);

CREATE TABLE IF NOT EXISTS sessions (
    value varchar(32) primary key,
    user_id bigserial not null unique references users(id),
    expire_date date
);