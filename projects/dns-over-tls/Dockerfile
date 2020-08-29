FROM golang:latest
# COPY . /app
WORKDIR /app
RUN go get github.com/githubnemo/CompileDaemon
RUN go get golang.org/x/net/dns/dnsmessage

ENTRYPOINT CompileDaemon --build="go build -o dns-proxy" --command="./dns-proxy"
