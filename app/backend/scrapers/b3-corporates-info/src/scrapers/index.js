import axios from "axios";
import cheerio from "cheerio";
import { of, from, merge, Observable } from "rxjs";
import { map, mergeMap, switchMap, toArray } from "rxjs/operators";
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

const LETTERS_RX = of(
  "A"
  // "B",
  // "C",
  // "D",
  // "E",
  // "F",
  // "G",
  // "H",
  // "I",
  // "J",
  // "L",
  // "M",
  // "N",
  // "O",
  // "P",
  // "Q",
  // "R",
  // "S",
  // "T",
  // "U",
  // "V",
  // "X",
  // "Z"
);

const REGEX_CVM_CODE = /.*codigoCvm=(\d+)/;

const TIMEOUT_DEFAULT_IN_MS = 10000;

const fetchHtml = async (url) => {
  // const { data } = await axios.get(url, { timeout: TIMEOUT_DEFAULT_IN_MS });
  const { data } = await axios.get(url);
  return data;
};

class Scraper {
  load() {
    return this.testingRx();
    //this.listCorporatesRx();
    // // Get general info of corporates
    // return this.listCorporates()
    //   .then((data) => {
    //     // Get detail info
    //     return this.detailCorporateInfo();
    //   })
    //   .catch((err) => {
    //     return new Promise((resolve, reject) => {
    //       reject(err);
    //     });
    //   });
  }

  testingRx() {
    return LETTERS_RX.pipe(
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
      switchMap((corporate) => {
        const filter = { cvm_code: corporate.cvm_code };
        return CorporateInfo.findOneAndUpdate(filter, corporate, {
          new: true,
          upsert: true,
        });
      }),
      toArray()
    );
  }

  // listCorporatesRx() {
  //   LETTERS_RX.pipe(
  //     map((letter) => {
  //       console.log(`Scraping corporates with letter ${letter}:  ${url}`);
  //       return from(fetchHtml(url));
  //     })
  //       .map((html) => {
  //         let $ = cheerio.load(html);
  //         let corporates_table = $(
  //           "#ctl00_contentPlaceHolderConteudo_BuscaNomeEmpresa1_grdEmpresa_ctl01"
  //         );

  //         let datas = [];

  //         $(".GridRow_SiteBmfBovespa", corporates_table).each((i, row) => {
  //           let name = $("td:nth-child(1)", row).text();
  //           let negotiation_name = $("td:nth-child(2)", row).text();
  //           let segment = $("td:nth-child(3)", row).text();
  //           let link_url = $("a", row).attr("href");
  //           var match = REGEX_CVM_CODE.exec(link_url);
  //           let cvm_code = match[1];

  //           const data = { name, cvm_code, negotiation_name, segment };
  //           datas.push(data);
  //         });
  //         returnd datas
  //       })
  //       .mergeMap((corporate) => {
  //         console.log(corporate);
  //         return corporate;
  //       })
  //     // .map(coporates => {
  //     //   return Promise.all(
  //     //     corporates.map((corporate) => {
  //     //       const filter = { cvm_code: corporate.cvm_code };
  //     //       return models.CorporateInfo.findOneAndUpdate(
  //     //         filter,
  //     //         corporate,
  //     //         {
  //     //           new: true,
  //     //           upsert: true,
  //     //         }
  //     //       );
  //     //     })
  //     //   );
  //     // })
  //   ).subscribe((value) => console.log(value));
  // }

  listCorporates() {
    return Promise.all(
      LETTERS.map((letter) => {
        let url = `http://bvmf.bmfbovespa.com.br/cias-listadas/empresas-listadas/BuscaEmpresaListada.aspx?Letra=${letter}&idioma=pt-br`;
        console.log(`Scraping corporates with letter ${letter}:  ${url}`);
        return fetchHtml(url)
          .then((html) => {
            return new Promise((resolve, reject) => {
              try {
                if (true) {
                  throw "Erro meu filho";
                }
                let $ = cheerio.load(html);
                let corporates_table = $(
                  "#ctl00_contentPlaceHolderConteudo_BuscaNomeEmpresa1_grdEmpresa_ctl01"
                );

                let datas = [];

                $(".GridRow_SiteBmfBovespa", corporates_table).each(
                  (i, row) => {
                    let name = $("td:nth-child(1)", row).text();
                    let negotiation_name = $("td:nth-child(2)", row).text();
                    let segment = $("td:nth-child(3)", row).text();
                    let link_url = $("a", row).attr("href");
                    var match = REGEX_CVM_CODE.exec(link_url);
                    let cvm_code = match[1];

                    const data = { name, cvm_code, negotiation_name, segment };
                    datas.push(data);
                  }
                );
                resolve(datas);
              } catch (err) {
                console.log("CAIU NO TRATAMENTOD E ERRO");
                reject(err);
              }
            });
          })
          .then((corporates) => {
            return Promise.all(
              corporates.map((corporate) => {
                const filter = { cvm_code: corporate.cvm_code };
                return models.CorporateInfo.findOneAndUpdate(
                  filter,
                  corporate,
                  {
                    new: true,
                    upsert: true,
                  }
                );
              })
            );
          });
      })
    );
  }

  detailCorporateInfo() {
    return models.CorporateInfo.find({}).then((corporates) => {
      return Promise.all(
        corporates.map((corporate) => {
          let url = `http://bvmf.bmfbovespa.com.br/pt-br/mercados/acoes/empresas/ExecutaAcaoConsultaInfoEmp.asp?CodCVM=${corporate.cvm_code}`;
          console.log(`Scraping corporate ${corporate.name} - ${url}`);
          return fetchHtml(url).then((html) => {
            return new Promise((resolve, reject) => {
              try {
                let $ = cheerio.load(html);
                let table_sumary_info = $("table.ficha");
                let assets_code = [];
                $("a.LinkCodNeg", table_sumary_info).each(
                  (i, link_asset_name) => {
                    assets_code.push($(link_asset_name).text());
                  }
                );
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

                let data = {
                  cnpj_number,
                  assets_code,
                  main_activity,
                  setorial_classes,
                  link: url,
                };
                const filter = { cvm_code: corporate.cvm_code };
                resolve({ filter, data });
              } catch (err) {
                reject(err);
              }
            }).then((filter_and_data) => {
              return models.CorporateInfo.updateOne(filter_and_data.filter, {
                $set: filter_and_data.data,
              });
            });
          });
        })
      );
    });
  }
}

module.exports = Scraper;
