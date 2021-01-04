const config = require("../config");
const utils = require("./utils");
const mongoose = config.mongoose;

const SelicSchema = new mongoose.Schema(
  {
    assembly_number: { type: Number, required: true },
    assembly_date: { type: Date, required: true },
    begin_date: { type: Date, required: true },
    final_date: { type: Date, required: false },
    goal_value: { type: Number, required: false },
    efective_value: { type: Number, required: false },
    annual_effective_value: { type: Number, required: false },
    created_at: Date,
    updated_at: Date,
  },
  { collection: "selic_values", autoIndex: false }
);

SelicSchema.pre("save", utils.beforeSave);

const Selic = mongoose.model("Selic", SelicSchema);

module.exports = { Selic };
