const { Wikipedia, Wikidata, DBpedia, InternetArchive } = require("./data");

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
        const duckDuckGoResults = await DuckDuckGo(term);
        results.push(...duckDuckGoResults);
    } catch (e) {
        console.error("DuckDuckGo hata:", e.message);
    }

    return results;
}

module.exports = deepSearch;