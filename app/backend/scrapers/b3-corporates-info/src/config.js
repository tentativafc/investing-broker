const mongoose = require("mongoose");

const url = "mongodb://mongouser:mongopass@localhost:27020/b3-corporates-info";

mongoose.connect(url, { useFindAndModify: false });

module.exports = { mongoose };
