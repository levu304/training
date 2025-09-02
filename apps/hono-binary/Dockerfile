#### BASE STAGE
#### Installs moon.

FROM oven/bun:slim AS base
WORKDIR /app

RUN apt-get update && apt-get install -y curl

# Install moon binary
RUN curl -fsSL https://moonrepo.dev/install/moon.sh | bash
ENV PATH="/root/.moon/bin:$PATH"

#### SKELETON STAGE
#### Scaffolds repository skeleton structures.

FROM base AS skeleton

# Copy entire repository and scaffold
COPY . .
RUN moon docker scaffold hono-binary

#### BUILD STAGE
#### Builds the project.

FROM base AS build

# Copy toolchain
COPY --from=skeleton /root/.proto /root/.proto

# Copy workspace configs
COPY --from=skeleton /app/.moon/docker/workspace .

# Install dependencies
RUN moon docker setup

# Copy project sources
COPY --from=skeleton /app/.moon/docker/sources .

# Build the project
RUN moon run hono-binary:build

#### RELEASE STAGE
FROM debian:stable-slim AS release
WORKDIR /app
COPY --from=build /app/apps/hono-binary/server .
ENTRYPOINT [ "./server" ]
