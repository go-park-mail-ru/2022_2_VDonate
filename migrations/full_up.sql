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

INSERT INTO users (username, email)
VALUES ('Music Geek', 'music@geek.mp4');
INSERT INTO user_info (user_id, password, is_author, avatar, about)
VALUES (5, 'musicgeek', true, '656f88ea-aaf8-4195-a1e6-40d7f8d6b169.12.28.jpeg', 'Самые свежие новости из мира музыки. Подписывайтесь и не пропустите ничего интересного!');
INSERT INTO author_subscriptions (author_id, img, tier, title, text, price)
VALUES (5, 'a6b2e71d-ff9e-4598-b4bf-8635136057bb.16.46.jpeg', 1, 'Топы', 'Подписка на музыкальные топы.', 1);
INSERT INTO author_subscriptions (author_id, img, tier, title, text, price)
VALUES (5, '4c6551fd-648f-440c-b43d-8fa7a7e54ce6.17.58.jpeg', 2, 'Обзоры', 'Подписка на обзоры.', 2);
INSERT INTO author_subscriptions (author_id, img, tier, title, text, price)
VALUES (5, 'a6b2e71d-ff9e-4598-b4bf-8635136057bb.16.46.jpeg', 3, 'Новости', 'Подписка на музыкальные новости.', 3);
INSERT INTO posts (user_id, content, date_created, tier)
VALUES (5, '<h1>Топ 10 лучших альбомов 2022 года</h1><br>
<p>Музыка – одна из немногих вещей, которая спасала от состояния тревоги. Какая бы не была нагрузка извне, какие бы шрамы не появлялись на нашей душе – песни ее исцелят. Вот с чем я провел 2022-й год. Ценю все перечисленное ниже.</p><br>
<h2>5. The Weeknd – dawn FM</h2><br><p>Фан факт: ни один из треков этого альбома не находится в десятке самых популярных треков певца в настоящее время, пока он занимает первую строчку в рейтинге артистов Spotify по месячной аудитории. Значит ли это, что альбом забыли или что он тупо не сработал? Вовсе нет. Просто это гиммиковый альбом-концепция, который намного ровнее звучит в формате прослушивания от и до. Практически нулевое промо, резкая смена стиля артиста и выход в самый спящий сезон года.</p><img src="https://cloud.vdonate.ml/1679091c5a880faf6fb5e6087eb1b2dc/4599ac2b-61ce-43b7-8808-45f6efdccb86.jpeg" class="post-content__image"><br>
<h2>4. Fontaines D.C. – Skinty Fia</h2><br><p>Британская меланхолия без слабых мест. Причем у этих ребят прям очень хорошо с сонграйтингом: некоторые песни прочно заседают в голове («I Love You», «Big Shot», проверено спустя полгода после выхода пластинки), другие работают исключительно на общую атмосферу, но филлерами их не назвать. Если бы фильм «Банши Инишерина» снимали в современной Ирландии в городских пейзажах, то эта музыка стала бы саундтреком.</p1><img src="https://cloud.vdonate.ml/4a8a08f09d37b73795649038408b5f33/73a6e52c-2f35-4853-8a17-28d5dae6421c.png" class="post-content__image"><br>
<h2>3. Black Country, New Road – Ants From Up There</h2><br><p>Если Black Midi рубят как не в себя (порой даже пугающе), то новый альбом BCNR разыгран очень аккуратно и по нотам. Всего по чуть-чуть: британская поэзия, будничные истории, что-то из фолка, что-то от Ника Кейва, что-то из неоклассики, проглядывается даже эмо-культура. Очень нравится, что треки «Snow Globes» и «Haldern» заканчиваются как фильм «Одержимость». Очень нравится, что у музыки существует кульминация. Очень нравится, что инструментам дают дышать и прерывать вокалиста, когда в голосе нет нужды. Единственное чего не хватило — нет сильных песен, которые можно было бы запомнить. «21st Century Schizoid Man» даже спустя 50 лет звучит как бэнгер. Здесь такого нет.</p><img src="https://cloud.vdonate.ml/8f14e45fceea167a5a36dedd4bea2543/2fce2a03-eee7-4a39-bbd3-cc0de57da267.png" class="post-content__image"><br>
<h2>2. The Smile – A Light For Attracting Attention</h2><p>Единственный альбом из списка, который мне удалось услышать вживую (наконец-то флекс). Некоторые утверждают, что The Smile это худший альбом Radiohead – это ошибочное мнение. Но только наполовину. Назвать группу сайд-проектом довольно сложно: здесь и привычная манера прятать высказывания в текст, все те же завывания, да и Гринвуд не пытался переизобрести себя. Тем не менее, переезд на вторую полосу участников любого фестиваля (Radiohead всегда достается главная сцена) будто бы сделал звук музыкантов более раскованным – без монументальности стало чуточку проще, но все так же очаровательно.</p><img src="https://cloud.vdonate.ml/e4da3b7fbbce2345d7772b0674a318d5/80bd14f9-4d38-4321-8a40-2ff7aed4c645.png" class="post-content__image"><br>
<h2>1. black midi – Hellfire</h2><br><p>Заслужили! В своем третьем альбоме экстравагантные британцы исправили все ошибки двух предыдущих. Они намешали все: от фри джаза до заходов в сторону фламенко, от нежного пения под кантри-мотив до истеричного речитатива под инструментальный визг. При этом их альбом получился в принципе очень легким для прослушивания. Любые стихийные перегрузы дополняют музыкальный нарратив, усиливая погружение слушателя. Это повышает ценность моментов для передышки — иногда создается такая нагрузка, когда ты прям физически выдыхаешь при переходе в более спокойную часть песни. И это очень круто! Потому что сегодня мало кто использует звук как психологическое давление. Black Midi сами по себе похожи на полноценный оркестр, когда получают карт-бланш на музыкальный хаос, поэтому в «Hellfire» звучат самые дикие вещи, которые создавали пост-брэкзитные рок-группы. Это боксерский поединок «Sugar/Tzu» с яростью саксофона под конец и «The Race Is About To Begin» (невероятная вещь под середину). При этом они каким-то образом без трудностей воспроизводят эту зарубу в живых перформансах. Как итог хочу сказать, что «Hellfire» — чуть ли не идеальный пример нового альбома с технически сложным исполнением, который способен развлекать слушателя, а не просто предоставлять тому для осмысления несвязные, но зато «прогрессивные» куски ломанных мелодий. Однозначно один из лучших релизов 2022 года.</p><img src="https://cloud.vdonate.ml/8277e0910d750195b448797616e091ad/64ed3cea-f5d4-4239-864e-139c52f03c3d.jpeg" class="post-content__image"><br>',
now(), 1);
INSERT INTO posts (user_id, content, date_created, tier)
VALUES (5, '<img src="https://cloud.vdonate.ml/e4da3b7fbbce2345d7772b0674a318d5/fdbb74cb-6351-4e12-9850-33a8cfdced35.jpeg" class="post-content__image"><br>
<p>«OK Computer» — третий студийный альбом английской рок-группы Radiohead, выпущенный в 1997 году. Альбом ознаменовал отход от прежнего гитарного звучания группы и часто упоминается как знаковый релиз в эволюции альтернативного рока.</p>
<p>Альбом исследует темы изоляции, технологии и современного общества, а его тексты затрагивают экзистенциальные и политические темы. Песни на «OK Computer» сложные и многослойные, с замысловатыми мелодиями и гитарной работой, а также с использованием электронных инструментов и манипуляций.</p>
<p>В музыкальном плане «OK Computer» представляет собой мастерское сочетание рока и электронных влияний, при этом участники группы расширяют границы своих инструментов, создавая уникальное и инновационное звучание. Тексты песен, написанные Томом Йорком, столь же амбициозны и заставляют задуматься, исследуя широкий спектр тем, включая отчуждение, дегуманизирующее влияние технологий и опасность беглого капитализма.</p>
<p>В целом,«OK Computer» — это знаковый альбом, амбициозный как в музыкальном, так и в лирическом плане. Это мощное исследование тревог и неопределенности современной жизни, и он остается непреходящим и актуальным произведением искусства.Текст написан компьютерной программой ChatGPT и переведен машинным переводчиком DeepL.</p>',
now(), 2);
INSERT INTO posts (user_id, content, date_created, tier)
VALUES (5, '<img src="https://cloud.vdonate.ml/92eb5ffee6ae2fec3ad71c777531578f/f1af4755-2bf0-41df-8505-f8ea482978bb.jpeg" class="post-content__image"><br>
<p>Дуа Липа завершила эру «Future Nostalgia» и заявила, что ближайшие месяцы проведет в студии с целью завершить работу над третьим альбомом</p>
<p>Одноименный релиз вышел в начале 2020 года — диско-звук превратил певицу в звезду мирового масштаба. На мой вкус, это одна из лучших поп-записей за последние десять лет, которая будет звучать актуально и в будущем. В рамках мирового тура Липа отыграла девять десятков выступлений и ориентировочно заработала 40 миллионов долларов, не говоря уже о деньгах со стриминговых платформ и физических продаж.</p>
<p>Буду сильно разочарован, если следующий альбом не удастся.</p>',
now(), 3);

INSERT INTO users (username, email)
VALUES ('GitHub Community', 'git@hub.com');
INSERT INTO user_info (user_id, password, is_author, avatar, about)
VALUES (6, 'githubcommunity', true, '3500bd0c-8eb5-4d84-ad3f-ff0ce3057bf6.jpeg', 'Сообщество пользователей GitHub');
INSERT INTO author_subscriptions (author_id, img, tier, title, text, price)
VALUES (6, 'b0e4a531-520e-4179-a428-229e71e8f812.png"', 1, 'Полезный софт', 'Делимся интересными проектами.', 1);
INSERT INTO posts (user_id, content, date_created, tier)
VALUES (6, '<h1><a href="https://github.com/kimlimjustin/xplorer">​​Xplorer</a> – красивый файловый менеджер, созданный с нуля для полной настройки.</h1><br>
<ul><li>Работает на Windows, GNU/Linux и MacOS.</li><li><li>Поддерживает предварительный просмотр не только изображений или документов, но и видео</ul><br>
<img src="https://cloud.vdonate.ml/cfcd208495d565ef66e7dff9f98764da/8b1c5c3f-eddc-479d-8842-0f7d219588c0.jpeg" class="post-content__image">', now(), 1);
INSERT INTO posts (user_id, content, date_created, tier)
VALUES (6, '<h1><a href="https://github.com/jmechner/Prince-of-Persia-Apple-II">​Prince of Persia Apple II</a> – Исходный код игры из 1985-го для компьютеров Apple 2</h1><br>
<p>Многие вспомнят культовую серию</p><br>
<img src="https://cloud.vdonate.ml/0cc175b9c0f1b6a831c399e269772661/0c1cd1f6-5d64-4b43-ab01-1030a2d6c84a.jpeg" class="post-content__image">', now(), 1);
INSERT INTO posts (user_id, content, date_created, tier)
VALUES (6, '<h1><a href="https://github.com/GyulyVGC/sniffnet">​Sniffnet</a> – полезное приложение, которое позволяет вам легко и привлекательно взглянуть на свой сетевой трафик в режиме реального времени</h1><br>
<p>Работает в Windows, GNU/Linux, Mac</p><br>
<img src="https://cloud.vdonate.ml/a87ff679a2f3e71d9181a67b7542122c/12439190-a769-4c22-916b-2e8c5f6fd954.jpeg" class="post-content__image">', now(), 1);

INSERT INTO users (username, email)
VALUES ('IT Юмор', 'it@umor.com');
INSERT INTO user_info (user_id, password, is_author, avatar, about)
VALUES (7, 'itumor', true, '9cf7fcdd-3f31-4b03-9162-7dacc3fecf86.jpeg', 'IT Мемы');
INSERT INTO author_subscriptions (author_id, img, tier, title, text, price)
VALUES (7, 'b0e4a531-520e-4179-a428-229e71e8f812.png"', 1, 'Полезный софт', 'Делимся интересными проектами.', 1);
INSERT INTO posts (user_id, content, date_created, tier)
VALUES (7, '​— А кем вы работало до того, как стать строителем?<br>— Фронтенд-разработчиком<br>— Так и думал...<br>
<img src="https://cloud.vdonate.ml/e4da3b7fbbce2345d7772b0674a318d5/4ecc7f60-bb0d-4b32-a250-74dce951c3d5.jpeg" class="post-content__image">', now(), 1);
INSERT INTO posts (user_id, content, date_created, tier)
VALUES (7, '<img src="https://cloud.vdonate.ml/8fa14cdd754f91cc6554c9e71929cce7/751d7c3e-4e01-4bf8-9e17-469a27dde80f.jpeg" class="post-content__image">', now(), 1);
INSERT INTO posts (user_id, content, date_created, tier)
VALUES (7, '<img src="https://cloud.vdonate.ml/8f14e45fceea167a5a36dedd4bea2543/23bcbc8f-5570-4083-997f-629316859597.jpeg" class="post-content__image">', now(), 1);
INSERT INTO posts (user_id, content, date_created, tier)
VALUES (7, '<img src="https://cloud.vdonate.ml/a87ff679a2f3e71d9181a67b7542122c/65f4cef8-dddd-4b88-9a6f-e1817902c0a4.jpeg" class="post-content__image">', now(), 1);
