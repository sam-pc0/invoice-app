const mysql = require("mysql");

const connection = mysql.createConnection({
  host: "64.227.28.17",
  user: "root",
  database: "invoiceappdb",
  password: "Ops.",
});

connection.connect();
module.exports = connection;
