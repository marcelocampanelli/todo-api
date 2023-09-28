CREATE TABLE IF NOT EXISTS tasks (
    id         varchar(36) NOT NULL PRIMARY KEY,
    name       varchar(255) NOT NULL,
    finished   boolean NOT NULL,
    created_at timestamp NOT NULL ,
    updated_at timestamp NOT NULL,
    user_id   varchar(36) NOT NULL REFERENCES users(id)
    )