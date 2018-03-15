DROP DATABASE IF EXISTS fitnessdb;

CREATE DATABASE fitnessdb;
USE fitnessdb;
CREATE TABLE users (
    user_id VARCHAR(37) NOT NULL UNIQUE PRIMARY KEY,
    username VARCHAR(35) NOT NULL UNIQUE,
    pass_hash VARCHAR(60) NOT NULL, 
    first_name VARCHAR(35) DEFAULT '',
    last_name VARCHAR(35) DEFAULT '',
    email VARCHAR(50) NOT NULL UNIQUE,
    account_status CHAR DEFAULT 'u'
);
-- Use Above This Line Only --
CREATE TABLE workouts (
    workout_id CHAR(36),
    workout_name VARCHAR(50),
    begin_date DATE,
    end_date DATE
);

CREATE TABLE exercise_catelog (
    ex_cat_id VARCHAR(25),  
    ex_cat_name VARCHAR(50)
);

CREATE TABLE exercise (
    ex_id VARCHAR(25),
    ex_cat_id VARCHAR(25),
    ex_name VARCHAR(50)
);

CREATE TABLE exercise_sets (
    set_id VARCHAR(50),
    reps INT,
    set_weight FLOAT,
    note VARCHAR(255)
);