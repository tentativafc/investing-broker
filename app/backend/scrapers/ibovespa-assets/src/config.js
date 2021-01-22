import mongoose from "mongoose";

const url = "mongodb://mongouser:mongopass@localhost:27017/ibovespa_assets";

mongoose.connect(url);

module.exports = { mongoose };
