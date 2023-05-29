create database if not exists Payroll;
use Payroll;

CREATE TABLE `record` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `date` date NOT NULL,
  `hours` float NOT NULL,
  `employeeId` int(11) NOT NULL,
  `jobGroup` varchar(1) NOT NULL,
  `sourceId` varchar(12) NOT NULL,
  `startDate` date DEFAULT NULL,
  `endDate` date DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=latin1;

CREATE TABLE `ratio` ( 
  `id` int(11) NOT NULL AUTO_INCREMENT, 
  `category` VARCHAR(10) DEFAULT '',  
  `price` INT,
  PRIMARY KEY (`id`) ) ENGINE=InnoDB;
insert into ratio (category,price) values ('A', 20);
insert into ratio (category,price) values ('B', 30);

CREATE TABLE `file` ( 
  `id` int(11) NOT NULL AUTO_INCREMENT, 
  `filename` VARCHAR(1000),
  PRIMARY KEY (`id`)) ENGINE=InnoDB;

create database  PayrollTest;
use PayrollTest;

CREATE TABLE `record` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `date` date NOT NULL,
  `hours` float NOT NULL,
  `employeeId` int(11) NOT NULL,
  `jobGroup` varchar(1) NOT NULL,
  `sourceId` varchar(12) NOT NULL,
  `startDate` date DEFAULT NULL,
  `endDate` date DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=latin1;

CREATE TABLE `ratio` ( 
  `id` int(11) NOT NULL AUTO_INCREMENT, 
  `category` VARCHAR(10) DEFAULT '',  
  `price` INT,
  PRIMARY KEY (`id`) ) ENGINE=InnoDB;
insert into ratio (category,price) values ('A', 20);
insert into ratio (category,price) values ('B', 30);

CREATE TABLE `file` ( 
  `id` int(11) NOT NULL AUTO_INCREMENT, 
  `filename` VARCHAR(1000),
  PRIMARY KEY (`id`)) ENGINE=InnoDB;
