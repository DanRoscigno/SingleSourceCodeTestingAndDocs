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
