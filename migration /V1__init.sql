CREATE TABLE note
(
    id          uuid primary key,
    title       varchar(255) not null,
    description varchar(255) not null
)