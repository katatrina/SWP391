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
    full_name       VARCHAR(50)  NOT NULL,
    email           VARCHAR(150) NOT NULL,
    phone           CHAR(10)     NOT NULL,
    address         VARCHAR(200) NOT NULL,
    role_id         INT          NOT NULL,
    hashed_password CHAR(60)     NOT NULL,
    created_at      timestamptz  NOT NULL DEFAULT NOW()
);

ALTER TABLE users
    ADD CONSTRAINT users_uc_phone UNIQUE (phone);

ALTER TABLE users
    ADD CONSTRAINT users_uc_email UNIQUE (email);

ALTER TABLE users
    ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("id");

CREATE TABLE providerDetails
(
    id           SERIAL PRIMARY KEY,
    user_id      INT          NOT NULL UNIQUE,
    company_name VARCHAR(100) NOT NULL,
    tax_code     VARCHAR(50)  NOT NULL,
    created_at   timestamptz  NOT NULL DEFAULT NOW()
);

ALTER TABLE providerDetails
    ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

CREATE TABLE sessions
(
    token  TEXT PRIMARY KEY,
    data   BYTEA       NOT NULL,
    expiry TIMESTAMPTZ NOT NULL
);

CREATE INDEX sessions_expiry_idx ON sessions (expiry);

CREATE TABLE category
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(100)
);

CREATE TABLE services
(
    id               SERIAL PRIMARY KEY,
    title            VARCHAR(350) NOT NULL,
    description      TEXT         NOT NULL,
    price            INT          NOT NULL,
    genre            VARCHAR(50)  NOT NULL,
    thumbnail_url    TEXT         NOT NULL,
    category_id      INT          NOT NULL,
    owned_by_user_id INT          NOT NULL,
    status           TEXT         NOT NULL DEFAULT 'inactive',
    created_at       timestamptz  NOT NULL DEFAULT NOW()
);

ALTER TABLE "services"
    ADD FOREIGN KEY ("category_id") REFERENCES "category" ("id");

ALTER TABLE "services"
    ADD FOREIGN KEY ("owned_by_user_id") REFERENCES "users" ("id");

CREATE TABLE feedbacks
(
    id         SERIAL PRIMARY KEY,
    service_id INT         NOT NULL,
    user_id    INT         NOT NULL,
    content    TEXT        NOT NULL,
    created_at timestamptz NOT NULL DEFAULT NOW()
);

ALTER TABLE "feedbacks"
    ADD FOREIGN KEY ("service_id") REFERENCES "services" ("id");

ALTER TABLE "feedbacks"
    ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

CREATE TABLE blogs
(
    id         SERIAL PRIMARY KEY,
    title      VARCHAR(350) NOT NULL,
    content    TEXT         NOT NULL,
    created_at timestamptz  NOT NULL DEFAULT NOW()
);

CREATE TABLE orders
(
    id            SERIAL PRIMARY KEY,
    buyer_id      INT         NOT NULL,
    seller_id     INT         NOT NULL,
    delivery_date timestamptz NOT NULL,
    delivered_to  TEXT        NOT NULL,
    status        TEXT        NOT NULL,
    total         INT         NOT NULL DEFAULT 0,
    created_at    timestamptz NOT NULL DEFAULT NOW()
);

ALTER TABLE "orders"
    ADD FOREIGN KEY ("buyer_id") REFERENCES "users" ("id");

ALTER TABLE "orders"
    ADD FOREIGN KEY ("seller_id") REFERENCES "users" ("id");

CREATE TABLE orderDetails
(
    id         SERIAL PRIMARY KEY,
    order_id   INT         NOT NULL,
    service_id INT         NOT NULL,
    quantity   INT         NOT NULL,
    price      INT         NOT NULL,
    created_at timestamptz NOT NULL DEFAULT NOW()
);