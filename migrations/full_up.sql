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
        UPDATE user_info SET balance = balance + (SELECT price FROM author_subscriptions as aSub WHERE aSub.id = new.sub_id) WHERE user_id = new.to_id;
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
/*
    https://vk.com/abstract_memes
    https://vk.com/tproger
    https://vk.com/mozgitreski
    https://vk.com/kinoartmag
*/
INSERT INTO users (username, email)
VALUES ('АМДЭВС', 'amdevs@mail.ru');
INSERT INTO user_info (user_id, password, is_author, avatar, about)
VALUES (1, 'amdevs', true, '430fd690-ad80-49bc-b46e-f3b547367781.jpg', 'у нас тут новая искренность');

INSERT INTO users (username, email)
VALUES ('Типичный программист', 'programist@mail.ru');
INSERT INTO user_info (user_id, password, is_author, avatar, about)
VALUES (2, 'programist', true, 'd42ab69a-e22b-4cad-8538-83bd4457a047.jpg', 'Tproger — издание о разработке и обо всём, что с ней связано.

Читайте нас в Telegram: https://t.me/tproger_official

Реклама на Tproger: https://tprg.ru/yCVP
Размещение вакансий: https://tprg.ru/poUM
HR-брендинг: https://tprg.ru/L2IJ

Яндекс.Дзен: https://zen.yandex.ru/tproger
');

INSERT INTO users (username, email)
VALUES ('Мозги трески', 'brains@mail.ru');
INSERT INTO user_info (user_id, password, is_author, avatar, about)
VALUES (3, 'brains', true, '4a28db1c-1696-4650-932f-f2b530c0af4d.jpg', 'Мозги трески - это смешные комикс-стрипы обо всём на свете: бытовые зарисовки, пародии на кино и игры, всякие безумные ситуации - в общем, все, что взбредет в голову автору.

Новый комикс каждый понедельник в 21:00 и иногда ВНЕЗАПНО в другие дни.

Почти 100 привидений - это история о старинном заброшенном доме, в котором обитает около сотни призраков - каждый со своим характером и историей.

Новый выпуск каждые среду и пятницу в 21:00.

А еще у меня вышел комикс-игра «Домашние монстры» про мальчика, обнаружившего в своей комнате монстров. Экземпляр можно купить в любых книжных магазинах и комикс-шопах страны.');

INSERT INTO users (username, email)
VALUES ('Искусство кино', 'art@mail.ru');
INSERT INTO user_info (user_id, password, is_author, avatar, about)
VALUES (4, 'art', true, '0c3b90e0-801d-4936-9e88-f51bfbee4bd8.jpg',
        '«Искусство кино» — журнал о кино, который издается с января 1931 года и остается одним из старейших в мире периодических изданий о кино. Каждый номер освещает актуальные вопросы кинематографа и визуальной культуры, а также включает в себя редкие архивные публикации, обзоры крупнейших международных кинофестивалей, русскую и зарубежную кинопрозу и другие материалы.');

INSERT INTO posts (user_id, content, date_created, tier)
VALUES (1,
        '<img src="https://wsrv.nl/?url=http://95.163.209.195:9000/e4da3b7fbbce2345d7772b0674a318d5/e97ca124-4237-4076-83b4-3484f93794f5.jpg" class="post-content__image">',
        now(), 1);
INSERT INTO posts (user_id, content, date_created, tier)
VALUES (1,
        '<img src="https://wsrv.nl/?url=http://95.163.209.195:9000/1679091c5a880faf6fb5e6087eb1b2dc/7dfae8ba-c70e-446d-b1b4-080b62fad7f6.jpg" class="post-content__image">',
        now(), 1);
INSERT INTO posts (user_id, content, date_created, tier)
VALUES (1,
        '<img src="https://wsrv.nl/?url=http://95.163.209.195:9000/4a8a08f09d37b73795649038408b5f33/c3c5d321-4028-4f21-afaa-572436d76c7c.jpg" class="post-content__image">',
        now(), 1);
INSERT INTO posts (user_id, content, date_created, tier)
VALUES (1,
        '<img src="https://wsrv.nl/?url=http://95.163.209.195:9000/0cc175b9c0f1b6a831c399e269772661/fc1e28e8-742e-4ae4-9ad5-c7eacf50824a.jpg" class="post-content__image">',
        now(), 1);
INSERT INTO posts (user_id, content, date_created, tier)
VALUES (1,
        '<img src="https://wsrv.nl/?url=http://95.163.209.195:9000/e1671797c52e15f763380b45e841ec32/35aae3f0-cfd5-4f73-b8da-33fc451f3a2e.jpg" class="post-content__image">',
        now(), 1);

INSERT INTO author_subscriptions (author_id, img, tier, title, text, price)
VALUES (1, 'a5293e04-74a4-4ca2-98b3-baf7425e1741.jpg', 1, 'Мемы', 'Обычная подписка на мемы АНДЭВС', 1990);

INSERT INTO posts (user_id, content, date_created, tier)
VALUES (2,
        'Каналы по конкретным направлениям разработки, подборки полезных ресурсов, канал для начинающих, новости, мемы — в экосистеме Tproger есть канал для каждого разработчика.<div><br></div><img src="https://wsrv.nl/?url=http://95.163.209.195:9000/92eb5ffee6ae2fec3ad71c777531578f/69176f42-a6b9-4305-80c0-667642b916db.jpg" class="post-content__image"><div><br></div>Если ещё не подписаны на какой-то из каналов по интересующей вас теме, то исправляйте это: https://t.me/tproger_channels<div><br></div>',
        now(), 1);
INSERT INTO posts (user_id, content, date_created, tier)
VALUES (2,
        '<h1>«Я попробовал, не получилось»: Mail отказался от собственного поискового движка, теперь за поиск отвечают алгоритмы «Яндекса».</h1><div><br></div><img src="https://wsrv.nl/?url=http://95.163.209.195:9000/e1671797c52e15f763380b45e841ec32/c9c10784-f6ea-41bd-b667-1dac9f0bba5e.jpg" class="post-content__image"><div><br></div>Mail с 2013 года пытался развивать собственные поисковые технологии. Но развитие поискового движка не вошло в новую стратегию компании. В Холдинге VK решили сделать упор на развитие контентных сервисов: «Мы постарались сохранить привычный для пользователей интерфейс с использованием поиска от „Яндекс“ и надеемся, что опыт использования нового решения будет результативным и приятным»

К слову, поиск Mail потерял не многое — его доля составляла всего 0,21% от российского рынка. Сейчас в лидерах — «Яндекс» (51,86%) и по-прежнему Google (45,1%).

Где теперь искать, как удалить браузер Амиго?',
        now(), 1);
INSERT INTO posts (user_id, content, date_created, tier)
VALUES (2,
        '<h1>Что делать, если у вас команде человек «Всё — г… но»?</h1><div><br></div><img src="https://wsrv.nl/?url=http://95.163.209.195:9000/45c48cce2e2d7fbdea1afc51c7c6ad26/a04462a2-5cde-49b5-80a5-b1b251c47b39.jpg" class="post-content__image"><div><br></div>Наверное, каждый из нас сталкивался с людьми, которым в компании почти ничего не нравится. Они выступают против большинства инициатив. А когда к ним обращаются за помощью, они топят встречными вопросами или просто очень медленно выполняют задачу. Да и вообще, «вокруг одни долбоящеры, процессы дебильные, а про менеджмент лучше промолчать».

Интересная статья на Хабре, в которой рассказали, как нейтрализовать такого коллегу или вовсе обратить его суперсилу на пользу делу. И что делать, когда вы узнали такого коллегу в себе: https://habr.com/ru/company/jetinfosystems/blog/699940/',
        now(), 2);
INSERT INTO posts (user_id, content, date_created, tier)
VALUES (2,
        '<h1>Главное, чтобы прямо там не отправил пилота на тренинги по повышению мотивации.</h1><div><br></div><img src="https://wsrv.nl/?url=http://95.163.209.195:9000/4a8a08f09d37b73795649038408b5f33/5d1caccb-ddd6-4759-b89f-fc62215aa5dc.jpg" class="post-content__image"><div><br></div>',
        now(), 2);
INSERT INTO posts (user_id, content, date_created, tier)
VALUES (2,
        '<h1>Подборка актуальных вакансий:</h1><div><br></div>

— Hadoop-администратор: https://tprg.ru/6qIu<div><br></div>
Где: Москва, можно удалённо<div><br></div>
Опыт: от 1 года<div><br></div>

— Java-разработчик: https://tprg.ru/bSVW<div><br></div>
Где: Москва, можно удалённо<div><br></div>
Опыт: от 3 лет<div><br></div>

— Прикладной администратор по поддержке фронтальных систем: https://tprg.ru/axXA<div><br></div>
Где: Москва, можно удалённо<div><br></div>
Опыт: от 3 лет<div><br></div>

— Middle DBA: https://tprg.ru/mx4R<div><br></div>
Где: Москва, можно удалённо<div><br></div>
Опыт: от 3 лет<div><br></div>

— Разработчик 1C (Senior / Lead): https://tprg.ru/TdXe<div><br></div>
Где: Москва, Санкт-Петербург, Ростов-на-Дону<div><br></div>
Опыт: от 3 лет<div><br></div>

— Senior Golang-разработчик: https://tprg.ru/AwNE<div><br></div>
Где: удалённо<div><br></div>
Опыт: от 3 лет<div><br></div>

— Ведущий Java-разработчик: https://tprg.ru/AR3q<div><br></div>
Где: удалённо<div><br></div>
Опыт: от 3 лет<div><br></div>

— Главный разработчик: https://tprg.ru/LEAF<div><br></div>
Где: Москва, можно удалённо<div><br></div>
Опыт: от 3 лет<div><br></div>

— Тимлид разработки: https://tprg.ru/kP9n<div><br></div>
Где: Москва, можно удалённо<div><br></div>
Опыт: от 3 лет<div><br></div>

— Руководитель разработки: https://tprg.ru/74xk<div><br></div>
Где: Москва, можно удалённо<div><br></div>
Опыт: от 4 лет<div><br></div>

— Senior Application Security Engineer: https://tprg.ru/CqRc<div><br></div>
Где: Москва, можно удалённо<div><br></div>
Опыт: от 5 лет<div><br></div>',
        now(), 3);

INSERT INTO author_subscriptions (author_id, img, tier, title, text, price)
VALUES (2, '4819996a-9f74-42bb-afcc-bf8ffddd85df.jpg', 1, 'Новости', 'Подписка на IT новости.', 100);

INSERT INTO author_subscriptions (author_id, img, tier, title, text, price)
VALUES (2, 'bbdef463-c60d-41d4-8a87-25c5526d3c15.jpg', 2, 'Мемы', 'Подписка на IT мемы.', 500);

INSERT INTO author_subscriptions (author_id, img, tier, title, text, price)
VALUES (2, 'bbdef463-c60d-41d4-8a87-25c5526d3c15.jpg', 3, 'Вакансии', 'Еженедельные подборки вакансий.', 1990);

CREATE TABLE IF NOT EXISTS comments
(
    id           bigserial     NOT NULL PRIMARY KEY,
    user_id      bigserial     NOT NULL REFERENCES users (id) ON DELETE NO ACTION,
    post_id      bigserial     NOT NULL REFERENCES posts (post_id) ON DELETE CASCADE,
    content      varchar(1024) NOT NULL,
    date_created TIMESTAMP     NOT NULL DEFAULT now()
);
