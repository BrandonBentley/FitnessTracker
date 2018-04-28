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
    workout_id CHAR(37) NOT NULL UNIQUE PRIMARY KEY,
    user_id VARCHAR(37) NOT NULL,
    num_ex INT DEFAULT 0,
    workout_name VARCHAR(50),
    date_complete DATE NOT NULL DEFAULT '1995-01-01'
);

CREATE TABLE exercises (
    ex_id VARCHAR(37) NOT NULL UNIQUE PRIMARY KEY,
    workout_id CHAR(37) NOT NULL,
    ex_order INT NOT NULL,
    num_sets INT DEFAULT 0,
    ex_name VARCHAR(50) DEFAULT ''
);

CREATE TABLE exercise_sets (
    set_id VARCHAR(37) NOT NULL UNIQUE PRIMARY KEY,
    ex_id VARCHAR(37) NOT NULL,
    set_order INT NOT NULL,
    reps INT DEFAULT 0,
    done BOOLEAN NOT NULL DEFAULT false,
    set_weight FLOAT DEFAULT 0,
    note VARCHAR(255) DEFAULT ''
);
-- Use Above This Line Only --
-- Nothing Below This Line Is Functional --
CREATE TABLE workout_templates (
    workout_id CHAR(37) NOT NULL UNIQUE PRIMARY KEY,
    user_id VARCHAR(37) NOT NULL,
    workout_name VARCHAR(50),
    begin_date DATE,
    end_date DATE
);

CREATE TABLE exercise_templates (
    ex_id VARCHAR(37) NOT NULL UNIQUE PRIMARY KEY,
    workout_id CHAR(37) NOT NULL,
    ex_cat_id VARCHAR(25),
    ex_name VARCHAR(50) DEFAULT ''
);

CREATE TABLE exercise_catelog (
    ex_cat_id VARCHAR(37) NOT NULL UNIQUE PRIMARY KEY, 
    ex_cat_name VARCHAR(50)
);

