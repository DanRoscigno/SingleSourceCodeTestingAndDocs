## To Do

1. Figure out why the curl command (see fail down at end of this file) fails
1. Need to write the workflow file

### Remember to export

```bash
export AWS_S3_SECRET_KEY=redacted
export AWS_S3_ACCESS_KEY=redacted
```

### Works

```bash
docker compose -f allin1-docker-compose.yml build
```

### Works

This starts the allin1 and the ginkgo container. The Ginkgo container just runs ash (shell) in the background and then I can use `docker compose exec` to run ginkgo commands.

```bash
docker compose -f allin1-docker-compose.yml up --detach --wait --wait-timeout 60
```

### This one works partially

The connection to the database succeeds, the DDL commands work, but the curl to populate the crash data fails.

```bash
docker compose -f allin1-docker-compose.yml exec test-harness ginkgo -v --focus-file=./quickstart_basic_test.go
```

### This command works

```bash
docker compose -f allin1-docker-compose.yml exec test-harness ginkgo -v --focus-file=./docs_test.go
```
