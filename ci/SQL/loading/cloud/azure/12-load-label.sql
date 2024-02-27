LOAD LABEL user_behavior
(
    DATA INFILE("abfss://samples@starrocksdocs.dfs.core.windows.net/user_behavior_ten_million_rows.parquet")
    INTO TABLE user_behavior
    FORMAT AS "parquet"
 )
 WITH BROKER
 (
    "azure.adls2.storage_account" = "starrocksdocs",
    "azure.adls2.shared_key" = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa/bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb+ccccccccc=="
 )
PROPERTIES
(
    "timeout" = "3600"
);
