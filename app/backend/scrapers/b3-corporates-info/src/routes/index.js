import Scraper from "../scrapers/index";

const API_PATH = "/scrapers/b3-corporates-info";

const get = async function (req, res, next) {
  models.CorporateInfo.find({})
    .then((corporates) => {
      res.json(corporates);
    })
    .catch((err) => {
      console.error(err);
      res.send(500);
    });
};

const post = async function (req, res, next) {
  let scraper = new Scraper();
  scraper.load().subscribe(
    (data) => console.log("DATAAAAA", data),
    (err) => {
      console.log(err);
      res.json(err);
    },
    () => res.json("Complete")
  );
};

function Routes(server) {
  server.get(API_PATH, get);
  server.post(API_PATH, post);
}

module.exports = Routes;
