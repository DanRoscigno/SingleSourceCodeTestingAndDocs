# docs

## Init golang

> note:
>
> ginkgo and gomega are installed, and my PATH includes `~/go/bin`

```
go mod init github.com/danroscigno/docs
ginkgo bootstrap
ginkgo generate
go get github.com/onsi/ginkgo/v2
go get -t github.com/danroscigno/docs
```

I needed a non-test go file, so I added a dummy `docs.go` that justs returns the text Hello World.

At this point I have not used the `gomega` import, so it is commented out.

Running `go test` returns the following, and it is correct. I have no specs defined, so none failed and we passed:

```go
Running Suite: Docs Suite - /Users/droscign/GitHub/docs
=======================================================
Random Seed: 1706751017

Will run 0 of 0 specs

Ran 0 of 0 Specs in 0.000 seconds
SUCCESS! -- 0 Passed | 0 Failed | 0 Pending | 0 Skipped
PASS
ok  	github.com/danroscigno/docs	0.193s
```

## StarRocks

To test against the database it needs to be running. For now I will run the `allin1` Docker container and expose port 9030:

```bash
docker run -p 9030:9030 -p 8030:8030 -p 8040:8040 -itd \
--name quickstart starrocks/allin1-ubuntu
```

## MySQL module

We will use the MySQL golang module(s) to connect and run queries

```bash
go get github.com/go-sql-driver/mysql
```

