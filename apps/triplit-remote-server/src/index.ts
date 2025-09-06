import { createServer, createTriplitStorageProvider } from "@triplit/server";

const port = +(process.env.PORT || 8080);

const startServer = await createServer({
  storage: await createTriplitStorageProvider("lmdb"),
  verboseLogs: true,
  jwtSecret: "secret",
});

const dbServer = startServer(port, () => {
  console.log("Server started");
  console.log("running on port", port);
});

process.on("SIGINT", function () {
  dbServer.close(() => {
    console.log("Shutting down server... ");
    process.exit();
  });
});
