import axios from "axios";
import { CurrencyPrice } from "../models";
import moment from "moment";
import { from } from "rxjs";
import { filter, mergeMap, toArray, retry } from "rxjs/operators";

const DOLAR = 61;
const EURO = 222;
const IENE = 101;
const IUAN = 178;

const CURRENCIES = [DOLAR, EURO, IENE, IUAN];

const fetchData = async (url) => {
  const { data } = await axios.get(url);
  return data;
};

class Scraper {
  load() {
    let begin_date = moment().subtract(1, "years").format("DD/MM/YYYY");
    let final_date = moment().format("DD/MM/YYYY");

    return from(CURRENCIES).pipe(
      mergeMap((currency_code) => {
        console.log("Mergemap inicial");
        if (true) {
          throw "Error xpto";
        }
        const url = `https://ptax.bcb.gov.br/ptax_internet/consultaBoletim.do?method=gerarCSVFechamentoMoedaNoPeriodo&ChkMoeda=${currency_code}&DATAINI=${begin_date}&DATAFIM=${final_date}`;
        return fetchData(url);
      }),
      //retry 2 times on erro
      retry(2),
      // Splits csv by row (\n)
      mergeMap((data_csv) => {
        return from(data_csv.split("\n")).pipe(filter((row) => row));
      }),
      mergeMap((row) => {
        let cols = row.split(";");
        let [date, code, type, symbol, buy_value, sell_value] = cols;

        date = moment(date, "DDMMYYYY").toDate();

        buy_value = buy_value.replace(/\,/g, ".");
        buy_value = parseFloat(buy_value);

        sell_value = sell_value.replace(/\,/g, ".");
        sell_value = parseFloat(sell_value);

        let filter = { date, symbol };
        let data = { date, code, type, symbol, buy_value, sell_value };
        return CurrencyPrice.findOneAndUpdate(filter, data, {
          new: true,
          upsert: true,
        });
      }),
      toArray()
    );
  }
}

module.exports = Scraper;
