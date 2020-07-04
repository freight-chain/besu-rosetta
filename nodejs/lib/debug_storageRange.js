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

let postData = JSON.stringify({"jsonrpc":"2.0","method":"debug_storageRangeAt","params":["0x2b76b3a2fc44c0e21ea183d06c846353279a7acf12abcc6fb9d5e8fb14ae2f8c",0,"0x0e0d2c8f7794e82164f11798276a188147fbd415","0x0000000000000000000000000000000000000000000000000000000000000000",1],"id":1});

req.write(postData);

req.end();