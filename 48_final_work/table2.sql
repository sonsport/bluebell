create table table2
(
    id             int auto_increment
        primary key,
    user_id   int unsigned                        not null auto_increment,
    username   varchar(64)                        not null,
    create_time    timestamp default CURRENT_TIMESTAMP not null
)
    collate = utf8mb4_general_ci;

INSERT INTO table2 (id, user_id, username, create_time) VALUES (1, 1, '七米', '2020-07-12 13:03:51');
INSERT INTO table2 (id, user_id, username, create_time) VALUES (2, 2, '七米', '2020-07-12 13:03:52');
INSERT INTO table2 (id, user_id, username, create_time) VALUES (3, 3, 'wanqingyun', '2020-07-12 13:03:52');
INSERT INTO table2 (id, user_id, username, create_time) VALUES (4, 4, 'wanqingyun', '2020-07-12 13:03:53');