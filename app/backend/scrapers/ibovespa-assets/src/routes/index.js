import { IbovespaAssets } from "../models/index";
import Scraper from "../scrapers/index";

const API_PATH = "/scrapers/ibovespa-assets";

const get = async (req, res, next) => {
  try {
    let ibovespaAssets = await IbovespaAssets.find({});
    res.json(ibovespaAssets);
  } catch (err) {
    res.json(500, { message: "Error get ibovespa assets", cause: err });
  }
};

const post = (req, res, next) => {
  let scraper = new Scraper();
  scraper.load().subscribe(
    (ibovespa_assets) => res.json(ibovespa_assets),
    (err) => {
      res.json(500, { message: "Error to scrap ibovespa assets", cause: err });
    }
  );
};

function Routes(server) {
  server.get(API_PATH, get);
  server.post(API_PATH, post);
}

module.exports = Routes;
