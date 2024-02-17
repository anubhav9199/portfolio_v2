FROM node:18-alpine

WORKDIR /nodejs_server

COPY package*.json ./
RUN npm install

COPY . .

EXPOSE 8003

# CMD ["npm", "start"]