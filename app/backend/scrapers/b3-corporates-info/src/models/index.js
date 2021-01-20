const config = require("../config");
const utils = require("./utils");
const mongoose = config.mongoose;

const CorporateInfoSchema = new mongoose.Schema({
  name: { type: String, required: true },
  cvm_code: { type: Number, required: true },
  negotiation_name: { type: String, required: true },
  segment: { type: String },
  cnpj_number: { type: String },
  main_activity: { type: String },
  assets_code: [
    {
      type: String,
    },
  ],
  setorial_classes: [
    {
      type: String,
    },
  ],
  link: { type: String },
});

CorporateInfoSchema.pre("save", utils.beforeSave);

const CorporateInfo = mongoose.model("CorporateInfo", CorporateInfoSchema);

module.exports = { CorporateInfo };
