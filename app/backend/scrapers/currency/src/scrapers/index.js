const axios = require("axios");
const cheerio = require("cheerio");
const models = require("../models/index");
const moment = require("moment");

const CurrencyEnum = {
  dolar: 61,
  euro: 222,
  iene: 101,
  iuan: 178,
};

const fetchData = async (url) => {
  const { data } = await axios.get(url);
  return data;
};

class Scraper {
  async load() {
    try {
      let begin_date = moment().subtract(1, "years").format("DD/MM/YYYY");
      let final_date = moment().format("DD/MM/YYYY");
      for (let currency in CurrencyEnum) {
        const url = `https://ptax.bcb.gov.br/ptax_internet/consultaBoletim.do?method=gerarCSVFechamentoMoedaNoPeriodo&ChkMoeda=${CurrencyEnum[currency]}&DATAINI=${begin_date}&DATAFIM=${final_date}`;
        const data_csv = await fetchData(url);
        data_csv.split("\n").forEach(async (row) => {
          if (!row) {
            return;
          }
          let cols = row.split(";");
          let [date, code, type, symbol, buy_value, sell_value] = cols;

          date = moment(date, "DDMMYYYY").toDate();

          buy_value = buy_value.replace(/\,/g, ".");
          buy_value = parseFloat(buy_value);

          sell_value = sell_value.replace(/\,/g, ".");
          sell_value = parseFloat(sell_value);

          let filter = { date, symbol };

          let data = { date, code, type, symbol, buy_value, sell_value };

          try {
            await models.CurrencyPrice.findOneAndUpdate(filter, data, {
              new: true,
              upsert: true,
            });
          } catch (error) {
            console.log("Error when save");
            throw error;
          }
        });
      }
    } catch (error) {
      throw new Error("Error to load currency prices.");
    }
  }
}
module.exports = Scraper;
