import mongoose from "mongoose";
import redis from "redis";
import config from "config";

const db = config.get("Configs.db");
const url = `mongodb://${db.user}:${db.password}@${db.host}:${db.port}/${db.database}`;

mongoose.connect(url, { useFindAndModify: false });

const AXIOS_TIMEOUT_MS = config.get("Configs.axios").timeout_in_ms;

const cache = redis.createClient({
  host: config.get("Configs.redis").host,
  port: config.get("Configs.redis").port,
  retry_strategy: () => config.get("Configs.redis").time_to_retry_in_ms,
});

cache.on("connect", () => {
  console.log("REDIS READY");
});

cache.on("error", (e) => {
  console.log("REDIS ERROR", e);
});

module.exports = { mongoose, cache, AXIOS_TIMEOUT_MS };
