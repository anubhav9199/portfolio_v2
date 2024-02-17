const { getAllData } = require("../services/portfolioService");
const portfolioService = require("../services/portfolioService");

exports.getAllData = async (req, res) => {
    try {
        const portfolioData = await portfolioService.getAllBlogs();
        res.json({ data: portfolioData, status: "success" });
    } catch (err) {
        res.status(500).json({ error: err.message });
    }
};
