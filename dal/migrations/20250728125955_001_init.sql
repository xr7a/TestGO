-- +goose Up
create table if not exists departments(
    id serial primary key,
    name varchar not null,
    phone varchar not null
);

create table if not exists users(
    id serial primary key,
    name varchar not null,
    surname varchar not null,
    phone varchar not null,
    company_id int not null,
    department_id int references departments(id)
);

create table if not exists passports(
    id serial primary key,
    user_id int references users(id) on delete cascade,
    type varchar not null,
    number varchar not null
);

-- +goose Down
drop table if exists users;
drop table if exists passports;
drop table if exists departments;
