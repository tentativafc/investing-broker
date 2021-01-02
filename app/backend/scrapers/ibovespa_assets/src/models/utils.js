// Arrow functions meaning self when translated by babel (https://github.com/Automattic/mongoose/issues/4816)
const beforeSave = function (next) {
  let currentDate = new Date();
  this.updated_at = currentDate;
  if (!this.created_at) {
    this.created_at = currentDate;
  }
  next();
};

module.exports = {
  beforeSave,
};
