import Scraper from "../scrapers/index";

const API_PATH = "/scrapers/b3-corporates-info";

const get = async (req, res, next) => {
  models.CorporateInfo.find({})
    .then((corporates) => {
      res.json(corporates);
    })
    .catch((err) => {
      res.json(500, {
        message: "Error to find b3 corporates info",
        cause: err,
      });
    });
};

const post = (req, res, next) => {
  let scraper = new Scraper();
  scraper.load().subscribe(
    (corporates) => res.json(corporates),
    (err) => {
      res.json(500, {
        message: "Error to scrap b3 corporates info",
        cause: err,
      });
    }
  );
};

function Routes(server) {
  server.get(API_PATH, get);
  server.post(API_PATH, post);
}

module.exports = Routes;
