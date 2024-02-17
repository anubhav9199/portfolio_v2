const connectionString = process.env.MONGO_URI || "mongodb://localhost/CRUD";
const mongoose = require("mongoose");

//configure mongoose
const conn =  await mongoose.connect(
    connectionString,
    {
        useNewUrlParser: true,
        useUnifiedTopology: true,
    },
    (err) => {
        if (err) {
            console.log(err);
        } else {
            console.log("Connected to MongoDB");
        }
    }
);


let db = conn.db("sample_training");

export default db;
