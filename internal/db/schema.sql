CREATE TABLE users
(
    id              SERIAL PRIMARY KEY,
    name            VARCHAR(50) NOT NULL,
    email           VARCHAR(50) NOT NULL,
    hashed_password CHAR(60)    NOT NULL,
    created_at      timestamptz NOT NULL
);

ALTER TABLE users
    ADD CONSTRAINT users_uc_email UNIQUE (email);

CREATE TABLE sessions
(
    token  TEXT PRIMARY KEY,
    data   BYTEA       NOT NULL,
    expiry TIMESTAMPTZ NOT NULL
);

CREATE INDEX sessions_expiry_idx ON sessions (expiry);