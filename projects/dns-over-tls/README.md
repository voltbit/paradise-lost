## DNS over TLS proxy
### Take-home project for a job interview

Task list
- [ ] Setup a different network namespace and start the process there
- [ ] Listen for TCP and UDP port 53 requests
- [ ] Send the DNS request via TLS

To run:
`sudo -E go run . # to preserve the GOPATH`

### Setup

- I am using a docker container to run the proxt with 53/udp and 53/tcp exposed on
the global namespace with different ports
- Using a hot-reload library for working inside the container (ContainerDaemon)

### About the process and failed attempts

I started the project with the intention to make the proxy work with UDP first because it was more challenging since
it required some setup as to not overlap with the global namespace what was already using 53/udp. I tried to make
a separate net namespace and make the proxy part of that namespace but after spending time setting up virtual interfaces,
bridges and raeding about fallacies of Setns() and Clone() in Golang I decided it takes too long and went for the Docker
setup.
