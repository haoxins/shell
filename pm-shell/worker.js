'use strict'

const http = require('http')

http.createServer(function(req, res) {
  res.end('end')
}).listen(3000)
