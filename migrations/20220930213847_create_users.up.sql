CREATE TABLE users (
                       id bigserial not null primary key,
                       username varchar not null unique,
                       first_name varchar,
                       last_name varchar,
                       avatar varchar,
                       email varchar not null unique,
                       password varchar not null,
                       phone text unique,
                       is_author boolean not null
);