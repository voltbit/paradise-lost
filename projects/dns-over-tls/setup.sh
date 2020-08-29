# https://blog.scottlowe.org/2013/09/04/introducing-linux-network-namespaces/

# create a new namespace
ip netns add blue
# add virtual interfaces to the namespace
ip link add blue0 type veth peer name blue1
# move blue1 in the blue namespace
ip link set blue1 netns blue
# add bridge
ip link add name blue-br type bridge
# bring bridge up
ip link set blue-br up
# bring blue0 (global ns) up
ip link set dev blue0 up
# bring blue1 (blue ns) up
ip netns exec blue ip link set dev blue1 up

ip addr add 10.1.0.1/24 dev blue0
ip netns exec blue ip addr add 10.0.0.1/24 dev blue1
