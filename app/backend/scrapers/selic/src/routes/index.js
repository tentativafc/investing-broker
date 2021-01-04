const moment = require("moment");
const models = require("../models/index");
const Scraper = require("../scrapers/index");

const API_PATH = "/api/selic";

const get = async function (req, res, next) {
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
      let selicValues = await models.Selic.find(filter);
      res.json(selicValues);
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
