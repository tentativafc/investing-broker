const mongoose = require("mongoose");

const url = "mongodb://mongo:123456@localhost:27017";

mongoose.connect(url);

module.exports = { mongoose };
