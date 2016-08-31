var elasticsearch = require('elasticsearch');
var client = new elasticsearch.Client({
  host: 'localhost:9200',
  log: 'trace'
});

var total = 0;
var count = 0;
var miss = 0;
var index = "logstash-*"
var totalTime = 0;
var totalQueries = 0;

var msg="12334_advbdasdfadsff"
client.search({
  "index": "logstash-*",
  "body": {
    "query": {
      "match":{
        "msg":msg
      }
    }
  }
}).then(function (body) {
  console.info(body.hits.hits);
  totalTime += body.took;
  totalQueries++;
  if(body.hits.total == 1) {
    count++;
  } else if(body.hits.total > 1) {
    console.log(msg, " has hits:", body.hits.total);
  } else {
    console.error("miss ", msg);
    miss++;
  }
}, function (error) {
  console.trace(error.message);
});

console.info("average time for executing the search:", totalTime/totalQueries, "ms")
