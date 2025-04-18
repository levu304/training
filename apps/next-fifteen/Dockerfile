FROM node:20-alpine as base

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
RUN moon docker scaffold next-fifteen

#### BUILD
FROM setup AS build

# Copy workspace skeleton
COPY --from=skeleton /app/.moon/docker/workspace .

# Install toolchain and dependencies
RUN moon docker setup

# Copy source files
COPY --from=skeleton /app/.moon/docker/sources .

RUN moon run next-fifteen:build

FROM base AS runner
WORKDIR /app

# Don't run production as root
RUN addgroup --system --gid 1001 nodejs
RUN adduser --system --uid 1001 nextjs
USER nextjs
 
# Automatically leverage output traces to reduce image size
# https://nextjs.org/docs/advanced-features/output-file-tracing
COPY --from=build --chown=nextjs:nodejs /app/apps/next-fifteen/.next/standalone ./
COPY --from=build --chown=nextjs:nodejs /app/apps/next-fifteen/.next/static ./.next/static
COPY --from=build --chown=nextjs:nodejs /app/apps/next-fifteen/public ./public

ENV HOST=0.0.0.0
ENV PORT=8080
 
CMD ["node", "server.js"]