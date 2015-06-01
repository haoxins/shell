'use strict'

const cluster = require('cluster')

cluster.setupMaster({
  exec: 'worker.js'
})

for (let i = 0; i < 4; i++) {
  cluster.fork()
}
