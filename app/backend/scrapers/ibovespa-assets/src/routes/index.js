const models = require("../models/index");
const Scraper = require("../scrapers/index");

const API_PATH = "/scrapers/ibovespa-assets";

const get = async function (req, res, next) {
  try {
    let ibovespaAssets = await models.IbovespaAssets.find({});
    res.json(ibovespaAssets);
  } catch (error) {
    console.error(error);
    res.send(500);
  }
};

const post = async function (req, res, next) {
  try {
    let scraper = new Scraper();
    res.json(await scraper.load());
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
