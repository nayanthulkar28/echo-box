FROM node:21-alpine
WORKDIR /workspace-ui
COPY public public
COPY src src
COPY package.json package.json
COPY .env.production .env.production
RUN npm install
RUN npm install -g serve
RUN npm run build
EXPOSE 4000
CMD ["serve", "-s", "build", "-l", "4000"]