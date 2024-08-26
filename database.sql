-- Active: 1723868453146@@127.0.0.1@5432@godof
CREATE DATABASE godof;

CREATE Table users (
    id SERIAL NOT NULL PRIMARY KEY,
    full_name VARCHAR(255) NOT NULL,
    phone VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    -- email_verified_at timestamp(6) without time zone
);

ALTER Table users
ADD COLUMN email_verified_at TIMESTAMP(0) without time zone;

ALTER Table users ADD COLUMN email VARCHAR(255) UNIQUE NOT NULL;

CREATE TABLE public.accounts (
    id SERIAL NOT NULL PRIMARY KEY,
    user_id integer NOT NULL,
    account_number character varying(100),
    balance numeric(19, 2)
);

CREATE TABLE public.transactions (
    id SERIAL NOT NULL PRIMARY KEY,
    sof_number character varying(100) NOT NULL,
    dof_number character varying(100) NOT NULL,
    amount numeric(19, 2),
    transaction_type character varying(1),
    account_id integer NOT NULL,
    transaction_datetime timestamp(0) without time zone
);

CREATE Table public.notifications (
    id SERIAL NOT NULL PRIMARY KEY,
    user_id integer NOT NULL,
    status INTEGER NOT NULL,
    title character varying(255) NOT NULL,
    body text NOT NULL,
    is_read INTEGER NOT NULL,
    created_at timestamp(0) without time zone
);

ALTER Table accounts ADD PRIMARY KEY (id);

ALTER Table transactions ADD PRIMARY KEY (id);

CREATE TABLE public.templates (
    code VARCHAR(100) NOT NULL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    body text NOT NULL
)

INSERT INTO
    public.templates (code, title, body)
VALUES (
        'TRANSFER',
        'Transfer berhasil',
        'Transfer berhasil senilai {{ .amount }} berhasil.'
    )

CREATE Table public.topup (
    id VARCHAR(100) NOT NULL PRIMARY KEY,
    user_id integer NOT NULL,
    amount BIGINT NOT NULL,
    status INTEGER NOT NULL DEFAULT (0),
    created_at timestamp(0) without time zone
) _