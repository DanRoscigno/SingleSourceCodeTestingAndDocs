SELECT * FROM FILES
(
    "path" = "s3://starrocks-examples/user-behavior-10-million-rows.parquet",
    "format" = "parquet",
    "aws.s3.region" = "us-east-1",
    -- highlight-start
    "aws.s3.access_key" = "AAAAAAAAAAAAAAAAAAAA",
    "aws.s3.secret_key" = "BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB"
    -- highlight-end
)
LIMIT 3;
