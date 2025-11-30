
ARG NODE_VERSION=22.14.0

FROM node:${NODE_VERSION}-alpine 

WORKDIR /nyg-app

COPY package*.json ./

RUN npm ci

COPY . .

RUN npm run build --application

EXPOSE 4200
CMD ["ng", "serve", "--host", "0.0.0.0"]