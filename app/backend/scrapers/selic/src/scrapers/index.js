import axios from "axios";
import { Selic } from "../models/index";
import moment from "moment";
import { from, of } from "rxjs";
import { map, mergeMap, toArray, retry } from "rxjs/operators";
import { AXIOS_TIMEOUT_MS } from "../config";

const fetchData = async (url) => {
  const { data } = await axios.get(url, { timeout: AXIOS_TIMEOUT_MS });
  return data;
};

const URL = "https://www.bcb.gov.br/api/servico/sitebcb/historicotaxasjuros";

class Scraper {
  load() {
    return of(URL).pipe(
      mergeMap((url) => {
        return fetchData(url);
      }),
      //retry 2 times on erro
      retry(2),
      mergeMap((data) => {
        console.log("Mergemap data", data);
        let selic_prices = [];
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

          let selic_price = {
            assembly_number,
            assembly_date,
            begin_date,
            final_date,
            goal_value,
            efective_value,
            annual_effective_value,
          };

          selic_prices.push(selic_price);
        });

        // Return new observer
        return from(selic_prices).pipe(map((selic_price) => selic_price));
      }),
      mergeMap((selic_price) => {
        let filter = { begin_date: selic_price.begin_date };
        return Selic.findOneAndUpdate(filter, selic_price, {
          new: true,
          upsert: true,
        });
      }),
      toArray()
    );
  }
}

module.exports = Scraper;
