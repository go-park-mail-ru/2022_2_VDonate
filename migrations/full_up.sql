CREATE TABLE IF NOT EXISTS users (
                                     id bigserial not null primary key,
                                     username varchar not null unique,
                                     avatar varchar,
                                     email varchar not null unique,
                                     password varchar not null,
                                     is_author boolean not null,
                                     about text
);

CREATE TABLE IF NOT EXISTS posts (
                                     post_id bigserial not null primary key,
                                     user_id bigserial not null references users(id),
                                     img text not null,
                                     title varchar(128) not null,
                                     text text not null,
                                     unique (user_id, title)
);

CREATE TABLE IF NOT EXISTS sessions (
                                        value varchar(36) primary key,
                                        user_id bigserial not null references users(id),
                                        expire_date date
);

CREATE TABLE IF NOT EXISTS author_subscriptions (
    id bigserial not null primary key,
    author_id bigserial not null references users(id),
    img varchar,
    tier integer not null,
    title varchar not null,
    text text not null,
    price bigserial not null,
    unique (author_id, tier),
    unique (author_id, title),
    unique (author_id, price)
);

CREATE TABLE IF NOT EXISTS subscriptions (
                                             author_id bigserial not null references users(id),
                                             subscriber_id bigserial not null references users(id),
                                             subscription_id bigserial not null references author_subscriptions(id),
                                             primary key (author_id, subscription_id, subscriber_id)
);
