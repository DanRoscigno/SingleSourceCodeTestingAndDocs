SHOW DATABASES;

USE nyc;

SHOW TABLES;

DESCRIBE greentaxis;
-- SQL = `select COLUMN_NAME, DATA_TYPE from INFORMATION_SCHEMA.COLUMNS where TABLE_NAME = 'greentaxis';`

SELECT lpep_pickup_datetime FROM greentaxis LIMIT 10;

