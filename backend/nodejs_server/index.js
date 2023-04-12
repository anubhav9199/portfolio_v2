const express = require("express");
const app = express();

//middleware
app.use(express.json());

app.listen(3001, () => {
    console.log("Server is running on port 3001");
});

const portfolioRouter = require("./routes/portfolioRoutes");

app.use("/api/portfolio", portfolioRouter);


module.exports = app;

