import mongoose from "mongoose";
import config from "config";

const db = config.get("Configs.db");
const url = `mongodb://${db.user}:${db.password}@${db.host}:${db.port}/${db.database}`;

mongoose.connect(url, { useFindAndModify: false });

const AXIOS_TIMEOUT_MS = config.get("Configs.axios").timeout_in_ms;

module.exports = { mongoose, AXIOS_TIMEOUT_MS };
