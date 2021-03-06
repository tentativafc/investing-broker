import moment from "moment";
import { Selic } from "../models";
import Scraper from "../scrapers";
import { cache } from "../config";

const API_PATH = "/scrapers/selic";

const get = async (req, res, next) => {
  try {
    let begin_date = req.params.begin_date
      ? moment(req.params.begin_date, "YYYYMMDD").toDate()
      : null;
    let final_date = req.params.final_date
      ? moment(req.params.final_date, "YYYYMMDD").toDate()
      : null;

    if (!begin_date) {
      res.send(400, {
        message: "Params expected: [begin_date, final_date (optional)]",
      });
    } else {
      let filter = { begin_date: { $gte: begin_date } };
      if (final_date) {
        filter = { ...filter, final_date: { $lte: final_date } };
      }
      let selicValues = await Selic.find(filter);
      res.json(selicValues);
    }
  } catch (err) {
    res.send(500, { message: "Error to find selic values", cause: err });
  }
};

const get_redis = (req, res, next) => {
  cache.get("selic_prices", (err, selic_prices_str) => {
    if (err) {
      res.json(500, {
        message: "Error to load cache.",
        cause: err,
      });
    } else {
      res.json(JSON.parse(selic_prices_str));
    }
  });
};

const post = (req, res, next) => {
  let scraper = new Scraper();
  scraper.load().subscribe(
    (selic_prices) => res.json(selic_prices),
    (err) => {      
      res.json(500, {
        message: "Error to scrap selic prices",
        cause: err,
      });
    }
  );
};

function Routes(server) {
  server.get(API_PATH, get);
  server.get(API_PATH + "/redis", get_redis);
  server.post(API_PATH, post);
}

module.exports = Routes;
