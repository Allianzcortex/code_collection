DROP TABLE if exists ARTICLE;
 CREATE TABLE  if not exists `ARTICLE` (
  `_id` int(11) NOT NULL,
  `title` text,
  `magazine` text,
 `volume_number` varchar(50) DEFAULT NULL,
  `year` int(11) DEFAULT NULL,
  `pages` varchar(50) DEFAULT NULL,
  `authors` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`_id`)
) ENGINE=InnoDB DEFAULT CHARSET='utf8mb4';

-- This is for web article
DROP TABLE if exists WEB_ARTICLE;
 CREATE TABLE  if not exists `WEB_ARTICLE` (
  `Article_ID` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(50),
  `volume` varchar(50),
  `volume_number` int(11) DEFAULT NULL,
  `page_number` int(11) DEFAULT NULL,
  `publication_year` varchar(20) DEFAULT NULL,
  `magazine_ID` varchar(50) DEFAULT NULL,
  `authors` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`Article_ID`)
) ENGINE=InnoDB DEFAULT CHARSET='utf8mb4';



DROP TABLE if exists WEB_AUTHOR;
CREATE TABLE if not exists `WEB_AUTHOR` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `lname` varchar(30) NOT NULL,
  `fname` varchar(30) DEFAULT NULL,
  `email` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET='utf8mb4';

DROP TABLE if exists ARTICLE_AUTHORS;
CREATE TABLE if not exists  `ARTICLE_AUTHORS` (
   `id` int(11) NOT NULL AUTO_INCREMENT,
  `article_id` int(11) DEFAULT NULL,
  `author_id` int(11) DEFAULT NULL,
  -- TODO what is the function of constraints
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

DROP TABLE if exists CUSTOMER;
 CREATE TABLE `CUSTOMER` (
  `CID` int(11) NOT NULL AUTO_INCREMENT,
  `first_name` varchar(250) DEFAULT NULL,
  `last_name` varchar(250) DEFAULT NULL,
  `telephone_number` varchar(250) DEFAULT NULL,
  `mailing_address` varchar(250) DEFAULT NULL,
  `discount_code` int(11) DEFAULT NULL,
  PRIMARY KEY (`CID`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

DROP TABLE if exists TRANSACTIONS;
CREATE TABLE `TRANSACTIONS` (
  `transaction_number` int(11) NOT NULL AUTO_INCREMENT,
  `transaction_date` date DEFAULT NULL,
  `total_purchase_price` float DEFAULT NULL,
  `customer_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`transaction_number`)
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=utf8mb4;

DROP TABLE if exists TRANSACTION_ITEMS;
CREATE TABLE `TRANSACTION_ITEMS` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `transaction_number` int(11) DEFAULT NULL,
  `item_id` int(11) DEFAULT NULL,
  `item_price` float DEFAULT NULL,                                       
  PRIMARY KEY (`id`)                                       
) ENGINE=InnoDB AUTO_INCREMENT=33 DEFAULT CHARSET=utf8mb4; 

ALTER TABLE WEB_ARTICLE AUTO_INCREMENT = 1;
ALTER TABLE WEB_AUTHOR AUTO_INCREMENT = 1;
ALTER TABLE article_authors AUTO_INCREMENT = 1;

insert into WEB_ARTICLE(magazine_ID,volume,volume_number,title,page_number,authors,publication_year) values('Magic World','earch',12,"Harry Potter",67,"JK Rolin",1998);
insert into WEB_AUTHOR(lname,fname,email) values("JK","Rolin","JKRolin@gmail.com");
insert into article_authors(article_id,author_id) values(1,1);
insert into WEB_ARTICLE(magazine_ID,volume,volume_number,title,page_number,authors,publication_year) values('Magic World','universe',14,"Three Body",109,"Cixin Liu,Shu Bao",2008);
insert into WEB_AUTHOR(lname,fname,email) values("Cixin","Liu","Cixin.Liu@gmail.com");
insert into WEB_AUTHOR(lname,fname,email) values("Shu","Bao","Shu.Bao@gmail.com");
insert into article_authors(article_id,author_id) values(2,2);
insert into article_authors(article_id,author_id) values(2,3);

-- insert into customer

ALTER TABLE customer AUTO_INCREMENT = 1;

insert into customer(first_name,last_name,telephone_number,mailing_address,discount_code) values("Tom","Cruise",13098456781,"Hollywood 1559 Vine Street.",0);
insert into customer(first_name,last_name,telephone_number,mailing_address,discount_code) values("Marco","Van basten",18908457890,"LA Beverly hills",1);

-- insert into transactions

ALTER TABLE transactions AUTO_INCREMENT = 1;
ALTER TABLE transaction_items AUTO_INCREMENT = 1;

insert into transactions(transaction_date,total_purchase_price,customer_id) values('2014-01-01',20,1); -- default
insert into transaction_items(transaction_number,item_id,item_price) values(1,1,20);
insert into transactions(transaction_date,total_purchase_price,customer_id) values('2019-04-10',40,1); -- default
insert into transaction_items(transaction_number,item_id,item_price) values(2,2,10);
insert into transaction_items(transaction_number,item_id,item_price) values(2,3,30);
insert into transactions(transaction_date,total_purchase_price,customer_id) values('2019-04-10',120,2); -- default
insert into transaction_items(transaction_number,item_id,item_price) values(3,4,40);
insert into transaction_items(transaction_number,item_id,item_price) values(3,5,80);

-- something needed to be noticed:

-- 1. we don't create volume table. Because volume attribute is only in article table. We will not do CRUD operations to volume
