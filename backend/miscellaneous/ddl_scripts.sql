drop table registered_users;

create table if not exists registered_users
(
    id         uuid
        constraint registered_users_pk
            primary key,
    first_name varchar(20) not null,
    last_name  varchar(30) not null,
    "email"  varchar(50) not null unique,
    "password"  varchar(100) not null
);
