FROM node:alpine
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
    && apk add --no-cache tzdata \
    && ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
WORKDIR /app
COPY ./package.json ./
RUN npm install --registry=https://registry.npm.taobao.org
COPY ./ ./
EXPOSE 3000
ENTRYPOINT ["node", "/app/index.js"]