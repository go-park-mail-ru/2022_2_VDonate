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
    avatar    varchar(64)           default null,
    password  varchar(128) not null,
    is_author boolean      not null,
    balance   integer      not null default 0,
    about     varchar(1024)         default null
);

/*
    {post_id} -> {img, title, text}
*/
CREATE TABLE IF NOT EXISTS posts
(
    post_id      bigserial not null primary key,
    user_id      bigserial not null references users (id) on delete cascade,
    content      text      not null,
    date_created timestamp not null default now(),
    tier         int                default 0
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

    unique (author_id, price, tier),
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
    subscription_id bigserial not null references author_subscriptions (id) on delete cascade,
    date_created    timestamp not null default now(),
    UNIQUE (author_id, subscriber_id),
    UNIQUE (subscriber_id, subscription_id)
);

/*
    Нет зависимостей
*/
CREATE TABLE IF NOT EXISTS followers
(
    author_id    bigserial not null references users (id) on delete cascade,
    follower_id  bigserial not null references users (id) on delete cascade,
    date_created timestamp not null default now(),
    UNIQUE (author_id, follower_id)
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
    {id} -> {tag_name}
*/
CREATE TABLE IF NOT EXISTS tags
(
    id       bigserial    NOT NULL PRIMARY KEY,
    tag_name VARCHAR(128) NOT NULL UNIQUE
);

/*
    No dependices
*/
CREATE TABLE IF NOT EXISTS post_tags
(
    post_id bigserial NOT NULL REFERENCES posts (post_id) ON DELETE CASCADE,
    tag_id  bigserial NOT NULL REFERENCES tags (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS comments
(
    id           bigserial     NOT NULL PRIMARY KEY,
    user_id      bigserial     NOT NULL REFERENCES users (id) ON DELETE NO ACTION,
    post_id      bigserial     NOT NULL REFERENCES posts (post_id) ON DELETE CASCADE,
    content      varchar(1024) NOT NULL,
    date_created TIMESTAMP     NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS notification
(
    name      varchar   NOT NULL,
    data      jsonb     NOT NULL,
    timestamp timestamp NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS payments
(
    id      varchar   NOT NULL PRIMARY KEY,
    from_id bigserial NOT NULL REFERENCES users (id) ON DELETE NO ACTION,
    to_id   bigserial NOT NULL REFERENCES users (id) ON DELETE NO ACTION,
    sub_id  bigserial NOT NULL REFERENCES author_subscriptions (id) ON DELETE NO ACTION,
    status  varchar            default 'WAITING',
    time    timestamp NOT NULL DEFAULT now()
);

CREATE OR REPLACE FUNCTION notify_event_like() RETURNS TRIGGER AS
$$

DECLARE
    data      jsonb;
    author_id bigint;
    name      varchar;

BEGIN
    SELECT posts.user_id FROM posts WHERE post_id = NEW.post_id INTO author_id;

    if new.user_id = author_id then
        return null;
    end if;

    SELECT username FROM users WHERE users.id = new.user_id INTO name;

    data = jsonb_build_object(
            'user_id', author_id,
            'username', name,
            'post_id', new.post_id
        );
    INSERT INTO notification(name, data) VALUES ('like', data);

    RETURN NULL;
END;

$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION notify_event_posts() RETURNS TRIGGER AS
$$

DECLARE
    data        jsonb;
    var         bigint;
    author_name varchar;

BEGIN
    FOR var IN SELECT subscriber_id FROM subscriptions WHERE author_id = NEW.user_id
        LOOP
            SELECT username FROM users WHERE id = new.user_id INTO author_name;
            data = jsonb_build_object(
                    'user_id', var,
                    'author_id', new.user_id,
                    'author_name', author_name
                );
            INSERT INTO notification(name, data) VALUES ('posts', data);
        END LOOP;

    RETURN NULL;
END;

$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION notify_event_subscriber() RETURNS TRIGGER AS
$$

DECLARE
    data            jsonb;
    subscriber_name varchar;

BEGIN
    SELECT username FROM users WHERE id = NEW.subscriber_id INTO subscriber_name;

    data = jsonb_build_object(
            'user_id', new.author_id,
            'subscriber_id', new.subscriber_id,
            'subscriberName', subscriber_name
        );
    INSERT INTO notification(name, data) VALUES ('subscriber', data);

    RETURN NULL;
END;

$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION notify_event_payment() RETURNS TRIGGER AS
$$

DECLARE
    data     jsonb;
    sub_name varchar;

BEGIN
    SELECT title FROM author_subscriptions as authSub WHERE authSub.id = new.sub_id INTO sub_name;

    data = jsonb_build_object(
            'user_id', new.from_id,
            'author_id', new.to_id,
            'sub_name', sub_name,
            'status', new.status
        );
    INSERT INTO notification(name, data) VALUES ('payment', data);

    IF new.status = 'PAID' THEN
        UPDATE user_info
        SET balance = balance + (SELECT price FROM author_subscriptions as aSub WHERE aSub.id = new.sub_id)
        WHERE user_id = new.to_id;
    END IF;

    RETURN NULL;
END;

$$ LANGUAGE plpgsql;

CREATE TRIGGER new_payment
    AFTER INSERT OR UPDATE
    ON payments
    FOR EACH ROW
EXECUTE PROCEDURE notify_event_payment();

CREATE TRIGGER new_like_notify
    AFTER INSERT
    ON likes
    FOR EACH ROW
EXECUTE PROCEDURE notify_event_like();

CREATE TRIGGER new_posts_notify
    AFTER INSERT
    ON posts
    FOR EACH ROW
EXECUTE PROCEDURE notify_event_posts();

CREATE TRIGGER new_subscribers_notify
    AFTER INSERT OR UPDATE
    ON subscriptions
    FOR EACH ROW
EXECUTE PROCEDURE notify_event_subscriber();
