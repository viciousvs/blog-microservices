CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
create table if not exists post(
    id uuid primary key,
    title varchar(255) not null,
    content text not null,
    created_at bigint not null,
    updated_at bigint not null
);

insert into post(id, title, content, created_at, updated_at)
values
(uuid_generate_v4(), 'title 1', 'content 1', 1, 1),
(uuid_generate_v4(), 'title 2', 'content 2', 1, 1),
(uuid_generate_v4(), 'title 3', 'content 3', 1, 1),
(uuid_generate_v4(), 'title 4', 'content 4', 1, 1),
(uuid_generate_v4(), 'title 5', 'content 5', 1, 1);