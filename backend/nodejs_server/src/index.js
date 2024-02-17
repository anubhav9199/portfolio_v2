const express = require("express");
const app = express();

//middleware
app.use(express.json());

app.listen(8003, () => {
    console.log("Server is running on port 8003");
});

const portfolioRouter = require("../src/routes/portfolioRoutes");

app.use("/api/portfolio", portfolioRouter);


module.exports = app;

