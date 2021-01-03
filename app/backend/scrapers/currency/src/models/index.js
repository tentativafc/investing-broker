const config = require("../config");
const utils = require("./utils");
const mongoose = config.mongoose;

const CurrencyPriceSchema = new mongoose.Schema(
  {
    symbol: { type: String, required: true },
    date: { type: Date, required: true },
    code: { type: Number, required: true },
    type: { type: String, required: true },
    buy_value: { type: Number, required: true },
    sell_value: { type: Number, required: true },
    created_at: Date,
    updated_at: Date,
  },
  { collection: "currency_prices", autoIndex: false }
);

CurrencyPriceSchema.pre("save", utils.beforeSave);

const CurrencyPrice = mongoose.model("CurrencyPrice", CurrencyPriceSchema);

module.exports = { CurrencyPrice };
