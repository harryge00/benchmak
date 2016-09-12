var elasticsearch = require('elasticsearch');
const LineByLineReader = require('line-by-line');
var client = new elasticsearch.Client({
  host: process.env.ES || 'localhost:9200'
  // log: 'trace'
});

var total = process.env.TOTAL;
var count = 0;
var miss = 0;
var index = process.env.INDEX || "logstash-*";
var totalTime = 0;
var totalQueries = 0;
if (!process.env.FILE || !process.env.POD_NAME) {
  console.log("Env FILE is not defined");
  process.exit(1);
}
console.log(process.env.FILE, total);
const lr = new LineByLineReader(process.env.FILE);
var pod_name = process.env.POD_NAME;

lr.on('error', function (err) {
     // 'err' contains error object
     console.error(err);
});
var concurrent = 0;
lr.on('line', function (line) {
  lr.pause();
  // if(concurrent > 5) {
  //   setTimeout(null, 60);
  // }
  // concurrent++;
  var log = JSON.parse(line).log;
  client.search({
    "index": index,
    "body": {
      "query": {
        "bool":{
          "must": [
            {"match":{"log": log}},
            {"match":{"pod_name": pod_name}}
          ]
        }
      }
    }
  }).then(function (body) {
    lr.resume();
    concurrent--;
    totalTime += body.took;
    totalQueries++;
    if(body.hits.total == 1) {
      count++;
    } else if(body.hits.total > 1) {
      console.log(log, " has hits:", body.hits.total);
      count++;
    } else {
      console.error("miss ", log);
      miss++;
    }
    if( (totalQueries > total - 10) || (totalQueries % 2000 == 0) ) {
      console.info("Total:", totalQueries, "Miss:", miss, "hits", count);
      console.info("average time for executing the search:", totalTime/totalQueries, "ms");
    }
  }, function (error) {
    console.error("search error!", error);
    console.trace(error);
    process.exit(1);
  });
});

lr.on('end', function () {
     // All lines are read, file is closed now.
     console.log("done reading");
});
