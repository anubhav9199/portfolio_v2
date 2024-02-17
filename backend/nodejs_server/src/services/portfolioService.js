const portfolioModel = require("../models/portfolio");

getAllData = async () => {
    return await portfolioModel.find();
};

module.exports = getAllData