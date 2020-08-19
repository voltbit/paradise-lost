## OUTDATED - check memdb2
## Memdb - in-memeory HA database - interview project

This is a 'take home' project assignment received as part of a Golang programmer interview.

---

### Tasks

- [ ] API
    - [x] Add text to the DB
    - [x] Retrieve number of occurences for a particular string
    - [ ] OpenAPI specification (swagger)
- [ ] Can run as system daemon, integrate with systemd
- [x] Activity log
- [ ] Data change logging
- [ ] Data dump to file and recovery on restart
- [ ] HA setup with eventual consistency
- [ ] Unit test coverage
- [ ] Performance tests

---

### HA setup - Source/Replica setup

- Model - asynchronous, single source node with multiple read-only replica nodes
- Source node (Leader) - receives all the requests from a client and replicates them on the Replica nodes
- Replica node - receives write requests from the Source node
- Gossip protocol: P2P gRPC; must determine the leader in case of failure; propagate the identity of all nodes in the cluster
- Replication using log sequence numbers
- Async communication and data replication between nodes
- When joining the cluster, a node will get as input a file with existing nodes (Redis CLUSTER MEET)[https://redis.io/commands/cluster-meet]

About design choices and reasoning

1. Why gRPC for the P2P gossip protocol?
To learn how the protocol works and because it is easy to map methods and objects to the data being transmitted. Probably less efficient than using plain UDP/TCP.

2. The data type used for storing data change log numbers is a plain int64 out of convenience, a more complete solution should look for another data type that is larger or used in combination with a reset mechanism.

https://www.alibabacloud.com/blog/in-depth-analysis-of-redis-cluster-gossip-protocol_594706
https://cristian.regolo.cc/2015/09/05/life-in-a-redis-cluster.html
