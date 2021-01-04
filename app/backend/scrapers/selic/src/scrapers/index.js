const axios = require("axios");
const models = require("../models/index");
const moment = require("moment");

const URL = "https://www.bcb.gov.br/api/servico/sitebcb/historicotaxasjuros";

const fetchData = async (url) => {
  const { data } = await axios.get(url);
  return data;
};

class Scraper {
  async load() {
    try {
      const data = await fetchData(URL);
      data.conteudo.forEach(async (selic_data) => {
        let assembly_number = selic_data.NumeroReuniaoCopom;
        let assembly_date = selic_data.DataReuniaoCopom
          ? moment(selic_data.DataReuniaoCopom)
          : null;

        let begin_date = moment(selic_data.DataInicioVigencia);
        let final_date = selic_data.DataFimVigencia
          ? moment(selic_data.DataFimVigencia)
          : null;

        let goal_value = selic_data.MetaSelic;
        let efective_value = selic_data.TaxaSelicEfetivaVigencia;
        let annual_effective_value = selic_data.TaxaSelicEfetivaAnualizada;

        let filter = { begin_date };

        let data_insert = {
          assembly_number,
          assembly_date,
          begin_date,
          final_date,
          goal_value,
          efective_value,
          annual_effective_value,
        };

        try {
          await models.Selic.findOneAndUpdate(filter, data_insert, {
            new: true,
            upsert: true,
          });
        } catch (error) {
          console.log("Error when save");
          throw error;
        }
      });
    } catch (error) {
      throw new Error("Error to load Selic.");
    }
  }
}
module.exports = Scraper;
