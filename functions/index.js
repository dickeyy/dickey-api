const functions = require("firebase-functions");

const express = require("express");
const app = express();

// Home page
app.get('/', (req, res) => {
    res.send('Hello')
})

// Firebase Function Export
exports.app = functions.https.onRequest(app)