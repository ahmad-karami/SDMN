# Linux Network Namespaces with Bridges and a Router

These shell scripts build a small virtual network on one Linux host from network namespaces, bridges, and veth pairs. Four node namespaces sit on two bridges: `ns1` and `ns2` on `br1` in the `172.0.0.0/24` subnet, and `ns3` and `ns4` on `br2` in the `10.10.0.0/24` subnet. A `router` namespace connects to both bridges and forwards packets between the two subnets. `create_topo.sh` builds the topology, `ping.sh` pings from one node to another node or to the router, and `del_ns.sh` removes the topology.

## How it works
`create_topo.sh` creates the namespaces and the bridges. It adds one veth pair per node and two for the router, moves one end of each pair into its namespace, and attaches the other end to a bridge. It then assigns the addresses, enables `net.ipv4.ip_forward` in `router`, and makes the router's address on each subnet the default gateway of that subnet's nodes. For example, `ping.sh node1 node3` sends four pings from `ns1` to `10.10.0.2` through `router`. `del_ns.sh` deletes the namespaces and the bridges. The scripts need root privileges.

### Design notes
Without the `router` namespace, the root namespace can route between the two subnets. IP forwarding is enabled in the root namespace, and one veth pair per subnet connects the root namespace to that subnet. The root-namespace end of each pair gets an address on its subnet, and the nodes on that subnet use it as their default gateway.

When the namespaces run on different servers that share a layer-2 network, each server connects its root namespace to its node namespaces in the same way. The root-namespace interfaces serve as the default gateways of the nodes, and each server enables IP forwarding. The layer-2 switch between the servers must handle ARP. Because the node subnets differ from the servers' own network, the servers also apply NAT: the source server translates packets and routes them to the destination server, which translates them for its node namespaces.
