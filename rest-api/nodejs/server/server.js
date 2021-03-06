const express = require('express');
const bodyParser = require('body-parser');
const dbConfig = require('./config/database.config.js');
const mongoose = require('mongoose');


mongoose.Promise = global.Promise;

// Connect to the database.
mongoose.connect(dbConfig.url, {
    useNewUrlParser: true
}).then(() => {
    console.log("Successfully connected to the database");
}).catch(err => {
    console.log("Could not connect to the database. Exiting now...", err);
    process.exit();
});


// Create express app.
const app = express();

// Parse requests of content-type - application/x-www-form-urlencoded.
app.use(bodyParser.urlencoded({ extended: true }));

// Parse requests of content-type - application/json.
app.use(bodyParser.json())

// Define a route.
app.get('/', (req, res) => {
    res.json({"message": "Welcome to the EasyNotes application."});
});


// Require notes routes.
require('./app/routes/note.routes.js')(app);


app.listen(3000, () => {
    console.log("Server is Listening on port 3000");
});
