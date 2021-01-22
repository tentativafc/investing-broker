import mongoose from "mongoose";

const url = "mongodb://mongouser:mongopass@mongo-selic:27017/selic";
mongoose.connect(url);

module.exports = { mongoose };
