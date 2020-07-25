FROM golang:alpine AS builder
RUN go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.cn,direct
WORKDIR /app/src
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY ./ ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/bin/app .

FROM alpine
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
    && apk add --no-cache tzdata \
    && ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=builder /app/bin/app /
EXPOSE 80
ENTRYPOINT ["/app"]
