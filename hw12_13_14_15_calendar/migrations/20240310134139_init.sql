-- +goose Up
CREATE table events (
                       id              uuid primary key,
                       user_id         uuid,
                       title           text,
                       date_and_time   timestamp with time zone,
                       duration        bigint,
                       description     text,
                       time_to_notify  bigint
);

-- +goose Down
drop table events;