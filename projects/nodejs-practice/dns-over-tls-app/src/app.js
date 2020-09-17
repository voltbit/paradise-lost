const dgram = require('dgram');
// const dnsServer = require('./proxy/dnsServer');
// const tlsServer = require('./proxy/tlsServer');
const apiServer = require('./api/apiServer')

function startDNSOverTLSProxy() {
  // dnsServer.startServer()
  // tlsServer.startServer()
  apiServer.startServer()
}

startDNSOverTLSProxy()
