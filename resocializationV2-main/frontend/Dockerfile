FROM node:20-alpine AS builder

WORKDIR /app

COPY package.json package-lock.json* ./
RUN npm ci

COPY . .

# URL da API no navegador — em produção (Railway) defina build-arg VITE_API_URL com a URL pública da API
ARG VITE_API_URL=http://localhost:8080
ENV VITE_API_URL=$VITE_API_URL

RUN npm run build

FROM nginx:alpine

RUN apk add --no-cache gettext

WORKDIR /usr/share/nginx/html

COPY --from=builder /app/dist .
COPY nginx/default.conf.template /etc/nginx/templates/default.conf.template
COPY docker-entrypoint.sh /docker-entrypoint.sh
RUN chmod +x /docker-entrypoint.sh && rm -f /etc/nginx/conf.d/default.conf

EXPOSE 80

ENTRYPOINT ["/docker-entrypoint.sh"]
