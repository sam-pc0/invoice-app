// index.js
const express = require("express");
const app = express();
const cors = require("cors");
// add routes
const router = require("./routes/router");

// set up port
const PORT = process.env.PORT || 3000;
app.use(express.json());

app.use(express.urlencoded({ extended: true }));

app.use(cors());

app.use("/api", router);
// run server
app.listen(PORT, () => console.log(`Server running on port ${PORT}`));
