CREATE DATABASE IF NOT EXISTS local;

CREATE TABLE
    urls(
        id serial PRIMARY KEY,
        surl varchar(8) UNIQUE NOT NULL,
        url varchar(255) UNIQUE NOT NULL
    );