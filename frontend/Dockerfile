# STEP 1 build static website
FROM node:16-alpine as build

WORKDIR /app

COPY package*.json .

RUN npm install

COPY .  .

RUN npm run build

FROM nginx:1.23.0-alpine

EXPOSE 4200

RUN rm -rf /usr/share/nginx/html/*

COPY ./nginx/default.conf /etc/nginx/conf.d/default.conf

COPY --from=build /app/dist/code /usr/share/nginx/html
