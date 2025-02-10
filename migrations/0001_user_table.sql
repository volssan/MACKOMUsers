-- +goose Up
-- +goose StatementBegin
create table users (
    id             serial8     not null primary key,
    first_name     text        not null,
    last_name      text        not null,
    age            int         not null,
    recording_date timestamptz not null default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists users;
-- +goose StatementEnd
