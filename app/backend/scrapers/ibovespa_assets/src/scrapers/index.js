const axios = require("axios");
const cheerio = require("cheerio");
var models = require("../models/index");

const url =
  "http://bvmf.bmfbovespa.com.br/indices/ResumoCarteiraTeorica.aspx?Indice=IBOV&idioma=pt-br";

const fethHtml = async (url) => {
  const { data } = await axios.get(url);
  return data;
};

class Scraper {
  async load() {
    try {
      const html = await fethHtml(url);

      let $ = cheerio.load(html);

      let assets = [];

      let rows = $(".rgMasterTable tbody tr");

      rows.each((i, row) => {
        let values = [];
        $(row)
          .find("td span")
          .each(function (i, span) {
            values.push($(span).text());
          });

        let [
          symbol,
          corporate_name,
          type,
          theoretical_quantity,
          percentage,
        ] = values;

        type = type.replace(/\s+/g, " ");

        theoretical_quantity = theoretical_quantity.replace(/\./g, "");
        theoretical_quantity = theoretical_quantity.replace(/\,/g, ".");

        theoretical_quantity = parseFloat(theoretical_quantity);

        percentage = percentage.replace(/\./g, "");
        percentage = percentage.replace(/\,/g, ".");

        percentage = parseFloat(percentage);

        const asset = new models.Asset({
          symbol,
          corporate_name,
          type,
          theoretical_quantity,
          percentage,
        });
        assets.push(asset);
      });

      let title = $("#ctl00_contentPlaceHolderConteudo_lblTitulo").text();

      let data = { title, assets };

      const filter = { title };

      return await models.IbovespaAssets.findOneAndUpdate(filter, data, {
        new: true,
        upsert: true, // Make this update into an upsert
      });
    } catch (error) {
      throw new Exception("Error to load Ibovespa Assets.");
    }
  }
}

module.exports = Scraper;
