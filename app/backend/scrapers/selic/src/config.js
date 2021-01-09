const mongoose = require("mongoose");

const url = "mongodb://mongouser:mongopass@mongo_selic:27017/selic";
mongoose.connect(url);

module.exports = { mongoose };
