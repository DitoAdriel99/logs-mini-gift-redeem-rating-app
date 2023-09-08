-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS logs_user
(
    id                           varchar        not null primary key,
    fullname                     varchar        not null default '-',
    email                        varchar        not null default '-',
    event                        varchar        not null default '-',
    user_agent                   TEXT           not null default '-',
    http_status_code             int            not null default 0,
    http_method                  varchar        not null default '-',
    client_request_data          json,
    client_response_data         json,
    created_at                   timestamptz    not null default now(),
    updated_at                   timestamptz    not null default now()
);

CREATE INDEX IF NOT EXISTS idx_logs_user ON logs_user
(
    fullname,
    email,
    event,
    http_status_code,
    http_method,
    created_at,
    updated_at
);

CREATE INDEX IF NOT EXISTS idx_logs_user_fullname           ON logs_user(fullname);
CREATE INDEX IF NOT EXISTS idx_logs_user_email              ON logs_user(email);
CREATE INDEX IF NOT EXISTS idx_logs_user_event              ON logs_user(event);
CREATE INDEX IF NOT EXISTS idx_logs_user_http_status_code   ON logs_user(http_status_code);
CREATE INDEX IF NOT EXISTS idx_logs_user_http_method        ON logs_user(http_method);
CREATE INDEX IF NOT EXISTS idx_logs_user_created_at         ON logs_user(created_at);
CREATE INDEX IF NOT EXISTS idx_logs_user_updated_at         ON logs_user(updated_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_logs_user_fullname;
DROP INDEX IF EXISTS idx_logs_user_email;
DROP INDEX IF EXISTS idx_logs_user_event;
DROP INDEX IF EXISTS idx_logs_user_http_method;
DROP INDEX IF EXISTS idx_logs_user_created_at;
DROP INDEX IF EXISTS idx_logs_user_updated_at;
DROP INDEX IF EXISTS idx_logs_user;

DROP TABLE IF EXISTS logs_user;
-- +goose StatementEnd
