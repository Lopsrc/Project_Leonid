CREATE TABLE IF NOT EXISTS auth
(
    id              SERIAL PRIMARY KEY,
    email           VARCHAR(255) NOT NULL UNIQUE,
    pass_hash       bytea  NOT NULL,
    del             BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS customers
(
    customer_id INTEGER REFERENCES auth(id),
    name        VARCHAR(255) NOT NULL,
    sex         VARCHAR(30) NOT NULL,
    birthdate   DATE NOT NULL,
    age         INTEGER NOT NULL,
    weight      INTEGER NOT NULL
);