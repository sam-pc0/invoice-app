const express = require("express");
const cors = require("cors");
const invoiceRouter = require("./routes/invoice.routes");
const loginRouter = require("./routes/login.routes");

const app = express();

const PORT = process.env.PORT || 3000;
app.set("port", PORT);

app.use(express.json());
app.use(express.urlencoded({ extended: true }));
app.use(cors());

app.use("/api/invoices", invoiceRouter);
app.use("/api/login", loginRouter);

module.exports = app;
