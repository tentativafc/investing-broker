const moment = require("moment");
const models = require("../models/index");
const Scraper = require("../scrapers/index");

const API_PATH = "/scrapers/currency";

const get = async function (req, res, next) {
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
      let currencyPrices = await models.CurrencyPrice.find(filter);
      res.json(currencyPrices);
    }
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
