create table users
(
    id         int unsigned primary key not null auto_increment,
    first_name varchar(255) not null,
    last_name  varchar(255) not null,
    created_at datetime     not null
);
