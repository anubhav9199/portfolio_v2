const express = require("express");
const { getAllData } = require("../controllers/portfolioController");

const router = express.Router();

router.route("/get_data").get(getAllData);

module.exports = router;
