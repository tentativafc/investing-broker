const config = require("../config");
const utils = require("./utils");
const mongoose = config.mongoose;

const AssetSchema = new mongoose.Schema({
  symbol: { type: String, required: true },
  corporate_name: { type: String, required: true },
  type: { type: String, required: true },
  theoretical_quantity: { type: Number, required: true },
  percentage: { type: Number, required: true },
  created_at: Date,
  updated_at: Date,
});

const IbovespaAssetsSchema = new mongoose.Schema(
  {
    title: { type: String, required: true },
    assets: [AssetSchema],
  },
  { collection: "ibovespa_assets", autoIndex: false }
);

IbovespaAssetsSchema.index({ symbol: 1, type: -1 });

IbovespaAssetsSchema.pre("save", utils.beforeSave);

const Asset = mongoose.model("Asset", AssetSchema);
const IbovespaAssets = mongoose.model("IbovespaAssets", IbovespaAssetsSchema);

module.exports = { Asset, IbovespaAssets };
