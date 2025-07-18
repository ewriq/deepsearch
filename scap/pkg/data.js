const axios = require("axios");

async function Wikidata(term) {
  const sparql = `
    SELECT ?description WHERE {
      ?item rdfs:label "${term}"@en.
      OPTIONAL { ?item schema:description ?description. FILTER (lang(?description) = "en") }
      FILTER (lang(?description) = "en")
    } LIMIT 5
  `;

  const url = `https://query.wikidata.org/sparql?format=json&query=${encodeURIComponent(sparql)}`;
  const { data } = await axios.get(url);

  return data.results.bindings
    .map(obj => obj.description?.value)
    .filter(Boolean); 
}


async function Wikipedia(term) {
  const url = `https://en.wikipedia.org/api/rest_v1/page/summary/${encodeURIComponent(term)}`;
  const { data } = await axios.get(url);

  return data.extract || "";
}


async function DBpedia(term) {
  const sparql = `
    SELECT ?abstract WHERE {
      <http://dbpedia.org/resource/${encodeURIComponent(term)}> dbo:abstract ?abstract .
      FILTER (lang(?abstract) = 'en')
    } LIMIT 1
  `;

  const url = `http://dbpedia.org/sparql?query=${encodeURIComponent(sparql)}&format=json`;
  const { data } = await axios.get(url);

  if (data.results.bindings.length > 0) {
    return data.results.bindings[0].abstract.value;
  }
  return "";
}


async function InternetArchive(term) {
  const url = `https://archive.org/advancedsearch.php?q=${encodeURIComponent(term)}&fl[]=description&rows=3&page=1&output=json`;
  const { data } = await axios.get(url);

  return data.response.docs
    .map(doc => doc.description)
    .filter(Boolean);
}


async function DuckDuckGo(term) {
    const url = `https://api.duckduckgo.com/?q=${encodeURIComponent(term)}&format=json&no_redirect=1&no_html=1&skip_disambig=1`;
    const { data } = await axios.get(url);
  
    const descriptions = [];
  
    if (data.AbstractText && data.AbstractText.length > 0) {
      descriptions.push(data.AbstractText);
    }

    if (Array.isArray(data.RelatedTopics)) {
      data.RelatedTopics.forEach(topic => {
        if (topic.Text) descriptions.push(topic.Text);
        else if (topic.Topics) {
          topic.Topics.forEach(subTopic => {
            if (subTopic.Text) descriptions.push(subTopic.Text);
          });
        }
      });
    }
  
    return descriptions.filter(Boolean);
  }
  

module.exports = {
  Wikidata,
  Wikipedia,
  DBpedia,
  InternetArchive,
  DuckDuckGo
};
