FROM node:18.14.1
WORKDIR /app
COPY . .
RUN npm install
COPY . .
EXPOSE 3000
CMD ["npm", "start"]