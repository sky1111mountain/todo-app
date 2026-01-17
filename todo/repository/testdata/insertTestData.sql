create table if not exists todolist (
    id integer unsigned auto_increment primary key,
    task varchar(100) not null,
    priority varchar(10) not null,
    status varchar(10) not null,
    username varchar(100) not null,
    created_at datetime
);