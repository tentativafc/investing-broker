const mongoose = require("mongoose");

const url = "mongodb://mongouser:mongopass@localhost:27019/selic";

mongoose.connect(url, { useFindAndModify: false });

module.exports = { mongoose };
