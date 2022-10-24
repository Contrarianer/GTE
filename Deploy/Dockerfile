FROM golang:1.18 as builder
ENV GOPROXY="https://goproxy.cn,direct" \
    GO111MODULE="on" \
    CGO_ENABLED="0" \
    GOOS="linux"
WORKDIR /go/src/github.com/gowebspider/gte/
COPY . .
RUN go mod tidy && \
    go build -o /tmp/example .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /tmp/example .
EXPOSE 9090
CMD ["./example"]
