# Delete the namespaces, then the bridges and the router's bridge-side veth ends.
ip netns del ns1
ip netns del ns2
ip netns del ns3
ip netns del ns4
ip netns del router


ip link set br1 down
brctl delbr br1
ip link set br2 down
brctl delbr br2
ip link set br3 down
brctl delbr br3
ip link set br4 down
brctl delbr br4
ip link delete veth_nsr_br1
ip link delete veth_nsr_br2
ip link delete veth_nsr_br3
ip link delete veth_nsr_br4

echo "all deleted"
