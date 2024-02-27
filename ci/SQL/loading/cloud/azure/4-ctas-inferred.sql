CREATE TABLE user_behavior_inferred AS
SELECT * FROM FILES
(
    "path" = "abfss://samples@starrocksdocs.dfs.core.windows.net/user_behavior_ten_million_rows.parquet",
    "format" = "parquet",
    "azure.adls2.storage_account" = "starrocksdocs",
    "azure.adls2.shared_key" = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa/bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb+ccccccccc=="
);
