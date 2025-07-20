const axios = require("axios");
const { SocksProxyAgent } = require("socks-proxy-agent");

const agent = new SocksProxyAgent("socks5h://127.0.0.1:9050"); // Tor default SOCKS5 proxy

async function TorDuckDuckGo(query) {
    const url = `http://duckduckgogg42xjoc72x3sjasowoarfbgcmvfimaftt6twagswzczad.onion/?q=${encodeURIComponent(query)}&format=json&no_redirect=1&no_html=1`;

    try {
        const response = await axios.get(url, {
            httpAgent: agent,
            httpsAgent: agent,
            timeout: 10000,
            headers: {
                "User-Agent": "TorSearchBot/1.0"
            }
        });

        return [{ source: "TorDuckDuckGo", data: response.data }];
    } catch (error) {
        throw new Error("Tor DuckDuckGo bağlantısı başarısız: " + error.message);
    }
}

module.exports = TorDuckDuckGo;
