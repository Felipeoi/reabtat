FROM node:18-alpine as builder

WORKDIR /app

COPY package.json package-lock.json* ./
RUN npm ci

COPY . .

# URL da API no navegador (host acessa backend na mesma máquina)
ARG VITE_API_URL=http://localhost:8080
ENV VITE_API_URL=$VITE_API_URL

RUN npm run build

FROM nginx:alpine

WORKDIR /usr/share/nginx/html

COPY --from=builder /app/dist .
COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]