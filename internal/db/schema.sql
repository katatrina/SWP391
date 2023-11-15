CREATE TABLE "roles"
(
    "id"   SERIAL PRIMARY KEY,
    "name" VARCHAR NOT NULL
);

CREATE TABLE "admin"
(
    id                SERIAL PRIMARY KEY,
    "email"           VARCHAR(150) UNIQUE NOT NULL,
    "hashed_password" CHAR(60)            NOT NULL
);

CREATE TABLE "users"
(
    "id"              SERIAL PRIMARY KEY,
    "full_name"       VARCHAR(50)         NOT NULL,
    "email"           VARCHAR(150) UNIQUE NOT NULL,
    "phone"           CHAR(10) UNIQUE     NOT NULL,
    "address"         VARCHAR(200)        NOT NULL,
    "role_id"         INTEGER             NOT NULL,
    "hashed_password" CHAR(60)            NOT NULL,
    "created_at"      timestamptz         NOT NULL DEFAULT (now())
);

CREATE TABLE "provider_details"
(
    "id"           SERIAL PRIMARY KEY,
    "provider_id"  INTEGER     NOT NULL,
    "company_name" VARCHAR(50) NOT NULL,
    "tax_code"     VARCHAR(50) NOT NULL,
    "created_at"   timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE sessions
(
    token  TEXT PRIMARY KEY,
    data   BYTEA       NOT NULL,
    expiry TIMESTAMPTZ NOT NULL
);

CREATE INDEX sessions_expiry_idx ON sessions (expiry);

CREATE TABLE "categories"
(
    "id"          SERIAL PRIMARY KEY,
    "name"        VARCHAR(50) NOT NULL,
    "slug"        VARCHAR(50) NOT NULL,
    "image_path"  TEXT        NOT NULL,
    "description" TEXT        NOT NULL
);

CREATE TABLE "services"
(
    "id"                   SERIAL PRIMARY KEY,
    "title"                VARCHAR(350) NOT NULL,
    "description"          VARCHAR      NOT NULL,
    "price"                INTEGER      NOT NULL,
    "image_path"           TEXT         NOT NULL,
    "category_id"          INTEGER      NOT NULL,
    "owned_by_provider_id" INTEGER      NOT NULL,
    "status"               varchar(20)  NOT NULL DEFAULT 'inactive',
    "reject_reason"        TEXT         NOT NULL DEFAULT '',
    "created_at"           timestamptz  NOT NULL DEFAULT (now())
);

CREATE TABLE "service_feedbacks"
(
    "id"         SERIAL PRIMARY KEY,
    "service_id" INTEGER     NOT NULL,
    "user_id"    INTEGER     NOT NULL,
    "rating"     INTEGER     NOT NULL,
    "content"    VARCHAR     NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "carts"
(
    "id"          SERIAL PRIMARY KEY,
    "user_id"     INTEGER NOT NULL,
    "grand_total" INTEGER NOT NULL DEFAULT 0
);

CREATE TABLE "cart_items"
(
    "uuid"       TEXT PRIMARY KEY,
    "cart_id"    INTEGER NOT NULL,
    "service_id" INTEGER NOT NULL,
    "quantity"   INTEGER NOT NULL,
    "sub_total"  INTEGER NOT NULL
);

CREATE TABLE "orders"
(
    "uuid"           TEXT PRIMARY KEY,
    "buyer_id"       INTEGER     NOT NULL,
    "seller_id"      INTEGER     NOT NULL,
    "status_id"      INTEGER     NOT NULL,
    "payment_method" VARCHAR     NOT NULL,
    "grand_total"    INTEGER     NOT NULL DEFAULT 0,
    "created_at"     timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "order_status"
(
    "id"     SERIAL PRIMARY KEY,
    "code"   VARCHAR(20) NOT NULL,
    "detail" VARCHAR(70) NOT NULL
);

CREATE TABLE "order_items"
(
    "uuid"       TEXT PRIMARY KEY,
    "order_id"   TEXT        NOT NULL,
    "service_id" INTEGER     NOT NULL,
    "quantity"   INTEGER     NOT NULL,
    "sub_total"  INTEGER     NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "order_item_details"
(
    "id"            SERIAL PRIMARY KEY,
    "order_item_id" TEXT         NOT NULL,
    "title"         VARCHAR(350) NOT NULL,
    "price"         INTEGER      NOT NULL,
    "image_path"    TEXT         NOT NULL,
    "created_at"    timestamptz  NOT NULL DEFAULT (now())
);

ALTER TABLE "users"
    ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("id");

ALTER TABLE "users"
    ADD CONSTRAINT users_uc_email UNIQUE ("email");

ALTER TABLE "users"
    ADD CONSTRAINT users_uc_phone UNIQUE ("phone");

ALTER TABLE "provider_details"
    ADD FOREIGN KEY ("provider_id") REFERENCES "users" ("id");

ALTER TABLE "services"
    ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

ALTER TABLE "services"
    ADD FOREIGN KEY ("owned_by_provider_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "carts"
    ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "cart_items"
    ADD FOREIGN KEY ("cart_id") REFERENCES "carts" ("id") ON DELETE CASCADE;

ALTER TABLE "cart_items"
    ADD FOREIGN KEY ("service_id") REFERENCES "services" ("id") ON DELETE CASCADE;

ALTER TABLE "service_feedbacks"
    ADD FOREIGN KEY ("service_id") REFERENCES "services" ("id");

ALTER TABLE "service_feedbacks"
    ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "orders"
    ADD FOREIGN KEY ("buyer_id") REFERENCES "users" ("id");

ALTER TABLE "orders"
    ADD FOREIGN KEY ("seller_id") REFERENCES "users" ("id");

ALTER TABLE "orders"
    ADD FOREIGN KEY ("status_id") REFERENCES "order_status" ("id");

ALTER TABLE "order_items"
    ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("uuid");

ALTER TABLE "order_item_details"
    ADD FOREIGN KEY ("order_item_id") REFERENCES "order_items" ("uuid");

INSERT INTO roles (name)
VALUES ('customer'),
       ('provider');

INSERT INTO categories (name, slug, image_path, description)
VALUES ('Phụ kiện', 'accessory', '/static/img/accessories-category.jpg',
        'Các dịch vụ cung cấp đầy đủ các sản phẩm cần thiết cho chim cảnh'),
       ('Dinh dưỡng và thức ăn', 'nutrition-and-food', '/static/img/nutrition-category.jpg',
        'Các dịch vụ cung cấp đầy đủ các dinh dưỡng cần thiết cho chim cảnh'),
       ('Y tế và chăm sóc sức khỏe', 'health-care', '/static/img/healthcare-category.jpg',
        'Đảm bảo sức khỏe và phòng ngừa bệnh tật cho chim cảnh'),
       ('Grooming', 'grooming', '/static/img/categories-grooming.jpg',
        'Chăm sóc, tạo phong cách và làm đẹp cho chim cảnh'),
       ('Đào tạo và huấn luyện', 'training', '/static/img/training-category.jpg',
        'Đào tạo và tương tác để cải thiện mối quan hệ và kỹ năng cho chim cảnh'),
       ('Khác', 'other', '/static/img/others-category.png',
        'Những dịch vụ khác nhằm đảm bảo pet yêu của bạn khỏe mạnh cũng như tăng thêm mối quan hệ thân thiết');

INSERT INTO order_status (code, detail)
VALUES ('pending', 'Đang chờ xác nhận'),
       ('confirmed', 'Đã xác nhận'),
       ('completed', 'Hoàn thành'),
       ('cancelled', 'Đã hủy');