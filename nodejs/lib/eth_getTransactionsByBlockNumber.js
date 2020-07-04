const http = require('follow-redirects').http;
const fs = require('fs');

let options = {
  'method': 'POST',
  'hostname': '127.0.0.1',
  'port': 8545,
  'path': '/',
  'headers': {
    'Content-Type': 'application/json'
  },
  'maxRedirects': 20
};

const req = http.request(options, (res) => {
  let chunks = [];

  res.on("data", (chunk) => {
    chunks.push(chunk);
  });

  res.on("end", (chunk) => {
    let body = Buffer.concat(chunks);
    console.log(body.toString());
  });

  res.on("error", (error) => {
    console.error(error);
  });
});

let postData = JSON.stringify({"jsonrpc":"2.0","method":"eth_getTransactionByBlockNumberAndIndex","params":["latest","0x0"],"id":1});

req.write(postData);

req.end();