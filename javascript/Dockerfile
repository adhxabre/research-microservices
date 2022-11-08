# syntax=docker/dockerfile:1

FROM node:16-alpine
WORKDIR /javascript
COPY . .
RUN npm i
RUN npm start
CMD ["npm", "start", "serve"]