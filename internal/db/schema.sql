CREATE TABLE roles
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL
);

INSERT INTO roles (name)
VALUES ('customer'),
       ('provider');

CREATE TABLE users
(
    id              SERIAL PRIMARY KEY,
    name            VARCHAR(50)  NOT NULL,
    phone           CHAR(10)     NOT NULL,
    address         VARCHAR(350) NOT NULL DEFAULT '',
    role_id         INT          NOT NULL,
    hashed_password CHAR(60)     NOT NULL,
    created_at      timestamptz  NOT NULL DEFAULT NOW()
);

ALTER TABLE users
    ADD CONSTRAINT users_uc_phone UNIQUE (phone);

ALTER TABLE users
    ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("id");

CREATE TABLE sessions
(
    token  TEXT PRIMARY KEY,
    data   BYTEA       NOT NULL,
    expiry TIMESTAMPTZ NOT NULL
);

CREATE INDEX sessions_expiry_idx ON sessions (expiry);

CREATE TABLE categorys
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR
);

CREATE TABLE genres
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR
);

CREATE TABLE services
(
    id               SERIAL PRIMARY KEY,
    title            VARCHAR(350) NOT NULL,
    description      TEXT         NOT NULL,
    price            INT          NOT NULL,
    genre_id         INT          NOT NULL,
    category_id      INT          NOT NULL,
    owned_by_user_id INT          NOT NULL,
    created_at       timestamptz  NOT NULL DEFAULT NOW()
);

ALTER TABLE "services"
    ADD FOREIGN KEY ("category_id") REFERENCES "categorys" ("id");

ALTER TABLE "services"
    ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

CREATE TABLE orders
(
    id            SERIAL PRIMARY KEY,
    buyer_id      INTEGER     NOT NULL,
    seller_id     INTEGER     NOT NULL,
    delivery_date timestamptz NOT NULL,
    delivered_to  TEXT        NOT NULL,
    status        TEXT        NOT NULL,
    total         INTEGER     NOT NULL DEFAULT 0,
    created_at    timestamptz NOT NULL DEFAULT NOW()
);

ALTER TABLE "orders"
    ADD FOREIGN KEY ("buyer_id") REFERENCES "users" ("id");

ALTER TABLE "orders"
    ADD FOREIGN KEY ("seller_id") REFERENCES "users" ("id");

CREATE TABLE orderDetails
(
    id         SERIAL PRIMARY KEY,
    order_id   INTEGER     NOT NULL,
    service_id INTEGER     NOT NULL,
    quantity   INTEGER     NOT NULL,
    price      INTEGER     NOT NULL,
    created_at timestamptz NOT NULL DEFAULT NOW()
);