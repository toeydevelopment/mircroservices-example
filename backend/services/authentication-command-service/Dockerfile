FROM node:lts-stretch As development

WORKDIR /usr/src/app

COPY package*.json ./

RUN npm update

RUN npm install -g node-gyp

RUN npm install node-pre-gyp

RUN npm install --only=development

COPY . .

RUN npm run build

CMD ["node", "dist/main"]