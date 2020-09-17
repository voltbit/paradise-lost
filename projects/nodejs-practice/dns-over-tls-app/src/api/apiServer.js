const express = require('express')

function startServer() {
  const port = 11222
  const addr = 'localhost'
  const app = express()
  app.listen(port, addr, () => {
    console.log(`API setup on {$localhost}:{$port}`)
  })
  app.get('/', (req, res) => {
    res.send(`Hello from the other side!!!`)
  })
};

module.exports.startServer = startServer;
