const express = require("express");
const app = express();
const cors = require("cors");
const db = require("../lib/db.js");
app.use(cors());

function getAllTempDat(req, res) {
  db.query("SELECT * FROM template_data", (error, filas) => {
    if (error) {
      throw error;
    } else {
      res.send(filas);
    }
  });
}

function getAllInvo(req, res) {
  db.query("SELECT * FROM invoice", (error, filas) => {
    if (error) {
      throw error;
    } else {
      res.send(filas);
    }
  });
}
function getAllMatRec(req, res) {
  db.query("SELECT * FROM material_record", (error, filas) => {
    if (error) {
      throw error;
    } else {
      res.send(filas);
    }
  });
}

function getAllMatRec(req, res) {
  db.query("SELECT * FROM bid_proposal", (error, filas) => {
    if (error) {
      throw error;
    } else {
      res.send(filas);
    }
  });
}

function getContInvo(req, res) {
  db.query("SELECT * FROM contract_invoice", (error, filas) => {
    if (error) {
      throw error;
    } else {
      res.send(filas);
    }
  });
}

function getTopTempDat(req, res) {
  db.query(
    "SELECT *  FROM template_data ORDER BY id_data DESC LIMIT 1",
    (error, filas) => {
      if (error) {
        throw error;
      } else {
        res.send(filas);
        filas.map((fruit, index) => {
          console.log(index, fruit.id_data);
          const dat = fruit.id_data;
          console.log(dat);
        });
      }
    }
  );
}

exports.getAllTempDat = getAllTempDat;
exports.getAllInvo = getAllInvo;
exports.getAllMatRec = getAllMatRec;
exports.getContInvo = getContInvo;
exports.getTopTempDat = getTopTempDat;
