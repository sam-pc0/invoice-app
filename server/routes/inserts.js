const express = require("express");
const app = express();
const cors = require("cors");
const db = require("../lib/db.js");
app.use(cors());

function insTempData(req, res) {
  db.query(
    `INSERT INTO template_data (temp_id, temp_name, temp_desc) VALUES ('${req.body.temp_id}', '${req.body.temp_name}', '${req.body.temp_desc}');`,
    //--insert the template data---
    (err, result) => {
      // user does not exists
      if (err) {
        throw err;
        return res.status(400).send({
          msg: err,
        });
      }
      return res.status(201).send({
        msg: "insertado!",
      });
    }
  );
}

function insBidProp(req, res) {
  db.query(
    `INSERT INTO bid_proposal (templateId_Fk) VALUES ('${req.body.temp_id}');`,
    //--insert the template data---
    (err, result) => {
      // user does not exists
      if (err) {
        throw err;
        return res.status(400).send({
          msg: err,
        });
      }
    }
  );
}

function insInvoice(req, res) {
  db.query(
    `INSERT INTO invoice (templateId_Fk) VALUES ('${req.body.temp_id}');`,
    //--insert the template data---
    (err, result) => {
      // user does not exists
      if (err) {
        throw err;
        return res.status(400).send({
          msg: err,
        });
      }
    }
  );
}

function insContInv(req, res) {
  db.query(
    `INSERT INTO contract_invoice (templateId_Fk) VALUES ('${req.body.temp_id}');`,
    //--insert the template data---
    (err, result) => {
      // user does not exists
      if (err) {
        throw err;
        return res.status(400).send({
          msg: err,
        });
      }
    }
  );
}

function matRec(req, res) {
  db.query(
    `INSERT INTO material_record (templateId_Fk) VALUES ('${req.body.temp_id}');`,
    //--insert the template data---
    (err, result) => {
      // user does not exists
      if (err) {
        throw err;
        return res.status(400).send({
          msg: err,
        });
      }
    }
  );
}


function instFirstOwn(req,res) {
    db.query(
        `INSERT INTO owner (ownersname) VALUES ('${req.body.temp_name}');`,
        //--insert the template data---
        (err, result) => {
          // user does not exists
          if (err) {
            throw err;
            return res.status(400).send({
              msg: err,
            });
          }
        }
      );
    
}

exports.insTempData = insTempData;
exports.insBidProp = insBidProp;
exports.insInvoice = insInvoice;
exports.insContInv = insContInv;
exports.matRec = matRec;
exports.instFirstOwn =instFirstOwn;
