const mongoose = require("mongoose");

const url = "mongodb://mongouser:mongopass@mongo-currency:27017/currency";

mongoose.connect(url, { useFindAndModify: false });

module.exports = { mongoose };
