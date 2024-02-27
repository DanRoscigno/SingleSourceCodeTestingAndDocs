SELECT * FROM FILES
(
    "path" = "abfss://samples@starrocksdocs.dfs.core.windows.net/user_behavior_ten_million_rows.parquet",
    "format" = "parquet",
    -- highlight-start
    "azure.adls2.storage_account" = "starrocksdocs",
    "azure.adls2.shared_key" = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa/bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb+ccccccccc=="
    -- highlight-end
)
LIMIT 3;
