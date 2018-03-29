DROP DATABASE IF EXISTS fitnessdb;

CREATE DATABASE fitnessdb;
USE fitnessdb;
CREATE TABLE users (
    user_id VARCHAR(37) NOT NULL UNIQUE PRIMARY KEY,
    username VARCHAR(35) NOT NULL UNIQUE,
    pass_hash VARCHAR(60) NOT NULL, 
    full_name VARCHAR(70) DEFAULT '',
    email VARCHAR(50) NOT NULL UNIQUE,
    account_status CHAR DEFAULT 'u'
);

CREATE TABLE workouts (
    workout_id CHAR(23),
    workout_name VARCHAR(50),
    begin_date DATE,
    end_date DATE
);

CREATE TABLE exercise (
    ex_id VARCHAR(23) NOT NULL UNIQUE PRIMARY KEY,
    workout_id CHAR(23) NOT NULL FOREIGN KEY REFERENCES workouts(workout_id),
    ex_cat_id VARCHAR(25),
    ex_name VARCHAR(50) DEFAULT ''
);
-- Use Above This Line Only --
CREATE TABLE exercise_catelog (
    ex_cat_id VARCHAR(25),  
    ex_cat_name VARCHAR(50)
);

CREATE TABLE exercise_sets (
    set_id VARCHAR(50),
    reps INT,
    set_weight FLOAT,
    note VARCHAR(255)
);