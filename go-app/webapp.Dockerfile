FROM golang:1.12.9-alpine3.10 AS builder
ENV GOPROXY https://goproxy.io
WORKDIR /app/src
# The purpose of the following 3 lines is to 
# use cache to avoid downloading dependencies 
# again if the go.mod and go.sum not be modified.
ADD ./go.mod .
ADD ./go.sum .
RUN go mod download
# In most cases we rebuild the image just because of
# that we modifed the source code but not changed
# the dependencies of them. So, we split the next 2 lines
# from the above 3 lines so if in case of that, the work
# of rebuild just begain in next line.
ADD . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/bin/app main.go

FROM alpine:3.10
# Timezone setting for Alpine: UTC -> CST
RUN apk add --no-cache tzdata \
    && ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=builder /app/bin/app /
EXPOSE 8080
ENTRYPOINT ["/app"]