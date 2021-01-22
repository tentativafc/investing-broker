import axios from "axios";
import cheerio from "cheerio";
import { from, of, combineLatest } from "rxjs";
import { map, mergeMap, toArray } from "rxjs/operators";
import { CorporateInfo } from "../models";
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

const REGEX_CVM_CODE = /.*codigoCvm=(\d+)/;

const TIMEOUT_DEFAULT_IN_MS = 60000;

const fetchHtml = async (url) => {
  const { data } = await axios.get(url, { timeout: TIMEOUT_DEFAULT_IN_MS });
  return data;
};

class Scraper {
  load() {
    return from(LETTERS).pipe(
      mergeMap((letter) => {
        let url = `http://bvmf.bmfbovespa.com.br/cias-listadas/empresas-listadas/BuscaEmpresaListada.aspx?Letra=${letter}&idioma=pt-br`;
        return fetchHtml(url);
      }),
      mergeMap((html) => {
        let $ = cheerio.load(html);
        let corporates_table = $(
          "#ctl00_contentPlaceHolderConteudo_BuscaNomeEmpresa1_grdEmpresa_ctl01"
        );

        let corporates = [];

        $(".GridRow_SiteBmfBovespa", corporates_table).each((i, row) => {
          let name = $("td:nth-child(1)", row).text();
          let negotiation_name = $("td:nth-child(2)", row).text();
          let segment = $("td:nth-child(3)", row).text();
          let link_url = $("a", row).attr("href");
          var match = REGEX_CVM_CODE.exec(link_url);
          let cvm_code = match[1];

          const corporate = { name, cvm_code, negotiation_name, segment };
          corporates.push(corporate);
        });
        // Return new observer
        return from(corporates).pipe(map((corporate) => corporate));
      }),
      mergeMap((corporate) => {
        const filter = { cvm_code: corporate.cvm_code };
        return CorporateInfo.findOneAndUpdate(filter, corporate, {
          new: true,
          upsert: true,
        });
      }),
      mergeMap((corporate) => {
        let url = `http://bvmf.bmfbovespa.com.br/pt-br/mercados/acoes/empresas/ExecutaAcaoConsultaInfoEmp.asp?CodCVM=${corporate.cvm_code}`;
        console.log(`Scraping corporate ${corporate.name} - ${url}`);
        return combineLatest(of(corporate), of(url), fetchHtml(url));
      }),
      mergeMap(([corporate, url, html]) => {
        let $ = cheerio.load(html);
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

        let update_data = {
          cnpj_number,
          assets_code,
          main_activity,
          setorial_classes,
          link: url,
        };
        let filter = { cvm_code: corporate.cvm_code };
        return CorporateInfo.updateOne(filter, {
          $set: update_data,
        });
      }),
      toArray()
    );
  }
}

module.exports = Scraper;
