---
sidebar_position: 1
---

# Embedding snippets of code from GitHub

This first block is a section of lines from a configuration file used in an integration
test.  The URL used is to a particular commit for two reasons:
- Line numbers can change, referring to a specific commit guarantees that the line numbers
will stay correct.
- At some point we may have versioned docs; this allows us to link to the docs at the appropriate point in time.

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

## Tutorial Intro

Let's discover **Docusaurus in less than 5 minutes**.

## Getting Started

Get started by **creating a new site**.

Or **try Docusaurus immediately** with **[docusaurus.new](https://docusaurus.new)**.

### What you'll need

- [Node.js](https://nodejs.org/en/download/) version 18.0 or above:
  - When installing Node.js, you are recommended to check all checkboxes related to dependencies.

## Generate a new site

Generate a new Docusaurus site using the **classic template**.

The classic template will automatically be added to your project after you run the command:

```bash
npm init docusaurus@latest my-website classic
```

You can type this command into Command Prompt, Powershell, Terminal, or any other integrated terminal of your code editor.

The command also installs all necessary dependencies you need to run Docusaurus.

## Start your site

Run the development server:

```bash
cd my-website
npm run start
```

The `cd` command changes the directory you're working with. In order to work with your newly created Docusaurus site, you'll need to navigate the terminal there.

The `npm run start` command builds your website locally and serves it through a development server, ready for you to view at http://localhost:3000/.

Open `docs/intro.md` (this page) and edit some lines: the site **reloads automatically** and displays your changes.
