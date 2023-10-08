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

CREATE TABLE categories
(
    id            SERIAL PRIMARY KEY,
    name          VARCHAR(100) NOT NULL,
    slug          VARCHAR(100) NOT NULL,
    thumbnail_url TEXT         NOT NULL,
    description   TEXT         NOT NULL
);

ALTER TABLE categories
    ADD CONSTRAINT categories_uc_slug UNIQUE (slug);

INSERT INTO categories (name, slug, thumbnail_url, description)
VALUES ('Phụ kiện', 'phu-kien', '/static/img/accessories-category.jpg',
        'Các dịch vụ cung cấp đầy đủ các sản phẩm cần thiết cho chim cảnh'),
       ('Dinh dưỡng và thức ăn', 'dinh-duong-va-thuc-an', '/static/img/nutrition-category.jpg',
        'Các dịch vụ cung cấp đầy đủ các dinh dưỡng cần thiết cho chim cảnh'),
       ('Y tế và chăm sóc sức khỏe', 'y-te-va-cham-soc-suc-khoe', '/static/img/healthcare-category.jpg',
        'Đảm bảo sức khỏe và phòng ngừa bệnh tật cho chim cảnh'),
       ('Grooming', 'grooming', '/static/img/grooming-category.jfif',
        'Chăm sóc, tạo phong cách và làm đẹp cho chim cảnh'),
       ('Đào tạo và huấn luyện', 'dao-tao-va-huan-luyen', '/static/img/training-category.jpg',
        'Đào tạo và tương tác để cải thiện mối quan hệ và kỹ năng cho chim cảnh'),
       ('Khác', 'khac', '/static/img/others-category.png',
        'Những dịch vụ khác nhằm đảm bảo pet yêu của bạn khỏe mạnh cũng như tăng thêm mối quan hệ thân thiết');

CREATE TABLE services
(
    id               SERIAL PRIMARY KEY,
    title            VARCHAR(350) NOT NULL,
    description      TEXT         NOT NULL,
    price            INT          NOT NULL,
    category_id      INT          NOT NULL,
    thumbnail_url    TEXT         NOT NULL,
    owned_by_user_id INT          NOT NULL,
    status           TEXT         NOT NULL DEFAULT 'inactive',
    created_at       timestamptz  NOT NULL DEFAULT NOW()
);

ALTER TABLE "services"
    ADD FOREIGN KEY ("owned_by_user_id") REFERENCES "users" ("id");

ALTER TABLE "services"
    ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

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