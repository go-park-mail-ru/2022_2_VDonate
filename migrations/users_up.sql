/*
    Создание и выдача прав пользователям для работы с БД
*/
CREATE USER auth_user WITH PASSWORD :auth_password;
CREATE USER notifications_user WITH PASSWORD :notifications_password;
CREATE USER posts_user WITH PASSWORD :posts_password;
CREATE USER subscriptions_user WITH PASSWORD :subscriptions_password;
CREATE USER author_subscriptions_user WITH PASSWORD :author_subscriptions_password;
CREATE USER users_user WITH PASSWORD :users_password;

GRANT SELECT, INSERT, DELETE, REFERENCES
    ON sessions, users TO auth_user;

GRANT SELECT, DELETE, TRIGGER
    ON notification TO notifications_user;

GRANT SELECT, INSERT, DELETE, REFERENCES
    ON posts, tags, post_tags, comments, likes TO posts_user;
GRANT SELECT
    ON subscriptions, author_subscriptions, followers TO posts_user;

GRANT SELECT, INSERT, DELETE, REFERENCES
    ON author_subscriptions TO author_subscriptions_user;
GRANT SELECT
    ON subscriptions, followers TO author_subscriptions_user;

GRANT SELECT, INSERT, DELETE, REFERENCES
    ON subscriptions, followers TO subscriptions_user;
GRANT INSERT, UPDATE
    ON payments TO subscriptions_user;

GRANT SELECT, INSERT, DELETE, REFERENCES
    ON users, user_info TO users_user;
GRANT SELECT
    ON posts, subscriptions, author_subscriptions, payments TO users_user;
