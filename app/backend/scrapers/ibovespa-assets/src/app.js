import regeneratorRuntime from "regenerator-runtime";

import restify from "restify";

const server = restify.createServer();

// Middleware
server.use((req, res, next) => {
  res.header("Access-Control-Allow-Origin", "*");
  res.header(
    "Access-Control-Allow-Methods",
    "GET, POST, PUT, DELETE, OPTIONS, HEAD"
  );
  return next();
});

server.use(
  restify.plugins.bodyParser({
    mapParams: true,
  })
);

server.opts("/.*/", (req, res, next) => {
  res.send(200);
  return next();
});

require("./routes/index")(server);

// Routes
server.listen(3000, function () {
  console.log("App listening on port 3000!");
});

module.exports = server;
