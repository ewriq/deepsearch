const express = require("express");
const router = express.Router();
const deepSearch = require("../pkg/utils");

router.get("/search/:term", async (req, res) => {
  const term = req.params.term;
  if (!term) return res.status(400).json({ error: "Arama terimi gerekli" });

  try {
    const descriptions = await deepSearch(term);
    res.json({ data: descriptions });
  } catch (err) {
    res.status(500).json({ error: "Arama sırasında hata oluştu", details: err.message });
  }
});

module.exports = router;
