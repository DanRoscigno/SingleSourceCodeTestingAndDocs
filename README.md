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
