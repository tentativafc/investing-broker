import { IbovespaAssets } from "../models/index";
import Scraper from "../scrapers/index";

const API_PATH = "/scrapers/ibovespa-assets";

const get = async function (req, res, next) {
  try {
    let ibovespaAssets = await IbovespaAssets.find({});
    res.json(ibovespaAssets);
  } catch (error) {
    console.error(error);
    res.send(500);
  }
};

const post = function (req, res, next) {
  let scraper = new Scraper();
  scraper.load().subscribe(
    (ibovespa_assets) => res.json(ibovespa_assets),
    (err) => {
      res.send(500);
      res.json(err);
    }
  );
};

function Routes(server) {
  server.get(API_PATH, get);
  server.post(API_PATH, post);
}

module.exports = Routes;
