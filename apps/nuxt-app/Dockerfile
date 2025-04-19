FROM node:20-alpine AS base

FROM base AS setup
WORKDIR /app
ENV MOON_TOOLCHAIN_FORCE_GLOBALS=true

# Add dependencies to get Bun working on Alpine
RUN apk --no-cache add ca-certificates wget
RUN wget https://raw.githubusercontent.com/athalonis/docker-alpine-rpi-glibc-builder/master/glibc-2.26-r1.apk
RUN apk add --allow-untrusted --force-overwrite glibc-2.26-r1.apk
RUN rm glibc-2.26-r1.apk

RUN npm install -g @moonrepo/cli
RUN npm i -g bun

FROM setup AS skeleton

# Copy entire repository and scaffold
COPY . .
RUN moon docker scaffold nuxt-app

#### BUILD
FROM setup AS build

# Copy workspace skeleton
COPY --from=skeleton /app/.moon/docker/workspace .

# Install toolchain and dependencies
RUN moon docker setup

# Copy source files
COPY --from=skeleton /app/.moon/docker/sources .

RUN moon run nuxt-app:build

FROM base AS runner
WORKDIR /app

COPY --from=build /app/apps/nuxt-app/.output/ ./

ENV PORT=80
ENV HOST=0.0.0.0

CMD ["node", "/app/server/index.mjs"]
