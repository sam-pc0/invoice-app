// routes/router.js

const express = require("express");
const router = express.Router();

const bcrypt = require("bcryptjs");
const uuid = require("uuid");
const jwt = require("jsonwebtoken");

const db = require("../lib/db.js");
const userMiddleware = require("../middleware/users.js");

router.post("/sign-up", userMiddleware.validateRegister, (req, res) => {
  db.query(
    `SELECT * FROM users WHERE LOWER(username) = LOWER(${db.escape(
      req.body.username
    )});`,
    (err, result) => {
      if (result.length) {
        return res.status(409).send({
          msg: "This username is already in use!",
        });
      } else {
        // username is available
        bcrypt.hash(req.body.password, 10, (err, hash) => {
          if (err) {
            return res.status(500).send({
              msg: err,
            });
          } else {
            // has hashed pw => add to database
            db.query(
              `INSERT INTO users (id, username, password, registered) VALUES ('${uuid.v4()}', ${db.escape(
                req.body.username
              )}, ${db.escape(hash)}, now())`,
              (err) => {
                if (err) {
                  throw err;
                  return res.status(400).send({
                    msg: err,
                  });
                }
                return res.status(201).send({
                  msg: "Registered!",
                });
              }
            );
          }
        });
      }
    }
  );
});

router.post("/login", (req, res) => {
  db.query(
    `SELECT * FROM users WHERE username = ${db.escape(req.body.username)};`,
    (err, result) => {
      // user does not exists
      if (err) {
        throw err;
        return res.status(400).send({
          msg: err,
        });
      }

      if (!result.length) {
        return res.status(401).send({
          msg: "Username or password is incorrect!",
        });
      }

      // check password
      bcrypt.compare(
        req.body.password,
        result[0]["password"],
        (bErr, bResult) => {
          // wrong password
          if (bErr) {
            throw bErr;
            return res.status(401).send({
              msg: "Username or password is incorrect!",
            });
          }

          if (bResult) {
            const token = jwt.sign(
              {
                username: result[0].username,
                userId: result[0].id,
              },
              "SECRETKEY",
              {
                expiresIn: "7d",
              }
            );

            db.query(
              `UPDATE users SET last_login = now() WHERE id = '${result[0].id}'`
            );
            return res.status(200).send({
              msg: "Logged in!",
              token,
              user: result[0],
            });
          }
          return res.status(401).send({
            msg: "Username or password is incorrect!",
          });
        }
      );
    }
  );
});

router.get("/secret-route", userMiddleware.isLoggedIn, (req, res, next) => {
  console.log(req.userData);
  res.send("This is the secret content. Only logged in users can see that!");
});

router.post("/owner", (req, res) => {
  const ownersname = req.body.owner.name;
  const ownerownersaddress = req.body.address;
  const ownerscity = req.body.location;
  const ownersstate = req.body.Precio;
  const ownerszip = req.body.CategoriaC_fk;
  const ownersphone = req.body.phone;
  const ownersalphone = req.body.altPhone;
  const ownersemail = req.body.email;

  let data = {
    ownersname,
    ownerownersaddress,
    ownerscity,
    ownersstate,
    ownerszip,
    ownersphone,
    ownersalphone,
    ownersemail,
  };
  let sql = "INSERT INTO owner ?";
  db.query(sql, data, function (error, results) {
    if (error) {
      throw error;
    } else {
      res.send(results);
    }
  });
});

router.post("/bid_proposal", (req, res) => {
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
});

router.post("/invoice_item", (req, res) => {
  const item = req.body.item;
  const description = req.body.description;
  const amount = req.body.amount;

  let data = {
    item,
    description,
    amount,
    data,
  };
  let sql = "INSERT INTO invoice_item SET ?";
  db.query(sql, data, function (error, results) {
    if (error) {
      throw error;
    } else {
      res.send(results);
    }
  });
});

router.post("/invoice", (req, res) => {
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
});

module.exports = router;
