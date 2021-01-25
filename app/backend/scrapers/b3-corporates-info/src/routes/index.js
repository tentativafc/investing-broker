import Scraper from "../scrapers/index";
import { cache } from "../config";
import { CorporateInfo } from "../models";
const API_PATH = "/scrapers/b3-corporates-info";

const get = async (req, res, next) => {
  CorporateInfo.find({})
    .then((corporates) => {
      res.json(corporates);
    })
    .catch((err) => {
      res.json(500, {
        message: "Error to find b3 corporates info.",
        cause: err,
      });
    });
};

const get_redis =  (req, res, next) => {
  cache.get("corporates_info", (err, corporates_info_str) => {
    if (err) {
      res.json(500, {
        message: "Error to load cache.",
        cause: err,
      });
    } else {
      res.json(JSON.parse(corporates_info_str));
    }    
  }); 
};

const post = (req, res, next) => {
  let scraper = new Scraper();
  scraper.load().subscribe(
    (corporates) => res.json(corporates),
    (err) => {
      res.json(500, {
        message: "Error to scrap b3 corporates info.",
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
