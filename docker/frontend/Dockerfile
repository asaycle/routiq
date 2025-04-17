# Dockerfile.frontend

# ビルドステージ
FROM node:23.3.0 AS builder

WORKDIR /app

COPY package.json package-lock.json ./
RUN npm ci

COPY . .
RUN npm run build

# serve stage
FROM nginx:alpine
COPY --from=builder /app/dist /usr/share/nginx/html