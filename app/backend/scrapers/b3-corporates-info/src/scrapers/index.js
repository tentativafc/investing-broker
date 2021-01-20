const axios = require("axios");
const cheerio = require("cheerio");
const models = require("../models/index");

const LETTERS = [
  "A",
  "B",
  "C",
  "D",
  "E",
  "F",
  "G",
  "H",
  "I",
  "J",
  "L",
  "M",
  "N",
  "O",
  "P",
  "Q",
  "R",
  "S",
  "T",
  "U",
  "V",
  "X",
  "Z",
];

const fetchHtml = async (url) => {
  const { data } = await axios.get(url);
  return data;
};

class Scraper {
  async load() {
    try {
      // Get general info of corporates
      await this.listCorporates();
      await this.detailCorporateInfo();
    } catch (error) {
      console.err(error);
      throw error;
    }
  }

  async listCorporates() {
    LETTERS.map(async (letter) => {
      let url = `http://bvmf.bmfbovespa.com.br/cias-listadas/empresas-listadas/BuscaEmpresaListada.aspx?Letra=${letter}&idioma=pt-br`;
      console.log(`Scraping corporates with letter ${letter}:  ${url}`);
      let html = await fetchHtml(url);
      let $ = cheerio.load(html);
      let corporates_table = $(
        "#ctl00_contentPlaceHolderConteudo_BuscaNomeEmpresa1_grdEmpresa_ctl01"
      );
      $(".GridRow_SiteBmfBovespa", corporates_table).each(async (i, row) => {
        let name = $("td:nth-child(1)", row).text();
        let negotiation_name = $("td:nth-child(2)", row).text();
        let segment = $("td:nth-child(3)", row).text();
        let link_url = $("a", row).attr("href");
        let regex_cvm_code = /.*codigoCvm=(\d+)/;
        var match = regex_cvm_code.exec(link_url);
        let cvm_code = match[1];

        let data = { name, cvm_code, negotiation_name, segment };
        const filter = { cvm_code };

        return await models.CorporateInfo.findOneAndUpdate(filter, data, {
          new: true,
          upsert: true,
        });
      });
      return "";
    });
  }

  async detailCorporateInfo() {
    let corporates = await models.CorporateInfo.find({});

    await corporates.forEach(async (corporate) => {
      let url = `http://bvmf.bmfbovespa.com.br/pt-br/mercados/acoes/empresas/ExecutaAcaoConsultaInfoEmp.asp?CodCVM=${corporate.cvm_code}`;

      console.log(`Scraping corporate ${corporate.name} - ${url}`);

      let html = await fetchHtml(url);
      let $ = cheerio.load(url);

      let table_sumary_info = $("table.ficha");

      let assets_code = [];
      $("a.LinkCodNeg", table_sumary_info).each((i, link_asset_name) => {
        assets_code.push($(link_asset_name).text());
      });

      let cnpj_number = $(
        "tr:nth-child(3) td:nth-child(2)",
        table_sumary_info
      ).text();

      let main_activity = $(
        "tr:nth-child(4) td:nth-child(2)",
        table_sumary_info
      ).text();

      let setorial_classes = $(
        "tr:nth-child(5) td:nth-child(2)",
        table_sumary_info
      )
        .text()
        .split("/");

      const filter = { cvm_code: corporate.cvm_code };
      await models.CorporateInfo.updateOne(filter, {
        $set: {
          cnpj_number,
          assets_code,
          main_activity,
          setorial_classes,
          link: url,
        },
      });
    });
  }
}

module.exports = Scraper;
