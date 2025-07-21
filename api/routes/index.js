const express = require("express");
const router = express.Router();
const deepSearch = require("../pkg/utils");

router.post("/search/", async (req, res) => {
  const term = req.body.term;
  if (!term) return res.status(400).json({ error: "Arama terimi gerekli" });

  try {
    const descriptions = await deepSearch(term);
    const texts = Array.isArray(descriptions)
      ? descriptions
          .map(item => item.content || item.description || JSON.stringify(item))
          .filter(Boolean)
          .join("\n\n")
      : descriptions;
      console.log(texts);
      
    res.type("text/plain").send(texts);
  } catch (err) {
    res.status(500).json({ error: "Arama sırasında hata oluştu", details: err.message });
  }
});

module.exports = router;
//  // const { browser } = require("../pkg/browser");