'use strict';

const http = require('http');

let app = http.createServer(function(req, res) {
  res.end('ok');
});

app.listen(3000);
