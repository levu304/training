import { join } from "node:path";
import viteFastify from "@fastify/vite/plugin";
import viteReact from "@vitejs/plugin-react-swc";
import tsConfigPaths from "vite-tsconfig-paths";
import tailwindcss from "@tailwindcss/vite";

export default {
  root: join(import.meta.dirname, "src", "client"),
  plugins: [
    tsConfigPaths({
      root: import.meta.dirname
    }),
    viteFastify({
      spa: true
    }),
    viteReact(),
    tailwindcss()
  ],
  build: {
    minify: true,
    emptyOutDir: true,
    outDir: join(import.meta.dirname, "dist", "client")
  }
};
