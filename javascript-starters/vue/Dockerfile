FROM node:24-alpine AS build
WORKDIR /app

COPY package.json package-lock.json ./
RUN npm ci

COPY . .
RUN npm run build-only

FROM nginx:1.29-alpine

# Where /api requests are forwarded; override with `docker run -e API_URL=...`
ENV API_URL=http://localhost:8084

COPY nginx.conf.template /etc/nginx/templates/default.conf.template
COPY --from=build /app/dist /usr/share/nginx/html

EXPOSE 8080
