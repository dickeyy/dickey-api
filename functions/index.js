const functions = require("firebase-functions");

// Express Server Stuff
const express = require("express");
const app = express();
const port = 3000;

// Random number
app.get('/randnum', (req, res) => {
    res.send({
        "number": Math.floor(Math.random() * 10000)
    })
})

// // Start server
// app.listen(port, () => {
//     console.log(`live on http://localhost:${port}`)
// })

// Firebase Function Export
exports.app = functions.https.onRequest(app)