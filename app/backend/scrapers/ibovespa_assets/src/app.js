const restify = require("restify");
const models = require("./models/models");
const Scraper = require("./scraper/Scraper");
const callback_crud = (req, res, next, transform_data_callback) => {
  return (err, document) => {
    if (err) {
      console.error(err);
      res.send(new Error(err));
    } else {
      var transformed_document = document;
      if (transform_data_callback) {
        transformed_document = transform_data_callback(document);
      }
      res.json(transformed_document);
    }
    return next();
  };
};

module.exports = callback_crud;

const server = restify.createServer();

// Middleware
server.use((req, res, next) => {
  res.header("Access-Control-Allow-Origin", "*");
  res.header(
    "Access-Control-Allow-Headers",
    "Origin, X-Requested-With, Content-Type, Accept"
  );
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

server.on("uncaughtException", function (req, res, route, err) {
  console.error("Erro não tratado", err);
});

server.opts("/.*/", (req, res, next) => {
  res.send(200);
  return next();
});

const API_PATH = "/api/ibovespa_assets";

const get = async function (req, res, next) {
  models.IbovespaAssets.find({}).exec(
    callback_crud(req, res, next, (assets) => {
      return assets;
    })
  );
};

const post = async function (req, res, next) {
  try {
    let scraper = new Scraper();
    res.json(scraper.load());
  } catch (error) {
    console.error(error);
    res.send(500);
  }
};

server.get(API_PATH, get);
server.post(API_PATH, post);

// Routes
server.listen(3000, function () {
  console.log("App listening on port 3000!");
});

module.exports = server;
