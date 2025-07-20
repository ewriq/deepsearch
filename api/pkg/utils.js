const { Wikipedia, Wikidata, DBpedia, InternetArchive } = require("./data");
const Tor = require("./tor");
const puppeteer = require('puppeteer-extra');
const AdblockerPlugin = require('puppeteer-extra-plugin-adblocker');
const { browser } = require('./browser');

puppeteer.use(AdblockerPlugin());


async function deepSearch(term) {
    const results = [];

    try {
        const wikidataResults = await Wikidata(term);
        results.push(...wikidataResults);
    } catch (e) {
        console.error("Wikidata hata:", e.message);
    }

    try {
        const wikiResult = await Wikipedia(term);
        if (wikiResult) results.push(wikiResult);
    } catch (e) {
        console.error("Wikipedia hata:", e.message);
    }

    try {
        const dbpediaResult = await DBpedia(term);
        if (dbpediaResult) results.push(dbpediaResult);
    } catch (e) {
        console.error("DBpedia hata:", e.message);
    }

    try {
        const internetArchiveResults = await InternetArchive(term);
        results.push(...internetArchiveResults);
    } catch (e) {
        console.error("Internet Archive hata:", e.message);
    }

    try {
        const duckDuckGoResults = await Tor(term);
        results.push(...duckDuckGoResults);
    } catch (e) {
        console.error("Tor DuckDuckGo hata:", e.message);
    }


(async () => {
  const term = 'how does photosynthesis work';

  const engines = ['google', 'bing', 'yandex', 'yahoo'];

  for (const engine of engines) {
    const result = await browser(engine, term);
    if (result) {
      console.log(`[${result.source}]`, result.snippet);
    } else {
      console.log(`[${engine}] açıklama bulunamadı.`);
    }
  }
})();


    return results;
}

module.exports = deepSearch;
