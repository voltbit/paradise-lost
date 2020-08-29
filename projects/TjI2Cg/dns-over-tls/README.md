## DNS over TLS proxy

### Results and overall description

- The proxy runs inside a Docker container with the 53 port exposed for both UDP and TCP
- It can handle multiple concurrent requests on both TCP and UDP
- To develop easily with a docker container I used a hot-swap package that reloads the changes on save (ContainerDaemon)
- From a design/archiectural perspective the proxy is made of two parts: a **DNS server** that listens for incomming DNS
requests on both TCP and UDP and a **TLS server** that receives the requests from the DNS server and handles the connection
to the external DNS resolver
- To handle multiple concurrent requests I linked the two servers with a channel, spawn a new goroutine for each request
on both servers and create a channel between the corresponding goroutines
- Used libraries - decided not to implement the DNS message management because I felt it took too long
and was not the relevent part of the exercies - used `golang.org/x/net/dns/dnsmessage` instead to parse the network byte
slices

Questions

- Security concerns when integrating with infrastructure

I did not thorowly handle the format of the DNS messages and made the assumption that they are limited to 512
bytes instead of the maximum length of 65535 when supporting DNSSEC, so this is something that must be addressed
before production-ready. Then there is also the issue of trust when it comes to relying on an external service (in this case CloudFlare) that provides support for TLS based DNS queries.

- Integration in a distributed, microservices-oriented architecture

The proxy should be easy to scale horizontally and ideally the load of the microservices should be spread evenly among multiple running instances of the proxy. In the context of Kubernetes for example it could be placed behind a Service and deployed with a ReplicaSet or even a DaemonSet.

---

### Usage

Start the proxy:

`docker-compose up`

Test:

```
dig www.google.com @localhost -p 11222
dig +tcp www.google.com @localhost -p 11222
```

or use the benchmark test from the `perf` directory

```
cd app/perf
go test -bench=.
```

### About the process and failed attempts

Initially I did not want to use Docker (didn't notice the required Dockerfile). Instead I tried to make a separate net namespace
and make the proxy part of that namespace but after spending time setting up virtual interfaces, bridges and reading about
fallacies of Setns() and Clone() in Golang I decided it takes too long and went for the Docker setup.

