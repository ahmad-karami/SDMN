# Usage: ping.sh <node1-4> <node1-4|router>
# Sends four pings from the source node's namespace to the destination address.
echo "ping from $1 to $2"

if [ "$1" = "node1" ]
then
	if [ "$2" = "router" ]
	then 
	 ip netns exec ns1 ping 172.0.0.1 -c4
	fi 
	if [ "$2" = "node2" ]
	then 
	 ip netns exec ns1 ping 172.0.0.3 -c4
	fi 	
	if [ "$2" = "node3" ]
	then 
	 ip netns exec ns1 ping 10.10.0.2 -c4
	fi 
	if [ "$2" = "node4" ]
	then 
	 ip netns exec ns1 ping 10.10.0.3 -c4
	fi	   
fi
if [ "$1" = "node2" ]
then
	if [ "$2" = "router" ]
	then 
	 ip netns exec ns2 ping 172.0.0.1 -c4
	fi 
	if [ "$2" = "node1" ]
	then 
	 ip netns exec ns2 ping 172.0.0.2 -c4
	fi 	
	if [ "$2" = "node3" ]
	then 
	 ip netns exec ns2 ping 10.10.0.2 -c4
	fi 
	if [ "$2" = "node4" ]
	then 
	 ip netns exec ns2 ping 10.10.0.3 -c4
	fi	   
fi
if [ "$1" = "node3" ]
then
	if [ "$2" = "router" ]
	then 
	 ip netns exec ns3 ping 10.10.0.1 -c4
	fi 
	if [ "$2" = "node1" ]
	then 
	 ip netns exec ns3 ping 172.0.0.2 -c4
	fi
	if [ "$2" = "node2" ]
	then 
	 ip netns exec ns3 ping 172.0.0.3 -c4
	fi 	
	if [ "$2" = "node4" ]
	then 
	 ip netns exec ns3 ping 10.10.0.3 -c4
	fi	   
fi
if [ "$1" = "node4" ]
then
	if [ "$2" = "router" ]
	then 
	 ip netns exec ns4 ping 10.10.0.1 -c4
	fi 
	if [ "$2" = "node1" ]
	then 
	 ip netns exec ns4 ping 172.0.0.2 -c4
	fi
	if [ "$2" = "node2" ]
	then 
	 ip netns exec ns4 ping 172.0.0.3 -c4
	fi 	
	if [ "$2" = "node3" ]
	then 
	 ip netns exec ns4 ping 10.10.0.2 -c4
	fi	   
fi