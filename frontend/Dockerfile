FROM node:20
#ENV NODE_ENV=production
ENV NODE_ENV=development
RUN npm install -g http-server
WORKDIR /app

COPY ["package.json", "package-lock.json*", "./"]
#RUN npm ci --production
RUN npm install

#COPY . .
#RUN wget https://github.com/justinfranco/mathjs/blob/add-Unit-type-definition-to-MathJsStatic/types/index.d.ts -O node_modules/mathjs/types/index.d.ts
#RUN npm run build

#EXPOSE 8080

#CMD [ "http-server", "dist" ]
#CMD [ "npm", "run", "dev" ]
