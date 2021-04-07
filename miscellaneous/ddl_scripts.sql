create table if not exists registered_users
(
    id         uuid
        constraint registered_users_pk
            primary key,
    first_name varchar(20) not null,
    last_name  varchar(30) not null
);
