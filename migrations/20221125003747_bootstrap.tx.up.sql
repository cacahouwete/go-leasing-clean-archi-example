SET statement_timeout = 0;

--bun:split

CREATE TABLE customer
(
    id   VARCHAR(255) PRIMARY KEY NOT NULL,
    name VARCHAR(255)             NOT NULL
);

CREATE TABLE car
(
    id   VARCHAR(255) PRIMARY KEY NOT NULL,
    name VARCHAR(255)             NOT NULL
);

CREATE TABLE schedule
(
    id          VARCHAR(255) PRIMARY KEY NOT NULL,
    begin_at    timestamptz              NOT NULL,
    end_at      timestamptz              NOT NULL,
    customer_id VARCHAR(255)             NOT NULL,
    CONSTRAINT fk_schedule_customer
        FOREIGN KEY (customer_id)
            REFERENCES customer (id),
    car_id      VARCHAR(255)             NOT NULL,
    CONSTRAINT fk_schedule_car
        FOREIGN KEY (car_id)
            REFERENCES car (id)
);
