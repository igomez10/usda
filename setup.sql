CREATE USER 'username'@'%' IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON *.* TO 'username'@'%';
/* UPDATE mysql.user SET Host='%' WHERE Host='localhost' AND User='username' */
ALTER USER 'username'@'%' IDENTIFIED WITH mysql_native_password BY 'password';
CREATE DATABASE dev;
USE dev;
CREATE TABLE `Reports` (
	`CommName` VARCHAR,
	`CityName` VARCHAR,
	`PackageDesc` VARCHAR,
	`VarietyName` VARCHAR,
	`SubvarName` VARCHAR,
	`GradeDesc` VARCHAR,
	`ReportDate` DATE,
	`LowPriceMin` FLOAT,
	`HighPriceMax` FLOAT,
	`MostlyLowMin` FLOAT,
	`MostlyHighMax` FLOAT,
	`OriginName` VARCHAR,
	`DistrictName` VARCHAR,
	`ItemSizeDesc` VARCHAR,
	`Color` VARCHAR,
	`EnvDesc` VARCHAR,
	`UnitSale` VARCHAR,
	`Quality` VARCHAR,
	`Condition` VARCHAR,
	`Appearance` VARCHAR,
	`Storage` VARCHAR,
	`Crop` VARCHAR,
	`RePack` VARCHAR,
	`Transmode` VARCHAR,
	`Offerings` VARCHAR,
	`MarketTone` VARCHAR,
	`PriceComment` VARCHAR,
	`Comment` VARCHAR,
);

