-- Create table from S3 using FILES() table function
CREATE TABLE DocsQA.user_behavior_inferred
AS SELECT * FROM FILES (
	"path" = "s3://starrocks-examples/user_behavior_ten_million_rows.parquet",
	"format" = "parquet",
	"aws.s3.region" = "us-east-1",
	"aws.s3.access_key" = "AAAAAAAAAAAAAAAAAAAA",
	"aws.s3.secret_key" = "BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB"
);
