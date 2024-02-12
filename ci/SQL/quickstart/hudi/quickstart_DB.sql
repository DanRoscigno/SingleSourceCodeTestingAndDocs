CREATE EXTERNAL CATALOG hudi_catalog_hms
PROPERTIES
(
    "type" = "hudi",
    "hive.metastore.type" = "hive",
    "hive.metastore.uris" = "thrift://hive-metastore:9083",
    "aws.s3.use_instance_profile" = "false",
    "aws.s3.access_key" = "admin",
    "aws.s3.secret_key" = "password",
    "aws.s3.region" = "us-east-1",
    "aws.s3.enable_ssl" = "false",
    "aws.s3.enable_path_style_access" = "true",
    "aws.s3.endpoint" = "http://minio:9000"
);

SET CATALOG hudi_catalog_hms;

SHOW DATABASES;

USE default;

SHOW tables;

SELECT * from hudi_coders_hive\G
