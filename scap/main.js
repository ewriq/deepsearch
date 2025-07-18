const express = require("express");
const app = express();
const indexRouter = require("./routes/index");
const logger = require("./middleware/logger");
const deepSearch = require("./pkg/utils");


app.use(express.json());
app.use(logger); 


app.use("/api", indexRouter);


app.use((err, req, res, next) => {
  console.error("Hata:", err.stack);
  res.status(500).json({ error: "Internal Server Error" });
});

const PORT = process.env.PORT || 5000;
app.listen(PORT, () => {
  console.log(`🚀 Server listening on port ${PORT}`);
});
