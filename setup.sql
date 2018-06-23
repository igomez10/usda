CREATE USER 'username'@'%' IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON *.* TO 'username'@'%';
/* UPDATE mysql.user SET Host='%' WHERE Host='localhost' AND User='username' */
ALTER USER 'username'@'%' IDENTIFIED WITH mysql_native_password BY 'password';
CREATE DATABASE dev;
USE dev;
CREATE TABLE `oranges` (
  `comm_name` varchar(255) DEFAULT NULL,
  `city_name` varchar(255) DEFAULT NULL,
  `package_desc` varchar(255) DEFAULT NULL,
  `variety_name` varchar(255) DEFAULT NULL,
  `subvar_name` varchar(255) DEFAULT NULL,
  `grade_desc` varchar(255) DEFAULT NULL,
  `report_date` varchar(255) DEFAULT NULL,
  `low_price_min` double DEFAULT NULL,
  `high_price_max` double DEFAULT NULL,
  `mostly_low_min` double DEFAULT NULL,
  `mostly_high_max` double DEFAULT NULL,
  `origin_name` varchar(255) DEFAULT NULL,
  `district_name` varchar(255) DEFAULT NULL,
  `item_size_desc` varchar(255) DEFAULT NULL,
  `color` varchar(255) DEFAULT NULL,
  `env_desc` varchar(255) DEFAULT NULL,
  `unit_sale` varchar(255) DEFAULT NULL,
  `quality` varchar(255) DEFAULT NULL,
  `condition` varchar(255) DEFAULT NULL,
  `appearance` varchar(255) DEFAULT NULL,
  `storage` varchar(255) DEFAULT NULL,
  `crop` varchar(255) DEFAULT NULL,
  `re_pack` varchar(255) DEFAULT NULL,
  `transmode` varchar(255) DEFAULT NULL,
  `offerings` varchar(255) DEFAULT NULL,
  `market_tone` varchar(255) DEFAULT NULL,
  `price_comment` varchar(255) DEFAULT NULL,
  `comment` varchar(255) DEFAULT NULL,
  `id` varchar(255) PRIMARY KEY
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
