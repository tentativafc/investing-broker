const axios = require("axios");
const cheerio = require("cheerio");

var models = require("../models/models");

const url =
  "http://bvmf.bmfbovespa.com.br/indices/ResumoCarteiraTeorica.aspx?Indice=IBOV&idioma=pt-br";

const fethHtml = async (url) => {
  try {
    const { data } = await axios.get(url);
    return data;
  } catch {
    console.error(
      `ERROR: An error occurred while trying to fetch the URL: ${url}`
    );
  }
};

class Scraper {
  async load() {
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

    let ibovespaAssets = new models.IbovespaAssets({ title, assets });

    ibovespaAssets.save(function (err, ibovespaAssetsPersisted) {
      if (err) {
        return console.error(err);
      }
      console.log(ibovespaAssetsPersisted);
    });
  }
}

module.exports = Scraper;
