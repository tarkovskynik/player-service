CREATE TABLE users
(
    id      int     not null unique,
    balance int     not null,
    token   varchar not null
);