const mysql = require("mysql");

const connection = mysql.createPool({
  host: "64.227.28.17",
  user: "root",
  database: "invoiceappdb",
  password: "Ops.",
});

connection.getConnection(function (error) {
  if (error) {
    throw error;
  } else {
    console.log("conexion exitosa a la base de datos");
  }
});
module.exports = connection;
