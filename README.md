# Static IPAM Driver
### IP address non-management for libnetwork

Static overlapping subnets are not a supported feature of the standard libnetwork IPAM driver. This driver is designed to do as little management as possible, giving the user full control over the IP addresses assigned to their containers. This is crucial when created a number of docker networks seperate on layer 2 and you want overlapping layer 3 (IP) domains.
