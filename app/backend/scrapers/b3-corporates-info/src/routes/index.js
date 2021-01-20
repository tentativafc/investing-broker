const Scraper = require("../scrapers/index");

const API_PATH = "/scrapers/b3-corporates-info";

const get = async function (req, res, next) {
  try {
    let corporates = await models.CorporateInfo.find({});
    res.json(corporates);
  } catch (error) {
    console.error(error);
    res.send(500);
  }
};

const post = async function (req, res, next) {
  try {
    let scraper = new Scraper();
    let result = await scraper.load();
    res.json(result);
  } catch (error) {
    console.error(error);
    res.send(500);
  }
};

function Routes(server) {
  server.get(API_PATH, get);
  server.post(API_PATH, post);
}

module.exports = Routes;
