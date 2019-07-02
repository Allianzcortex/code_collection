
/**
 * tips:
 * add more description inside the function
 */
const port = 8171;
const mongo_connection = {
    'user': 'XX',
    'password': 'XXX',
    'db': 'XXX',
    'mongo_host':'127.0.0.1',
    // 'mongo_host': 'localhost',
    'mongo_port': '27017',
    'mongo_collection': 'university'
}
const express = require('express')
const app = express()


var allowCrossDomain = function (req, res, next) {
    res.header('Access-Control-Allow-Origin', '*');
    res.header('Access-Control-Allow-Methods', 'GET,PUT,POST,DELETE');
    res.header('Access-Control-Allow-Headers', 'Content-Type');
    next();
};

app.use(express.bodyParser());
app.use(allowCrossDomain);

app.use('/scripts', express.static(__dirname + '/scripts'));
app.use('/css', express.static(__dirname + '/css'));
app.use(express.static(__dirname));

// create mongodb connection
var mongo = require('mongodb');
var MongoClient = mongo.MongoClient;

//var mongo_string='mongodb://'+mongo_connection['user']+':'+mongo_connection['password']+'@'+mongo_connection['mongo_host']+':'+mongo_connection['mongo_port']
var mongo_string = 'mongodb://' + mongo_connection['mongo_host'] + ':' + mongo_connection['mongo_port']
mongo_string="mongodb://XX:XXXXXX@127.0.0.1:27017"
console.log(mongo_string)

const client = new MongoClient(mongo_string)

// a tricky way to save collection to global variable

var u_collection;
var a;

// TODO know why connection string cannot work but must hard-coded
MongoClient.connect("mongodb://h_wang:A00431268@127.0.0.1:27017/h_wang", function (err, db) {
    if (err)
        console.log(err);
    else {
        console.log('Mongo Conn....');
        this.u_collection = db.collection("university")
        this.a = 1
    }
});


app.get('/', (req, res) => res.send('Hello World!'))

var server = app.listen(port, () => console.log(`Example app listening on port ${port}!`))

app.post('/saveUniversity', function (request, response) {
    console.log(request.body);
    console.log("a = " + this.a);
    /**
     * parse all data . 
     * 1. Check whether it exists : https://stackoverflow.com/questions/8389811/how-to-query-mongodb-to-test-if-an-item-exists
     * 2. if not , insert one to collection
     */

    this.u_collection.count(request.body, function (err, count) {
        if (count >= 1) {
            response.send(200, "Records Already Exixts")
        } else {
            this.u_collection.insertOne(request.body, function (err, res) {
                if (err) throw err;
                console.log("1 document inserted");
                response.send(200, "Save University Successfully")
            })
        }
    })

})

app.post('/deleteUniversity', function (request, response) {
    console.log(request.body);

    this.u_collection.count(request.body, function (err, count) {
        if (count < 1) {
            response.send(200, "Records doesn't Exixts")
        } else {
            this.u_collection.remove(request.body, function (err, res) {
                if (err) throw err;
                console.log("1 document deleted");
                response.send(200, "Delete University Successfully")
            })
        }
    })

})


app.post('/searchUniversity', function (request, response) {

    this.u_collection.count({
        "Name": request.body.Name
    }, function (err, count) {
        if (count < 1) {
            response.send(400, "Records don't Exixts")
        } else {
            console.log(request.body.Name)
            this.u_collection.find({
                "Name": request.body.Name
            }).toArray(function (err, res) {
                console.log("result is " + res)
                response.send(200,res)
            })
        }
    })
})


app.post('/displayUniversity', function (request, response) {

    this.u_collection.count({
    }, function (err, count) {
        if (count < 1) {
            response.send(400, "Records don't Exixts")
        } else {
            console.log(request.body.Name)
            this.u_collection.find({
            }).toArray(function (err, res) {
                console.log("result is " + res)
                response.send(200,res)
            })
        }
    })
})