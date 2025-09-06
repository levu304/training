#### BASE STAGE
#### Installs moon.

FROM node:22-slim AS base
WORKDIR /app

RUN apt-get update && apt-get install -y curl unzip

# Install moon binary
RUN curl -fsSL https://moonrepo.dev/install/moon.sh | bash
ENV PATH="/root/.moon/bin:/root/.proto/bin:$PATH"

RUN curl -fsSL https://bun.com/install | bash
ENV PATH="/root/.bun/bin:$PATH"

#### SKELETON STAGE
#### Scaffolds repository skeleton structures.

FROM base AS skeleton

# Copy entire repository and scaffold
COPY . .
RUN moon docker scaffold triplit-remote-server

#### BUILD STAGE
#### Builds the project.

FROM base AS build

# Copy workspace configs
COPY --from=skeleton /app/.moon/docker/workspace .

# Install dependencies
RUN moon docker setup

# Copy project sources
COPY --from=skeleton /app/.moon/docker/sources .

# Build the project
RUN moon run triplit-remote-server:build

# Prune extraneous dependencies
RUN moon docker prune

RUN moon run triplit-remote-server:compile

#### RELEASE STAGE
FROM debian:stable-slim AS release
WORKDIR /app
COPY --from=build /app/apps/triplit-remote-server/build/server .
ENTRYPOINT [ "./server" ]

