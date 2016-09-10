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
if (!process.env.FILE) {
  console.log("Env FILE is not defined");
  process.exit(1);
}
console.log(process.env.FILE, total);
const lr = new LineByLineReader(process.env.FILE);

lr.on('error', function (err) {
     // 'err' contains error object
     console.error(err);
});
lr.on('line', function (line) {
  // lr.pause();
  var log = JSON.parse(line).log;
  client.search({
    "index": index,
    "body": {
      "query": {
        "match":{
          "log": log
        }
      }
    }
  }).then(function (body) {
    totalTime += body.took;
    totalQueries++;
    if(body.hits.total == 1) {
      count++;
    } else if(body.hits.total > 1) {
      // console.log(log, " has hits:", body.hits.total);
      count++;
    } else {
      console.error("miss ", log);
      miss++;
    }
    // lr.resume();
    if( (totalQueries > total - 10) || (totalQueries % 5000 == 0) ) {
      console.info("Total:", totalQueries, "Miss:", miss, "hits", count);
      console.info("average time for executing the search:", totalTime/totalQueries, "ms");
    }
  }, function (error) {
    console.error("search error!");
    console.trace(error);
    // lr.resume();
  });
});

lr.on('end', function () {
     // All lines are read, file is closed now.
     console.log("done reading");
});
