const portfolioModel = require("../models/portfolio");

exports.getAllData = async () => {
    return await portfolioModel.find();
};
