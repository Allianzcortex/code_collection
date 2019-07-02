

const express = require('express')
const app = express()

var MongoClient = require('mongodb').MongoClient;
MongoClient.connect('mongodb://h_wang:A00431268@127.0.0.1:27017/h_wang', function(err, db) {
    if (err) {
        console.error(err);
    }
    var collection = db.collection('university');
    collection.find().toArray(function(err, docs) {
        console.log(docs);
    });
    console.log("connect successfully")
});

app.get('/', (req, res) => res.send('Hello World!'))