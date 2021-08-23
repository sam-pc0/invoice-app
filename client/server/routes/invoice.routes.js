const { Router } = require("express");
const invoiceController = require("../controllers/invoice.controller");

const router = Router();

router.post("/", invoiceController.createInvoice);
router.get("/", invoiceController.allInvoices);

module.exports = router;
