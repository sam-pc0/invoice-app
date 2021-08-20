// routes/router.js

const express = require("express");
const router = express.Router();
const app = express();
const bcrypt = require("bcryptjs");
const uuid = require("uuid");
const jwt = require("jsonwebtoken");

const db = require("../lib/db.js");
const userMiddleware = require("../middleware/users.js");
const cors = require("cors");
const insert = require("./inserts.js");
const get = require("./gets");

app.use(cors());

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

//---------------------------------------inser first temp dat an owner---------------

router.post("/invoice", (req, res) => {
  insert.insTempData(req, res);
  insert.instFirstOwn(req, res);
});

//----------------------------------------------switch to chose a template----------------
router.post("/:menu", (req, res) => {
  switch (req.params.menu) {
    case "1":
      insert.insTempData(req, res);
      insert.insBidProp(req, res);
      insert.instFirstOwn(req, res);
      break;
    case "2":
      insert.insTempData(req, res);
      insert.insInvoice(req, res);
      break;
    case "3":
      insert.insTempData(req, res);
      insert.insContInv(req, res);
      break;
    case "4":
      insert.insTempData(req, res);
      insert.matRec(req, res);
      break;
  }
});

//----------------------------------------------------------------------------------------
router.get("/:menu", (req, res) => {
  switch (req.params.menu) {
    case "5":
      get.getAllTempDat(req, res);
      break;
    case "6":
      get.getAllInvo(req, res);
      break;
    case "7":
      get.getAllMatRec(req, res);
      break;
    case "8":
      get.getContInvo(req, res);
      break;
    case "9":
      get.getTopTempDat(req, res);
      break;
  }
});
//----------------------------------------------------------------------------------------

//this is for templates
module.exports = router;
