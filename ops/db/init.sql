create database if not exists local_dev_database;
use local_dev_database;
create table if not exists user (id varchar(50), email varchar(255), password varchar(255), name varchar(100), address varchar(100), age INT(3), created_at datetime, updated_at datetime, PRIMARY KEY (`id`)) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;