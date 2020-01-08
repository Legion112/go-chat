create database gochat;
use gochat;
drop table message;
create table message (
    id int not null primary key AUTO_INCREMENT,
    username varchar(255) not null,
    text text,
    date int UNSIGNED not null
)