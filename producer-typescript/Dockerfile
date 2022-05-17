FROM node:18 as build

WORKDIR /src
COPY package*.json .
RUN npm install
COPY . .
RUN node --test
RUN npm run build

FROM node:18-alpine
WORKDIR /usr/app
COPY --from=build /src/dist .
COPY --from=build /src/package.json .
RUN npm install --omit=dev
EXPOSE 8080
CMD ["node", "index"]
