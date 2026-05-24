FROM node:22-alpine AS dev
WORKDIR /app

RUN corepack enable && corepack prepare pnpm@11.1.2 --activate

# Keep the pnpm store outside /app — /app is a bind mount in dev,
# so a project-local store would leak 350MB+ onto the host.
RUN pnpm config set store-dir /pnpm/store

COPY . .
RUN pnpm install --frozen-lockfile
EXPOSE 3000
CMD ["pnpm", "dev", "--host", "0.0.0.0"]
FROM dev AS build
RUN pnpm build

FROM node:22-alpine AS prod
WORKDIR /app
ENV NODE_ENV=production
ENV NITRO_PORT=3000
ENV NITRO_HOST=0.0.0.0
COPY --from=build /app/.output ./.output
EXPOSE 3000
CMD ["node", ".output/server/index.mjs"]
