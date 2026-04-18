FROM oven/bun:1 AS builder
WORKDIR /app

COPY package.json bun.lockb turbo.json ./
COPY apps/frontend/package.json ./apps/frontend/
COPY packages/ ./packages/
RUN bun install --frozen-lockfile

COPY apps/frontend ./apps/frontend
RUN bun run build --filter=@aropi/frontend

FROM nginx:alpine AS runner
COPY --from=builder /app/apps/frontend/dist /usr/share/nginx/html

COPY apps/frontend/nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
