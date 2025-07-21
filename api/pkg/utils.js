const { Wikipedia, Wikidata, DBpedia, InternetArchive } = require("./data");
const puppeteer = require('puppeteer-extra');
const AdblockerPlugin = require('puppeteer-extra-plugin-adblocker');
const { browser } = require('./browser');
const { redditSearch } = require('./social');
const axios = require('axios');

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
        const redditResults = await redditSearch(term);
        results.push(...redditResults);
    } catch (e) {
        console.error("Reddit hata:", e.message);
    }

    const engines = ['bing', 'yandex', 'yahoo'];
    for (const engine of engines) {
        try {
            const result = await browser(engine, term);
            if (result) {
                results.push({
                    source: engine,
                    title: result.title || '',
                    snippet: result.snippet || '',
                    link: result.link || '',
                });
            }
        } catch (e) {
            console.error(`${engine} hata:`, e.message);
        }
    }

    try {
        const news = await axios.get(`https://news.google.com/rss/search?q=${encodeURIComponent(term)}`);
        if (news.data) {
            results.push({ source: 'Google News', snippet: '[RSS verisi alındı]', raw: news.data });
        }
    } catch (e) {
        console.error("Google News hata:", e.message);
    }

    try {
        const quora = await browser('bing', `${term} site:quora.com`);
        if (quora) results.push({ source: 'Quora', snippet: quora.snippet });
    } catch (e) {
        console.error("Quora hata:", e.message);
    }

    try {
        const medium = await browser('bing', `${term} site:medium.com`);
        if (medium) results.push({ source: 'Medium', snippet: medium.snippet });
    } catch (e) {
        console.error("Medium hata:", e.message);
    }

    try {
        const substack = await browser('bing', `${term} site:substack.com`);
        if (substack) results.push({ source: 'Substack', snippet: substack.snippet });
    } catch (e) {
        console.error("Substack hata:", e.message);
    }

    try {
        const hunt = await browser('bing', `${term} site:producthunt.com`);
        if (hunt) results.push({ source: 'ProductHunt', snippet: hunt.snippet });
    } catch (e) {
        console.error("ProductHunt hata:", e.message);
    }

    try {
        const hn = await axios.get(`https://hn.algolia.com/api/v1/search?query=${encodeURIComponent(term)}`);
        const hits = hn.data.hits.slice(0, 3).map(h => ({
            source: 'HackerNews',
            title: h.title,
            snippet: h.comment_text || h.title,
            link: h.url || `https://news.ycombinator.com/item?id=${h.objectID}`
        }));
        results.push(...hits);
    } catch (e) {
        console.error("HackerNews hata:", e.message);
    }

    try {
        const so = await axios.get(`https://api.stackexchange.com/2.3/search/advanced?order=desc&sort=relevance&q=${encodeURIComponent(term)}&site=stackoverflow`);
        const items = so.data.items.slice(0, 3).map(i => ({
            source: 'StackOverflow',
            title: i.title,
            snippet: i.title,
            link: i.link
        }));
        results.push(...items);
    } catch (e) {
        console.error("StackOverflow hata:", e.message);
    }

    try {
        const gh = await axios.get(`https://api.github.com/search/repositories?q=${encodeURIComponent(term)}`);
        const repos = gh.data.items.slice(0, 3).map(r => ({
            source: 'GitHub',
            title: r.full_name,
            snippet: r.description || '',
            link: r.html_url
        }));
        results.push(...repos);
    } catch (e) {
        console.error("GitHub hata:", e.message);
    }

    return results;
}

module.exports = deepSearch;
