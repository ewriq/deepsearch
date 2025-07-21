const puppeteer = require('puppeteer-extra');
const AdblockerPlugin = require('puppeteer-extra-plugin-adblocker');

puppeteer.use(AdblockerPlugin({ blockTrackers: true }));

const urls = {
  bing: (q) => `https://www.bing.com/search?q=${encodeURIComponent(q)}`,
  yandex: (q) => `https://yandex.com/search/?text=${encodeURIComponent(q)}`,
  yahoo: (q) => `https://search.yahoo.com/search?p=${encodeURIComponent(q)}`
};

const selectors = {
  bing: '.b_caption p',
  yandex: '.OrganicSnippet-Content span',
  yahoo: '.compText p'
};

async function browser(engine, term) {
  if (!urls[engine]) throw new Error(`Desteklenmeyen arama motoru: ${engine}`);

  const browser = await puppeteer.launch({
    headless: 'new',
    args: ['--no-sandbox', '--disable-setuid-sandbox']
  });

  const page = await browser.newPage();
  await page.setUserAgent(
    'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114 Safari/537.36'
  );

  try {
    await page.goto(urls[engine](term), { waitUntil: 'domcontentloaded', timeout: 50000 });

    await page.waitForSelector(selectors[engine], { timeout: 5000 });
    const snippet = await page.$eval(selectors[engine], el => el.innerText);
    await browser.close();
    return { source: engine, snippet };
  } catch (e) {
    console.warn(`[${engine}] açıklama alınamadı:`, e.message);
    await browser.close();
    return null;
  }
}


module.exports = { browser };
