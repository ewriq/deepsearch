async function redditSearch(term) {
    const res = await fetch(`https://www.reddit.com/search.json?q=${encodeURIComponent(term)}&limit=3`);
    const json = await res.json();
  
    return json.data.children.map(post => ({
      source: 'reddit',
      title: post.data.title,
      snippet: post.data.selftext || post.data.title,
      link: `https://www.reddit.com${post.data.permalink}`
    }));
  }

  
  module.exports = {    
    redditSearch
  };