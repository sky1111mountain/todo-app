create table if not exists todolist (
    id integer unsigned auto_increment primary key,
    task varchar(100) not null,
    priority varchar(10) not null,
    status varchar(10) not null,
    username varchar(100) not null,
    created_at datetime
);

insert into todolist (task, priority, status, username, created_at) 
    values("Test Data 1", "high", "not_done", "rainbow777", now());

insert into todolist (task, priority, status, username, created_at) 
    values("Test Data 2", "medium", "done", "rainbow777", now());

insert into todolist (task, priority, status, username, created_at) 
    values("Test Data 3", "low", "not_done", "rainbow777", now());