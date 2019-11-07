# 简单的 Go web app 的 Docker 容器化示例

## 组成

- Nginx 反代
- Go web app

## 启动

使用 docker-compose 启动

```
$ docker-compose up -d --build
```

## 文件

- `docker-compose.yml` docker-compose 模版
- `webapp.Dockerfile` Go web app 的镜像构建
- `nginx.Dockerfile` Nginx 镜像构建
- `nginx.conf` Nginx 容器/镜像中的反向代理配置
- `main.go`、`go.mod`、`go.sum` Go web app 源码