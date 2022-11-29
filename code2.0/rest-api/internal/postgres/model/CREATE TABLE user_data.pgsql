-- CREATE TABLE UserAuth(
--     id   SERIAL PRIMARY KEY,
--     login VARCHAR(30) UNIQUE,   --сделать уникальным, нельзя чтобы они повторялись (not complete) (-)
--     state BOOLEAN,
--     access_token VARCHAR UNIQUE,
--     refresh_token VARCHAR UNIQUE
-- );
-- CREATE TYPE sex AS ENUM('man', 'woman');
-- CREATE TABLE UserData(
--     id   SERIAL PRIMARY KEY,
--     user_name VARCHAR(30) NOT NULL,
--     sex sex ,           --НЕ ИСПОЛЬЗОВАТЬ NULL, ТК ПОЛЕ СТАНОВИТСЯ ОБЯЗАТЕЛЬНЫМ ДЛЯ ЗАПОЛНЕНИЯ
--     birthdate DATE ,
--     weight SMALLINT 
-- );
--DROP TABLE UserAuth;
--SELECT * FROM UserAuth;
-- DROP TABLE UserData;

-- INSERT INTO UserData (user_name, sex, birthdate, weight) VALUES ('Sergey', 'man', '2002-04-04', 70);

-- INSERT INTO UserData (user_name) VALUES ('Ivan');


-- SELECT * FROM UserData;
--INSERT INTO UserAUTH (login, state, access_token, refresh_token) VALUES ('ivanIO', FALSE, 'access_token_02', 'refresh_token_02');