const mongoose = require("mongoose");

const url =
  "mongodb://mongouser:mongopass@mongo-b3-corporates-info:27017/b3-corporates-info";

mongoose.connect(url, { useFindAndModify: false });

module.exports = { mongoose };
