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


CREATE TABLE `coordinates` (
    `name` varchar(50),
    `latitude` float,
    `longitude` float
);

INSERT INTO coordinates VALUES ('SAN FRANCISCO', 37.774929, -122.419416 );
INSERT INTO coordinates VALUES ('NEW YORK', 40.712775, -74.005973 );
INSERT INTO coordinates VALUES ('CHICAGO', 41.878114, -87.629798 );
INSERT INTO coordinates VALUES ('BOSTON', 42.360082, -71.058880 );
INSERT INTO coordinates VALUES ('ROTTERDAM', 51.924420, 4.477733 );
INSERT INTO coordinates VALUES ('ATLANTA', 33.748995,-84.387982 );
INSERT INTO coordinates VALUES ('LOS ANGELES', 34.052234,-118.243685);
INSERT INTO coordinates VALUES ('BALTIMORE', 39.290385,-76.612189 );
INSERT INTO coordinates VALUES ('DALLAS', 32.776664,-96.796988);
INSERT INTO coordinates VALUES ('ST. LOUIS', 38.627003,-90.199404);
INSERT INTO coordinates VALUES ('DETROIT', 42.331427,-83.045754  );
INSERT INTO coordinates VALUES ('MONTREAL, QUE', 45.501689,-73.567256);
INSERT INTO coordinates VALUES ('PHILADELPHIA',39.952584,-75.165222 );
INSERT INTO coordinates VALUES ('COLUMBIA', );
INSERT INTO coordinates VALUES ('PHILADELPHIA',39.952584,-75.165222 );
INSERT INTO coordinates VALUES ('PHILADELPHIA',39.952584,-75.165222 );
INSERT INTO coordinates VALUES ('PHILADELPHIA',39.952584,-75.165222 );
