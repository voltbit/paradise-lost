const dgram = require('dgram');
const dnsServer = require('.dnsServer');
const tlsServer = require('.tlsServer');

function startDNSOverTLSProxy() {
  dnsServer.startServer()
  tlsServer.startServer()
}

startDNSOverTLSProxy()
