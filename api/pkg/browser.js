const puppeteer = require('puppeteer-extra');
const AdblockerPlugin = require('puppeteer-extra-plugin-adblocker');

puppeteer.use(AdblockerPlugin({ blockTrackers: true }));

const urls = {
  google: (q) => `https://www.google.com/search?q=${encodeURIComponent(q)}`,
  bing: (q) => `https://www.bing.com/search?q=${encodeURIComponent(q)}`,
  yandex: (q) => `https://yandex.com/search/?text=${encodeURIComponent(q)}`,
  yahoo: (q) => `https://search.yahoo.com/search?p=${encodeURIComponent(q)}`
};

const selectors = {
  google: '#search .VwiC3b',
  bing: '.b_caption p',
  yandex: '.OrganicSnippet-Content span',
  yahoo: '.compText p'
};

async function browser(engine, term) {
  if (!urls[engine]) throw new Error(`Desteklenmeyen arama motoru: ${engine}`);

  const browser = await puppeteer.launch({ headless: 'new' });
  const page = await browser.newPage();

  try {
    await page.goto(urls[engine](term), { waitUntil: 'domcontentloaded' });

    const snippet = await page.$eval(selectors[engine], el => el.innerText);
    await browser.close();
    return { source: engine, snippet };
  } catch (e) {
    console.warn(`${engine} açıklama alınamadı:`, e.message);
    await browser.close();
    return null;
  }
}

module.exports = { browser };
