#### BASE STAGE
#### Installs moon.

FROM golang:1.24.5 AS base
WORKDIR /app

# Install moon binary
RUN curl -fsSL https://moonrepo.dev/install/moon.sh | bash
ENV PATH="/root/.moon/bin:$PATH"

#### SKELETON STAGE
#### Scaffolds repository skeleton structures.

FROM base AS skeleton

# Copy entire repository and scaffold
COPY . .
RUN moon docker scaffold go-api-server

#### BUILD STAGE
#### Builds the project.

FROM base AS build

# Copy toolchain
# COPY --from=skeleton /root/.proto /root/.proto

# Copy workspace configs
COPY --from=skeleton /app/.moon/docker/workspace .

# Install dependencies
RUN moon docker setup

# Copy project sources
COPY --from=skeleton /app/.moon/docker/sources .

ENV CGO_ENABLED=0 GOOS=linux

# Build the project
RUN moon run go-api-server:build

# Prune extraneous dependencies
RUN moon docker prune

# ENTRYPOINT [ "tail", "-f", "/dev/null" ]

FROM scratch
COPY --from=build /app/apps/go-api-server/bin/main .
CMD [ "./main" ]