#### BASE
FROM node:alpine as base

ENV MOON_TOOLCHAIN_FORCE_GLOBALS=true
WORKDIR /app

# Add dependencies to get Bun working on Alpine
RUN apk --no-cache add ca-certificates wget
RUN wget https://raw.githubusercontent.com/athalonis/docker-alpine-rpi-glibc-builder/master/glibc-2.26-r1.apk
RUN apk add --allow-untrusted --force-overwrite glibc-2.26-r1.apk
RUN rm glibc-2.26-r1.apk

# Install moon binary
RUN npm install -g @moonrepo/cli
RUN npm i -g bun

#### SKELETON
FROM base AS skeleton

# Copy entire repository and scaffold
COPY . .
RUN moon docker scaffold waku-project

#### BUILD
FROM base AS build

# Copy workspace skeleton
COPY --from=skeleton /app/.moon/docker/workspace .

# Install toolchain and dependencies
RUN moon docker setup

# Copy source files
COPY --from=skeleton /app/.moon/docker/sources .

RUN moon run waku-project:build

# Prune workspace
RUN moon docker prune