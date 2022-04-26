CREATE DATABASE IF NOT EXISTS maxxer;
USE maxxer;

DROP TABLE IF EXISTS funds;

CREATE TABLE funds(
    nickname varchar(50) primary key not null,
    currency varchar(50) not null,
    amount DOUBLE UNSIGNED not null 
) ENGINE=INNODB;

CREATE TABLE funds_history(
    id int primary key auto_increment,
    nickname varchar(50) not null,
    currency varchar(50) not null,
    amount DOUBLE UNSIGNED not null, 
    date TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=INNODB;


ALTER TABLE `funds_history`
ADD COLUMN `transaction_type` VARCHAR(15) AFTER currency;

ALTER TABLE funds
DROP PRIMARY KEY,
ADD PRIMARY KEY (nickname, currency)

