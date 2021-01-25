import { IbovespaAssets } from "../models/index";
import Scraper from "../scrapers/index";
import { cache } from "../config";

const API_PATH = "/scrapers/ibovespa-assets";

const get = async (req, res, next) => {
  try {
    let ibovespaAssets = await IbovespaAssets.find({});
    res.json(ibovespaAssets);
  } catch (err) {
    res.json(500, { message: "Error get ibovespa assets", cause: err });
  }
};

const get_redis = (req, res, next) => {
  cache.get("ibovespa_assets", (err, ibovespa_assets_str) => {
    if (err) {
      res.json(500, {
        message: "Error to load cache.",
        cause: err,
      });
    } else {
      res.json(JSON.parse(ibovespa_assets_str));
    }
  });
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
  server.get(API_PATH + "/redis", get_redis);
  server.post(API_PATH, post);
}

module.exports = Routes;
