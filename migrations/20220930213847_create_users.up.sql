CREATE TABLE users (
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

CREATE TABLE posts (
    post_id bigserial not null primary key,
    user_id bigserial not null references users(id),
    title text not null

)