---
sidebar_position: 20
---

# Embedding snippets of code from GitHub

This first block is a section of lines from a configuration file used in an integration
test.  The URL used is to a particular commit for two reasons:
- Line numbers can change, referring to a specific commit guarantees that the line numbers
will stay correct.
- As we have versioned docs; this supports linking to the docs at the appropriate point in history.

The plan is to have dedicated tests for code that is included in the docs so that we know before the community when:
- a change happens to a dataset that is used in the docs
- a breaking change breaks a tutorial

```sql reference title="Create table from S3 using FILES() table function"
https://github.com/DanRoscigno/docs/blob/6d6fcf905162adf80bd094cb9dd133a5c557bdd3/SQL/files_table_fxn.sql#L1-L11
```

```sql
-- Create table from S3 using FILES() table function
CREATE TABLE DocsQA.user_behavior_inferred
AS SELECT * FROM FILES (
	"path" = "s3://starrocks-examples/user_behavior_ten_million_rows.parquet",
	"format" = "parquet",
	"aws.s3.region" = "us-east-1",
	-- highlight-start
	"aws.s3.access_key" = "AAAAAAAAAAAAAAAAAAAA",
	"aws.s3.secret_key" = "BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB"
	-- highlight-end
);
```

```js reference useHTTPS
https://github.com/webdriverio/example-recipes/blob/fc362f2f8dd823d294b9bb5f92bd5991339d4591/getting-started/run-in-script.js#L2-L19
```

```js reference title="Example"
https://github.com/saucelabs/docusaurus-theme-github-codeblock/blob/main/src/theme/ReferenceCodeBlock/index.tsx#L105-L108
```

```js reference title="be.conf"
https://github.com/StarRocks/starrocks/blob/bed907747e07cd529a3b11e372997a14fd03436e/conf/be.conf#L21-L26
```

```sql reference title="JOIN from Quick Start"
https://github.com/StarRocks/starrocks/blob/878ed8ae56ded2324c90f52c5ea99e2d2abcb51a/docs/en/assets/quick-start/_SQL.mdx#L136-L147
```
