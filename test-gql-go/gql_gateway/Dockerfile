FROM node:16

WORKDIR /app
COPY package.json ./
RUN npm i

COPY server.js ./

CMD [ "node", "server.js" ]