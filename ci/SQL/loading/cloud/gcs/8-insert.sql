INSERT INTO user_behavior_declared
SELECT * FROM FILES
(
    "path" = "gs://starrocks-samples/user_behavior_ten_million_rows.parquet",
    "format" = "parquet",
    -- highlight-start
    "gcp.gcs.service_account_email" = "sampledatareader@xxxxx-xxxxxx-000000.iam.gserviceaccount.com",
    "gcp.gcs.service_account_private_key_id" = "baaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
    "gcp.gcs.service_account_private_key" = "-----BEGIN PRIVATE KEY----- ----END PRIVATE KEY-----"
    -- highlight-end
);
