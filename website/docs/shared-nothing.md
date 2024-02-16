---
sidebar_position: 30
---

import DDL from './_assets/quick-start/_DDL.mdx'
import Clients from './_assets/quick-start/_clientsAllin1.mdx'
import SQL from './_assets/quick-start/_SQL.mdx'
import Curl from './_assets/quick-start/_curl.mdx'

# StarRocks basics

This tutorial covers:

- Running StarRocks in a single Docker container
- Loading two public datasets including basic transformation of the data
- Analyzing the data with SELECT and JOIN
- Basic data transformation (the **T** in ETL)

The data used is provided by NYC OpenData and the National Centers for Environmental Information.

Both of these datasets are very large, and because this tutorial is intended to help you get exposed to working with StarRocks we are not going to load data for the past 120 years. You can run the Docker image and load this data on a machine with 4 GB RAM assigned to Docker. For larger fault-tolerant and scalable deployments we have other documentation and will provide that later.

There is a lot of information in this document, and it is presented with the step by step content at the beginning, and the technical details at the end. This is done to serve these purposes in this order:

1. Allow the reader to load data in StarRocks and analyze that data.
2. Explain the basics of data transformation during loading.

---

## Prerequisites

### Docker

- [Docker](https://docs.docker.com/engine/install/)
- 4 GB RAM assigned to Docker
- 10 GB free disk space assigned to Docker

### SQL client

You can use the SQL client provided in the Docker environment, or use one on your system. Many MySQL compatible clients will work, and this guide covers the configuration of DBeaver and MySQL WorkBench.

### curl

`curl` is used to issue the data load job to StarRocks, and to download the datasets. Check to see if you have it installed by running `curl` or `curl.exe` at your OS prompt. If curl is not installed, [get curl here](https://curl.se/dlwiz/?type=bin).

---
## Terminology

### FE
Frontend nodes are responsible for metadata management, client connection management, query planning, and query scheduling. Each FE stores and maintains a complete copy of metadata in its memory, which guarantees indiscriminate services among the FEs.

### BE
Backend nodes are responsible for both data storage and executing query plans.

---

## Launch StarRocks

```bash
docker run -p 9030:9030 -p 8030:8030 -p 8040:8040 -itd \
--name quickstart starrocks/allin1-ubuntu
```
---
## SQL clients

<Clients />

---

## Download the data

Download these two datasets to your machine. You can download them to the host machine where you are running Docker, they do not need to be downloaded inside the container.

### New York City crash data

```bash reference title="Download crash data"
https://github.com/DanRoscigno/SingleSourceCodeTestingAndDocs/blob/4482ab7e53871fb531cc311a76f26edf60658c07/ci/SHELL/quickstart/basic/NYPD_download#L3-L5
```

### Weather data

```bash reference title="Download weather data"
https://github.com/DanRoscigno/SingleSourceCodeTestingAndDocs/blob/f11fcd0f1e80526c09fa838c0fa31c315bc0598c/ci/SHELL/quickstart/basic/Weather_download#L3-L5
```

---

### Connect to StarRocks with a SQL client

:::tip

If you are using a client other than the mysql CLI, open that now.
:::

This command will run the `mysql` command in the Docker container:

```sql
docker exec -it quickstart \
mysql -P 9030 -h 127.0.0.1 -u root --prompt="StarRocks > "
```

---

## Create some tables

<DDL />

---

## Load two datasets
There are many ways to load data into StarRocks. For this tutorial the simplest way is to use curl and StarRocks Stream Load.

:::tip
Open a new shell as these curl commands are run at the operating system prompt, not in the `mysql` client. The commands refer to the datasets that you downloaded, so run them from the directory where you downloaded the files.

You will be prompted for a password. You probably have not assigned a password to the MySQL `root` user, so just hit enter.
:::

The `curl` commands look complex, but they are explained in detail at the end of the tutorial. For now, we recommend running the commands and running some SQL to analyze the data, and then reading about the data loading details at the end.

### New York City collision data - Crashes

```bash reference title="Stream crash data into StarRocks"
https://github.com/DanRoscigno/SingleSourceCodeTestingAndDocs/blob/0f0e41e5f46af64f3b566d465aa2797aa46bf142/ci/SHELL/quickstart/basic/NYPD_stream_load#L3-L12
```

Here is the output of the preceding command. The first highlighted section shows what you should expect to see (OK and all but one row inserted). One row was filtered out because it does not contain the correct number of columns.

```bash
Enter host password for user 'root':
{
    "TxnId": 2,
    "Label": "crashdata-0",
    "Status": "Success",
    # highlight-start
    "Message": "OK",
    "NumberTotalRows": 423726,
    "NumberLoadedRows": 423725,
    # highlight-end
    "NumberFilteredRows": 1,
    "NumberUnselectedRows": 0,
    "LoadBytes": 96227746,
    "LoadTimeMs": 1013,
    "BeginTxnTimeMs": 21,
    "StreamLoadPlanTimeMs": 63,
    "ReadDataTimeMs": 563,
    "WriteDataTimeMs": 870,
    "CommitAndPublishTimeMs": 57,
    # highlight-start
    "ErrorURL": "http://127.0.0.1:8040/api/_load_error_log?file=error_log_da41dd88276a7bfc_739087c94262ae9f"
    # highlight-end
}%
```

If there was an error the output provides a URL to see the error messages. Open this in a browser to find out what happened. Expand the detail to see the error message:

<details>

<summary>Reading error messages in the browser</summary>

```bash
Error: Value count does not match column count. Expect 29, but got 32.

Column delimiter: 44,Row delimiter: 10.. Row: 09/06/2015,14:15,,,40.6722269,-74.0110059,"(40.6722269, -74.0110059)",,,"R/O 1 BEARD ST. ( IKEA'S 
09/14/2015,5:30,BRONX,10473,40.814551,-73.8490955,"(40.814551, -73.8490955)",TORRY AVENUE                    ,NORTON AVENUE                   ,,0,0,0,0,0,0,0,0,Driver Inattention/Distraction,Unspecified,,,,3297457,PASSENGER VEHICLE,PASSENGER VEHICLE,,,
```

</details>

### Weather data

Load the weather dataset in the same manner as you loaded the crash data.

```bash reference title="Stream weather data into StarRocks"
https://github.com/DanRoscigno/SingleSourceCodeTestingAndDocs/blob/9f5fb21071124d9814fbf8cbb828b4c93a3d12da/ci/SHELL/quickstart/basic/Weather_stream_load#L3-L12
```

---

## Answer some questions

<SQL />

---

## Summary

In this tutorial you:

- Deployed StarRocks in Docker
- Loaded crash data provided by New York City and weather data provided by NOAA
- Analyzed the data using SQL JOINs to find out that driving in low visibility or icy streets is a bad idea

There is more to learn; we intentionally glossed over the data transformation done during the Stream Load. The details on that are in the notes on the curl commands below.

---

## Notes on the curl commands

<Curl />

---

## More information

[StarRocks table design](../table_design/StarRocks_table_design.md)

[Materialized views](../cover_pages/mv_use_cases.mdx)

[Stream Load](../sql-reference/sql-statements/data-manipulation/STREAM_LOAD.md)

The [Motor Vehicle Collisions - Crashes](https://data.cityofnewyork.us/Public-Safety/Motor-Vehicle-Collisions-Crashes/h9gi-nx95) dataset is provided by New York City subject to these [terms of use](https://www.nyc.gov/home/terms-of-use.page) and [privacy policy](https://www.nyc.gov/home/privacy-policy.page).

The [Local Climatological Data](https://www.ncdc.noaa.gov/cdo-web/datatools/lcd)(LCD) is provided by NOAA with this [disclaimer](https://www.noaa.gov/disclaimer) and this [privacy policy](https://www.noaa.gov/protecting-your-privacy).
