const mongoose = require("mongoose");

const url = "mongodb://mongouser:mongopass@localhost:27018/currency";

mongoose.connect(url, { useFindAndModify: false });

module.exports = { mongoose };
