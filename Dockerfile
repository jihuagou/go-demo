FROM golang:latest

#ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/tmaio/go-gin-example
COPY . $GOPATH/src/github.com/tmaio/go-gin-example
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./go-gin-example"]