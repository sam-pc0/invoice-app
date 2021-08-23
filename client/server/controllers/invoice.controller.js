const db = require("../lib/db");

const createInvoice = (req, res) => {
  console.log(req.body);
};

const allInvoices = (req, res) => {
  db.query("select * from invoice", (err, rows) => {
    if (err) {
      return res.stauts(500).send({
        message: err.message,
      });
    } else {
      res.send(rows);
    }
  });
};

module.exports = {
  createInvoice,
  allInvoices,
};
