FROM golang:1.12.9-alpine3.10 AS builder
ENV GOPROXY https://goproxy.io
WORKDIR /app/src
ADD ./go.mod .
ADD ./go.sum .
RUN go mod download
ADD . .
RUN go build -o /app/bin/app main.go

FROM alpine:3.10
COPY --from=builder /app/bin/app /
EXPOSE 8080
ENTRYPOINT ["/app"]