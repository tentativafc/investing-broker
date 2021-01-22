import mongoose from "mongoose";

const url = "mongodb://mongouser:mongopass@localhost:27019/selic";
mongoose.connect(url);

module.exports = { mongoose };
