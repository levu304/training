import Fastify from "fastify";
import FastifyVite from "@fastify/vite";
import { resolve } from "node:path";

const server = Fastify({
  logger: {
    transport: {
      target: "@fastify/one-line-logger",
    },
  },
});

await server.register(FastifyVite, {
  root: resolve(import.meta.dirname, "../"),
  dev: process.argv.includes("--dev"),
  spa: true,
});

server.get("/", (req, reply) => {
  return reply.html();
});

async function main() {
  await server.vite.ready();
  await server.listen({ port: 3000 });
}

main();
