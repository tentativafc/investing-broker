const mongoose = require("mongoose");

const url =
  "mongodb://mongouser:mongopass@mongo_ibovespa_assets:27017/ibovespa_assets";

mongoose.connect(url);

module.exports = { mongoose };
