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
                                        value varchar(36) primary key,
                                        user_id bigserial not null references users(id),
                                        expire_date date
);

CREATE TABLE IF NOT EXISTS author_subscriptions (
    id bigserial not null primary key,
    author_id bigserial not null references users(id),
    tier integer not null DEFAULT 0 UNIQUE,
    text text not null,
    price bigserial not null
);

CREATE TABLE IF NOT EXISTS subscriptions (
                                             author_id bigserial not null references users(id),
                                             subscriber_id bigserial not null references users(id),
                                             subscription_id bigserial not null references author_subscriptions(id)
);

CREATE TABLE IF NOT EXISTS donates (
    id bigserial NOT NULL PRIMARY KEY,
    author_id bigserial not null references users(id),
    user_id bigserial not null references users(id),
    price bigserial not null
);
