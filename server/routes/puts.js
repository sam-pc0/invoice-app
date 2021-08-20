const express = require("express");
const app = express();
const cors = require("cors");
const db = require("../lib/db.js");
app.use(cors());

function putOwner(req, res) {
  db.query(
    `UPDATE owner SET (templateId, ownersname, ownerownersaddress, ownerscity, ownersphone, ownersemail) VALUES ('${req.body.templateId}', '${req.body.owner.name}', '${req.body.address}','${req.body.location}','${req.body.phone}','${req.body.altPhone}','${req.body.email}');`,
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

function putBidProposal(req, res) {
  const templateId = req.body.templateId;
  const name_bid = req.body.name;
  const desc_bid = req.body.description;
  const number = req.body.number;
  const proyectnameaddr = req.body.projectNameNAddress;
  const scopeofwork = req.body.specificationNStimates;
  const not_included = req.body.notIncluded;
  const weproposeprice = req.body.totalSum;
  const weproposedays = req.body.withdrawnDays;
  const weproposedate = req.body.withdrawnDate;

  let data = {
    templateId,
    name_bid,
    desc_bid,
    number,
    proyectnameaddr,
    scopeofwork,
    not_included,
    weproposeprice,
    weproposedays,
    weproposedate,
  };
  let sql = "INSERT INTO bid_proposal SET ?";
  db.query(sql, data, function (error, results) {
    if (error) {
      throw error;
    } else {
      res.send(results);
    }
  });
}

function putInvoItem(req, res) {
  const item = req.body.item;
  const description = req.body.description;
  const amount = req.body.amount;

  let data = {
    item,
    description,
    amount,
  };
  let sql = "INSERT INTO invoice_item SET ?";
  db.query(sql, data, function (error, results) {
    if (error) {
      throw error;
    } else {
      res.send(results);
    }
  });
}
function putInvo(req, res) {
  const inv_proyectname = req.body.name;
  const inv_proyectaddress = req.body.projectNameNAddress;
  const ownerid_Fk_inv = req.body.owner;
  const intem_C_fk = req.body.itemsList;
  const totalinvoice = req.body.total;
  const inv_date = req.body.dateSubmitted;

  let data = {
    inv_proyectname,
    inv_proyectaddress,
    ownerid_Fk_inv,
    intem_C_fk,
    totalinvoice,
    inv_date,
  };
  let sql = "INSERT INTO invoice SET ?";
  db.query(sql, data, function (error, results) {
    if (error) {
      throw error;
    } else {
      res.send(results);
    }
  });
}

exports.putOwner = putOwner;
exports.putBidProposal = putBidProposal;
exports.putInvoItem = putInvoItem;
exports.putInvo = putInvo;
