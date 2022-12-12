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
    post_id          bigserial not null primary key,
    user_id          bigserial not null references users (id) on delete cascade,
    content          text      not null,
    date_created     timestamp not null default now(),
    tier             int default 0
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
    primary key (author_id, subscriber_id)
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


/*
    https://vk.com/abstract_memes
    https://vk.com/tproger
    https://vk.com/mozgitreski
    https://vk.com/kinoartmag
*/
INSERT INTO users (username, email) VALUES ('АМДЭВС', 'amdevs@mail.ru');
INSERT INTO user_info (user_id, password, is_author, avatar, about) VALUES (1, 'amdevs', true, '430fd690-ad80-49bc-b46e-f3b547367781.jpg', 'у нас тут новая искренность');

INSERT INTO users (username, email) VALUES ('Типичный программист', 'programist@mail.ru');
INSERT INTO user_info (user_id, password, is_author, avatar, about) VALUES (2, 'programist', true, 'd42ab69a-e22b-4cad-8538-83bd4457a047.jpg', 'Tproger — издание о разработке и обо всём, что с ней связано.

Читайте нас в Telegram: https://t.me/tproger_official

Реклама на Tproger: https://tprg.ru/yCVP
Размещение вакансий: https://tprg.ru/poUM
HR-брендинг: https://tprg.ru/L2IJ

Яндекс.Дзен: https://zen.yandex.ru/tproger
');

INSERT INTO users (username, email) VALUES ('Мозги трески', 'brains@mail.ru');
INSERT INTO user_info (user_id, password, is_author, avatar, about) VALUES (3, 'brains', true, '4a28db1c-1696-4650-932f-f2b530c0af4d.jpg', 'Мозги трески - это смешные комикс-стрипы обо всём на свете: бытовые зарисовки, пародии на кино и игры, всякие безумные ситуации - в общем, все, что взбредет в голову автору.

Новый комикс каждый понедельник в 21:00 и иногда ВНЕЗАПНО в другие дни.

Почти 100 привидений - это история о старинном заброшенном доме, в котором обитает около сотни призраков - каждый со своим характером и историей.

Новый выпуск каждые среду и пятницу в 21:00.

А еще у меня вышел комикс-игра «Домашние монстры» про мальчика, обнаружившего в своей комнате монстров. Экземпляр можно купить в любых книжных магазинах и комикс-шопах страны.');

INSERT INTO users (username, email) VALUES ('Искусство кино', 'art@mail.ru');
INSERT INTO user_info (user_id, password, is_author, avatar, about) VALUES (4, 'art', true, '0c3b90e0-801d-4936-9e88-f51bfbee4bd8.jpg', '«Искусство кино» — журнал о кино, который издается с января 1931 года и остается одним из старейших в мире периодических изданий о кино. Каждый номер освещает актуальные вопросы кинематографа и визуальной культуры, а также включает в себя редкие архивные публикации, обзоры крупнейших международных кинофестивалей, русскую и зарубежную кинопрозу и другие материалы.');

INSERT INTO posts (user_id, content, date_created, tier) VALUES (1, '<img src="https://wsrv.nl/?url=http://95.163.209.195:9000/e4da3b7fbbce2345d7772b0674a318d5/e97ca124-4237-4076-83b4-3484f93794f5.jpg" class="post-content__image">', now(), 1);
INSERT INTO posts (user_id, content, date_created, tier) VALUES (1, '<img src="https://wsrv.nl/?url=http://95.163.209.195:9000/1679091c5a880faf6fb5e6087eb1b2dc/7dfae8ba-c70e-446d-b1b4-080b62fad7f6.jpg" class="post-content__image">', now(), 1);
INSERT INTO posts (user_id, content, date_created, tier) VALUES (1, '<img src="https://wsrv.nl/?url=http://95.163.209.195:9000/4a8a08f09d37b73795649038408b5f33/c3c5d321-4028-4f21-afaa-572436d76c7c.jpg" class="post-content__image">', now(), 1);
INSERT INTO posts (user_id, content, date_created, tier) VALUES (1, '<img src="https://wsrv.nl/?url=http://95.163.209.195:9000/0cc175b9c0f1b6a831c399e269772661/fc1e28e8-742e-4ae4-9ad5-c7eacf50824a.jpg" class="post-content__image">', now(), 1);
INSERT INTO posts (user_id, content, date_created, tier) VALUES (1, '<img src="https://wsrv.nl/?url=http://95.163.209.195:9000/e1671797c52e15f763380b45e841ec32/35aae3f0-cfd5-4f73-b8da-33fc451f3a2e.jpg" class="post-content__image">', now(), 1);

INSERT INTO author_subscriptions (author_id, img, tier, title, text, price) VALUES (1, 'a5293e04-74a4-4ca2-98b3-baf7425e1741.jpg', 1, 'Мемы', 'Обычная подписка на мемы АНДЭВС', 1990);

INSERT INTO posts (user_id, content, date_created, tier) VALUES (2, '<img src="https://wsrv.nl/?url=http://95.163.209.195:9000/e4da3b7fbbce2345d7772b0674a318d5/e97ca124-4237-4076-83b4-3484f93794f5.jpg" class="post-content__image">', now(), 1);
INSERT INTO posts (user_id, content, date_created, tier) VALUES (2, '<img src="https://wsrv.nl/?url=http://95.163.209.195:9000/1679091c5a880faf6fb5e6087eb1b2dc/7dfae8ba-c70e-446d-b1b4-080b62fad7f6.jpg" class="post-content__image">', now(), 1);
INSERT INTO posts (user_id, content, date_created, tier) VALUES (2, '<img src="https://wsrv.nl/?url=http://95.163.209.195:9000/4a8a08f09d37b73795649038408b5f33/c3c5d321-4028-4f21-afaa-572436d76c7c.jpg" class="post-content__image">', now(), 2);
INSERT INTO posts (user_id, content, date_created, tier) VALUES (2, '<img src="https://wsrv.nl/?url=http://95.163.209.195:9000/0cc175b9c0f1b6a831c399e269772661/fc1e28e8-742e-4ae4-9ad5-c7eacf50824a.jpg" class="post-content__image">', now(), 2);
INSERT INTO posts (user_id, content, date_created, tier) VALUES (2, '<img src="https://wsrv.nl/?url=http://95.163.209.195:9000/e1671797c52e15f763380b45e841ec32/35aae3f0-cfd5-4f73-b8da-33fc451f3a2e.jpg" class="post-content__image">', now(), 3);

INSERT INTO author_subscriptions (author_id, img, tier, title, text, price) VALUES (2, 'a5293e04-74a4-4ca2-98b3-baf7425e1741.jpg', 1, 'Мемы', 'Обычная подписка на мемы АНДЭВС', 1990);

