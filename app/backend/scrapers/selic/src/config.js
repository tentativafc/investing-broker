const mongoose = require("mongoose");

const url = "mongodb://mongouser:mongopass@mongo_selic:27019/selic";

mongoose.connect(url, { useFindAndModify: false });

module.exports = { mongoose };
