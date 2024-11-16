FROM node:21-alpine
WORKDIR /workspace-ui
COPY public public
COPY src src
COPY package.json package.json
COPY .env .env
RUN npm install
RUN npm install -g serve
RUN npm run build
CMD ["serve", "-s", "build"]