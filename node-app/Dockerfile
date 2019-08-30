FROM node:12.9.1
WORKDIR /app
ADD ./package.json .
RUN npm install --registry=https://registry.npm.taobao.org
ADD . .
EXPOSE 3000
ENTRYPOINT ["node", "/app/index.js"]