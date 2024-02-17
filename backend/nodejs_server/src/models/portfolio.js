const mongoose = require("mongoose");
const Schema = mongoose.Schema;

const somethingSchema = new Schema({});

module.exports = mongoose.model("", somethingSchema);
