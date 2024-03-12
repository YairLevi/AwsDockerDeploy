FROM node:slim

ENV NODE_ENV development

WORKDIR /app

COPY package*.json ./
COPY . .

RUN npm install

CMD ["npm", "start"]