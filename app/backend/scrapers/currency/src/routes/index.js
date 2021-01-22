import moment from "moment";
import { CurrencyPrice } from "../models";
import Scraper from "../scrapers";

const API_PATH = "/scrapers/currency";

const get = async (req, res, next) => {
  try {
    let symbol = req.params.symbol;
    let begin_date = req.params.begin_date
      ? moment(req.params.begin_date, "YYYYMMDD").toDate()
      : "";
    let final_date = req.params.final_date
      ? moment(req.params.final_date, "YYYYMMDD").toDate()
      : "";

    if (!symbol || !begin_date || !final_date) {
      res.send(400, {
        message: "Params expected: [symbol, begin_date, final_date]",
      });
    } else {
      let filter = { symbol, date: { $gte: begin_date, $lt: final_date } };
      let currencyPrices = await CurrencyPrice.find(filter);
      res.json(currencyPrices);
    }
  } catch (err) {
    res.json({ message: "Error to get currencies prices", cause: err });
  }
};

const post = (req, res, next) => {
  let scraper = new Scraper();
  scraper.load().subscribe(
    (prices) => {
      console.log("Price received");
      res.json(prices);
    },
    (err) => {
      res.send(500, { message: "Error to scrap currencies", cause: err });
    },
    () => {
      console.log("Complete");
    }
  );
};

function Routes(server) {
  server.get(API_PATH, get);
  server.post(API_PATH, post);
}

module.exports = Routes;
