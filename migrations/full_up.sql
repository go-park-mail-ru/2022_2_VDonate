/*
    Подходит под 3НФ. Доказательство:

    1НФ:
    - Все атрибуты атомарны;
    - У отношения присутствует ключ.

    2НФ:
    - 1НФ;
    - Каждый неключевой атрибут неприводимо зависит от ПК.

    3НФ:
    - 2НФ;
    - Каждый не ключевой атрибут нетранзитивно зависит от ПК.
*/

/*
    {id} -> {username, email}
*/

CREATE TABLE IF NOT EXISTS users
(
    id       bigserial    not null primary key,
    username varchar(32)  not null unique,
    email    varchar(254) not null unique
);

/*
    {user_id} -> {avatar, password, is_author, about}
*/
CREATE TABLE IF NOT EXISTS user_info
(
    user_id   bigserial    not null primary key references users (id),
    avatar    varchar(64)   default null,
    password  varchar(128) not null,
    is_author boolean      not null,
    about     varchar(1024) default null
);

/*
    {post_id} -> {img, title, text}
*/
CREATE TABLE IF NOT EXISTS posts
(
    post_id bigserial     not null primary key,
    user_id bigserial     not null references users (id) on delete cascade,
    img     varchar(64)   not null,
    title   varchar(128)  not null,
    text    varchar(2048) not null
);

/*
    {value} -> {user_id, expire_date}
*/
CREATE TABLE IF NOT EXISTS sessions
(
    value       varchar(60) primary key,
    user_id     bigserial not null references users (id) on delete cascade,
    expire_date date      not null
);

/*
    {id} -> {author_id, img, tier, title, text, price}
*/
CREATE TABLE IF NOT EXISTS author_subscriptions
(
    id        bigserial    not null primary key,
    author_id bigserial    not null references users (id) on delete cascade,
    img       varchar(64) default null,
    tier      smallint     not null,
    title     varchar(128) not null,
    text      varchar(128) not null,
    price     integer      not null check ( price >= 0 ),

    unique (author_id, tier),
    unique (author_id, title),
    unique (author_id, price)
);

/*
    Нет зависимостей
*/
CREATE TABLE IF NOT EXISTS subscriptions
(
    author_id       bigserial not null references users (id) on delete cascade,
    subscriber_id   bigserial not null references users (id) on delete cascade,
    subscription_id bigserial not null references author_subscriptions (id) on delete restrict,
    primary key (author_id, subscription_id, subscriber_id)
);

/*
    Нет зависимостей
*/
CREATE TABLE IF NOT EXISTS likes
(
    user_id bigserial not null references users (id) on delete cascade,
    post_id bigserial not null references posts (post_id) on delete cascade,
    primary key (user_id, post_id)
);

/*
    {id} -> {author_id, user_id, price}
*/
CREATE TABLE IF NOT EXISTS donates
(
    id        bigserial NOT NULL PRIMARY KEY,
    author_id bigserial not null references users (id) on delete cascade,
    user_id   bigserial not null references users (id) on delete cascade,
    price     integer   not null default 0 check ( price >= 0 )
);
